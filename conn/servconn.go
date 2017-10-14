package conn

import(
    "github.com/xiaominfc/xiny/base"
    "github.com/xiaominfc/xiny/utils"
    "math/rand"
)


const CACHESIZE = 1024

type ServConnManager interface {
    HandleData(b []byte) error
    OnServConnAdd(server *ServConn)
    GetServConn() *ServConn
}

type PduHandler interface {
    HandlePdu(pdu *base.Pdu)
}


type ServConn struct {
    conn *base.BConn
    host string
    port int
    manager ServConnManager
    cacheData []byte
}



func NewConnect(host string, port int,manager ServConnManager) *ServConn {
    server := &ServConn{host:host, port:port,manager: manager}
    server.conn,_ = base.Connect(host,port)
    return server
}

func AddNewServConnFor(host string, port int, manager ServConnManager) *ServConn{
    server := NewConnect(host, port, manager)
    server.Run()
    if server != nil && manager != nil {
        manager.OnServConnAdd(server)
    }
    return server
}


func (this *ServConn)handleData(b []byte) {
    if this.manager != nil {
        if this.cacheData != nil {
            println("cach not nil")
            this.cacheData = append(this.cacheData,b...)
        }else {
            println("cache is nil:", len(b))
            this.cacheData = b
        }
        err := (this.manager).HandleData(this.cacheData)
        if err == nil {
            println("finish work")
            this.cacheData = nil
        }
    }
}

func (this *ServConn) OnDataReaded(data []byte, err error) {
    this.handleData(data)
}

func (this *ServConn) OnRead() {
    base.OnRead(this.conn,this)
}

func (this *ServConn) Run() {
    go this.OnRead()
}

func (this *ServConn) Send(b []byte) {
    go func() {
      this.conn.Send(b)
    }()
}

type DefaultManager struct{
    ServConnList utils.Array
    handler   PduHandler
}

func NewDefaultManager(handler PduHandler) *DefaultManager {
    manager := &DefaultManager{handler:handler}
    manager.ServConnList = utils.NewArray()
    return manager
}

func (manager *DefaultManager) HandleData(b []byte) error {
    pdu := base.ReadPdu(b)
    if pdu != nil {
        manager.handler.HandlePdu(pdu)
    }
    return nil
}

func (manager *DefaultManager) OnServConnAdd(servconn *ServConn) {
    manager.ServConnList.Add(servconn)
}

func (manager *DefaultManager) GetServConn() *ServConn {
    size := manager.ServConnList.Size()
    index := rand.Intn(size)
    server,err := manager.ServConnList.Get(index)
    if err!=nil {

    }
    return server.(*ServConn)
}


