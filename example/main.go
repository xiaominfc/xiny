package main

import (
    "github.com/xiaominfc/xiny/base"
    "github.com/xiaominfc/xiny/utils"
    "github.com/xiaominfc/xiny/conn"
    "github.com/xiaominfc/xiny/pb/IM_Login"
    "github.com/xiaominfc/xiny/pb/IM_BaseDefine"
    "net"
    "fmt"
	"time"
    "math/rand"
)

type Manager struct{
    ServConnList utils.Array
}

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

func (manager *Manager) OnServConnAdd(servconn *conn.ServConn) {
    println("add new servConn")
    manager.ServConnList.Add(servconn)
}

func (manager *Manager) GetServConn() *conn.ServConn {
    size := manager.ServConnList.Size()
    index := rand.Intn(size)
    server,err := manager.ServConnList.Get(index)
    if err!=nil {

    }
    return server.(*conn.ServConn)
}


func main() {

    loginReq := &IM_Login.IMLoginReq{
        UserName:utils.NewString("xiaominfc"),
        Password:utils.NewString("test"),
        OnlineStatus:&IM_BaseDefine.UserStatType.UserStatType_USER_STATUS_ONLINE}

    println(loginReq)

    test := func (){
        println("hello")
    }
    utils.AddTask(2*time.Second, test)
    manager := &Manager{}
    manager.ServConnList = utils.NewArray()
    server := base.NewTcpServer("0.0.0.0",9090, manager)
    go server.Start()
    time.Sleep(2 * time.Second)
    conn.AddNewServFor("127.0.0.1",9090,manager)

    println(time.Second)
	client,_ := base.Connect("127.0.0.1", 9090)
	client.Send([]byte{'h','e','r'})

    for i:=0; i < 20 ; i++ {
        time.Sleep(2 * time.Second)
        client.Send([]byte{'h','e','r'})
        manager.GetServConn().Send([]byte{'h','e','r'})
    }
	//client.Reciv(make([]byte, 100))
}
