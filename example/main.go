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
    "log"
)

import proto "github.com/golang/protobuf/proto"

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
   pdu := base.ReadPdu(b)
   if pdu != nil {
        println("serviceId:",pdu.GetServiceId(), "  commandId:",pdu.GetCommandId())
        if pdu.GetCommandId() == 260 {
            loginRes := &IM_Login.IMLoginRes{}
            proto.Unmarshal(pdu.GetBodyData(), loginRes)
            println(*loginRes.ResultCode)
        }
   }
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
    manager.ServConnList = utils.NewArray()
    server := base.NewTcpServer("0.0.0.0",9090, manager)
    go server.Start()
    time.Sleep(2 * time.Second)
    conn.AddNewServConnFor("im.xiaominfc.com",8000,manager)    
//  println(time.Second)
	client,_ := base.Connect("127.0.0.1", 9090)
    manager.GetServConn().Send(pdu.GetBufferData())

    for i:=0; i < 20 ; i++ {
        time.Sleep(2 * time.Second)
        client.Send([]byte{'h','e','r'})
        //manager.GetServConn().Send([]byte{'h','e','r'})
    }
	//client.Reciv(make([]byte, 100))
}
