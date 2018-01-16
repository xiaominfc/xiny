package conn

import(
    "github.com/xiaominfc/xiny/base"
    "github.com/xiaominfc/xiny/utils"
    "math/rand"
    "time"
    "errors"
)


const CACHESIZE = 1024

const RECONNECTMAXTIME = 5

type ServConnManager interface {
    HandleData(b []byte) error
    OnServConnAdd(server *ServConn)
    OnServConnClose(server *ServConn)
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
    needReconnet bool
    reconnectTime int
    task   *utils.Task
}



func NewConnect(host string, port int,manager ServConnManager) *ServConn {
    server := &ServConn{host:host, port:port,manager: manager,needReconnet:true, reconnectTime:0}
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


func (this *ServConn)reconnect(){
    for this.reconnectTime < RECONNECTMAXTIME {
        time.Sleep(time.Second * 5)
        conn, err := base.Connect(this.host,this.port)
        if err != nil {
            println("reconnect ok:  ",this.host,":",this.port)
            this.conn = conn
            this.reconnectTime = 0
            this.Run()
            this.manager.OnServConnAdd(this)
            break
        }else {
            this.reconnectTime ++
        }    
    }
    println("reconnect failed for:   ",this.host,":", this.port)
    this.manager = nil
    this.task.Stop()
    this.task = nil
}

func (this *ServConn)handleData(b []byte) {
    if this.manager != nil {
        if this.cacheData != nil {
            this.cacheData = append(this.cacheData,b...)
        }else {
            this.cacheData = b
        }
        err := (this.manager).HandleData(this.cacheData)
        if err == nil {
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

func (this *ServConn) Close() {
    this.conn.Close()
    this.conn = nil
    this.manager.OnServConnClose(this) 
    this.task.Pause()
    if(this.needReconnet) {
        go this.reconnect()
    }
}

func (this *ServConn) Run() {
    go this.OnRead()
    if this.task != nil {
        this.task.Start()
    } else {
        this.task = utils.AddTask(5*time.Second, this.OnTime)    
    }
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
    if pdu != nil {
        if(this.handler != nil) {
            go this.handler.HandlePdu(pdu,this)
        }
        return nil
    }
    return errors.New("")
}

func (manager *DefaultManager) OnServConnAdd(servconn *ServConn) {
    println("new server connect ok")
    manager.ServConnList.Add(servconn)
}

func (manager *DefaultManager) OnServConnClose(servconn *ServConn) {
    manager.ServConnList.Remove(servconn)
}


func (this *DefaultManager) OnTimeForServConn(servconn *ServConn) {
    if this.handler != nil {
        this.handler.OnTimeWork(servconn, this)
    }
}

func (this *DefaultManager) GetServConn() *ServConn {
    size := this.ServConnList.Size()
    if(size == 0) {
        return nil
    }
    index := rand.Intn(size)
    server,err := this.ServConnList.Get(index)
    if err != nil {
        return nil
    }
    return server.(*ServConn)
}


