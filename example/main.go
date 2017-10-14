package main

import (
    "github.com/xiaominfc/xiny/base"
    "github.com/xiaominfc/xiny/utils"
    "github.com/xiaominfc/xiny/conn"
    "github.com/xiaominfc/xiny/pb/IM_Login"
    "github.com/xiaominfc/xiny/pb/IM_BaseDefine"
    "net"
    // "fmt"
	"time"
    "log"
)

import proto "github.com/golang/protobuf/proto"

type Manager struct{

}

func (manager *Manager) NewConn(tcpConn *net.TCPConn){
    println("new connect");
    client := conn.NewClientConn(&base.BConn{Conn:tcpConn},manager)
    client.Run()
}

func (manager *Manager)HandlePdu(pdu *base.Pdu){
    println("HandlePdu")
}

func (manager *Manager)HandleData(data []byte,conn *conn.ClientConn) error{
    println("HandleData:",len(data))
    return nil
}


func main() {
    onlineStatus := IM_BaseDefine.UserStatType_USER_STATUS_ONLINE
    clientType := IM_BaseDefine.ClientType_CLIENT_TYPE_WINDOWS
    loginReq := &IM_Login.IMLoginReq{
        UserName:proto.String("xiaominfc"),
        Password:proto.String("test"),
        OnlineStatus:&onlineStatus,
        ClientType:&clientType}

    out, err := proto.Marshal(loginReq)
    if err != nil {
        log.Fatalln("Failed to encode:", err)
    } else {
        println("out size:",len(out))
    }

    pdu := base.NewPdu();
    pdu.SetServiceId(int16(IM_BaseDefine.ServiceID_SID_LOGIN))
    pdu.SetCommandId(int16(IM_BaseDefine.LoginCmdID_CID_LOGIN_REQ_USERLOGIN))
    pdu.SetSeqNum(100)
    pdu.Write(out)

    test := func (){
        println("hello")
    }

    utils.AddTask(2*time.Second, test)
    manager := &Manager{}
    connManager := conn.NewDefaultManager(manager)
    server := base.NewTcpServer("0.0.0.0",9090, manager)
    go server.Start()
    time.Sleep(2 * time.Second)
    conn.AddNewServConnFor("im.xiaominfc.com",8000,connManager)
//  println(time.Second)
	client,_ := base.Connect("127.0.0.1", 9090)
    connManager.GetServConn().Send(pdu.GetBufferData())

    for i:=0; i < 20 ; i++ {
        time.Sleep(2 * time.Second)
        client.Send([]byte{'h','e','r'})
        //manager.GetServConn().Send([]byte{'h','e','r'})
    }
	//client.Reciv(make([]byte, 100))
}
