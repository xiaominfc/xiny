package conn

import(
    "github.com/xiaominfc/xiny/base"
)

type ClientConn struct {
    connection *base.BConn
    manager ClientConnManager
}

type ClientConnManager interface {
    HandleData(b []byte,conn *ClientConn)error
    OnClose(conn *ClientConn)
}

func NewClientConn(connection *base.BConn,manager ClientConnManager) *ClientConn{
    return &ClientConn{connection:connection, manager: manager}
}

func (this *ClientConn)GetConn() *base.BConn {
    return this.connection
}

func (this *ClientConn)handleData(data []byte) {
    this.manager.HandleData(data,this)
}

func (this *ClientConn) OnDataReaded(data []byte, err error) {
    this.handleData(data)
}

func (this *ClientConn) Send(b []byte) {
    if this.connection == nil {
        this.Close()
        return
    }

    _,err := this.connection.Send(b)
    if err != nil {
        this.Close()
    }
}

func (this *ClientConn) OnRead() {
    base.OnRead(this.connection,this)
}

func (this *ClientConn) GetSocketFd() int64 {
    fd,_ := this.connection.GetFd()
    return fd
}

func (this *ClientConn) Run() {
    go this.OnRead()
}

func (this *ClientConn) Close() {
    if this.connection != nil {
        this.connection.Close()    
    }
    this.manager.OnClose(this)
}
