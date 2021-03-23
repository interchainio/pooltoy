// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pooltoy/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgCreateUser struct {
	User    *User  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Creator string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *MsgCreateUser) Reset()         { *m = MsgCreateUser{} }
func (m *MsgCreateUser) String() string { return proto.CompactTextString(m) }
func (*MsgCreateUser) ProtoMessage()    {}
func (*MsgCreateUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_161a2114346f3644, []int{0}
}
func (m *MsgCreateUser) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateUser.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateUser.Merge(m, src)
}
func (m *MsgCreateUser) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateUser) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateUser.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateUser proto.InternalMessageInfo

func (m *MsgCreateUser) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *MsgCreateUser) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

type MsgCreateUserResponse struct {
}

func (m *MsgCreateUserResponse) Reset()         { *m = MsgCreateUserResponse{} }
func (m *MsgCreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateUserResponse) ProtoMessage()    {}
func (*MsgCreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_161a2114346f3644, []int{1}
}
func (m *MsgCreateUserResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateUserResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateUserResponse.Merge(m, src)
}
func (m *MsgCreateUserResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateUserResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateUser)(nil), "pooltoy.v1beta1.MsgCreateUser")
	proto.RegisterType((*MsgCreateUserResponse)(nil), "pooltoy.v1beta1.MsgCreateUserResponse")
}

func init() { proto.RegisterFile("pooltoy/v1beta1/tx.proto", fileDescriptor_161a2114346f3644) }

var fileDescriptor_161a2114346f3644 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0xbb, 0x55, 0x14, 0x57, 0x4a, 0x21, 0x58, 0x4c, 0x83, 0x2c, 0x25, 0x07, 0xa9, 0x97,
	0x2c, 0x6d, 0xdf, 0x40, 0xcf, 0xbd, 0x14, 0xbd, 0x78, 0x91, 0x4d, 0x58, 0xb7, 0x8b, 0xe9, 0x4e,
	0xd8, 0xdd, 0x48, 0x73, 0xf5, 0x09, 0x04, 0x5f, 0xca, 0x63, 0xc1, 0x8b, 0x47, 0x49, 0x7c, 0x10,
	0xc9, 0xbf, 0x62, 0x2b, 0x78, 0x9b, 0x99, 0xdf, 0x37, 0xdf, 0x7c, 0xbb, 0xd8, 0x4d, 0x00, 0x62,
	0x0b, 0x19, 0x7d, 0x9e, 0x84, 0xdc, 0xb2, 0x09, 0xb5, 0xeb, 0x20, 0xd1, 0x60, 0xc1, 0xe9, 0x37,
	0x24, 0x68, 0x88, 0x37, 0x8c, 0xc0, 0xac, 0xc0, 0x3c, 0x54, 0x98, 0xd6, 0x4d, 0xad, 0xf5, 0xce,
	0x04, 0x08, 0xa8, 0xe7, 0x65, 0xd5, 0x4c, 0x87, 0x02, 0x40, 0xc4, 0x9c, 0x56, 0x5d, 0x98, 0x3e,
	0x52, 0xa6, 0xb2, 0x06, 0x5d, 0x34, 0x88, 0x25, 0x92, 0x32, 0xa5, 0xc0, 0x32, 0x2b, 0x41, 0xb5,
	0x76, 0xde, 0x7e, 0xa8, 0xd4, 0x70, 0x5d, 0x33, 0xff, 0x16, 0xf7, 0xe6, 0x46, 0xdc, 0x68, 0xce,
	0x2c, 0xbf, 0x33, 0x5c, 0x3b, 0x57, 0xf8, 0xb0, 0xc4, 0x2e, 0x1a, 0xa1, 0xf1, 0xe9, 0x74, 0x10,
	0xec, 0xc5, 0x0e, 0x4a, 0xd1, 0xa2, 0x92, 0x38, 0x2e, 0x3e, 0x8e, 0xca, 0x45, 0xd0, 0x6e, 0x77,
	0x84, 0xc6, 0x27, 0x8b, 0xb6, 0xf5, 0xcf, 0xf1, 0x60, 0xc7, 0x75, 0xc1, 0x4d, 0x02, 0xca, 0xf0,
	0xa9, 0xc6, 0x07, 0x73, 0x23, 0x9c, 0x27, 0x8c, 0x7f, 0x9d, 0x24, 0x7f, 0x8e, 0xec, 0x2c, 0x7b,
	0x97, 0xff, 0xf3, 0xd6, 0xdc, 0x1f, 0xbc, 0x7c, 0x7c, 0xbf, 0x75, 0xfb, 0x7e, 0x8f, 0xb6, 0x0f,
	0x2e, 0x63, 0x5e, 0xcf, 0xdf, 0x73, 0x82, 0x36, 0x39, 0x41, 0x5f, 0x39, 0x41, 0xaf, 0x05, 0xe9,
	0x6c, 0x0a, 0xd2, 0xf9, 0x2c, 0x48, 0xe7, 0x7e, 0x26, 0xa4, 0x5d, 0xa6, 0x61, 0x10, 0xc1, 0x8a,
	0x4a, 0x65, 0xb9, 0x8e, 0x96, 0x4c, 0xaa, 0x90, 0xeb, 0x58, 0xaa, 0xad, 0xc7, 0x7a, 0x5b, 0xd9,
	0x2c, 0xe1, 0x26, 0x3c, 0xaa, 0x3e, 0x6e, 0xf6, 0x13, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x16, 0xbc,
	0x9d, 0xeb, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CreateUser(ctx context.Context, in *MsgCreateUser, opts ...grpc.CallOption) (*MsgCreateUserResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateUser(ctx context.Context, in *MsgCreateUser, opts ...grpc.CallOption) (*MsgCreateUserResponse, error) {
	out := new(MsgCreateUserResponse)
	err := c.cc.Invoke(ctx, "/pooltoy.v1beta1.Msg/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateUser(context.Context, *MsgCreateUser) (*MsgCreateUserResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateUser(ctx context.Context, req *MsgCreateUser) (*MsgCreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pooltoy.v1beta1.Msg/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateUser(ctx, req.(*MsgCreateUser))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pooltoy.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _Msg_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pooltoy/v1beta1/tx.proto",
}

func (m *MsgCreateUser) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateUser) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateUser) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if m.User != nil {
		{
			size, err := m.User.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateUserResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateUserResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateUserResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateUser) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.User != nil {
		l = m.User.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateUserResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateUser) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateUser: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateUser: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.User == nil {
				m.User = &User{}
			}
			if err := m.User.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreateUserResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateUserResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateUserResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)