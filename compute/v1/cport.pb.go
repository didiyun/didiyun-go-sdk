// Code generated by protoc-gen-go. DO NOT EDIT.
// source: compute/v1/cport.proto

// dc2 Service
//
// dc2 Service API consists of a single service which returns
// a message.

package compute

import (
	context "context"
	fmt "fmt"
	v1 "github.com/didiyun/didiyun-go-sdk/base/v1"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetCportByUuidRequest struct {
	Header               *v1.Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	CportUuid            string     `protobuf:"bytes,2,opt,name=cportUuid,proto3" json:"cportUuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetCportByUuidRequest) Reset()         { *m = GetCportByUuidRequest{} }
func (m *GetCportByUuidRequest) String() string { return proto.CompactTextString(m) }
func (*GetCportByUuidRequest) ProtoMessage()    {}
func (*GetCportByUuidRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c309c866379f7da8, []int{0}
}

func (m *GetCportByUuidRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCportByUuidRequest.Unmarshal(m, b)
}
func (m *GetCportByUuidRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCportByUuidRequest.Marshal(b, m, deterministic)
}
func (m *GetCportByUuidRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCportByUuidRequest.Merge(m, src)
}
func (m *GetCportByUuidRequest) XXX_Size() int {
	return xxx_messageInfo_GetCportByUuidRequest.Size(m)
}
func (m *GetCportByUuidRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCportByUuidRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCportByUuidRequest proto.InternalMessageInfo

func (m *GetCportByUuidRequest) GetHeader() *v1.Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *GetCportByUuidRequest) GetCportUuid() string {
	if m != nil {
		return m.CportUuid
	}
	return ""
}

type GetCportByUuidResponse struct {
	Error                *v1.Error    `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Data                 []*CportInfo `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetCportByUuidResponse) Reset()         { *m = GetCportByUuidResponse{} }
func (m *GetCportByUuidResponse) String() string { return proto.CompactTextString(m) }
func (*GetCportByUuidResponse) ProtoMessage()    {}
func (*GetCportByUuidResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c309c866379f7da8, []int{1}
}

func (m *GetCportByUuidResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCportByUuidResponse.Unmarshal(m, b)
}
func (m *GetCportByUuidResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCportByUuidResponse.Marshal(b, m, deterministic)
}
func (m *GetCportByUuidResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCportByUuidResponse.Merge(m, src)
}
func (m *GetCportByUuidResponse) XXX_Size() int {
	return xxx_messageInfo_GetCportByUuidResponse.Size(m)
}
func (m *GetCportByUuidResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCportByUuidResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCportByUuidResponse proto.InternalMessageInfo

func (m *GetCportByUuidResponse) GetError() *v1.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *GetCportByUuidResponse) GetData() []*CportInfo {
	if m != nil {
		return m.Data
	}
	return nil
}

type ListCportData struct {
	Count                int32        `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Items                []*CportInfo `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListCportData) Reset()         { *m = ListCportData{} }
func (m *ListCportData) String() string { return proto.CompactTextString(m) }
func (*ListCportData) ProtoMessage()    {}
func (*ListCportData) Descriptor() ([]byte, []int) {
	return fileDescriptor_c309c866379f7da8, []int{2}
}

func (m *ListCportData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCportData.Unmarshal(m, b)
}
func (m *ListCportData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCportData.Marshal(b, m, deterministic)
}
func (m *ListCportData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCportData.Merge(m, src)
}
func (m *ListCportData) XXX_Size() int {
	return xxx_messageInfo_ListCportData.Size(m)
}
func (m *ListCportData) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCportData.DiscardUnknown(m)
}

var xxx_messageInfo_ListCportData proto.InternalMessageInfo

func (m *ListCportData) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ListCportData) GetItems() []*CportInfo {
	if m != nil {
		return m.Items
	}
	return nil
}

type ListCportResponse struct {
	Error                *v1.Error      `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Data                 *ListCportData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListCportResponse) Reset()         { *m = ListCportResponse{} }
func (m *ListCportResponse) String() string { return proto.CompactTextString(m) }
func (*ListCportResponse) ProtoMessage()    {}
func (*ListCportResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c309c866379f7da8, []int{3}
}

func (m *ListCportResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCportResponse.Unmarshal(m, b)
}
func (m *ListCportResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCportResponse.Marshal(b, m, deterministic)
}
func (m *ListCportResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCportResponse.Merge(m, src)
}
func (m *ListCportResponse) XXX_Size() int {
	return xxx_messageInfo_ListCportResponse.Size(m)
}
func (m *ListCportResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCportResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListCportResponse proto.InternalMessageInfo

func (m *ListCportResponse) GetError() *v1.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ListCportResponse) GetData() *ListCportData {
	if m != nil {
		return m.Data
	}
	return nil
}

type ListCportCondition struct {
	CportUuids           []string `protobuf:"bytes,1,rep,name=cportUuids,proto3" json:"cportUuids,omitempty"`
	CportExclude         bool     `protobuf:"varint,2,opt,name=cportExclude,proto3" json:"cportExclude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCportCondition) Reset()         { *m = ListCportCondition{} }
func (m *ListCportCondition) String() string { return proto.CompactTextString(m) }
func (*ListCportCondition) ProtoMessage()    {}
func (*ListCportCondition) Descriptor() ([]byte, []int) {
	return fileDescriptor_c309c866379f7da8, []int{4}
}

func (m *ListCportCondition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCportCondition.Unmarshal(m, b)
}
func (m *ListCportCondition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCportCondition.Marshal(b, m, deterministic)
}
func (m *ListCportCondition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCportCondition.Merge(m, src)
}
func (m *ListCportCondition) XXX_Size() int {
	return xxx_messageInfo_ListCportCondition.Size(m)
}
func (m *ListCportCondition) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCportCondition.DiscardUnknown(m)
}

var xxx_messageInfo_ListCportCondition proto.InternalMessageInfo

func (m *ListCportCondition) GetCportUuids() []string {
	if m != nil {
		return m.CportUuids
	}
	return nil
}

func (m *ListCportCondition) GetCportExclude() bool {
	if m != nil {
		return m.CportExclude
	}
	return false
}

type ListCportRequest struct {
	Header               *v1.Header          `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Offset               int32               `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int32               `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Simplify             bool                `protobuf:"varint,4,opt,name=simplify,proto3" json:"simplify,omitempty"`
	Condition            *ListCportCondition `protobuf:"bytes,5,opt,name=condition,proto3" json:"condition,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ListCportRequest) Reset()         { *m = ListCportRequest{} }
func (m *ListCportRequest) String() string { return proto.CompactTextString(m) }
func (*ListCportRequest) ProtoMessage()    {}
func (*ListCportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c309c866379f7da8, []int{5}
}

func (m *ListCportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCportRequest.Unmarshal(m, b)
}
func (m *ListCportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCportRequest.Marshal(b, m, deterministic)
}
func (m *ListCportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCportRequest.Merge(m, src)
}
func (m *ListCportRequest) XXX_Size() int {
	return xxx_messageInfo_ListCportRequest.Size(m)
}
func (m *ListCportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCportRequest proto.InternalMessageInfo

func (m *ListCportRequest) GetHeader() *v1.Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ListCportRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListCportRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListCportRequest) GetSimplify() bool {
	if m != nil {
		return m.Simplify
	}
	return false
}

func (m *ListCportRequest) GetCondition() *ListCportCondition {
	if m != nil {
		return m.Condition
	}
	return nil
}

type CreateCportRequest struct {
	Header               *v1.Header                `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Cport                *CreateCportRequest_Input `protobuf:"bytes,2,opt,name=cport,proto3" json:"cport,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *CreateCportRequest) Reset()         { *m = CreateCportRequest{} }
func (m *CreateCportRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCportRequest) ProtoMessage()    {}
func (*CreateCportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c309c866379f7da8, []int{6}
}

func (m *CreateCportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCportRequest.Unmarshal(m, b)
}
func (m *CreateCportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCportRequest.Marshal(b, m, deterministic)
}
func (m *CreateCportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCportRequest.Merge(m, src)
}
func (m *CreateCportRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCportRequest.Size(m)
}
func (m *CreateCportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCportRequest proto.InternalMessageInfo

func (m *CreateCportRequest) GetHeader() *v1.Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CreateCportRequest) GetCport() *CreateCportRequest_Input {
	if m != nil {
		return m.Cport
	}
	return nil
}

type CreateCportRequest_Input struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	VpcUuid              string   `protobuf:"bytes,2,opt,name=vpcUuid,proto3" json:"vpcUuid,omitempty"`
	SubnetUuid           string   `protobuf:"bytes,3,opt,name=subnetUuid,proto3" json:"subnetUuid,omitempty"`
	PodNs                string   `protobuf:"bytes,4,opt,name=podNs,proto3" json:"podNs,omitempty"`
	PodName              string   `protobuf:"bytes,5,opt,name=podName,proto3" json:"podName,omitempty"`
	ContainerId          string   `protobuf:"bytes,6,opt,name=containerId,proto3" json:"containerId,omitempty"`
	VmIp                 string   `protobuf:"bytes,7,opt,name=vmIp,proto3" json:"vmIp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCportRequest_Input) Reset()         { *m = CreateCportRequest_Input{} }
func (m *CreateCportRequest_Input) String() string { return proto.CompactTextString(m) }
func (*CreateCportRequest_Input) ProtoMessage()    {}
func (*CreateCportRequest_Input) Descriptor() ([]byte, []int) {
	return fileDescriptor_c309c866379f7da8, []int{6, 0}
}

func (m *CreateCportRequest_Input) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCportRequest_Input.Unmarshal(m, b)
}
func (m *CreateCportRequest_Input) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCportRequest_Input.Marshal(b, m, deterministic)
}
func (m *CreateCportRequest_Input) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCportRequest_Input.Merge(m, src)
}
func (m *CreateCportRequest_Input) XXX_Size() int {
	return xxx_messageInfo_CreateCportRequest_Input.Size(m)
}
func (m *CreateCportRequest_Input) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCportRequest_Input.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCportRequest_Input proto.InternalMessageInfo

func (m *CreateCportRequest_Input) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *CreateCportRequest_Input) GetVpcUuid() string {
	if m != nil {
		return m.VpcUuid
	}
	return ""
}

func (m *CreateCportRequest_Input) GetSubnetUuid() string {
	if m != nil {
		return m.SubnetUuid
	}
	return ""
}

func (m *CreateCportRequest_Input) GetPodNs() string {
	if m != nil {
		return m.PodNs
	}
	return ""
}

func (m *CreateCportRequest_Input) GetPodName() string {
	if m != nil {
		return m.PodName
	}
	return ""
}

func (m *CreateCportRequest_Input) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

func (m *CreateCportRequest_Input) GetVmIp() string {
	if m != nil {
		return m.VmIp
	}
	return ""
}

type CreateCportResponse struct {
	Error                *v1.Error     `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Data                 []*v1.JobInfo `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateCportResponse) Reset()         { *m = CreateCportResponse{} }
func (m *CreateCportResponse) String() string { return proto.CompactTextString(m) }
func (*CreateCportResponse) ProtoMessage()    {}
func (*CreateCportResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c309c866379f7da8, []int{7}
}

func (m *CreateCportResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCportResponse.Unmarshal(m, b)
}
func (m *CreateCportResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCportResponse.Marshal(b, m, deterministic)
}
func (m *CreateCportResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCportResponse.Merge(m, src)
}
func (m *CreateCportResponse) XXX_Size() int {
	return xxx_messageInfo_CreateCportResponse.Size(m)
}
func (m *CreateCportResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCportResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCportResponse proto.InternalMessageInfo

func (m *CreateCportResponse) GetError() *v1.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *CreateCportResponse) GetData() []*v1.JobInfo {
	if m != nil {
		return m.Data
	}
	return nil
}

type DeleteCportRequest struct {
	Header               *v1.Header                  `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Cport                []*DeleteCportRequest_Input `protobuf:"bytes,2,rep,name=cport,proto3" json:"cport,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *DeleteCportRequest) Reset()         { *m = DeleteCportRequest{} }
func (m *DeleteCportRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteCportRequest) ProtoMessage()    {}
func (*DeleteCportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c309c866379f7da8, []int{8}
}

func (m *DeleteCportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCportRequest.Unmarshal(m, b)
}
func (m *DeleteCportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCportRequest.Marshal(b, m, deterministic)
}
func (m *DeleteCportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCportRequest.Merge(m, src)
}
func (m *DeleteCportRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteCportRequest.Size(m)
}
func (m *DeleteCportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCportRequest proto.InternalMessageInfo

func (m *DeleteCportRequest) GetHeader() *v1.Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DeleteCportRequest) GetCport() []*DeleteCportRequest_Input {
	if m != nil {
		return m.Cport
	}
	return nil
}

type DeleteCportRequest_Input struct {
	CportUuid            string   `protobuf:"bytes,1,opt,name=cportUuid,proto3" json:"cportUuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCportRequest_Input) Reset()         { *m = DeleteCportRequest_Input{} }
func (m *DeleteCportRequest_Input) String() string { return proto.CompactTextString(m) }
func (*DeleteCportRequest_Input) ProtoMessage()    {}
func (*DeleteCportRequest_Input) Descriptor() ([]byte, []int) {
	return fileDescriptor_c309c866379f7da8, []int{8, 0}
}

func (m *DeleteCportRequest_Input) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCportRequest_Input.Unmarshal(m, b)
}
func (m *DeleteCportRequest_Input) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCportRequest_Input.Marshal(b, m, deterministic)
}
func (m *DeleteCportRequest_Input) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCportRequest_Input.Merge(m, src)
}
func (m *DeleteCportRequest_Input) XXX_Size() int {
	return xxx_messageInfo_DeleteCportRequest_Input.Size(m)
}
func (m *DeleteCportRequest_Input) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCportRequest_Input.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCportRequest_Input proto.InternalMessageInfo

func (m *DeleteCportRequest_Input) GetCportUuid() string {
	if m != nil {
		return m.CportUuid
	}
	return ""
}

type DeleteCportResponse struct {
	Error                *v1.Error     `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Data                 []*v1.JobInfo `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DeleteCportResponse) Reset()         { *m = DeleteCportResponse{} }
func (m *DeleteCportResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteCportResponse) ProtoMessage()    {}
func (*DeleteCportResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c309c866379f7da8, []int{9}
}

func (m *DeleteCportResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCportResponse.Unmarshal(m, b)
}
func (m *DeleteCportResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCportResponse.Marshal(b, m, deterministic)
}
func (m *DeleteCportResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCportResponse.Merge(m, src)
}
func (m *DeleteCportResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteCportResponse.Size(m)
}
func (m *DeleteCportResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCportResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCportResponse proto.InternalMessageInfo

func (m *DeleteCportResponse) GetError() *v1.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *DeleteCportResponse) GetData() []*v1.JobInfo {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetCportByUuidRequest)(nil), "didi.cloud.compute.v1.GetCportByUuidRequest")
	proto.RegisterType((*GetCportByUuidResponse)(nil), "didi.cloud.compute.v1.GetCportByUuidResponse")
	proto.RegisterType((*ListCportData)(nil), "didi.cloud.compute.v1.ListCportData")
	proto.RegisterType((*ListCportResponse)(nil), "didi.cloud.compute.v1.ListCportResponse")
	proto.RegisterType((*ListCportCondition)(nil), "didi.cloud.compute.v1.ListCportCondition")
	proto.RegisterType((*ListCportRequest)(nil), "didi.cloud.compute.v1.ListCportRequest")
	proto.RegisterType((*CreateCportRequest)(nil), "didi.cloud.compute.v1.CreateCportRequest")
	proto.RegisterType((*CreateCportRequest_Input)(nil), "didi.cloud.compute.v1.CreateCportRequest.Input")
	proto.RegisterType((*CreateCportResponse)(nil), "didi.cloud.compute.v1.CreateCportResponse")
	proto.RegisterType((*DeleteCportRequest)(nil), "didi.cloud.compute.v1.DeleteCportRequest")
	proto.RegisterType((*DeleteCportRequest_Input)(nil), "didi.cloud.compute.v1.DeleteCportRequest.Input")
	proto.RegisterType((*DeleteCportResponse)(nil), "didi.cloud.compute.v1.DeleteCportResponse")
}

func init() {
	proto.RegisterFile("compute/v1/cport.proto", fileDescriptor_c309c866379f7da8)
}

var fileDescriptor_c309c866379f7da8 = []byte{
	// 808 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xc1, 0x6e, 0xdb, 0x46,
	0x10, 0x05, 0x25, 0x51, 0x36, 0x47, 0xad, 0xd1, 0xae, 0x6b, 0x9b, 0x66, 0x0d, 0x57, 0x20, 0xdc,
	0x56, 0x2e, 0x6a, 0x12, 0x52, 0x8d, 0xa2, 0x40, 0x6f, 0x96, 0x0c, 0x57, 0x45, 0x51, 0x14, 0x04,
	0x6a, 0x14, 0x85, 0x73, 0xa0, 0xc8, 0x95, 0xbc, 0xb0, 0xc8, 0x65, 0xc8, 0xa5, 0x1c, 0x5f, 0x12,
	0x20, 0x08, 0x90, 0x4b, 0x90, 0x4b, 0xfe, 0x22, 0x1f, 0x90, 0x9c, 0xf2, 0x01, 0xc9, 0x35, 0xa7,
	0xdc, 0xf3, 0x11, 0x39, 0x06, 0x1c, 0x52, 0x32, 0x69, 0x4b, 0x91, 0xe0, 0x04, 0x39, 0x89, 0x33,
	0xf3, 0x76, 0xde, 0x9b, 0x99, 0x1d, 0x91, 0xb0, 0xee, 0x70, 0x2f, 0x88, 0x05, 0x35, 0x47, 0x4d,
	0xd3, 0x09, 0x78, 0x28, 0x8c, 0x20, 0xe4, 0x82, 0x93, 0x35, 0x97, 0xb9, 0xcc, 0x70, 0x86, 0x3c,
	0x76, 0x8d, 0x0c, 0x62, 0x8c, 0x9a, 0x1a, 0xe9, 0xd9, 0x11, 0x62, 0x93, 0xdf, 0x14, 0xaa, 0x6d,
	0xe4, 0x53, 0x70, 0xcf, 0xe3, 0x7e, 0x16, 0xd8, 0x1a, 0x70, 0x3e, 0x18, 0x52, 0xd3, 0x0e, 0x98,
	0x69, 0xfb, 0x3e, 0x17, 0xb6, 0x60, 0xdc, 0x8f, 0xd2, 0xa8, 0xce, 0x60, 0xed, 0x88, 0x8a, 0x76,
	0xc2, 0x79, 0x70, 0xf1, 0x6f, 0xcc, 0x5c, 0x8b, 0xde, 0x8e, 0x69, 0x24, 0x48, 0x0b, 0xaa, 0xa7,
	0xd4, 0x76, 0x69, 0xa8, 0x4a, 0x75, 0xa9, 0x51, 0x6b, 0x69, 0x46, 0x4e, 0x0b, 0xf2, 0x8e, 0x9a,
	0xc6, 0x1f, 0x88, 0xb0, 0x32, 0x24, 0xd9, 0x02, 0x05, 0xd5, 0x27, 0x79, 0xd4, 0x52, 0x5d, 0x6a,
	0x28, 0xd6, 0xa5, 0x43, 0xbf, 0x07, 0xeb, 0x57, 0xa9, 0xa2, 0x80, 0xfb, 0x11, 0x25, 0x26, 0xc8,
	0x34, 0x0c, 0xf9, 0x98, 0x6a, 0x73, 0x1a, 0xd5, 0x61, 0x02, 0xb0, 0x52, 0x1c, 0xd9, 0x87, 0x8a,
	0x6b, 0x0b, 0x5b, 0x2d, 0xd5, 0xcb, 0x8d, 0x5a, 0xab, 0x6e, 0x4c, 0x6d, 0x93, 0x81, 0x54, 0x5d,
	0xbf, 0xcf, 0x2d, 0x44, 0xeb, 0xb7, 0xe0, 0xcb, 0xbf, 0x58, 0x94, 0x2a, 0xe8, 0xd8, 0xc2, 0x26,
	0xdf, 0x80, 0xec, 0xf0, 0xd8, 0x17, 0xc8, 0x2b, 0x5b, 0xa9, 0x41, 0x7e, 0x05, 0x99, 0x09, 0xea,
	0x45, 0x0b, 0x67, 0x4f, 0xe1, 0xfa, 0x5d, 0xf8, 0x7a, 0x92, 0xfe, 0xe6, 0xa5, 0xfd, 0x36, 0x29,
	0x2d, 0xc1, 0xef, 0xcc, 0x20, 0x2f, 0xd4, 0x91, 0x95, 0xf7, 0x1f, 0x90, 0x89, 0xbb, 0xcd, 0x7d,
	0x97, 0x25, 0x73, 0x26, 0xdb, 0x00, 0x93, 0x11, 0x44, 0xaa, 0x54, 0x2f, 0x37, 0x14, 0x2b, 0xe7,
	0x21, 0x3a, 0x7c, 0x81, 0xd6, 0xe1, 0x1d, 0x67, 0x18, 0xbb, 0x14, 0x79, 0x97, 0xad, 0x82, 0x4f,
	0x7f, 0x23, 0xc1, 0x57, 0xb9, 0xd2, 0x6e, 0x7e, 0x41, 0xd6, 0xa1, 0xca, 0xfb, 0xfd, 0x88, 0x0a,
	0xa4, 0x91, 0xad, 0xcc, 0x4a, 0x06, 0x31, 0x64, 0x1e, 0x13, 0x6a, 0x39, 0x1d, 0x04, 0x1a, 0x44,
	0x83, 0xe5, 0x88, 0x79, 0xc1, 0x90, 0xf5, 0x2f, 0xd4, 0x0a, 0xca, 0x9a, 0xd8, 0xe4, 0x08, 0x14,
	0x67, 0x5c, 0xa3, 0x2a, 0xa3, 0x80, 0xdd, 0x79, 0xbd, 0x9a, 0x34, 0xc5, 0xba, 0x3c, 0xab, 0xbf,
	0x2c, 0x01, 0x69, 0x87, 0xd4, 0x16, 0xf4, 0xa3, 0xab, 0x3b, 0x04, 0x19, 0xdb, 0x96, 0xcd, 0xce,
	0x9c, 0x75, 0x71, 0xae, 0xb1, 0x19, 0x5d, 0x3f, 0x88, 0x85, 0x95, 0x9e, 0xd6, 0x9e, 0x4b, 0x20,
	0xa3, 0x83, 0xac, 0x40, 0x89, 0x05, 0x28, 0x40, 0xb1, 0x4a, 0x2c, 0x20, 0x2a, 0x2c, 0x8d, 0x02,
	0x27, 0xb7, 0x5d, 0x63, 0x33, 0x99, 0x72, 0x14, 0xf7, 0x7c, 0x9a, 0xae, 0x5e, 0x19, 0x83, 0x39,
	0x4f, 0xd2, 0xe0, 0x80, 0xbb, 0x7f, 0x47, 0xd8, 0x47, 0xc5, 0x4a, 0x8d, 0x24, 0x5f, 0xf2, 0x60,
	0x7b, 0x14, 0x5b, 0xa8, 0x58, 0x63, 0x93, 0xd4, 0xa1, 0xe6, 0x70, 0x5f, 0xd8, 0xcc, 0xa7, 0x61,
	0xd7, 0x55, 0xab, 0x18, 0xcd, 0xbb, 0x08, 0x81, 0xca, 0xc8, 0xeb, 0x06, 0xea, 0x12, 0x86, 0xf0,
	0x59, 0x3f, 0x87, 0xd5, 0x42, 0x71, 0x37, 0xdd, 0x01, 0xb3, 0xb0, 0xde, 0xdf, 0x4e, 0xc3, 0xff,
	0xc9, 0x7b, 0xb9, 0xcd, 0x7e, 0x26, 0x01, 0xe9, 0xd0, 0x21, 0xfd, 0xb4, 0x43, 0x2c, 0x7f, 0x60,
	0x88, 0xd7, 0xd9, 0x8a, 0x43, 0xfc, 0x7e, 0x3c, 0xc3, 0xc2, 0x7f, 0xa2, 0x74, 0xf5, 0x3f, 0xf1,
	0x1c, 0x56, 0x0b, 0x99, 0x3e, 0x57, 0xc7, 0x5a, 0x2f, 0x2a, 0x20, 0x23, 0x27, 0x79, 0x20, 0x81,
	0x32, 0x59, 0x11, 0xf2, 0xe3, 0xbc, 0x25, 0xca, 0xaa, 0xd5, 0x1a, 0xf3, 0x81, 0x69, 0x31, 0xfa,
	0xce, 0xfd, 0xd7, 0x6f, 0x9f, 0x94, 0xb6, 0xf5, 0x2d, 0x97, 0x21, 0xda, 0x64, 0xa6, 0x4f, 0xc5,
	0x39, 0x0f, 0xcf, 0xd2, 0x77, 0x9d, 0x39, 0x64, 0x91, 0x20, 0x8f, 0x25, 0x58, 0x29, 0xbe, 0x1e,
	0xc8, 0xcf, 0x33, 0x28, 0xa6, 0xbe, 0xb0, 0xb4, 0xbd, 0x05, 0xd1, 0x99, 0xaa, 0xef, 0x50, 0xd5,
	0x26, 0xd9, 0x98, 0xa1, 0x8a, 0x3c, 0x92, 0xa0, 0x96, 0xbb, 0xcd, 0x64, 0x77, 0xe1, 0x75, 0xd6,
	0x7e, 0x5a, 0x04, 0x9a, 0xe9, 0xf8, 0x01, 0x75, 0xd4, 0xf5, 0xed, 0x59, 0xdd, 0x71, 0xf0, 0x10,
	0xca, 0xc9, 0x5d, 0x95, 0x99, 0x72, 0xae, 0x5f, 0xcc, 0x99, 0x72, 0xa6, 0xdc, 0xbc, 0xf9, 0x72,
	0x5c, 0x3c, 0x74, 0xf0, 0x50, 0xc2, 0x8f, 0x16, 0xcc, 0x7c, 0x11, 0xfb, 0xb9, 0xb4, 0xff, 0x48,
	0xff, 0xef, 0x0f, 0x98, 0x38, 0x8d, 0x7b, 0x89, 0xd3, 0xcc, 0x00, 0xe3, 0xdf, 0xbd, 0x01, 0xdf,
	0x8b, 0xdc, 0x33, 0xf3, 0xf2, 0x6b, 0xe5, 0xf7, 0xec, 0xf1, 0x9d, 0x24, 0x3d, 0x2d, 0x6d, 0x74,
	0x12, 0xa5, 0x9d, 0x94, 0xdb, 0x68, 0x67, 0x49, 0x8f, 0x9b, 0xaf, 0xd2, 0xc8, 0x49, 0x16, 0x39,
	0xc9, 0x22, 0x27, 0xc7, 0xcd, 0x5e, 0x15, 0x3f, 0x64, 0x7e, 0x79, 0x1f, 0x00, 0x00, 0xff, 0xff,
	0xc2, 0x92, 0x00, 0xf1, 0x44, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CportClient is the client API for Cport service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CportClient interface {
	// 获取cport列表
	ListCport(ctx context.Context, in *ListCportRequest, opts ...grpc.CallOption) (*ListCportResponse, error)
	// 根据Uuid查询cport
	GetCportByUuid(ctx context.Context, in *GetCportByUuidRequest, opts ...grpc.CallOption) (*GetCportByUuidResponse, error)
	// 创建cport
	CreateCport(ctx context.Context, in *CreateCportRequest, opts ...grpc.CallOption) (*CreateCportResponse, error)
	// 删除cport
	DeleteCport(ctx context.Context, in *DeleteCportRequest, opts ...grpc.CallOption) (*DeleteCportResponse, error)
}

type cportClient struct {
	cc grpc.ClientConnInterface
}

func NewCportClient(cc grpc.ClientConnInterface) CportClient {
	return &cportClient{cc}
}

func (c *cportClient) ListCport(ctx context.Context, in *ListCportRequest, opts ...grpc.CallOption) (*ListCportResponse, error) {
	out := new(ListCportResponse)
	err := c.cc.Invoke(ctx, "/didi.cloud.compute.v1.Cport/ListCport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cportClient) GetCportByUuid(ctx context.Context, in *GetCportByUuidRequest, opts ...grpc.CallOption) (*GetCportByUuidResponse, error) {
	out := new(GetCportByUuidResponse)
	err := c.cc.Invoke(ctx, "/didi.cloud.compute.v1.Cport/GetCportByUuid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cportClient) CreateCport(ctx context.Context, in *CreateCportRequest, opts ...grpc.CallOption) (*CreateCportResponse, error) {
	out := new(CreateCportResponse)
	err := c.cc.Invoke(ctx, "/didi.cloud.compute.v1.Cport/CreateCport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cportClient) DeleteCport(ctx context.Context, in *DeleteCportRequest, opts ...grpc.CallOption) (*DeleteCportResponse, error) {
	out := new(DeleteCportResponse)
	err := c.cc.Invoke(ctx, "/didi.cloud.compute.v1.Cport/DeleteCport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CportServer is the server API for Cport service.
type CportServer interface {
	// 获取cport列表
	ListCport(context.Context, *ListCportRequest) (*ListCportResponse, error)
	// 根据Uuid查询cport
	GetCportByUuid(context.Context, *GetCportByUuidRequest) (*GetCportByUuidResponse, error)
	// 创建cport
	CreateCport(context.Context, *CreateCportRequest) (*CreateCportResponse, error)
	// 删除cport
	DeleteCport(context.Context, *DeleteCportRequest) (*DeleteCportResponse, error)
}

// UnimplementedCportServer can be embedded to have forward compatible implementations.
type UnimplementedCportServer struct {
}

func (*UnimplementedCportServer) ListCport(ctx context.Context, req *ListCportRequest) (*ListCportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCport not implemented")
}
func (*UnimplementedCportServer) GetCportByUuid(ctx context.Context, req *GetCportByUuidRequest) (*GetCportByUuidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCportByUuid not implemented")
}
func (*UnimplementedCportServer) CreateCport(ctx context.Context, req *CreateCportRequest) (*CreateCportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCport not implemented")
}
func (*UnimplementedCportServer) DeleteCport(ctx context.Context, req *DeleteCportRequest) (*DeleteCportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCport not implemented")
}

func RegisterCportServer(s *grpc.Server, srv CportServer) {
	s.RegisterService(&_Cport_serviceDesc, srv)
}

func _Cport_ListCport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CportServer).ListCport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/didi.cloud.compute.v1.Cport/ListCport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CportServer).ListCport(ctx, req.(*ListCportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cport_GetCportByUuid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCportByUuidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CportServer).GetCportByUuid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/didi.cloud.compute.v1.Cport/GetCportByUuid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CportServer).GetCportByUuid(ctx, req.(*GetCportByUuidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cport_CreateCport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CportServer).CreateCport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/didi.cloud.compute.v1.Cport/CreateCport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CportServer).CreateCport(ctx, req.(*CreateCportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cport_DeleteCport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CportServer).DeleteCport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/didi.cloud.compute.v1.Cport/DeleteCport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CportServer).DeleteCport(ctx, req.(*DeleteCportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cport_serviceDesc = grpc.ServiceDesc{
	ServiceName: "didi.cloud.compute.v1.Cport",
	HandlerType: (*CportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCport",
			Handler:    _Cport_ListCport_Handler,
		},
		{
			MethodName: "GetCportByUuid",
			Handler:    _Cport_GetCportByUuid_Handler,
		},
		{
			MethodName: "CreateCport",
			Handler:    _Cport_CreateCport_Handler,
		},
		{
			MethodName: "DeleteCport",
			Handler:    _Cport_DeleteCport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "compute/v1/cport.proto",
}
