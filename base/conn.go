package base

import(
    "net"
    "fmt"
    "log"
)

type BConn struct{
    Conn *net.TCPConn
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

func (this *BConn) Reciv(b []byte) (int, error) {
    return this.Conn.Read(b)
}

