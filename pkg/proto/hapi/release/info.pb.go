// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hapi/release/info.proto

package release

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Info describes release information.
type Info struct {
	Status        *Status                    `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	FirstDeployed *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=first_deployed,json=firstDeployed" json:"first_deployed,omitempty"`
	LastDeployed  *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=last_deployed,json=lastDeployed" json:"last_deployed,omitempty"`
	// Deleted tracks when this object was deleted.
	Deleted *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=deleted" json:"deleted,omitempty"`
	// Description is human-friendly "log entry" about this release.
	Description string `protobuf:"bytes,5,opt,name=Description" json:"Description,omitempty"`
	// Username is the authenticated user who performed this release.
	Username string `protobuf:"bytes,6,opt,name=Username" json:"Username,omitempty"`
}

func (m *Info) Reset()                    { *m = Info{} }
func (m *Info) String() string            { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()               {}
func (*Info) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Info) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Info) GetFirstDeployed() *google_protobuf.Timestamp {
	if m != nil {
		return m.FirstDeployed
	}
	return nil
}

func (m *Info) GetLastDeployed() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastDeployed
	}
	return nil
}

func (m *Info) GetDeleted() *google_protobuf.Timestamp {
	if m != nil {
		return m.Deleted
	}
	return nil
}

func (m *Info) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Info) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func init() {
	proto.RegisterType((*Info)(nil), "hapi.release.Info")
}

func init() { proto.RegisterFile("hapi/release/info.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x95, 0x52, 0x52, 0xea, 0xb6, 0x0c, 0x16, 0x12, 0x26, 0x0b, 0x11, 0x53, 0x07, 0xe4,
	0x48, 0xc0, 0x8e, 0x40, 0x5d, 0x58, 0x03, 0x2c, 0x2c, 0xc8, 0x25, 0x97, 0x62, 0xc9, 0xc9, 0x59,
	0xf6, 0x75, 0xe0, 0x3f, 0xf1, 0x23, 0x51, 0x1d, 0xa7, 0x0a, 0x53, 0xc6, 0xe4, 0x7b, 0xdf, 0xbb,
	0x27, 0xb3, 0xcb, 0x6f, 0x65, 0x75, 0xe1, 0xc0, 0x80, 0xf2, 0x50, 0xe8, 0xb6, 0x46, 0x69, 0x1d,
	0x12, 0xf2, 0xe5, 0x01, 0xc8, 0x08, 0xb2, 0xeb, 0x1d, 0xe2, 0xce, 0x40, 0x11, 0xd8, 0x76, 0x5f,
	0x17, 0xa4, 0x1b, 0xf0, 0xa4, 0x1a, 0xdb, 0xc5, 0xb3, 0xab, 0x7f, 0x3d, 0x9e, 0x14, 0xed, 0x7d,
	0x87, 0x6e, 0x7e, 0x27, 0x6c, 0xfa, 0xd2, 0xd6, 0xc8, 0x6f, 0x59, 0xda, 0x01, 0x91, 0xe4, 0xc9,
	0x7a, 0x71, 0x77, 0x21, 0x87, 0x37, 0xe4, 0x6b, 0x60, 0x65, 0xcc, 0xf0, 0x27, 0x76, 0x5e, 0x6b,
	0xe7, 0xe9, 0xb3, 0x02, 0x6b, 0xf0, 0x07, 0x2a, 0x31, 0x09, 0x56, 0x26, 0xbb, 0x2d, 0xb2, 0xdf,
	0x22, 0xdf, 0xfa, 0x2d, 0xe5, 0x2a, 0x18, 0x9b, 0x28, 0xf0, 0x47, 0xb6, 0x32, 0x6a, 0xd8, 0x70,
	0x32, 0xda, 0xb0, 0x3c, 0x08, 0xc7, 0x82, 0x07, 0x36, 0xab, 0xc0, 0x00, 0x41, 0x25, 0xa6, 0xa3,
	0x6a, 0x1f, 0xe5, 0x39, 0x5b, 0x6c, 0xc0, 0x7f, 0x39, 0x6d, 0x49, 0x63, 0x2b, 0x4e, 0xf3, 0x64,
	0x3d, 0x2f, 0x87, 0xbf, 0x78, 0xc6, 0xce, 0xde, 0x3d, 0xb8, 0x56, 0x35, 0x20, 0xd2, 0x80, 0x8f,
	0xdf, 0xcf, 0xf3, 0x8f, 0x59, 0x7c, 0x91, 0x6d, 0x1a, 0xae, 0xdc, 0xff, 0x05, 0x00, 0x00, 0xff,
	0xff, 0xf4, 0x9e, 0xd9, 0x9e, 0xa5, 0x01, 0x00, 0x00,
}
