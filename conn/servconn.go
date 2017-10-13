package conn

import(
    "github.com/xiaominfc/xiny/base"
    "github.com/xiaominfc/xiny/utils"
    "math/rand"
)

var servList = utils.NewArray()

const CACHESIZE = 1024

type ServConnHandler interface {
    HandleData(b []byte) error
}

type ServConn struct {
    conn *base.BConn
    host string
    port int
    handler ServConnHandler
    cacheData []byte
}



func NewConnect(host string, port int,handler ServConnHandler) *ServConn {
    server := &ServConn{host:host, port:port,handler: handler}
    server.conn,_ = base.Connect(host,port)
    return server
}

func AddNewServFor(host string, port int, handler ServConnHandler) *ServConn{
    server := NewConnect(host, port, handler)
    server.Run()
    if server != nil {
        servList.Add(server)
    }
    return server
}

func GetServConn() *ServConn{
   size := servList.Size()
   index := rand.Intn(size)
   server,err := servList.Get(index)
   if err!=nil {

   }
   return server.(*ServConn)
}

func (this *ServConn)handleData(b []byte) {
    if this.handler != nil {
        if this.cacheData != nil {
            println("cach not nil")
            this.cacheData = append(this.cacheData,b...)
        }else {
            println("cache is nil:", len(b))
            this.cacheData = b
        }
        err := (this.handler).HandleData(this.cacheData)
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



