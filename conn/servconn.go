package conn

import(
    "github.com/xiaominfc/xiny/base"
)


const CACHESIZE = 1024

type ServConnManager interface {
    HandleData(b []byte) error
    OnServConnAdd(server *ServConn)
    GetServConn() *ServConn
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

func AddNewServFor(host string, port int, manager ServConnManager) *ServConn{
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

func (this *ServConn) OnRead() {
    tmp_buf :=  make([]byte, CACHESIZE)
    buf := make([]byte, 0, 4096)
    buffer_size := 0
    for {
        count,err := this.conn.Reciv(tmp_buf)
        if count > 0 {
            buf = append(buf, tmp_buf[:count]...)
            buffer_size = buffer_size + count;
        }

        if err != nil || count < 4096 {
            if buffer_size > 0 {
                data := append([]byte(nil), buf...)
                this.handleData(data)
                buf = buf[:0]
                buffer_size = 0

            }
        }
    }
}

func (this *ServConn) Run() {
    go this.OnRead()
}

func (this *ServConn) Send(b []byte) {
    go func() {
      this.conn.Send(b)
    }()
}



