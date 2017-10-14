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

func (this *ClientConn) OnRead() {
    base.OnRead(this.connection,this)
}

func (this *ClientConn) Run() {
    go this.OnRead()
}