package base

import (
    "net"
    "fmt"
	"github.com/xiaominfc/xiny/iface"
)


type Server struct {
    Host string
    Port int
	manager iface.Manager
}

func NewTcpServer(host string, port int,manager iface.Manager) *Server {
   s := &Server{host, port , manager}
   return s
}

func (ser *Server) handleConn(conn *net.TCPConn) {
   if ser.manager != nil {
       ser.manager.NewConn(conn)
   }
}


func (ser *Server) Start (){
    tcpAddr, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:%d", ser.Host, ser.Port))

    ln, err := net.ListenTCP("tcp",tcpAddr)
    if err != nil {

    }

    println("start server " ,ser.Host , ":" , ser.Port)
    for {
        con, err := ln.AcceptTCP()
        if err == nil {
            go ser.handleConn(con)
        }
    }
}
