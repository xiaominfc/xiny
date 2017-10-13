package iface

import(
  "net"
)

type Manager interface{
	NewConn(conn *net.TCPConn)
}

type ServConn interface {
	Send(data []byte)
	OnRead()
	HandlePdu();
}
