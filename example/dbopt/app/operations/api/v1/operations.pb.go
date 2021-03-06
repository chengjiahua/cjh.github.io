// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: operations.proto

package v1

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type DeploymentErrCode int32

const (
	DeploymentErrCode_NoUse          DeploymentErrCode = 0
	DeploymentErrCode_NameIsEmptyErr DeploymentErrCode = 16409
)

var DeploymentErrCode_name = map[int32]string{
	0:     "NoUse",
	16409: "NameIsEmptyErr",
}

var DeploymentErrCode_value = map[string]int32{
	"NoUse":          0,
	"NameIsEmptyErr": 16409,
}

func (x DeploymentErrCode) String() string {
	return proto.EnumName(DeploymentErrCode_name, int32(x))
}

func (DeploymentErrCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1b4a5877375e491e, []int{0}
}

type HelloReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" form:"name" validate:"required"`
	Age                  string   `protobuf:"bytes,2,opt,name=age,proto3" json:"age,omitempty" form:"age"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReq) Reset()         { *m = HelloReq{} }
func (m *HelloReq) String() string { return proto.CompactTextString(m) }
func (*HelloReq) ProtoMessage()    {}
func (*HelloReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b4a5877375e491e, []int{0}
}
func (m *HelloReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HelloReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HelloReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HelloReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReq.Merge(m, src)
}
func (m *HelloReq) XXX_Size() int {
	return m.Size()
}
func (m *HelloReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReq.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReq proto.InternalMessageInfo

type HelloResp struct {
	Content              string   `protobuf:"bytes,1,opt,name=Content,proto3" json:"content"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResp) Reset()         { *m = HelloResp{} }
func (m *HelloResp) String() string { return proto.CompactTextString(m) }
func (*HelloResp) ProtoMessage()    {}
func (*HelloResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b4a5877375e491e, []int{1}
}
func (m *HelloResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HelloResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HelloResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HelloResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResp.Merge(m, src)
}
func (m *HelloResp) XXX_Size() int {
	return m.Size()
}
func (m *HelloResp) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResp.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResp proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("service.operations.v1.DeploymentErrCode", DeploymentErrCode_name, DeploymentErrCode_value)
	proto.RegisterType((*HelloReq)(nil), "service.operations.v1.HelloReq")
	proto.RegisterType((*HelloResp)(nil), "service.operations.v1.HelloResp")
}

func init() { proto.RegisterFile("operations.proto", fileDescriptor_1b4a5877375e491e) }

var fileDescriptor_1b4a5877375e491e = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4f, 0x6b, 0x1a, 0x41,
	0x18, 0xc6, 0x5d, 0xfb, 0xc7, 0xee, 0x58, 0xac, 0x1d, 0x5a, 0x10, 0x5b, 0x76, 0x97, 0x29, 0x85,
	0x52, 0xe8, 0x2e, 0xda, 0x42, 0xc1, 0xa3, 0x56, 0x68, 0x69, 0x30, 0xb0, 0xc1, 0x4b, 0x6e, 0xa3,
	0xbe, 0x4e, 0x16, 0x76, 0xf7, 0x5d, 0x67, 0xc7, 0x05, 0x6f, 0x21, 0x5f, 0x21, 0x97, 0xe4, 0x1b,
	0x79, 0x0c, 0xe4, 0x2e, 0x89, 0xc9, 0x29, 0x47, 0x3f, 0x41, 0x70, 0x5c, 0x13, 0x0f, 0x21, 0xb7,
	0x79, 0xdf, 0xdf, 0xf3, 0xcc, 0x33, 0xf3, 0x90, 0x2a, 0x26, 0x20, 0xb9, 0x0a, 0x30, 0x4e, 0xdd,
	0x44, 0xa2, 0x42, 0xfa, 0x31, 0x05, 0x99, 0x05, 0x43, 0x70, 0x77, 0x48, 0xd6, 0xa8, 0xff, 0x10,
	0x81, 0x3a, 0x9a, 0x0e, 0xdc, 0x21, 0x46, 0x9e, 0x40, 0x81, 0x9e, 0x56, 0x0f, 0xa6, 0x63, 0x3d,
	0xe9, 0x41, 0x9f, 0x36, 0xb7, 0xd4, 0x3f, 0x09, 0x44, 0x11, 0xc2, 0xa3, 0x0a, 0xa2, 0x44, 0xcd,
	0x72, 0xf8, 0x39, 0x87, 0x3c, 0x09, 0x3c, 0x1e, 0xc7, 0xa8, 0x76, 0x1f, 0xc0, 0x80, 0xbc, 0xf9,
	0x0b, 0x61, 0x88, 0x3e, 0x4c, 0xe8, 0x6f, 0xf2, 0x32, 0xe6, 0x11, 0xd4, 0x0c, 0xc7, 0xf8, 0x66,
	0xb6, 0xbf, 0xac, 0x16, 0xb6, 0x3d, 0x46, 0x19, 0xb5, 0xd8, 0x7a, 0xcb, 0x9c, 0x8c, 0x87, 0xc1,
	0x88, 0x2b, 0x68, 0x31, 0x09, 0x93, 0x69, 0x20, 0x61, 0xc4, 0x7c, 0x6d, 0xa0, 0x0e, 0x79, 0xc1,
	0x05, 0xd4, 0x8a, 0xda, 0x57, 0x59, 0x2d, 0x6c, 0xb2, 0xf1, 0x71, 0x01, 0xcc, 0x5f, 0x23, 0xd6,
	0x24, 0x66, 0x1e, 0x93, 0x26, 0xf4, 0x2b, 0x29, 0x75, 0x30, 0x56, 0x10, 0xab, 0x3c, 0xaa, 0x7c,
	0xb7, 0xb0, 0x4b, 0xc3, 0xcd, 0xca, 0xdf, 0xb2, 0xef, 0xbf, 0xc8, 0xfb, 0x3f, 0x90, 0x84, 0x38,
	0x8b, 0x20, 0x56, 0x5d, 0x29, 0x3b, 0x38, 0x02, 0x6a, 0x92, 0x57, 0x3d, 0xec, 0xa7, 0x50, 0x2d,
	0xd0, 0x0f, 0xa4, 0xd2, 0xe3, 0x11, 0xfc, 0x4b, 0xbb, 0xeb, 0xdf, 0x76, 0xa5, 0xac, 0x9e, 0x1f,
	0x1b, 0xcd, 0x09, 0x79, 0xfb, 0x1f, 0xf7, 0x1f, 0xda, 0xa4, 0x9c, 0x94, 0x0f, 0xf8, 0x4c, 0x87,
	0xf7, 0xfd, 0x3d, 0x6a, 0xbb, 0x4f, 0x36, 0xee, 0x6e, 0x4b, 0xa8, 0x3b, 0xcf, 0x0b, 0xd2, 0x84,
	0xbd, 0x3b, 0xb9, 0xbc, 0x3d, 0x2d, 0x9a, 0xb4, 0xe4, 0x89, 0x54, 0x71, 0xa9, 0xda, 0xb5, 0xf9,
	0xb5, 0x55, 0x98, 0x2f, 0x2d, 0xe3, 0x62, 0x69, 0x19, 0x57, 0x4b, 0xcb, 0x38, 0xbb, 0xb1, 0x0a,
	0x87, 0xc5, 0xac, 0x31, 0x78, 0xad, 0x4b, 0xfe, 0x79, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x75, 0x9f,
	0xbb, 0xc3, 0xf9, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KoOperationsClient is the client API for KoOperations service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KoOperationsClient interface {
	SayHelloURL(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloResp, error)
}

type koOperationsClient struct {
	cc *grpc.ClientConn
}

func NewKoOperationsClient(cc *grpc.ClientConn) KoOperationsClient {
	return &koOperationsClient{cc}
}

func (c *koOperationsClient) SayHelloURL(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloResp, error) {
	out := new(HelloResp)
	err := c.cc.Invoke(ctx, "/service.operations.v1.KoOperations/SayHelloURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KoOperationsServer is the server API for KoOperations service.
type KoOperationsServer interface {
	SayHelloURL(context.Context, *HelloReq) (*HelloResp, error)
}

// UnimplementedKoOperationsServer can be embedded to have forward compatible implementations.
type UnimplementedKoOperationsServer struct {
}

func (*UnimplementedKoOperationsServer) SayHelloURL(ctx context.Context, req *HelloReq) (*HelloResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHelloURL not implemented")
}

func RegisterKoOperationsServer(s *grpc.Server, srv KoOperationsServer) {
	s.RegisterService(&_KoOperations_serviceDesc, srv)
}

func _KoOperations_SayHelloURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KoOperationsServer).SayHelloURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.operations.v1.KoOperations/SayHelloURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KoOperationsServer).SayHelloURL(ctx, req.(*HelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _KoOperations_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.operations.v1.KoOperations",
	HandlerType: (*KoOperationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHelloURL",
			Handler:    _KoOperations_SayHelloURL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "operations.proto",
}

func (m *HelloReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HelloReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HelloReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Age) > 0 {
		i -= len(m.Age)
		copy(dAtA[i:], m.Age)
		i = encodeVarintOperations(dAtA, i, uint64(len(m.Age)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintOperations(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HelloResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HelloResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HelloResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintOperations(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOperations(dAtA []byte, offset int, v uint64) int {
	offset -= sovOperations(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HelloReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovOperations(uint64(l))
	}
	l = len(m.Age)
	if l > 0 {
		n += 1 + l + sovOperations(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HelloResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovOperations(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovOperations(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOperations(x uint64) (n int) {
	return sovOperations(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HelloReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOperations
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
			return fmt.Errorf("proto: HelloReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HelloReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
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
				return ErrInvalidLengthOperations
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOperations
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Age", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
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
				return ErrInvalidLengthOperations
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOperations
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Age = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOperations(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOperations
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOperations
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HelloResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOperations
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
			return fmt.Errorf("proto: HelloResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HelloResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
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
				return ErrInvalidLengthOperations
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOperations
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOperations(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOperations
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOperations
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipOperations(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOperations
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
					return 0, ErrIntOverflowOperations
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowOperations
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
				return 0, ErrInvalidLengthOperations
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOperations
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOperations
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOperations        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOperations          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOperations = fmt.Errorf("proto: unexpected end of group")
)
