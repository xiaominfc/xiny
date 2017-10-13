package main

import (
    "github.com/xiaominfc/xiny/base"
    "github.com/xiaominfc/xiny/utils"
    "github.com/xiaominfc/xiny/conn"
    "net"
    "fmt"
	"time"
)

type Manager struct{}

func (manager *Manager) NewConn(conn *net.TCPConn){
    println("new connect");
    tmp_buf :=  make([]byte, 256)
    buf := make([]byte, 0, 4096)
    client := &base.BConn{Conn:conn}
    for {
        count,err := client.Reciv(tmp_buf)
        println("reciv count:", count)
        buf = append(buf, tmp_buf[:count]...)
        if err != nil || count < 256 {
            println("read error")
            client.Send(buf)
            buf = buf[:0]
 //           break
        }

    }
    fmt.Println("read size:",len(buf))
    fmt.Println(string(buf[:len(buf)]))
}


func (manager *Manager) HandleData(b []byte) error {
   fmt.Println(len(b))
   return nil
}


func main() {

    test := func (){
        println("hello")
    }
    utils.AddTask(2*time.Second, test)
    manager := &Manager{}
    server := base.NewTcpServer("0.0.0.0",9090, manager)
    go server.Start()
    time.Sleep(2 * time.Second)
    servConn := conn.AddNewServFor("127.0.0.1",9090,manager)

    println(time.Second)
	client,_ := base.Connect("127.0.0.1", 9090)
	client.Send([]byte{'h','e','r'})

    for i:=0; i < 20 ; i++ {
        time.Sleep(2 * time.Second)
        client.Send([]byte{'h','e','r'})
        servConn.Send([]byte{'h','e','r'})
    }
	//client.Reciv(make([]byte, 100))
}
