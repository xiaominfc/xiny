package base

import(
    "net"
    "fmt"
    "log"
)
const CACHESIZE = 1024

type BConn struct{
    Conn *net.TCPConn
}

type IConnIO interface {
    OnDataReaded(b []byte, err error)
    Close()
}


func Connect(host string, port int) (*BConn,error) {
    laddr, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:%d", host,port))
    if err != nil {
        log.Fatal(err)
        return nil, err
    }
    tcp, err := net.DialTCP("tcp", nil, laddr)
    if err != nil {
        log.Fatal(err)
        return nil, err
    }
    return &BConn{Conn:tcp},nil
}

func (this *BConn) Send(b []byte) (int, error){
    return this.Conn.Write(b)
}

func (this *BConn) Close(){
    this.Conn.Close()
}

func (this *BConn) Reciv(b []byte) (int, error) {
    return this.Conn.Read(b)
}

func (this *BConn) GetFd() (int64 , error){
    file,err := this.Conn.File()
    if err != nil {
        return 0, err
    }
    return int64(file.Fd()) , nil
}

func OnRead(conn *BConn, connIO IConnIO) {
    tmp_buf :=  make([]byte, CACHESIZE)
    buf := make([]byte, 0, 4096)
    buffer_size := 0
    for {
        count,err := conn.Reciv(tmp_buf)
        if count > 0 {
            buf = append(buf, tmp_buf[:count]...)
            buffer_size = buffer_size + count;
        }

        if err != nil || count < 4096 {
            if buffer_size > 0 {
                data := append([]byte(nil), buf...)
                //this.handleData(data)
                connIO.OnDataReaded(data,nil)
                buf = buf[:0]
                buffer_size = 0
            }else {

            }
        }

        if err != nil{
            connIO.Close()
            return
        }
    }
}
