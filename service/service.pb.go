// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: slinky/oracle/v1/service.proto

package service

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryPricesRequest defines the request type for the the Prices method.
type QueryPricesRequest struct {
	// optional tickers to fetch.
	Tickers []string `protobuf:"bytes,1,rep,name=tickers,proto3" json:"tickers,omitempty"`
	// optional provider to query against.
	Provider string `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
}

func (m *QueryPricesRequest) Reset()         { *m = QueryPricesRequest{} }
func (m *QueryPricesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryPricesRequest) ProtoMessage()    {}
func (*QueryPricesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cb616c17b07d393, []int{0}
}
func (m *QueryPricesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryPricesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryPricesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryPricesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryPricesRequest.Merge(m, src)
}
func (m *QueryPricesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryPricesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryPricesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryPricesRequest proto.InternalMessageInfo

func (m *QueryPricesRequest) GetTickers() []string {
	if m != nil {
		return m.Tickers
	}
	return nil
}

func (m *QueryPricesRequest) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

// QueryPricesResponse defines the response type for the Prices method.
type QueryPricesResponse struct {
	// prices defines the list of prices.
	Prices    map[string]string `protobuf:"bytes,1,rep,name=prices,proto3" json:"prices" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Timestamp time.Time         `protobuf:"bytes,2,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
}

func (m *QueryPricesResponse) Reset()         { *m = QueryPricesResponse{} }
func (m *QueryPricesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryPricesResponse) ProtoMessage()    {}
func (*QueryPricesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cb616c17b07d393, []int{1}
}
func (m *QueryPricesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryPricesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryPricesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryPricesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryPricesResponse.Merge(m, src)
}
func (m *QueryPricesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryPricesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryPricesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryPricesResponse proto.InternalMessageInfo

func (m *QueryPricesResponse) GetPrices() map[string]string {
	if m != nil {
		return m.Prices
	}
	return nil
}

func (m *QueryPricesResponse) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*QueryPricesRequest)(nil), "slinky.oracle.v1.QueryPricesRequest")
	proto.RegisterType((*QueryPricesResponse)(nil), "slinky.oracle.v1.QueryPricesResponse")
	proto.RegisterMapType((map[string]string)(nil), "slinky.oracle.v1.QueryPricesResponse.PricesEntry")
}

func init() { proto.RegisterFile("slinky/oracle/v1/service.proto", fileDescriptor_5cb616c17b07d393) }

var fileDescriptor_5cb616c17b07d393 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xbd, 0xce, 0xd3, 0x30,
	0x14, 0x8d, 0x5b, 0x08, 0xad, 0xbb, 0x54, 0xa6, 0x43, 0x88, 0x50, 0x5a, 0x45, 0x20, 0x75, 0xc1,
	0x56, 0xcb, 0x02, 0x88, 0x29, 0x12, 0x0b, 0x4b, 0x21, 0x62, 0x62, 0x41, 0x69, 0x30, 0xc1, 0xca,
	0x8f, 0x83, 0xed, 0x04, 0x65, 0x84, 0x27, 0xa8, 0xc4, 0x4b, 0x75, 0xac, 0xc4, 0xc2, 0x04, 0xa8,
	0x85, 0xf7, 0x40, 0xb5, 0x93, 0xf2, 0xfd, 0x49, 0xdf, 0xb7, 0xdd, 0xe3, 0x73, 0xef, 0xf1, 0x3d,
	0x47, 0x17, 0x7a, 0x32, 0x63, 0x45, 0xda, 0x10, 0x2e, 0xa2, 0x38, 0xa3, 0xa4, 0x5e, 0x10, 0x49,
	0x45, 0xcd, 0x62, 0x8a, 0x4b, 0xc1, 0x15, 0x47, 0x63, 0xc3, 0x63, 0xc3, 0xe3, 0x7a, 0xe1, 0x4e,
	0x12, 0x9e, 0x70, 0x4d, 0x92, 0x63, 0x65, 0xfa, 0xdc, 0xfb, 0x09, 0xe7, 0x49, 0x46, 0x49, 0x54,
	0x32, 0x12, 0x15, 0x05, 0x57, 0x91, 0x62, 0xbc, 0x90, 0x2d, 0x3b, 0x6d, 0x59, 0x8d, 0xd6, 0xd5,
	0x07, 0xa2, 0x58, 0x4e, 0xa5, 0x8a, 0xf2, 0xb2, 0x6d, 0xb8, 0x17, 0x73, 0x99, 0x73, 0xf9, 0xce,
	0xe8, 0x1a, 0x60, 0x28, 0xff, 0x25, 0x44, 0xaf, 0x2b, 0x2a, 0x9a, 0x57, 0x82, 0xc5, 0x54, 0x86,
	0xf4, 0x53, 0x45, 0xa5, 0x42, 0x0e, 0xbc, 0xa3, 0x58, 0x9c, 0x52, 0x21, 0x1d, 0x30, 0xeb, 0xcf,
	0x87, 0x61, 0x07, 0x91, 0x0b, 0x07, 0xa5, 0xe0, 0x35, 0x7b, 0x4f, 0x85, 0xd3, 0x9b, 0x81, 0xf9,
	0x30, 0x3c, 0x61, 0xff, 0x2f, 0x80, 0x77, 0xcf, 0x89, 0xc9, 0x92, 0x17, 0x92, 0xa2, 0x15, 0xb4,
	0x4b, 0xfd, 0xa2, 0xc5, 0x46, 0xcb, 0x05, 0xbe, 0x68, 0x1b, 0x5f, 0x31, 0x86, 0x0d, 0x7c, 0x51,
	0x28, 0xd1, 0x04, 0xb7, 0xb6, 0x3f, 0xa7, 0x56, 0xd8, 0xca, 0xa0, 0x00, 0x0e, 0x4f, 0x16, 0xf5,
	0x16, 0xa3, 0xa5, 0x8b, 0x4d, 0x08, 0xb8, 0x0b, 0x01, 0xbf, 0xe9, 0x3a, 0x82, 0xc1, 0x71, 0x78,
	0xf3, 0x6b, 0x0a, 0xc2, 0xff, 0x63, 0xee, 0x53, 0x38, 0x3a, 0xf3, 0x01, 0x1a, 0xc3, 0x7e, 0x4a,
	0x1b, 0x07, 0x68, 0x4b, 0xc7, 0x12, 0x4d, 0xe0, 0xed, 0x3a, 0xca, 0x2a, 0xda, 0xda, 0x34, 0xe0,
	0x59, 0xef, 0x09, 0x58, 0x7e, 0x01, 0xd0, 0x5e, 0xe9, 0xd5, 0xd1, 0x67, 0x68, 0x1b, 0x15, 0xf4,
	0xe0, 0x1a, 0x53, 0x3a, 0x58, 0xf7, 0xe1, 0x8d, 0xac, 0xfb, 0xb3, 0xaf, 0xdf, 0xff, 0x7c, 0xeb,
	0xb9, 0xc8, 0x21, 0x97, 0x0e, 0xc8, 0x44, 0x10, 0x3c, 0xdf, 0xee, 0x3d, 0xb0, 0xdb, 0x7b, 0xe0,
	0xf7, 0xde, 0x03, 0x9b, 0x83, 0x67, 0xed, 0x0e, 0x9e, 0xf5, 0xe3, 0xe0, 0x59, 0x6f, 0xfd, 0x84,
	0xa9, 0x8f, 0xd5, 0x1a, 0xc7, 0x3c, 0x27, 0x32, 0x65, 0xe5, 0xa3, 0x9c, 0xd6, 0x9d, 0x4c, 0x7b,
	0x7d, 0x6b, 0x5b, 0xa7, 0xf4, 0xf8, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x53, 0xd6, 0x91, 0xbd,
	0xa0, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OracleClient is the client API for Oracle service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OracleClient interface {
	// Prices defines a method for fetching the latest prices.
	Prices(ctx context.Context, in *QueryPricesRequest, opts ...grpc.CallOption) (*QueryPricesResponse, error)
}

type oracleClient struct {
	cc grpc1.ClientConn
}

func NewOracleClient(cc grpc1.ClientConn) OracleClient {
	return &oracleClient{cc}
}

func (c *oracleClient) Prices(ctx context.Context, in *QueryPricesRequest, opts ...grpc.CallOption) (*QueryPricesResponse, error) {
	out := new(QueryPricesResponse)
	err := c.cc.Invoke(ctx, "/slinky.oracle.v1.Oracle/Prices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OracleServer is the server API for Oracle service.
type OracleServer interface {
	// Prices defines a method for fetching the latest prices.
	Prices(context.Context, *QueryPricesRequest) (*QueryPricesResponse, error)
}

// UnimplementedOracleServer can be embedded to have forward compatible implementations.
type UnimplementedOracleServer struct {
}

func (*UnimplementedOracleServer) Prices(ctx context.Context, req *QueryPricesRequest) (*QueryPricesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prices not implemented")
}

func RegisterOracleServer(s grpc1.Server, srv OracleServer) {
	s.RegisterService(&_Oracle_serviceDesc, srv)
}

func _Oracle_Prices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPricesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OracleServer).Prices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slinky.oracle.v1.Oracle/Prices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OracleServer).Prices(ctx, req.(*QueryPricesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Oracle_serviceDesc = grpc.ServiceDesc{
	ServiceName: "slinky.oracle.v1.Oracle",
	HandlerType: (*OracleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Prices",
			Handler:    _Oracle_Prices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "slinky/oracle/v1/service.proto",
}

func (m *QueryPricesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryPricesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryPricesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintService(dAtA, i, uint64(len(m.Provider)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Tickers) > 0 {
		for iNdEx := len(m.Tickers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Tickers[iNdEx])
			copy(dAtA[i:], m.Tickers[iNdEx])
			i = encodeVarintService(dAtA, i, uint64(len(m.Tickers[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryPricesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryPricesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryPricesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintService(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	if len(m.Prices) > 0 {
		for k := range m.Prices {
			v := m.Prices[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintService(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintService(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintService(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintService(dAtA []byte, offset int, v uint64) int {
	offset -= sovService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryPricesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Tickers) > 0 {
		for _, s := range m.Tickers {
			l = len(s)
			n += 1 + l + sovService(uint64(l))
		}
	}
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func (m *QueryPricesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Prices) > 0 {
		for k, v := range m.Prices {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovService(uint64(len(k))) + 1 + len(v) + sovService(uint64(len(v)))
			n += mapEntrySize + 1 + sovService(uint64(mapEntrySize))
		}
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovService(uint64(l))
	return n
}

func sovService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozService(x uint64) (n int) {
	return sovService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryPricesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: QueryPricesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryPricesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tickers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tickers = append(m.Tickers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func (m *QueryPricesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: QueryPricesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryPricesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Prices == nil {
				m.Prices = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowService
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowService
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthService
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthService
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowService
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthService
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthService
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipService(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthService
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Prices[mapkey] = mapvalue
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func skipService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowService
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
					return 0, ErrIntOverflowService
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
					return 0, ErrIntOverflowService
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
				return 0, ErrInvalidLengthService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupService = fmt.Errorf("proto: unexpected end of group")
)
