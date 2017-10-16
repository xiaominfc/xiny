package main

import (
    "github.com/xiaominfc/xiny/base"
    //"github.com/xiaominfc/xiny/utils"
    "github.com/xiaominfc/xiny/conn"
    "github.com/xiaominfc/xiny/ylog"
    "github.com/xiaominfc/xiny/pb/IM_Login"
    "github.com/xiaominfc/xiny/pb/IM_BaseDefine"
    "github.com/xiaominfc/xiny/pb/IM_Group"
    "github.com/xiaominfc/xiny/pb/IM_Other"
    "net"
    "net/http"
    "bufio"
    "bytes"
    "time"
    "log"
    "fmt"
    "io/ioutil"
    "encoding/json"
    "github.com/antonholmquist/jason"
)

import proto "github.com/golang/protobuf/proto"

const HTTP_QUEYR_HEADER = "HTTP/1.1 200 OK\r\nCache-Control:no-cache\r\nConnection:close\r\nContent-Length:%d\r\nContent-Type:text/html;charset=utf-8\r\n\r\n%s"

type Manager struct{
    clientMap map[int64]*conn.ClientConn
    DbConnManager *conn.DefaultManager
    RConnManager *conn.DefaultManager
}

func (manager *Manager) NewConn(tcpConn *net.TCPConn){
    client := conn.NewClientConn(&base.BConn{Conn:tcpConn},manager)
    client.Run()
    fd := client.GetSocketFd()
    if fd > 0 {
        println("add client:",fd)
        manager.clientMap[fd] = client
    }
}

func (manager *Manager)OnTimeWork(servConn *conn.ServConn, servConnManager conn.ServConnManager) {
    ylog.ILog("OnTimeWork")
    beat := &IM_Other.IMHeartBeat{}
    out, err := proto.Marshal(beat)
    if err == nil {
        pdu := base.NewPdu();
        pdu.SetServiceId(int16(IM_BaseDefine.ServiceID_SID_OTHER))
        pdu.SetCommandId(int16(IM_BaseDefine.OtherCmdID_CID_OTHER_HEARTBEAT))
        //pdu.SetSeqNum(100)
        pdu.Write(out)
        servConn.Send(pdu.GetBufferData())

    }else {
        ylog.ILog(err.Error())
    }
}

func (manager *Manager)HandlePdu(pdu *base.Pdu, connManager conn.ServConnManager){
    //println("pdu for:",pdu.GetCommandId())
    switch int32(pdu.GetCommandId()) {
    case int32(IM_BaseDefine.OtherCmdID_CID_OTHER_HEARTBEAT):
        ylog.ILog("heartbeat pdu")
        break
    case int32(IM_BaseDefine.GroupCmdID_CID_GROUP_CHANGE_MEMBER_RESPONSE):
        manager.DoResponseForChangeMember(pdu)
        break
    case int32(IM_BaseDefine.GroupCmdID_CID_GROUP_CREATE_RESPONSE):
        manager.DoResponseForCreateGroup(pdu)
        break
    default:
        ylog.ILog("no anwser for pdu:",pdu.GetCommandId())
    }
}

func NewManager() *Manager{
    manager := &Manager{}
    manager.clientMap = make(map[int64]*conn.ClientConn)
    return manager
}

type CreateGroupReq struct {
    GroupName string `json:"group_name"`
    GroupType IM_BaseDefine.GroupType `json:"group_type"`
    GroupAvatar string `json:"group_avatar"`
    MemberIdList []uint32 `json:"user_id_list"`
}

func (manager *Manager)DoCreateGroup(postData []byte, client *conn.ClientConn){
    var reqData CreateGroupReq
    err := json.Unmarshal(postData,&reqData)
    if err != nil {
        ylog.ILog(err.Error())
        return
    }

    dbConn := manager.DbConnManager.GetServConn()
    req := &IM_Group.IMGroupCreateReq{UserId:proto.Uint32(0),GroupType:&reqData.GroupType,GroupName:proto.String(reqData.GroupName),GroupAvatar:proto.String(reqData.GroupAvatar),MemberIdList:reqData.MemberIdList}
    attach := base.NewAttachData(3,uint32(client.GetSocketFd()),0)
    req.AttachData = attach.GetBufferData()
    out, err := proto.Marshal(req)
    if err == nil {
        pdu := base.NewPdu();
        pdu.SetServiceId(int16(IM_BaseDefine.ServiceID_SID_GROUP))
        pdu.SetCommandId(int16(IM_BaseDefine.GroupCmdID_CID_GROUP_CREATE_REQUEST))
        //pdu.SetSeqNum(100)
        pdu.Write(out)
        dbConn.Send(pdu.GetBufferData())

    }else {
        ylog.ILog(err.Error())
    }
}


func (manager *Manager)DoResponseForCreateGroup(pdu *base.Pdu){
    ylog.ILog("response for DoResponseForCreateGroup")
    var res IM_Group.IMGroupCreateRsp
    err := proto.Unmarshal(pdu.GetBodyData(), &res)
    if err != nil {
        ylog.ILog(err.Error())
        return
    }

    attach := base.NewAttachDataForData(res.AttachData)
    resultCode := *res.ResultCode
    result := fmt.Sprintf("{\"error_code\":%d,\"error_msg\":\"ok\"}",resultCode)
    outData := fmt.Sprintf(HTTP_QUEYR_HEADER, len(result) , result);
    //println(outData)
    clientConn := manager.clientMap[int64(attach.GetHandle()) - 1]
    data := []byte(outData);
    clientConn.Send(data)
    clientConn.Close()
}

type ChangeMemberReq struct {
    GroupId uint32 `json:"group_id"`
    ChangeType IM_BaseDefine.GroupModifyType `json:"modify_type"`
    MemberIdList []uint32 `json:"user_id_list"`
}

func (manager *Manager)DoChangeMembers(postData []byte, client *conn.ClientConn) {
    ylog.ILog("DoChangeMembers")
    var reqData ChangeMemberReq
    err := json.Unmarshal(postData,&reqData)
    if err != nil {
        ylog.ILog(err.Error())
        return
    }
    dbConn := manager.DbConnManager.GetServConn()
    req := &IM_Group.IMGroupChangeMemberReq{UserId:proto.Uint32(0),ChangeType:&reqData.ChangeType,GroupId:proto.Uint32(reqData.GroupId),MemberIdList:reqData.MemberIdList}
    attach := base.NewAttachData(3,uint32(client.GetSocketFd()),0)
    req.AttachData = attach.GetBufferData()
    //println("user_id_list:",reqData.MemberIdList[0], "  group_id:",reqData.GroupId);
    out, err := proto.Marshal(req)
    if err == nil {
        pdu := base.NewPdu();
        pdu.SetServiceId(int16(IM_BaseDefine.ServiceID_SID_GROUP))
        pdu.SetCommandId(int16(IM_BaseDefine.GroupCmdID_CID_GROUP_CHANGE_MEMBER_REQUEST))
        //pdu.SetSeqNum(100)
        pdu.Write(out)
        dbConn.Send(pdu.GetBufferData())
    }else {
        ylog.ILog(err.Error())
    }
}


func (manager *Manager)DoResponseForChangeMember(pdu *base.Pdu){
    ylog.ILog("DoResponseForChangeMember")
    var res IM_Group.IMGroupChangeMemberRsp
    err := proto.Unmarshal(pdu.GetBodyData(), &res)
    if err != nil {
        ylog.ILog(err.Error())
        return
    }
    attach := base.NewAttachDataForData(res.AttachData)
    resultCode := *res.ResultCode
    result := fmt.Sprintf("{\"error_code\":%d,\"error_msg\":\"ok\"}",resultCode)
    outData := fmt.Sprintf(HTTP_QUEYR_HEADER, len(result) , result);
    println(outData)
    clientConn := manager.clientMap[int64(attach.GetHandle()) - 1]
    data := []byte(outData);
    clientConn.Send(data)
    clientConn.Close()

    if resultCode == 0 {
        req := &IM_Group.IMGroupChangeMemberNotify{UserId:res.UserId, ChangeType:res.ChangeType, GroupId:res.GroupId,CurUserIdList:res.CurUserIdList,ChgUserIdList:res.ChgUserIdList}
        out, err := proto.Marshal(req)
        rserverConn := manager.RConnManager.GetServConn()
        if err == nil && rserverConn != nil {
            rpdu := base.NewPdu();
            rpdu.SetServiceId(int16(IM_BaseDefine.ServiceID_SID_GROUP))
            rpdu.SetCommandId(int16(IM_BaseDefine.GroupCmdID_CID_GROUP_CHANGE_MEMBER_NOTIFY))
            rpdu.Write(out)
            rserverConn.Send(rpdu.GetBufferData())
        }else {
            ylog.ILog(err.Error())
        }


    }
}


func (manager *Manager)DispatchQuery(url string,postData []byte,client *conn.ClientConn) {
    _, err := jason.NewObjectFromBytes(postData)
    if err == nil {
    switch url {
        case "/query/CreateGroup":
            manager.DoCreateGroup(postData, client)
            break
        case "/query/ChangeMembers":
            manager.DoChangeMembers(postData, client)
            break
        default:
            println("url:",url)
        }
    }
}

func (manager *Manager)OnClose(client *conn.ClientConn) {
   fd := client.GetSocketFd()
   delete(manager.clientMap, fd)
}

func (manager *Manager)HandleData(data []byte,client *conn.ClientConn) error{
    // println("HandleData:",len(data))
    r := bytes.NewReader(data)
    io := bufio.NewReader(r)
    request ,err:= http.ReadRequest(io)
    if err != nil {

    } else {
        bodyData,_ := ioutil.ReadAll(request.Body)
        url := request.URL
        manager.DispatchQuery(url.Path,bodyData,client)
    }
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

    manager := NewManager()
    connManager := conn.NewDefaultManager(manager)
    rConnManager := conn.NewDefaultManager(manager)
    manager.DbConnManager = connManager
    manager.RConnManager = rConnManager
    server := base.NewTcpServer("0.0.0.0",9090, manager)
    go server.Start()
    time.Sleep(2 * time.Second)
    conn.AddNewServConnFor("im.xiaominfc.com",10600,connManager)
    conn.AddNewServConnFor("im.xiaominfc.com",8200,rConnManager)
//  println(time.Second)
    //client,_ := base.Connect("127.0.0.1", 9090)
    //connManager.GetServConn().Send(pdu.GetBufferData())

    for i:=0; i < 200 ; i++ {
        time.Sleep(2 * time.Second)
         //client.Send([]byte{'h','e','r'})
        //manager.GetServConn().Send([]byte{'h','e','r'})
    }
    //client.Reciv(make([]byte, 100))
}
