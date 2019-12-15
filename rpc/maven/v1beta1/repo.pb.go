// Code generated by protoc-gen-go. DO NOT EDIT.
// source: maven/v1beta1/repo.proto

package mavenv1beta1

import (
	fmt "fmt"
	v1beta1 "github.com/enseadaio/enseada/rpc/meta/v1beta1"
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

// A Maven repository.
type Repo struct {
	Metadata             *v1beta1.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	GroupId              string            `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	ArtifactId           string            `protobuf:"bytes,3,opt,name=artifact_id,json=artifactId,proto3" json:"artifact_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Repo) Reset()         { *m = Repo{} }
func (m *Repo) String() string { return proto.CompactTextString(m) }
func (*Repo) ProtoMessage()    {}
func (*Repo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fa16ca33d8b98a2, []int{0}
}

func (m *Repo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Repo.Unmarshal(m, b)
}
func (m *Repo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Repo.Marshal(b, m, deterministic)
}
func (m *Repo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Repo.Merge(m, src)
}
func (m *Repo) XXX_Size() int {
	return xxx_messageInfo_Repo.Size(m)
}
func (m *Repo) XXX_DiscardUnknown() {
	xxx_messageInfo_Repo.DiscardUnknown(m)
}

var xxx_messageInfo_Repo proto.InternalMessageInfo

func (m *Repo) GetMetadata() *v1beta1.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Repo) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *Repo) GetArtifactId() string {
	if m != nil {
		return m.ArtifactId
	}
	return ""
}

func init() {
	proto.RegisterType((*Repo)(nil), "maven.v1beta1.Repo")
}

func init() { proto.RegisterFile("maven/v1beta1/repo.proto", fileDescriptor_1fa16ca33d8b98a2) }

var fileDescriptor_1fa16ca33d8b98a2 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc8, 0x4d, 0x2c, 0x4b,
	0xcd, 0xd3, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x2f, 0x4a, 0x2d, 0xc8, 0xd7, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x05, 0xcb, 0xe8, 0x41, 0x65, 0xa4, 0xa4, 0x73, 0x53, 0x4b,
	0x12, 0xe1, 0xea, 0x40, 0x9c, 0x94, 0xc4, 0x92, 0x44, 0x88, 0x5a, 0xa5, 0x32, 0x2e, 0x96, 0xa0,
	0xd4, 0x82, 0x7c, 0x21, 0x23, 0x2e, 0x0e, 0x98, 0x8c, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91,
	0x98, 0x1e, 0x48, 0x00, 0x66, 0x8a, 0x9e, 0x2f, 0x54, 0x36, 0x08, 0xae, 0x4e, 0x48, 0x92, 0x8b,
	0x23, 0xbd, 0x28, 0xbf, 0xb4, 0x20, 0x3e, 0x33, 0x45, 0x82, 0x49, 0x81, 0x51, 0x83, 0x33, 0x88,
	0x1d, 0xcc, 0xf7, 0x4c, 0x11, 0x92, 0xe7, 0xe2, 0x4e, 0x2c, 0x2a, 0xc9, 0x4c, 0x4b, 0x4c, 0x2e,
	0x01, 0xc9, 0x32, 0x83, 0x65, 0xb9, 0x60, 0x42, 0x9e, 0x29, 0x4e, 0xd1, 0x5c, 0x12, 0x99, 0xf9,
	0x7a, 0xa9, 0x79, 0xc5, 0xa9, 0x89, 0x29, 0x89, 0x7a, 0x28, 0x0e, 0x76, 0xe2, 0x04, 0xb9, 0x28,
	0x00, 0xe4, 0xbc, 0x00, 0xc6, 0x28, 0x1e, 0xb0, 0x1c, 0x54, 0x6a, 0x11, 0x13, 0xb3, 0x6f, 0x44,
	0xc4, 0x2a, 0x26, 0x5e, 0x5f, 0xb0, 0x86, 0x30, 0x43, 0x27, 0x90, 0xe8, 0x29, 0x28, 0x3f, 0x06,
	0xca, 0x4f, 0x62, 0x03, 0xfb, 0xcd, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x70, 0xde, 0x8e,
	0x23, 0x01, 0x00, 0x00,
}
