// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: auth.proto

package auth

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import empty "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Allgemeine Anmeldeinformationen
type Credentials struct {
	// Der eindeutige Name des Benutzers.
	// Könnte auch die emailadresse sein
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// Das Kennwort des Benutzers
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// not-implemented
	// Der zweite Faktor.
	// Könnte TOPT Verfahren sein.
	SecondFactor         string   `protobuf:"bytes,3,opt,name=secondFactor,proto3" json:"secondFactor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Credentials) Reset()         { *m = Credentials{} }
func (m *Credentials) String() string { return proto.CompactTextString(m) }
func (*Credentials) ProtoMessage()    {}
func (*Credentials) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_a8026597ddcbe4ed, []int{0}
}
func (m *Credentials) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Credentials) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Credentials.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Credentials) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Credentials.Merge(dst, src)
}
func (m *Credentials) XXX_Size() int {
	return m.Size()
}
func (m *Credentials) XXX_DiscardUnknown() {
	xxx_messageInfo_Credentials.DiscardUnknown(m)
}

var xxx_messageInfo_Credentials proto.InternalMessageInfo

func (m *Credentials) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Credentials) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Credentials) GetSecondFactor() string {
	if m != nil {
		return m.SecondFactor
	}
	return ""
}

// Benutzer Credentials mit ID
type User struct {
	// Die ulid des Benutzers
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" db:"id,pk,omitempty"`
	// Der _neue_ Benutzername
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty" db:"username,omitempty"`
	// Das _neue_ Kennwort des Benutzers
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty" db:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_a8026597ddcbe4ed, []int{1}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_User.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return m.Size()
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// Request Daten für die Anmeldung
type CredentialsRequest struct {
	// enthält im mindesten:  username &amp;&amp; (password || secondFactor)
	Credentials          *Credentials `protobuf:"bytes,1,opt,name=credentials" json:"credentials,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CredentialsRequest) Reset()         { *m = CredentialsRequest{} }
func (m *CredentialsRequest) String() string { return proto.CompactTextString(m) }
func (*CredentialsRequest) ProtoMessage()    {}
func (*CredentialsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_a8026597ddcbe4ed, []int{2}
}
func (m *CredentialsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CredentialsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CredentialsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *CredentialsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialsRequest.Merge(dst, src)
}
func (m *CredentialsRequest) XXX_Size() int {
	return m.Size()
}
func (m *CredentialsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialsRequest proto.InternalMessageInfo

func (m *CredentialsRequest) GetCredentials() *Credentials {
	if m != nil {
		return m.Credentials
	}
	return nil
}

func init() {
	proto.RegisterType((*Credentials)(nil), "example.package.Credentials")
	proto.RegisterType((*User)(nil), "example.package.User")
	proto.RegisterType((*CredentialsRequest)(nil), "example.package.CredentialsRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	// Ausloggen, setzt nur den cookie auf leer, kann später die jwt.id revoken
	Logout(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	// Der Login setzt Header und Cookies, darum kein Body in der Antwort.
	// Falls du ein Profil erwartest, sollte dies über ProfileServices angeboten werden.
	Login(ctx context.Context, in *CredentialsRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Logout(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/example.package.AuthService/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Login(ctx context.Context, in *CredentialsRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/example.package.AuthService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	// Ausloggen, setzt nur den cookie auf leer, kann später die jwt.id revoken
	Logout(context.Context, *empty.Empty) (*empty.Empty, error)
	// Der Login setzt Header und Cookies, darum kein Body in der Antwort.
	// Falls du ein Profil erwartest, sollte dies über ProfileServices angeboten werden.
	Login(context.Context, *CredentialsRequest) (*empty.Empty, error)
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.package.AuthService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Logout(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.package.AuthService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*CredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.package.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Logout",
			Handler:    _AuthService_Logout_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}

func (m *Credentials) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Credentials) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Username) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAuth(dAtA, i, uint64(len(m.Username)))
		i += copy(dAtA[i:], m.Username)
	}
	if len(m.Password) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAuth(dAtA, i, uint64(len(m.Password)))
		i += copy(dAtA[i:], m.Password)
	}
	if len(m.SecondFactor) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintAuth(dAtA, i, uint64(len(m.SecondFactor)))
		i += copy(dAtA[i:], m.SecondFactor)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *User) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *User) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAuth(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.Username) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAuth(dAtA, i, uint64(len(m.Username)))
		i += copy(dAtA[i:], m.Username)
	}
	if len(m.Password) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintAuth(dAtA, i, uint64(len(m.Password)))
		i += copy(dAtA[i:], m.Password)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *CredentialsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CredentialsRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Credentials != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAuth(dAtA, i, uint64(m.Credentials.Size()))
		n1, err := m.Credentials.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintAuth(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Credentials) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovAuth(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovAuth(uint64(l))
	}
	l = len(m.SecondFactor)
	if l > 0 {
		n += 1 + l + sovAuth(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *User) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovAuth(uint64(l))
	}
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovAuth(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovAuth(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CredentialsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Credentials != nil {
		l = m.Credentials.Size()
		n += 1 + l + sovAuth(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAuth(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAuth(x uint64) (n int) {
	return sovAuth(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Credentials) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuth
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Credentials: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Credentials: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecondFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SecondFactor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuth
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
func (m *User) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuth
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: User: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: User: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuth
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
func (m *CredentialsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuth
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CredentialsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CredentialsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Credentials", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Credentials == nil {
				m.Credentials = &Credentials{}
			}
			if err := m.Credentials.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuth
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
func skipAuth(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuth
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
					return 0, ErrIntOverflowAuth
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
					return 0, ErrIntOverflowAuth
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthAuth
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAuth
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
				next, err := skipAuth(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthAuth = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuth   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("auth.proto", fileDescriptor_auth_a8026597ddcbe4ed) }

var fileDescriptor_auth_a8026597ddcbe4ed = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcf, 0x8a, 0xd3, 0x40,
	0x18, 0x27, 0xe9, 0x1f, 0x75, 0xa2, 0x08, 0x43, 0xd1, 0x90, 0x96, 0x56, 0xc6, 0x4b, 0x91, 0x3a,
	0x81, 0x7a, 0x28, 0xf4, 0x20, 0x58, 0x51, 0x2f, 0x3d, 0x55, 0xbd, 0x78, 0x72, 0x92, 0x8c, 0xd3,
	0xa1, 0x4d, 0xbe, 0x98, 0x99, 0xf8, 0xe7, 0xea, 0x2b, 0xf8, 0x06, 0x3e, 0x85, 0x8f, 0xb0, 0xc7,
	0x85, 0xbd, 0x97, 0xa5, 0xec, 0x13, 0xf4, 0x09, 0x96, 0x4c, 0x9a, 0x6d, 0xba, 0xd0, 0xbd, 0xe5,
	0xf7, 0xef, 0xcb, 0xef, 0xcb, 0x17, 0x84, 0x58, 0xae, 0x97, 0x34, 0xcd, 0x40, 0x03, 0x7e, 0xcc,
	0x7f, 0xb1, 0x38, 0x5d, 0x73, 0x9a, 0xb2, 0x70, 0xc5, 0x04, 0xf7, 0x7a, 0x02, 0x40, 0xac, 0xb9,
	0xcf, 0x52, 0xe9, 0xb3, 0x24, 0x01, 0xcd, 0xb4, 0x84, 0x44, 0x95, 0x76, 0xaf, 0xbb, 0x57, 0x0d,
	0x0a, 0xf2, 0x6f, 0x3e, 0x8f, 0x53, 0xfd, 0x7b, 0x2f, 0xbe, 0x14, 0x52, 0x2f, 0xf3, 0x80, 0x86,
	0x10, 0xfb, 0x02, 0x04, 0x1c, 0x5c, 0x05, 0x32, 0xc0, 0x3c, 0x95, 0x76, 0x22, 0x91, 0xf3, 0x36,
	0xe3, 0x11, 0x4f, 0xb4, 0x64, 0x6b, 0x85, 0x3d, 0x74, 0x3f, 0x57, 0x3c, 0x4b, 0x58, 0xcc, 0x5d,
	0xeb, 0x99, 0x35, 0x7c, 0xb0, 0xb8, 0xc1, 0x85, 0x96, 0x32, 0xa5, 0x7e, 0x42, 0x16, 0xb9, 0x76,
	0xa9, 0x55, 0x18, 0x13, 0xf4, 0x50, 0xf1, 0x10, 0x92, 0xe8, 0x3d, 0x0b, 0x35, 0x64, 0x6e, 0xc3,
	0xe8, 0x47, 0x1c, 0xf9, 0x67, 0xa1, 0xe6, 0x67, 0xc5, 0x33, 0x3c, 0x44, 0xb6, 0x8c, 0xca, 0xf1,
	0x33, 0x77, 0xb7, 0x19, 0x74, 0xa2, 0x60, 0x4a, 0x64, 0x34, 0x4a, 0x57, 0x23, 0x88, 0xa5, 0x36,
	0xdb, 0x90, 0x85, 0x2d, 0x23, 0x3c, 0xa9, 0xd5, 0x31, 0xaf, 0x9c, 0x75, 0x77, 0x9b, 0xc1, 0xd3,
	0xc2, 0x5f, 0xf1, 0xf5, 0xc8, 0xa1, 0xeb, 0xa4, 0xd6, 0xb5, 0x71, 0x1c, 0xac, 0xf8, 0xa3, 0x60,
	0x45, 0x92, 0x4f, 0x08, 0xd7, 0xbe, 0xc7, 0x82, 0x7f, 0xcf, 0xb9, 0xd2, 0xf8, 0x35, 0x72, 0xc2,
	0x03, 0x6b, 0xaa, 0x3b, 0xe3, 0x1e, 0xbd, 0x75, 0x36, 0x5a, 0x4f, 0xd6, 0x03, 0xe3, 0xff, 0x16,
	0x72, 0xde, 0xe4, 0x7a, 0xf9, 0x91, 0x67, 0x3f, 0x64, 0xc8, 0xf1, 0x07, 0xd4, 0x9e, 0x83, 0x80,
	0x5c, 0xe3, 0x27, 0xb4, 0x3c, 0x26, 0xad, 0xce, 0x44, 0xdf, 0x15, 0x95, 0xbc, 0x13, 0x3c, 0x79,
	0xf4, 0xe7, 0xe2, 0xea, 0xaf, 0x7d, 0xef, 0x45, 0xcb, 0x2f, 0xfe, 0x1f, 0xfc, 0x15, 0xb5, 0xe6,
	0x20, 0x64, 0x82, 0x9f, 0xdf, 0x59, 0xa6, 0x5c, 0xe3, 0xe4, 0x50, 0xcf, 0x0c, 0xed, 0x90, 0x72,
	0xe8, 0xb4, 0x5e, 0x7d, 0x86, 0xcf, 0xb6, 0x7d, 0xeb, 0x7c, 0xdb, 0xb7, 0x2e, 0xb7, 0x7d, 0xeb,
	0x4b, 0xb3, 0x30, 0x04, 0x6d, 0x93, 0x7f, 0x75, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x04, 0x55,
	0x98, 0xc4, 0x02, 0x00, 0x00,
}
