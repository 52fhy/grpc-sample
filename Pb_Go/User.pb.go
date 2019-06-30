// Code generated by protoc-gen-go. DO NOT EDIT.
// source: User.proto

package Sample_Model

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type User struct {
	Id                   int64             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Avatar               string            `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Address              string            `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Mobile               string            `protobuf:"bytes,5,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Ext                  map[string]string `protobuf:"bytes,6,rep,name=ext,proto3" json:"ext,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_979821478719c248, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *User) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *User) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *User) GetExt() map[string]string {
	if m != nil {
		return m.Ext
	}
	return nil
}

type UserList struct {
	List                 []*User  `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Page                 int32    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit                int32    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserList) Reset()         { *m = UserList{} }
func (m *UserList) String() string { return proto.CompactTextString(m) }
func (*UserList) ProtoMessage()    {}
func (*UserList) Descriptor() ([]byte, []int) {
	return fileDescriptor_979821478719c248, []int{1}
}

func (m *UserList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserList.Unmarshal(m, b)
}
func (m *UserList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserList.Marshal(b, m, deterministic)
}
func (m *UserList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserList.Merge(m, src)
}
func (m *UserList) XXX_Size() int {
	return xxx_messageInfo_UserList.Size(m)
}
func (m *UserList) XXX_DiscardUnknown() {
	xxx_messageInfo_UserList.DiscardUnknown(m)
}

var xxx_messageInfo_UserList proto.InternalMessageInfo

func (m *UserList) GetList() []*User {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *UserList) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *UserList) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "Sample.Model.User")
	proto.RegisterMapType((map[string]string)(nil), "Sample.Model.User.ExtEntry")
	proto.RegisterType((*UserList)(nil), "Sample.Model.UserList")
}

func init() { proto.RegisterFile("User.proto", fileDescriptor_979821478719c248) }

var fileDescriptor_979821478719c248 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x31, 0x4b, 0x04, 0x31,
	0x10, 0x85, 0xc9, 0x66, 0x77, 0xbd, 0x1b, 0x45, 0x64, 0x10, 0x09, 0xda, 0x2c, 0x57, 0xc8, 0x36,
	0x6e, 0xa1, 0x20, 0x62, 0x7f, 0x9d, 0x36, 0x11, 0x3b, 0x9b, 0x1c, 0x09, 0x12, 0xcc, 0x5e, 0x96,
	0x24, 0x1e, 0x77, 0xbf, 0xd5, 0x3f, 0x23, 0x99, 0xb8, 0x20, 0x5c, 0x37, 0xdf, 0xbc, 0xcc, 0xbc,
	0x37, 0x01, 0x78, 0x8f, 0x26, 0x0c, 0x53, 0xf0, 0xc9, 0xe3, 0xd9, 0x9b, 0x1a, 0x27, 0x67, 0x86,
	0x57, 0xaf, 0x8d, 0x5b, 0xfd, 0x30, 0xa8, 0xb3, 0x88, 0xe7, 0x50, 0x59, 0x2d, 0x58, 0xc7, 0x7a,
	0x2e, 0x2b, 0xab, 0x11, 0xa1, 0xde, 0xaa, 0xd1, 0x88, 0xaa, 0x63, 0xfd, 0x52, 0x52, 0x8d, 0x57,
	0xd0, 0xaa, 0x9d, 0x4a, 0x2a, 0x08, 0x4e, 0xdd, 0x3f, 0x42, 0x01, 0x27, 0x4a, 0xeb, 0x60, 0x62,
	0x14, 0x35, 0x09, 0x33, 0xe6, 0x89, 0xd1, 0x6f, 0xac, 0x33, 0xa2, 0x29, 0x13, 0x85, 0xf0, 0x0e,
	0xb8, 0xd9, 0x27, 0xd1, 0x76, 0xbc, 0x3f, 0xbd, 0xbf, 0x19, 0xfe, 0x47, 0x1a, 0x28, 0xeb, 0x7a,
	0x9f, 0xd6, 0xdb, 0x14, 0x0e, 0x32, 0xbf, 0xbb, 0x7e, 0x84, 0xc5, 0xdc, 0xc0, 0x0b, 0xe0, 0x5f,
	0xe6, 0x40, 0x49, 0x97, 0x32, 0x97, 0x78, 0x09, 0xcd, 0x4e, 0xb9, 0xef, 0x39, 0x6b, 0x81, 0xe7,
	0xea, 0x89, 0xad, 0x3e, 0x60, 0x91, 0xb7, 0xbd, 0xd8, 0x98, 0xf0, 0x16, 0x6a, 0x67, 0x63, 0x12,
	0x8c, 0x3c, 0xf1, 0xd8, 0x53, 0x92, 0x9e, 0x0f, 0x9f, 0xd4, 0x67, 0x59, 0xd6, 0x48, 0xaa, 0xb3,
	0x83, 0xb3, 0xa3, 0x4d, 0x74, 0x77, 0x23, 0x0b, 0x6c, 0x5a, 0xfa, 0xd0, 0x87, 0xdf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x22, 0x56, 0xc1, 0x62, 0x5e, 0x01, 0x00, 0x00,
}
