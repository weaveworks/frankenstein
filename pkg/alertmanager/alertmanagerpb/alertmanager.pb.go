// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: alertmanager.proto

package alertmanagerpb

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	httpgrpc "github.com/weaveworks/common/httpgrpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strconv "strconv"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Status int32

const (
	OK            Status = 0
	ERROR         Status = 1
	SHUTTING_DOWN Status = 2
)

var Status_name = map[int32]string{
	0: "OK",
	1: "ERROR",
	2: "SHUTTING_DOWN",
}

var Status_value = map[string]int32{
	"OK":            0,
	"ERROR":         1,
	"SHUTTING_DOWN": 2,
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e60437b6e0c74c9a, []int{0}
}

type WriteRequest struct {
	UserID      string                `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	HttpRequest *httpgrpc.HTTPRequest `protobuf:"bytes,2,opt,name=httpRequest,proto3" json:"httpRequest,omitempty"`
}

func (m *WriteRequest) Reset()      { *m = WriteRequest{} }
func (*WriteRequest) ProtoMessage() {}
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e60437b6e0c74c9a, []int{0}
}
func (m *WriteRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WriteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WriteRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WriteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteRequest.Merge(m, src)
}
func (m *WriteRequest) XXX_Size() int {
	return m.Size()
}
func (m *WriteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WriteRequest proto.InternalMessageInfo

func (m *WriteRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *WriteRequest) GetHttpRequest() *httpgrpc.HTTPRequest {
	if m != nil {
		return m.HttpRequest
	}
	return nil
}

type WriteResponse struct {
	Status       Status                 `protobuf:"varint,1,opt,name=status,proto3,enum=alertmanagerpb.Status" json:"status,omitempty"`
	Error        string                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	HttpResponse *httpgrpc.HTTPResponse `protobuf:"bytes,3,opt,name=httpResponse,proto3" json:"httpResponse,omitempty"`
}

func (m *WriteResponse) Reset()      { *m = WriteResponse{} }
func (*WriteResponse) ProtoMessage() {}
func (*WriteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e60437b6e0c74c9a, []int{1}
}
func (m *WriteResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WriteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WriteResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WriteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteResponse.Merge(m, src)
}
func (m *WriteResponse) XXX_Size() int {
	return m.Size()
}
func (m *WriteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WriteResponse proto.InternalMessageInfo

func (m *WriteResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return OK
}

func (m *WriteResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *WriteResponse) GetHttpResponse() *httpgrpc.HTTPResponse {
	if m != nil {
		return m.HttpResponse
	}
	return nil
}

func init() {
	proto.RegisterEnum("alertmanagerpb.Status", Status_name, Status_value)
	proto.RegisterType((*WriteRequest)(nil), "alertmanagerpb.WriteRequest")
	proto.RegisterType((*WriteResponse)(nil), "alertmanagerpb.WriteResponse")
}

func init() { proto.RegisterFile("alertmanager.proto", fileDescriptor_e60437b6e0c74c9a) }

var fileDescriptor_e60437b6e0c74c9a = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4f, 0x6b, 0x1a, 0x41,
	0x1c, 0x9d, 0xb1, 0xb8, 0xe0, 0xf8, 0x07, 0x3b, 0xb4, 0x22, 0xd2, 0x0e, 0xe2, 0x49, 0x0a, 0xdd,
	0x05, 0x5b, 0x28, 0xf4, 0xd6, 0x62, 0xa9, 0xd2, 0xa2, 0x65, 0x5c, 0x11, 0x4a, 0x40, 0x76, 0xcd,
	0x64, 0x95, 0xb8, 0x3b, 0x9b, 0xd9, 0xd9, 0x78, 0xcd, 0x47, 0x48, 0xbe, 0x45, 0x3e, 0x4a, 0x8e,
	0x1e, 0x3d, 0xc6, 0xf5, 0x92, 0xa3, 0x1f, 0x21, 0x64, 0x66, 0x91, 0x55, 0xc8, 0xed, 0xf7, 0xe6,
	0xbd, 0xdf, 0xef, 0xbd, 0xdf, 0xcc, 0x20, 0xec, 0x2c, 0x99, 0x90, 0xbe, 0x13, 0x38, 0x1e, 0x13,
	0x66, 0x28, 0xb8, 0xe4, 0xb8, 0x92, 0x3d, 0x0b, 0xdd, 0xc6, 0x67, 0x6f, 0x21, 0xe7, 0xb1, 0x6b,
	0xce, 0xb8, 0x6f, 0x79, 0xdc, 0xe3, 0x96, 0x92, 0xb9, 0xf1, 0x85, 0x42, 0x0a, 0xa8, 0x4a, 0xb7,
	0x37, 0xbe, 0x66, 0xe4, 0x2b, 0xe6, 0x5c, 0xb3, 0x15, 0x17, 0x97, 0x91, 0x35, 0xe3, 0xbe, 0xcf,
	0x03, 0x6b, 0x2e, 0x65, 0xe8, 0x89, 0x70, 0x76, 0x28, 0x74, 0x57, 0x6b, 0x8a, 0x4a, 0x13, 0xb1,
	0x90, 0x8c, 0xb2, 0xab, 0x98, 0x45, 0x12, 0xd7, 0x90, 0x31, 0x8e, 0x98, 0xe8, 0x77, 0xeb, 0xb0,
	0x09, 0xdb, 0x05, 0x9a, 0x22, 0xfc, 0x0d, 0x15, 0x5f, 0x3a, 0x53, 0x59, 0x3d, 0xd7, 0x84, 0xed,
	0x62, 0xe7, 0xbd, 0x79, 0x98, 0xd6, 0xb3, 0xed, 0x7f, 0x29, 0x49, 0xb3, 0xca, 0xd6, 0x1d, 0x44,
	0xe5, 0xd4, 0x21, 0x0a, 0x79, 0x10, 0x31, 0x6c, 0x22, 0x23, 0x92, 0x8e, 0x8c, 0x23, 0x65, 0x51,
	0xe9, 0xd4, 0xcc, 0xe3, 0xc5, 0xcd, 0x91, 0x62, 0x69, 0xaa, 0xc2, 0xef, 0x50, 0x9e, 0x09, 0xc1,
	0x85, 0x32, 0x2d, 0x50, 0x0d, 0xf0, 0x77, 0x54, 0xd2, 0x36, 0x7a, 0x6a, 0xfd, 0x8d, 0x4a, 0x54,
	0x3b, 0x4d, 0xa4, 0x59, 0x7a, 0xa4, 0xfd, 0x64, 0x22, 0x43, 0x7b, 0x60, 0x03, 0xe5, 0x86, 0x7f,
	0xaa, 0x00, 0x17, 0x50, 0xfe, 0x17, 0xa5, 0x43, 0x5a, 0x85, 0xf8, 0x2d, 0x2a, 0x8f, 0x7a, 0x63,
	0xdb, 0xee, 0x0f, 0x7e, 0x4f, 0xbb, 0xc3, 0xc9, 0xa0, 0x9a, 0xeb, 0x9c, 0xa1, 0xd2, 0x8f, 0x4c,
	0x44, 0xfc, 0x17, 0x15, 0x7b, 0x4e, 0x70, 0xbe, 0x64, 0x6a, 0x31, 0xfc, 0xe1, 0x74, 0x81, 0xec,
	0x8d, 0x36, 0x3e, 0xbe, 0xc2, 0xea, 0x2c, 0x2d, 0xf0, 0xb3, 0xbb, 0xde, 0x12, 0xb0, 0xd9, 0x12,
	0xb0, 0xdf, 0x12, 0x78, 0x93, 0x10, 0x78, 0x9f, 0x10, 0xf8, 0x90, 0x10, 0xb8, 0x4e, 0x08, 0x7c,
	0x4c, 0x08, 0x7c, 0x4a, 0x08, 0xd8, 0x27, 0x04, 0xde, 0xee, 0x08, 0x58, 0xef, 0x08, 0xd8, 0xec,
	0x08, 0xf8, 0x7f, 0xf2, 0x5b, 0x5c, 0x43, 0xbd, 0xe7, 0x97, 0xe7, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xed, 0x53, 0x5b, 0x78, 0x5a, 0x02, 0x00, 0x00,
}

func (x Status) String() string {
	s, ok := Status_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *WriteRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WriteRequest)
	if !ok {
		that2, ok := that.(WriteRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.UserID != that1.UserID {
		return false
	}
	if !this.HttpRequest.Equal(that1.HttpRequest) {
		return false
	}
	return true
}
func (this *WriteResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WriteResponse)
	if !ok {
		that2, ok := that.(WriteResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	if this.Error != that1.Error {
		return false
	}
	if !this.HttpResponse.Equal(that1.HttpResponse) {
		return false
	}
	return true
}
func (this *WriteRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&alertmanagerpb.WriteRequest{")
	s = append(s, "UserID: "+fmt.Sprintf("%#v", this.UserID)+",\n")
	if this.HttpRequest != nil {
		s = append(s, "HttpRequest: "+fmt.Sprintf("%#v", this.HttpRequest)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *WriteResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&alertmanagerpb.WriteResponse{")
	s = append(s, "Status: "+fmt.Sprintf("%#v", this.Status)+",\n")
	s = append(s, "Error: "+fmt.Sprintf("%#v", this.Error)+",\n")
	if this.HttpResponse != nil {
		s = append(s, "HttpResponse: "+fmt.Sprintf("%#v", this.HttpResponse)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringAlertmanager(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AlertmanagerClient is the client API for Alertmanager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AlertmanagerClient interface {
	HandleWrite(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error)
}

type alertmanagerClient struct {
	cc *grpc.ClientConn
}

func NewAlertmanagerClient(cc *grpc.ClientConn) AlertmanagerClient {
	return &alertmanagerClient{cc}
}

func (c *alertmanagerClient) HandleWrite(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := c.cc.Invoke(ctx, "/alertmanagerpb.Alertmanager/HandleWrite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlertmanagerServer is the server API for Alertmanager service.
type AlertmanagerServer interface {
	HandleWrite(context.Context, *WriteRequest) (*WriteResponse, error)
}

// UnimplementedAlertmanagerServer can be embedded to have forward compatible implementations.
type UnimplementedAlertmanagerServer struct {
}

func (*UnimplementedAlertmanagerServer) HandleWrite(ctx context.Context, req *WriteRequest) (*WriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleWrite not implemented")
}

func RegisterAlertmanagerServer(s *grpc.Server, srv AlertmanagerServer) {
	s.RegisterService(&_Alertmanager_serviceDesc, srv)
}

func _Alertmanager_HandleWrite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertmanagerServer).HandleWrite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alertmanagerpb.Alertmanager/HandleWrite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertmanagerServer).HandleWrite(ctx, req.(*WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Alertmanager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "alertmanagerpb.Alertmanager",
	HandlerType: (*AlertmanagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleWrite",
			Handler:    _Alertmanager_HandleWrite_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "alertmanager.proto",
}

func (m *WriteRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WriteRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WriteRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.HttpRequest != nil {
		{
			size, err := m.HttpRequest.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAlertmanager(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.UserID) > 0 {
		i -= len(m.UserID)
		copy(dAtA[i:], m.UserID)
		i = encodeVarintAlertmanager(dAtA, i, uint64(len(m.UserID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WriteResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WriteResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WriteResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.HttpResponse != nil {
		{
			size, err := m.HttpResponse.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAlertmanager(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintAlertmanager(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0x12
	}
	if m.Status != 0 {
		i = encodeVarintAlertmanager(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAlertmanager(dAtA []byte, offset int, v uint64) int {
	offset -= sovAlertmanager(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WriteRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UserID)
	if l > 0 {
		n += 1 + l + sovAlertmanager(uint64(l))
	}
	if m.HttpRequest != nil {
		l = m.HttpRequest.Size()
		n += 1 + l + sovAlertmanager(uint64(l))
	}
	return n
}

func (m *WriteResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovAlertmanager(uint64(m.Status))
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovAlertmanager(uint64(l))
	}
	if m.HttpResponse != nil {
		l = m.HttpResponse.Size()
		n += 1 + l + sovAlertmanager(uint64(l))
	}
	return n
}

func sovAlertmanager(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAlertmanager(x uint64) (n int) {
	return sovAlertmanager(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *WriteRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&WriteRequest{`,
		`UserID:` + fmt.Sprintf("%v", this.UserID) + `,`,
		`HttpRequest:` + strings.Replace(fmt.Sprintf("%v", this.HttpRequest), "HTTPRequest", "httpgrpc.HTTPRequest", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *WriteResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&WriteResponse{`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`Error:` + fmt.Sprintf("%v", this.Error) + `,`,
		`HttpResponse:` + strings.Replace(fmt.Sprintf("%v", this.HttpResponse), "HTTPResponse", "httpgrpc.HTTPResponse", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringAlertmanager(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *WriteRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAlertmanager
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WriteRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WriteRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlertmanager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAlertmanager
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAlertmanager
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HttpRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlertmanager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAlertmanager
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAlertmanager
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HttpRequest == nil {
				m.HttpRequest = &httpgrpc.HTTPRequest{}
			}
			if err := m.HttpRequest.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAlertmanager(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAlertmanager
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAlertmanager
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *WriteResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAlertmanager
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WriteResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WriteResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlertmanager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlertmanager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAlertmanager
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAlertmanager
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HttpResponse", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlertmanager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAlertmanager
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAlertmanager
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HttpResponse == nil {
				m.HttpResponse = &httpgrpc.HTTPResponse{}
			}
			if err := m.HttpResponse.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAlertmanager(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAlertmanager
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAlertmanager
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAlertmanager(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAlertmanager
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAlertmanager
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAlertmanager
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthAlertmanager
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthAlertmanager
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAlertmanager
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipAlertmanager(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthAlertmanager
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthAlertmanager = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAlertmanager   = fmt.Errorf("proto: integer overflow")
)
