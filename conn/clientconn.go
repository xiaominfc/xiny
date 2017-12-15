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

func (clientConn *ClientConn)GetConn() *base.BConn {
    return clientConn.connection
}

func (clientConn *ClientConn)handleData(data []byte) {
    clientConn.manager.HandleData(data,clientConn)
}

func (clientConn *ClientConn) OnDataReaded(data []byte, err error) {
    clientConn.handleData(data)
}

func (clientConn *ClientConn) Send(b []byte) {
    if clientConn == nil {
        return
    }

    if clientConn.connection == nil {
        clientConn.Close()
        return
    }

    _,err := clientConn.connection.Send(b)
    if err != nil {
        clientConn.Close()
    }
}

func (clientConn *ClientConn) OnRead() {
    base.OnRead(clientConn.connection,clientConn)
}

func (clientConn *ClientConn) GetSocketFd() int64 {
    fd,_ := clientConn.connection.GetFd()
    return fd
}

func (clientConn *ClientConn) Run() {
    go clientConn.OnRead()
}

func (clientConn *ClientConn) Close() {
    println("close", clientConn.GetSocketFd())
    if clientConn == nil {
        return
    }

    if clientConn.connection != nil {
        clientConn.connection.Close()
    }
    clientConn.manager.OnClose(clientConn)
}
