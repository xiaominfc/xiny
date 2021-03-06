// Code generated by protoc-gen-go. DO NOT EDIT.
// source: IM.Group.proto

/*
Package IM_Group is a generated protocol buffer package.

It is generated from these files:
	IM.Group.proto

It has these top-level messages:
	IMNormalGroupListReq
	IMNormalGroupListRsp
	IMGroupInfoListReq
	IMGroupInfoListRsp
	IMGroupCreateReq
	IMGroupCreateRsp
	IMGroupChangeMemberReq
	IMGroupChangeMemberRsp
	IMGroupShieldReq
	IMGroupShieldRsp
	IMGroupChangeMemberNotify
*/
package IM_Group

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import IM_BaseDefine "github.com/xiaominfc/xiny/pb/IM_BaseDefine"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type IMNormalGroupListReq struct {
	// cmd id:			0x0401
	UserId           *uint32 `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id,omitempty"`
	AttachData       []byte  `protobuf:"bytes,20,opt,name=attach_data,json=attachData" json:"attach_data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *IMNormalGroupListReq) Reset()                    { *m = IMNormalGroupListReq{} }
func (m *IMNormalGroupListReq) String() string            { return proto.CompactTextString(m) }
func (*IMNormalGroupListReq) ProtoMessage()               {}
func (*IMNormalGroupListReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IMNormalGroupListReq) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMNormalGroupListReq) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMNormalGroupListRsp struct {
	// cmd id:			0x0402
	UserId           *uint32                           `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id,omitempty"`
	GroupVersionList []*IM_BaseDefine.GroupVersionInfo `protobuf:"bytes,2,rep,name=group_version_list,json=groupVersionList" json:"group_version_list,omitempty"`
	AttachData       []byte                            `protobuf:"bytes,20,opt,name=attach_data,json=attachData" json:"attach_data,omitempty"`
	XXX_unrecognized []byte                            `json:"-"`
}

func (m *IMNormalGroupListRsp) Reset()                    { *m = IMNormalGroupListRsp{} }
func (m *IMNormalGroupListRsp) String() string            { return proto.CompactTextString(m) }
func (*IMNormalGroupListRsp) ProtoMessage()               {}
func (*IMNormalGroupListRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IMNormalGroupListRsp) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMNormalGroupListRsp) GetGroupVersionList() []*IM_BaseDefine.GroupVersionInfo {
	if m != nil {
		return m.GroupVersionList
	}
	return nil
}

func (m *IMNormalGroupListRsp) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMGroupInfoListReq struct {
	// cmd id:			0x0403
	UserId           *uint32                           `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id,omitempty"`
	GroupVersionList []*IM_BaseDefine.GroupVersionInfo `protobuf:"bytes,2,rep,name=group_version_list,json=groupVersionList" json:"group_version_list,omitempty"`
	AttachData       []byte                            `protobuf:"bytes,20,opt,name=attach_data,json=attachData" json:"attach_data,omitempty"`
	XXX_unrecognized []byte                            `json:"-"`
}

func (m *IMGroupInfoListReq) Reset()                    { *m = IMGroupInfoListReq{} }
func (m *IMGroupInfoListReq) String() string            { return proto.CompactTextString(m) }
func (*IMGroupInfoListReq) ProtoMessage()               {}
func (*IMGroupInfoListReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *IMGroupInfoListReq) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMGroupInfoListReq) GetGroupVersionList() []*IM_BaseDefine.GroupVersionInfo {
	if m != nil {
		return m.GroupVersionList
	}
	return nil
}

func (m *IMGroupInfoListReq) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMGroupInfoListRsp struct {
	// cmd id:			0x0404
	UserId           *uint32                    `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id,omitempty"`
	GroupInfoList    []*IM_BaseDefine.GroupInfo `protobuf:"bytes,2,rep,name=group_info_list,json=groupInfoList" json:"group_info_list,omitempty"`
	AttachData       []byte                     `protobuf:"bytes,20,opt,name=attach_data,json=attachData" json:"attach_data,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *IMGroupInfoListRsp) Reset()                    { *m = IMGroupInfoListRsp{} }
func (m *IMGroupInfoListRsp) String() string            { return proto.CompactTextString(m) }
func (*IMGroupInfoListRsp) ProtoMessage()               {}
func (*IMGroupInfoListRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *IMGroupInfoListRsp) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMGroupInfoListRsp) GetGroupInfoList() []*IM_BaseDefine.GroupInfo {
	if m != nil {
		return m.GroupInfoList
	}
	return nil
}

func (m *IMGroupInfoListRsp) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMGroupCreateReq struct {
	// cmd id:			0x0405
	UserId           *uint32                  `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id,omitempty"`
	GroupType        *IM_BaseDefine.GroupType `protobuf:"varint,2,req,name=group_type,json=groupType,enum=IM.BaseDefine.GroupType,def=2" json:"group_type,omitempty"`
	GroupName        *string                  `protobuf:"bytes,3,req,name=group_name,json=groupName" json:"group_name,omitempty"`
	GroupAvatar      *string                  `protobuf:"bytes,4,req,name=group_avatar,json=groupAvatar" json:"group_avatar,omitempty"`
	MemberIdList     []uint32                 `protobuf:"varint,5,rep,name=member_id_list,json=memberIdList" json:"member_id_list,omitempty"`
	AttachData       []byte                   `protobuf:"bytes,20,opt,name=attach_data,json=attachData" json:"attach_data,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *IMGroupCreateReq) Reset()                    { *m = IMGroupCreateReq{} }
func (m *IMGroupCreateReq) String() string            { return proto.CompactTextString(m) }
func (*IMGroupCreateReq) ProtoMessage()               {}
func (*IMGroupCreateReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

const Default_IMGroupCreateReq_GroupType IM_BaseDefine.GroupType = IM_BaseDefine.GroupType_GROUP_TYPE_TMP

func (m *IMGroupCreateReq) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMGroupCreateReq) GetGroupType() IM_BaseDefine.GroupType {
	if m != nil && m.GroupType != nil {
		return *m.GroupType
	}
	return Default_IMGroupCreateReq_GroupType
}

func (m *IMGroupCreateReq) GetGroupName() string {
	if m != nil && m.GroupName != nil {
		return *m.GroupName
	}
	return ""
}

func (m *IMGroupCreateReq) GetGroupAvatar() string {
	if m != nil && m.GroupAvatar != nil {
		return *m.GroupAvatar
	}
	return ""
}

func (m *IMGroupCreateReq) GetMemberIdList() []uint32 {
	if m != nil {
		return m.MemberIdList
	}
	return nil
}

func (m *IMGroupCreateReq) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMGroupCreateRsp struct {
	// cmd id:			0x0406
	UserId           *uint32  `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id,omitempty"`
	ResultCode       *uint32  `protobuf:"varint,2,req,name=result_code,json=resultCode" json:"result_code,omitempty"`
	GroupId          *uint32  `protobuf:"varint,3,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
	GroupName        *string  `protobuf:"bytes,4,req,name=group_name,json=groupName" json:"group_name,omitempty"`
	UserIdList       []uint32 `protobuf:"varint,5,rep,name=user_id_list,json=userIdList" json:"user_id_list,omitempty"`
	AttachData       []byte   `protobuf:"bytes,20,opt,name=attach_data,json=attachData" json:"attach_data,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *IMGroupCreateRsp) Reset()                    { *m = IMGroupCreateRsp{} }
func (m *IMGroupCreateRsp) String() string            { return proto.CompactTextString(m) }
func (*IMGroupCreateRsp) ProtoMessage()               {}
func (*IMGroupCreateRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *IMGroupCreateRsp) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMGroupCreateRsp) GetResultCode() uint32 {
	if m != nil && m.ResultCode != nil {
		return *m.ResultCode
	}
	return 0
}

func (m *IMGroupCreateRsp) GetGroupId() uint32 {
	if m != nil && m.GroupId != nil {
		return *m.GroupId
	}
	return 0
}

func (m *IMGroupCreateRsp) GetGroupName() string {
	if m != nil && m.GroupName != nil {
		return *m.GroupName
	}
	return ""
}

func (m *IMGroupCreateRsp) GetUserIdList() []uint32 {
	if m != nil {
		return m.UserIdList
	}
	return nil
}

func (m *IMGroupCreateRsp) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMGroupChangeMemberReq struct {
	// cmd id:			0x0407
	UserId           *uint32                        `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id,omitempty"`
	ChangeType       *IM_BaseDefine.GroupModifyType `protobuf:"varint,2,req,name=change_type,json=changeType,enum=IM.BaseDefine.GroupModifyType" json:"change_type,omitempty"`
	GroupId          *uint32                        `protobuf:"varint,3,req,name=group_id,json=groupId" json:"group_id,omitempty"`
	MemberIdList     []uint32                       `protobuf:"varint,4,rep,name=member_id_list,json=memberIdList" json:"member_id_list,omitempty"`
	AttachData       []byte                         `protobuf:"bytes,20,opt,name=attach_data,json=attachData" json:"attach_data,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *IMGroupChangeMemberReq) Reset()                    { *m = IMGroupChangeMemberReq{} }
func (m *IMGroupChangeMemberReq) String() string            { return proto.CompactTextString(m) }
func (*IMGroupChangeMemberReq) ProtoMessage()               {}
func (*IMGroupChangeMemberReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *IMGroupChangeMemberReq) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMGroupChangeMemberReq) GetChangeType() IM_BaseDefine.GroupModifyType {
	if m != nil && m.ChangeType != nil {
		return *m.ChangeType
	}
	return IM_BaseDefine.GroupModifyType_GROUP_MODIFY_TYPE_ADD
}

func (m *IMGroupChangeMemberReq) GetGroupId() uint32 {
	if m != nil && m.GroupId != nil {
		return *m.GroupId
	}
	return 0
}

func (m *IMGroupChangeMemberReq) GetMemberIdList() []uint32 {
	if m != nil {
		return m.MemberIdList
	}
	return nil
}

func (m *IMGroupChangeMemberReq) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMGroupChangeMemberRsp struct {
	// cmd id:			0x0408
	UserId           *uint32                        `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id,omitempty"`
	ChangeType       *IM_BaseDefine.GroupModifyType `protobuf:"varint,2,req,name=change_type,json=changeType,enum=IM.BaseDefine.GroupModifyType" json:"change_type,omitempty"`
	ResultCode       *uint32                        `protobuf:"varint,3,req,name=result_code,json=resultCode" json:"result_code,omitempty"`
	GroupId          *uint32                        `protobuf:"varint,4,req,name=group_id,json=groupId" json:"group_id,omitempty"`
	CurUserIdList    []uint32                       `protobuf:"varint,5,rep,name=cur_user_id_list,json=curUserIdList" json:"cur_user_id_list,omitempty"`
	ChgUserIdList    []uint32                       `protobuf:"varint,6,rep,name=chg_user_id_list,json=chgUserIdList" json:"chg_user_id_list,omitempty"`
	AttachData       []byte                         `protobuf:"bytes,20,opt,name=attach_data,json=attachData" json:"attach_data,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *IMGroupChangeMemberRsp) Reset()                    { *m = IMGroupChangeMemberRsp{} }
func (m *IMGroupChangeMemberRsp) String() string            { return proto.CompactTextString(m) }
func (*IMGroupChangeMemberRsp) ProtoMessage()               {}
func (*IMGroupChangeMemberRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *IMGroupChangeMemberRsp) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMGroupChangeMemberRsp) GetChangeType() IM_BaseDefine.GroupModifyType {
	if m != nil && m.ChangeType != nil {
		return *m.ChangeType
	}
	return IM_BaseDefine.GroupModifyType_GROUP_MODIFY_TYPE_ADD
}

func (m *IMGroupChangeMemberRsp) GetResultCode() uint32 {
	if m != nil && m.ResultCode != nil {
		return *m.ResultCode
	}
	return 0
}

func (m *IMGroupChangeMemberRsp) GetGroupId() uint32 {
	if m != nil && m.GroupId != nil {
		return *m.GroupId
	}
	return 0
}

func (m *IMGroupChangeMemberRsp) GetCurUserIdList() []uint32 {
	if m != nil {
		return m.CurUserIdList
	}
	return nil
}

func (m *IMGroupChangeMemberRsp) GetChgUserIdList() []uint32 {
	if m != nil {
		return m.ChgUserIdList
	}
	return nil
}

func (m *IMGroupChangeMemberRsp) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMGroupShieldReq struct {
	// cmd id:			0x0409
	UserId           *uint32 `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id,omitempty"`
	GroupId          *uint32 `protobuf:"varint,2,req,name=group_id,json=groupId" json:"group_id,omitempty"`
	ShieldStatus     *uint32 `protobuf:"varint,3,req,name=shield_status,json=shieldStatus" json:"shield_status,omitempty"`
	AttachData       []byte  `protobuf:"bytes,20,opt,name=attach_data,json=attachData" json:"attach_data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *IMGroupShieldReq) Reset()                    { *m = IMGroupShieldReq{} }
func (m *IMGroupShieldReq) String() string            { return proto.CompactTextString(m) }
func (*IMGroupShieldReq) ProtoMessage()               {}
func (*IMGroupShieldReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *IMGroupShieldReq) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMGroupShieldReq) GetGroupId() uint32 {
	if m != nil && m.GroupId != nil {
		return *m.GroupId
	}
	return 0
}

func (m *IMGroupShieldReq) GetShieldStatus() uint32 {
	if m != nil && m.ShieldStatus != nil {
		return *m.ShieldStatus
	}
	return 0
}

func (m *IMGroupShieldReq) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMGroupShieldRsp struct {
	// cmd id:			0x040a
	UserId           *uint32 `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id,omitempty"`
	GroupId          *uint32 `protobuf:"varint,2,req,name=group_id,json=groupId" json:"group_id,omitempty"`
	ResultCode       *uint32 `protobuf:"varint,3,req,name=result_code,json=resultCode" json:"result_code,omitempty"`
	AttachData       []byte  `protobuf:"bytes,20,opt,name=attach_data,json=attachData" json:"attach_data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *IMGroupShieldRsp) Reset()                    { *m = IMGroupShieldRsp{} }
func (m *IMGroupShieldRsp) String() string            { return proto.CompactTextString(m) }
func (*IMGroupShieldRsp) ProtoMessage()               {}
func (*IMGroupShieldRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *IMGroupShieldRsp) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMGroupShieldRsp) GetGroupId() uint32 {
	if m != nil && m.GroupId != nil {
		return *m.GroupId
	}
	return 0
}

func (m *IMGroupShieldRsp) GetResultCode() uint32 {
	if m != nil && m.ResultCode != nil {
		return *m.ResultCode
	}
	return 0
}

func (m *IMGroupShieldRsp) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMGroupChangeMemberNotify struct {
	// cmd id: 			0x040b
	UserId           *uint32                        `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id,omitempty"`
	ChangeType       *IM_BaseDefine.GroupModifyType `protobuf:"varint,2,req,name=change_type,json=changeType,enum=IM.BaseDefine.GroupModifyType" json:"change_type,omitempty"`
	GroupId          *uint32                        `protobuf:"varint,3,req,name=group_id,json=groupId" json:"group_id,omitempty"`
	CurUserIdList    []uint32                       `protobuf:"varint,4,rep,name=cur_user_id_list,json=curUserIdList" json:"cur_user_id_list,omitempty"`
	ChgUserIdList    []uint32                       `protobuf:"varint,5,rep,name=chg_user_id_list,json=chgUserIdList" json:"chg_user_id_list,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *IMGroupChangeMemberNotify) Reset()                    { *m = IMGroupChangeMemberNotify{} }
func (m *IMGroupChangeMemberNotify) String() string            { return proto.CompactTextString(m) }
func (*IMGroupChangeMemberNotify) ProtoMessage()               {}
func (*IMGroupChangeMemberNotify) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *IMGroupChangeMemberNotify) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMGroupChangeMemberNotify) GetChangeType() IM_BaseDefine.GroupModifyType {
	if m != nil && m.ChangeType != nil {
		return *m.ChangeType
	}
	return IM_BaseDefine.GroupModifyType_GROUP_MODIFY_TYPE_ADD
}

func (m *IMGroupChangeMemberNotify) GetGroupId() uint32 {
	if m != nil && m.GroupId != nil {
		return *m.GroupId
	}
	return 0
}

func (m *IMGroupChangeMemberNotify) GetCurUserIdList() []uint32 {
	if m != nil {
		return m.CurUserIdList
	}
	return nil
}

func (m *IMGroupChangeMemberNotify) GetChgUserIdList() []uint32 {
	if m != nil {
		return m.ChgUserIdList
	}
	return nil
}

func init() {
	proto.RegisterType((*IMNormalGroupListReq)(nil), "IM.Group.IMNormalGroupListReq")
	proto.RegisterType((*IMNormalGroupListRsp)(nil), "IM.Group.IMNormalGroupListRsp")
	proto.RegisterType((*IMGroupInfoListReq)(nil), "IM.Group.IMGroupInfoListReq")
	proto.RegisterType((*IMGroupInfoListRsp)(nil), "IM.Group.IMGroupInfoListRsp")
	proto.RegisterType((*IMGroupCreateReq)(nil), "IM.Group.IMGroupCreateReq")
	proto.RegisterType((*IMGroupCreateRsp)(nil), "IM.Group.IMGroupCreateRsp")
	proto.RegisterType((*IMGroupChangeMemberReq)(nil), "IM.Group.IMGroupChangeMemberReq")
	proto.RegisterType((*IMGroupChangeMemberRsp)(nil), "IM.Group.IMGroupChangeMemberRsp")
	proto.RegisterType((*IMGroupShieldReq)(nil), "IM.Group.IMGroupShieldReq")
	proto.RegisterType((*IMGroupShieldRsp)(nil), "IM.Group.IMGroupShieldRsp")
	proto.RegisterType((*IMGroupChangeMemberNotify)(nil), "IM.Group.IMGroupChangeMemberNotify")
}

func init() { proto.RegisterFile("IM.Group.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 610 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x96, 0xcf, 0x8e, 0x93, 0x5e,
	0x14, 0xc7, 0x03, 0x74, 0xfe, 0x9d, 0x96, 0xfe, 0x1a, 0x7e, 0x13, 0x87, 0xd1, 0x68, 0x71, 0x34,
	0x91, 0x15, 0x0b, 0x97, 0x6e, 0xd4, 0x99, 0x31, 0x95, 0x44, 0x6a, 0xc3, 0xb4, 0x26, 0xae, 0xc8,
	0x1d, 0xb8, 0x50, 0x4c, 0xe1, 0x56, 0xb8, 0x4c, 0xd2, 0x37, 0x70, 0xe1, 0xc2, 0x17, 0x30, 0xc6,
	0xc7, 0x71, 0xef, 0x13, 0xf8, 0x1e, 0x26, 0x86, 0x7b, 0x69, 0x86, 0xb6, 0x08, 0x55, 0x13, 0x5d,
	0x72, 0xee, 0x99, 0x73, 0x3f, 0xdf, 0xfb, 0x39, 0x93, 0x14, 0xba, 0xa6, 0x65, 0x0c, 0x12, 0x92,
	0xcd, 0x8d, 0x79, 0x42, 0x28, 0x51, 0xf6, 0x97, 0xdf, 0x37, 0xff, 0x37, 0x2d, 0xe3, 0x14, 0xa5,
	0xf8, 0x1c, 0xfb, 0x61, 0x8c, 0xf9, 0xf1, 0xc9, 0x08, 0x0e, 0x4d, 0x6b, 0x48, 0x92, 0x08, 0xcd,
	0x58, 0xd7, 0x8b, 0x30, 0xa5, 0x36, 0x7e, 0xab, 0x1c, 0xc1, 0x5e, 0x96, 0xe2, 0xc4, 0x09, 0x3d,
	0x55, 0xd0, 0x44, 0x5d, 0xb6, 0x77, 0xf3, 0x4f, 0xd3, 0x53, 0xfa, 0xd0, 0x46, 0x94, 0x22, 0x77,
	0xea, 0x78, 0x88, 0x22, 0xf5, 0x50, 0x13, 0xf4, 0x8e, 0x0d, 0xbc, 0x74, 0x8e, 0x28, 0x3a, 0xf9,
	0x24, 0x54, 0x8d, 0x4c, 0xe7, 0x3f, 0x1f, 0x69, 0x81, 0x12, 0xe4, 0x8d, 0xce, 0x15, 0x4e, 0xd2,
	0x90, 0xc4, 0xce, 0x2c, 0x4c, 0xa9, 0x2a, 0x6a, 0x92, 0xde, 0x7e, 0xd8, 0x37, 0x56, 0xa9, 0xd9,
	0xc4, 0x57, 0xbc, 0xcf, 0x8c, 0x7d, 0x62, 0xf7, 0x82, 0x52, 0x25, 0xbf, 0xaa, 0x99, 0xf0, 0xa3,
	0x00, 0x8a, 0x69, 0xb1, 0x49, 0xf9, 0x88, 0xc6, 0xc8, 0x7f, 0x9b, 0xef, 0x43, 0x05, 0x5f, 0xdd,
	0xfb, 0x3d, 0x81, 0xff, 0x38, 0x5f, 0x18, 0xfb, 0xa4, 0x0c, 0xa7, 0x56, 0xc1, 0x31, 0x2a, 0x39,
	0x28, 0x4f, 0x6f, 0x46, 0xfa, 0x2e, 0x40, 0xaf, 0x40, 0x3a, 0x4b, 0x30, 0xa2, 0xb8, 0xf6, 0xc1,
	0x06, 0x00, 0x1c, 0x88, 0x2e, 0xe6, 0x58, 0x15, 0x35, 0x51, 0xef, 0x56, 0xb3, 0x8c, 0x17, 0x73,
	0xfc, 0xa8, 0x3b, 0xb0, 0x5f, 0x4e, 0x46, 0xce, 0xf8, 0xf5, 0xe8, 0x99, 0x33, 0xb6, 0x46, 0xf6,
	0x41, 0xb0, 0x3c, 0x52, 0x6e, 0x2f, 0x07, 0xc5, 0x28, 0xc2, 0xaa, 0xa4, 0x89, 0xfa, 0x41, 0x71,
	0x3c, 0x44, 0x11, 0x56, 0xee, 0x42, 0x87, 0x1f, 0xa3, 0x2b, 0x44, 0x51, 0xa2, 0xb6, 0x58, 0x43,
	0x9b, 0xd5, 0x9e, 0xb2, 0x92, 0x72, 0x1f, 0xba, 0x11, 0x8e, 0x2e, 0x19, 0x25, 0x7f, 0x9a, 0x1d,
	0x4d, 0xd2, 0x65, 0xbb, 0xc3, 0xab, 0xa6, 0xb7, 0x5d, 0xfe, 0x2f, 0x1b, 0xf9, 0xeb, 0x84, 0xf4,
	0xa1, 0x9d, 0xe0, 0x34, 0x9b, 0x51, 0xc7, 0x25, 0x1e, 0x7f, 0x00, 0xd9, 0x06, 0x5e, 0x3a, 0x23,
	0x1e, 0x56, 0x8e, 0x61, 0xbf, 0x30, 0xe6, 0xa9, 0x92, 0x26, 0xe8, 0xb2, 0xbd, 0xc7, 0x85, 0x78,
	0x6b, 0x91, 0x5b, 0xeb, 0x91, 0x35, 0xe8, 0x14, 0x77, 0x96, 0xd3, 0x00, 0xbf, 0x78, 0xbb, 0x2c,
	0x5f, 0x05, 0xb8, 0xb1, 0xcc, 0x32, 0x45, 0x71, 0x80, 0x2d, 0xf6, 0x14, 0xb5, 0x46, 0x1f, 0x43,
	0xdb, 0x65, 0xbd, 0x65, 0xa5, 0x77, 0xaa, 0x94, 0x5a, 0xc4, 0x0b, 0xfd, 0x45, 0x6e, 0xcf, 0x06,
	0xfe, 0x27, 0xcc, 0xe4, 0x6a, 0x62, 0xb1, 0x9c, 0x78, 0x53, 0x51, 0xeb, 0x77, 0x14, 0x7d, 0x16,
	0xab, 0x63, 0xd5, 0x89, 0xfa, 0xe3, 0x58, 0x6b, 0xa6, 0xa5, 0x5a, 0xd3, 0xad, 0xd5, 0xdc, 0x0f,
	0xa0, 0xe7, 0x66, 0x89, 0x53, 0xa1, 0x53, 0x76, 0xb3, 0x64, 0x72, 0x6d, 0x34, 0x6f, 0x9c, 0x06,
	0xab, 0x8d, 0xbb, 0x45, 0xe3, 0x34, 0x98, 0xfc, 0x82, 0xfa, 0xf7, 0xd7, 0x6b, 0x7c, 0x31, 0x0d,
	0xf1, 0xcc, 0xab, 0x95, 0x5e, 0x66, 0x17, 0x57, 0xd9, 0xef, 0x81, 0x9c, 0xb2, 0x01, 0x4e, 0x4a,
	0x11, 0xcd, 0xd2, 0x22, 0x79, 0x87, 0x17, 0x2f, 0x58, 0xad, 0x19, 0xe7, 0xdd, 0x06, 0x4e, 0x9d,
	0xac, 0x1a, 0x9c, 0x46, 0x0d, 0x8d, 0x28, 0xdf, 0x04, 0x38, 0xae, 0xd8, 0x9e, 0x21, 0xa1, 0xa1,
	0xbf, 0xf8, 0x37, 0xff, 0x17, 0x55, 0xfb, 0xd1, 0xda, 0x76, 0x3f, 0x76, 0x2a, 0xf6, 0xe3, 0xf4,
	0x16, 0x1c, 0xb9, 0x24, 0x32, 0x22, 0x12, 0x64, 0x6f, 0x42, 0x6c, 0x50, 0xca, 0x7f, 0x04, 0x5c,
	0x66, 0xfe, 0x73, 0xe9, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x5b, 0xab, 0x16, 0x39, 0x08,
	0x00, 0x00,
}
