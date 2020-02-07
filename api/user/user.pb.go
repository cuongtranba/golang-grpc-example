// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Model struct {
	Id                   string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string           `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string           `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Address              []*Model_Address `protobuf:"bytes,5,rep,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Model) Reset()         { *m = Model{} }
func (m *Model) String() string { return proto.CompactTextString(m) }
func (*Model) ProtoMessage()    {}
func (*Model) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *Model) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Model.Unmarshal(m, b)
}
func (m *Model) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Model.Marshal(b, m, deterministic)
}
func (m *Model) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Model.Merge(m, src)
}
func (m *Model) XXX_Size() int {
	return xxx_messageInfo_Model.Size(m)
}
func (m *Model) XXX_DiscardUnknown() {
	xxx_messageInfo_Model.DiscardUnknown(m)
}

var xxx_messageInfo_Model proto.InternalMessageInfo

func (m *Model) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Model) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Model) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Model) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Model) GetAddress() []*Model_Address {
	if m != nil {
		return m.Address
	}
	return nil
}

type Model_Address struct {
	Street               string   `protobuf:"bytes,1,opt,name=street,proto3" json:"street,omitempty"`
	City                 string   `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	State                string   `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Country              string   `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Model_Address) Reset()         { *m = Model_Address{} }
func (m *Model_Address) String() string { return proto.CompactTextString(m) }
func (*Model_Address) ProtoMessage()    {}
func (*Model_Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0, 0}
}

func (m *Model_Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Model_Address.Unmarshal(m, b)
}
func (m *Model_Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Model_Address.Marshal(b, m, deterministic)
}
func (m *Model_Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Model_Address.Merge(m, src)
}
func (m *Model_Address) XXX_Size() int {
	return xxx_messageInfo_Model_Address.Size(m)
}
func (m *Model_Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Model_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Model_Address proto.InternalMessageInfo

func (m *Model_Address) GetStreet() string {
	if m != nil {
		return m.Street
	}
	return ""
}

func (m *Model_Address) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Model_Address) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Model_Address) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

type GetRequest struct {
	Keyword              string   `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

type FindRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindRequest) Reset()         { *m = FindRequest{} }
func (m *FindRequest) String() string { return proto.CompactTextString(m) }
func (*FindRequest) ProtoMessage()    {}
func (*FindRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *FindRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindRequest.Unmarshal(m, b)
}
func (m *FindRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindRequest.Marshal(b, m, deterministic)
}
func (m *FindRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindRequest.Merge(m, src)
}
func (m *FindRequest) XXX_Size() int {
	return xxx_messageInfo_FindRequest.Size(m)
}
func (m *FindRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindRequest proto.InternalMessageInfo

func (m *FindRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CreateGenRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateGenRequest) Reset()         { *m = CreateGenRequest{} }
func (m *CreateGenRequest) String() string { return proto.CompactTextString(m) }
func (*CreateGenRequest) ProtoMessage()    {}
func (*CreateGenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *CreateGenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateGenRequest.Unmarshal(m, b)
}
func (m *CreateGenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateGenRequest.Marshal(b, m, deterministic)
}
func (m *CreateGenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateGenRequest.Merge(m, src)
}
func (m *CreateGenRequest) XXX_Size() int {
	return xxx_messageInfo_CreateGenRequest.Size(m)
}
func (m *CreateGenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateGenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateGenRequest proto.InternalMessageInfo

type CreateRequest struct {
	User                 *Model   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetUser() *Model {
	if m != nil {
		return m.User
	}
	return nil
}

type GenResponse struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenResponse) Reset()         { *m = GenResponse{} }
func (m *GenResponse) String() string { return proto.CompactTextString(m) }
func (*GenResponse) ProtoMessage()    {}
func (*GenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *GenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenResponse.Unmarshal(m, b)
}
func (m *GenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenResponse.Marshal(b, m, deterministic)
}
func (m *GenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenResponse.Merge(m, src)
}
func (m *GenResponse) XXX_Size() int {
	return xxx_messageInfo_GenResponse.Size(m)
}
func (m *GenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenResponse proto.InternalMessageInfo

func (m *GenResponse) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type GetResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Users                []*Model `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *GetResponse) GetUsers() []*Model {
	if m != nil {
		return m.Users
	}
	return nil
}

type FindResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	User                 *Model   `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindResponse) Reset()         { *m = FindResponse{} }
func (m *FindResponse) String() string { return proto.CompactTextString(m) }
func (*FindResponse) ProtoMessage()    {}
func (*FindResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7}
}

func (m *FindResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindResponse.Unmarshal(m, b)
}
func (m *FindResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindResponse.Marshal(b, m, deterministic)
}
func (m *FindResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindResponse.Merge(m, src)
}
func (m *FindResponse) XXX_Size() int {
	return xxx_messageInfo_FindResponse.Size(m)
}
func (m *FindResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindResponse proto.InternalMessageInfo

func (m *FindResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *FindResponse) GetUser() *Model {
	if m != nil {
		return m.User
	}
	return nil
}

type CreateResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{8}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *CreateResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Model)(nil), "user.Model")
	proto.RegisterType((*Model_Address)(nil), "user.Model.Address")
	proto.RegisterType((*GetRequest)(nil), "user.GetRequest")
	proto.RegisterType((*FindRequest)(nil), "user.FindRequest")
	proto.RegisterType((*CreateGenRequest)(nil), "user.CreateGenRequest")
	proto.RegisterType((*CreateRequest)(nil), "user.CreateRequest")
	proto.RegisterType((*GenResponse)(nil), "user.GenResponse")
	proto.RegisterType((*GetResponse)(nil), "user.GetResponse")
	proto.RegisterType((*FindResponse)(nil), "user.FindResponse")
	proto.RegisterType((*CreateResponse)(nil), "user.CreateResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xcd, 0x6a, 0xdb, 0x40,
	0x10, 0x46, 0xb2, 0x64, 0xb9, 0xa3, 0xd6, 0xb8, 0x5b, 0x63, 0x16, 0x41, 0xa9, 0x2b, 0x68, 0xf1,
	0xa5, 0xae, 0xeb, 0x52, 0x7a, 0x6e, 0x03, 0x51, 0x72, 0xc8, 0x45, 0x90, 0x07, 0x90, 0xad, 0x21,
	0x11, 0xb1, 0x56, 0xce, 0xee, 0x8a, 0xe0, 0x57, 0xcd, 0x3b, 0xe4, 0x1d, 0xc2, 0xfe, 0x45, 0x72,
	0x20, 0xf8, 0x36, 0xdf, 0x37, 0xf3, 0xed, 0x7c, 0x33, 0x23, 0x01, 0xb4, 0x02, 0xf9, 0x72, 0xcf,
	0x1b, 0xd9, 0x90, 0x40, 0xc5, 0xe9, 0x93, 0x07, 0xe1, 0x55, 0x53, 0xe2, 0x8e, 0x8c, 0xc1, 0xaf,
	0x4a, 0xea, 0xcd, 0xbd, 0xc5, 0xbb, 0xdc, 0xaf, 0x4a, 0x42, 0x20, 0x60, 0x45, 0x8d, 0xd4, 0xd7,
	0x8c, 0x8e, 0xc9, 0x14, 0x42, 0xac, 0x8b, 0x6a, 0x47, 0x07, 0x9a, 0x34, 0x40, 0xb1, 0xfb, 0xdb,
	0x86, 0x21, 0x0d, 0x0c, 0xab, 0x01, 0xf9, 0x01, 0x51, 0x51, 0x96, 0x1c, 0x85, 0xa0, 0xe1, 0x7c,
	0xb0, 0x88, 0xd7, 0x9f, 0x96, 0xba, 0xbb, 0xee, 0xb6, 0xfc, 0x67, 0x52, 0xb9, 0xab, 0x49, 0x10,
	0x22, 0xcb, 0x91, 0x19, 0x0c, 0x85, 0xe4, 0x88, 0xd2, 0xba, 0xb1, 0x48, 0x39, 0xda, 0x56, 0xf2,
	0xe0, 0x1c, 0xa9, 0x58, 0xf5, 0x16, 0xb2, 0x90, 0xe8, 0x1c, 0x69, 0x40, 0x28, 0x44, 0xdb, 0xa6,
	0x65, 0x92, 0x1f, 0xac, 0x27, 0x07, 0xd3, 0xef, 0x00, 0x19, 0xca, 0x1c, 0xef, 0x5b, 0x14, 0x52,
	0xd5, 0xdd, 0xe1, 0xe1, 0xa1, 0xe1, 0x6e, 0x70, 0x07, 0xd3, 0xcf, 0x10, 0x9f, 0x57, 0xac, 0x74,
	0x85, 0xaf, 0x96, 0x93, 0x12, 0x98, 0x9c, 0x71, 0x2c, 0x24, 0x66, 0xc8, 0x6c, 0x4d, 0xba, 0x82,
	0x0f, 0x86, 0x73, 0xa2, 0x2f, 0xa0, 0x77, 0xac, 0x65, 0xf1, 0x3a, 0xee, 0x8d, 0x9f, 0x9b, 0xe5,
	0x7f, 0x83, 0x58, 0xeb, 0xc5, 0xbe, 0x61, 0x02, 0xd5, 0xdc, 0xac, 0xad, 0x37, 0x56, 0x11, 0xe6,
	0x16, 0xa5, 0x17, 0xaa, 0x4c, 0xf6, 0xcb, 0xd4, 0x94, 0xad, 0xe8, 0xd6, 0xa3, 0x10, 0xf9, 0x0a,
	0xa1, 0x7a, 0x55, 0x50, 0x5f, 0xaf, 0xfb, 0xa8, 0x9f, 0xc9, 0xa4, 0x19, 0xbc, 0x37, 0x53, 0x9d,
	0x78, 0xca, 0x39, 0xf7, 0xdf, 0x72, 0xfe, 0x1f, 0xc6, 0x6e, 0xd6, 0x13, 0x4f, 0x51, 0x88, 0x6a,
	0x14, 0xa2, 0xb8, 0x71, 0x5f, 0x92, 0x83, 0xeb, 0x47, 0x0f, 0x82, 0x6b, 0x81, 0x9c, 0xfc, 0x84,
	0x51, 0x86, 0x52, 0x85, 0x82, 0x4c, 0x4c, 0xaf, 0xee, 0x46, 0xc9, 0xc7, 0x1e, 0x63, 0x7b, 0xfd,
	0x82, 0x91, 0x1a, 0x43, 0x8b, 0x6d, 0xba, 0x77, 0xac, 0x84, 0xf4, 0x29, 0x2b, 0xf9, 0x0b, 0x60,
	0x0c, 0x6b, 0x91, 0xfd, 0x14, 0x8f, 0xce, 0x95, 0x4c, 0x8f, 0x49, 0x2b, 0xfc, 0x03, 0xc3, 0x0c,
	0xd9, 0x25, 0x93, 0x64, 0xd6, 0xcf, 0x77, 0x77, 0xef, 0x0c, 0xbe, 0x5c, 0x72, 0xe5, 0x6d, 0x86,
	0xfa, 0x27, 0xfb, 0xfd, 0x1c, 0x00, 0x00, 0xff, 0xff, 0xfb, 0xcc, 0xa0, 0xb0, 0x72, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	GetUsers(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	FindUser(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error)
	CreateUser(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	GenInt(ctx context.Context, in *CreateGenRequest, opts ...grpc.CallOption) (User_GenIntClient, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetUsers(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/user.User/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) FindUser(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error) {
	out := new(FindResponse)
	err := c.cc.Invoke(ctx, "/user.User/FindUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CreateUser(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/user.User/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GenInt(ctx context.Context, in *CreateGenRequest, opts ...grpc.CallOption) (User_GenIntClient, error) {
	stream, err := c.cc.NewStream(ctx, &_User_serviceDesc.Streams[0], "/user.User/GenInt", opts...)
	if err != nil {
		return nil, err
	}
	x := &userGenIntClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type User_GenIntClient interface {
	Recv() (*GenResponse, error)
	grpc.ClientStream
}

type userGenIntClient struct {
	grpc.ClientStream
}

func (x *userGenIntClient) Recv() (*GenResponse, error) {
	m := new(GenResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	GetUsers(context.Context, *GetRequest) (*GetResponse, error)
	FindUser(context.Context, *FindRequest) (*FindResponse, error)
	CreateUser(context.Context, *CreateRequest) (*CreateResponse, error)
	GenInt(*CreateGenRequest, User_GenIntServer) error
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) GetUsers(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (*UnimplementedUserServer) FindUser(ctx context.Context, req *FindRequest) (*FindResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUser not implemented")
}
func (*UnimplementedUserServer) CreateUser(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUserServer) GenInt(req *CreateGenRequest, srv User_GenIntServer) error {
	return status.Errorf(codes.Unimplemented, "method GenInt not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUsers(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_FindUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).FindUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/FindUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).FindUser(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateUser(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GenInt_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CreateGenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServer).GenInt(m, &userGenIntServer{stream})
}

type User_GenIntServer interface {
	Send(*GenResponse) error
	grpc.ServerStream
}

type userGenIntServer struct {
	grpc.ServerStream
}

func (x *userGenIntServer) Send(m *GenResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUsers",
			Handler:    _User_GetUsers_Handler,
		},
		{
			MethodName: "FindUser",
			Handler:    _User_FindUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _User_CreateUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GenInt",
			Handler:       _User_GenInt_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "user.proto",
}
