package conn

import(
    "github.com/xiaominfc/xiny/base"
    "github.com/xiaominfc/xiny/utils"
    "math/rand"
    "time"
)


const CACHESIZE = 1024

type ServConnManager interface {
    HandleData(b []byte) error
    OnServConnAdd(server *ServConn)
    GetServConn() *ServConn
    OnTimeForServConn(server *ServConn)
}

type PduHandler interface {
    HandlePdu(pdu *base.Pdu, servConnManager ServConnManager)
    OnTimeWork(servConn *ServConn, servConnManager ServConnManager)
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

func (this *ServConn) OnTime() {
    this.manager.OnTimeForServConn(this)
}

func (this *ServConn) Run() {
    go this.OnRead()
    utils.AddTask(5*time.Second, this.OnTime)
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

func (this *DefaultManager) HandleData(b []byte) error {
    pdu := base.ReadPdu(b)
    if pdu != nil && this.handler != nil {
        this.handler.HandlePdu(pdu,this)
    }
    return nil
}

func (manager *DefaultManager) OnServConnAdd(servconn *ServConn) {
    println("new server connect ok")
    manager.ServConnList.Add(servconn)
}

func (this *DefaultManager) OnTimeForServConn(servconn *ServConn) {
    if this.handler != nil {
        this.handler.OnTimeWork(servconn, this)    
    }
}

func (this *DefaultManager) GetServConn() *ServConn {
    size := this.ServConnList.Size()
    index := rand.Intn(size)
    server,err := this.ServConnList.Get(index)
    if err!=nil {

    }
    return server.(*ServConn)
}


