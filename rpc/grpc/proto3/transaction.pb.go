// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rpc/grpc/proto3/transaction.proto

package proto3

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import github_com_gallactic_gallactic_txs "github.com/gallactic/gallactic/txs"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type TransactRequest struct {
	TxEnvelope           *github_com_gallactic_gallactic_txs.Envelope `protobuf:"bytes,1,opt,name=txEnvelope,customtype=github.com/gallactic/gallactic/txs.Envelope" json:"txEnvelope,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *TransactRequest) Reset()         { *m = TransactRequest{} }
func (m *TransactRequest) String() string { return proto.CompactTextString(m) }
func (*TransactRequest) ProtoMessage()    {}
func (*TransactRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_260589f6c138e39f, []int{0}
}
func (m *TransactRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransactRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransactRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TransactRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactRequest.Merge(dst, src)
}
func (m *TransactRequest) XXX_Size() int {
	return m.Size()
}
func (m *TransactRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransactRequest proto.InternalMessageInfo

func (*TransactRequest) XXX_MessageName() string {
	return "proto3.TransactRequest"
}

type ReceiptResponse struct {
	TxReceipt            *github_com_gallactic_gallactic_txs.Receipt `protobuf:"bytes,1,opt,name=TxReceipt,proto3,customtype=github.com/gallactic/gallactic/txs.Receipt" json:"TxReceipt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *ReceiptResponse) Reset()         { *m = ReceiptResponse{} }
func (m *ReceiptResponse) String() string { return proto.CompactTextString(m) }
func (*ReceiptResponse) ProtoMessage()    {}
func (*ReceiptResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_260589f6c138e39f, []int{1}
}
func (m *ReceiptResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReceiptResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReceiptResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ReceiptResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiptResponse.Merge(dst, src)
}
func (m *ReceiptResponse) XXX_Size() int {
	return m.Size()
}
func (m *ReceiptResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiptResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiptResponse proto.InternalMessageInfo

func (*ReceiptResponse) XXX_MessageName() string {
	return "proto3.ReceiptResponse"
}

type UnconfirmedTxsRequest struct {
	MaxTxs               int32    `protobuf:"varint,1,opt,name=maxTxs,proto3" json:"maxTxs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnconfirmedTxsRequest) Reset()         { *m = UnconfirmedTxsRequest{} }
func (m *UnconfirmedTxsRequest) String() string { return proto.CompactTextString(m) }
func (*UnconfirmedTxsRequest) ProtoMessage()    {}
func (*UnconfirmedTxsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_260589f6c138e39f, []int{2}
}
func (m *UnconfirmedTxsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UnconfirmedTxsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UnconfirmedTxsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *UnconfirmedTxsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnconfirmedTxsRequest.Merge(dst, src)
}
func (m *UnconfirmedTxsRequest) XXX_Size() int {
	return m.Size()
}
func (m *UnconfirmedTxsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnconfirmedTxsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnconfirmedTxsRequest proto.InternalMessageInfo

func (m *UnconfirmedTxsRequest) GetMaxTxs() int32 {
	if m != nil {
		return m.MaxTxs
	}
	return 0
}

func (*UnconfirmedTxsRequest) XXX_MessageName() string {
	return "proto3.UnconfirmedTxsRequest"
}

type UnconfirmTxsResponse struct {
	Count                int32                                         `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
	TxEnvelopes          []github_com_gallactic_gallactic_txs.Envelope `protobuf:"bytes,2,rep,name=txEnvelopes,customtype=github.com/gallactic/gallactic/txs.Envelope" json:"txEnvelopes"`
	XXX_NoUnkeyedLiteral struct{}                                      `json:"-"`
	XXX_unrecognized     []byte                                        `json:"-"`
	XXX_sizecache        int32                                         `json:"-"`
}

func (m *UnconfirmTxsResponse) Reset()         { *m = UnconfirmTxsResponse{} }
func (m *UnconfirmTxsResponse) String() string { return proto.CompactTextString(m) }
func (*UnconfirmTxsResponse) ProtoMessage()    {}
func (*UnconfirmTxsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_260589f6c138e39f, []int{3}
}
func (m *UnconfirmTxsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UnconfirmTxsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UnconfirmTxsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *UnconfirmTxsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnconfirmTxsResponse.Merge(dst, src)
}
func (m *UnconfirmTxsResponse) XXX_Size() int {
	return m.Size()
}
func (m *UnconfirmTxsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UnconfirmTxsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UnconfirmTxsResponse proto.InternalMessageInfo

func (m *UnconfirmTxsResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (*UnconfirmTxsResponse) XXX_MessageName() string {
	return "proto3.UnconfirmTxsResponse"
}
func init() {
	proto.RegisterType((*TransactRequest)(nil), "proto3.TransactRequest")
	golang_proto.RegisterType((*TransactRequest)(nil), "proto3.TransactRequest")
	proto.RegisterType((*ReceiptResponse)(nil), "proto3.ReceiptResponse")
	golang_proto.RegisterType((*ReceiptResponse)(nil), "proto3.ReceiptResponse")
	proto.RegisterType((*UnconfirmedTxsRequest)(nil), "proto3.UnconfirmedTxsRequest")
	golang_proto.RegisterType((*UnconfirmedTxsRequest)(nil), "proto3.UnconfirmedTxsRequest")
	proto.RegisterType((*UnconfirmTxsResponse)(nil), "proto3.UnconfirmTxsResponse")
	golang_proto.RegisterType((*UnconfirmTxsResponse)(nil), "proto3.UnconfirmTxsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TransactionClient is the client API for Transaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransactionClient interface {
	BroadcastTx(ctx context.Context, in *TransactRequest, opts ...grpc.CallOption) (*ReceiptResponse, error)
	GetUnconfirmedTxs(ctx context.Context, in *UnconfirmedTxsRequest, opts ...grpc.CallOption) (*UnconfirmTxsResponse, error)
}

type transactionClient struct {
	cc *grpc.ClientConn
}

func NewTransactionClient(cc *grpc.ClientConn) TransactionClient {
	return &transactionClient{cc}
}

func (c *transactionClient) BroadcastTx(ctx context.Context, in *TransactRequest, opts ...grpc.CallOption) (*ReceiptResponse, error) {
	out := new(ReceiptResponse)
	err := c.cc.Invoke(ctx, "/proto3.Transaction/BroadcastTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionClient) GetUnconfirmedTxs(ctx context.Context, in *UnconfirmedTxsRequest, opts ...grpc.CallOption) (*UnconfirmTxsResponse, error) {
	out := new(UnconfirmTxsResponse)
	err := c.cc.Invoke(ctx, "/proto3.Transaction/GetUnconfirmedTxs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServer is the server API for Transaction service.
type TransactionServer interface {
	BroadcastTx(context.Context, *TransactRequest) (*ReceiptResponse, error)
	GetUnconfirmedTxs(context.Context, *UnconfirmedTxsRequest) (*UnconfirmTxsResponse, error)
}

func RegisterTransactionServer(s *grpc.Server, srv TransactionServer) {
	s.RegisterService(&_Transaction_serviceDesc, srv)
}

func _Transaction_BroadcastTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).BroadcastTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto3.Transaction/BroadcastTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).BroadcastTx(ctx, req.(*TransactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transaction_GetUnconfirmedTxs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnconfirmedTxsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).GetUnconfirmedTxs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto3.Transaction/GetUnconfirmedTxs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).GetUnconfirmedTxs(ctx, req.(*UnconfirmedTxsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Transaction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto3.Transaction",
	HandlerType: (*TransactionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BroadcastTx",
			Handler:    _Transaction_BroadcastTx_Handler,
		},
		{
			MethodName: "GetUnconfirmedTxs",
			Handler:    _Transaction_GetUnconfirmedTxs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/grpc/proto3/transaction.proto",
}

func (m *TransactRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransactRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.TxEnvelope != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTransaction(dAtA, i, uint64(m.TxEnvelope.Size()))
		n1, err := m.TxEnvelope.MarshalTo(dAtA[i:])
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

func (m *ReceiptResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReceiptResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.TxReceipt != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTransaction(dAtA, i, uint64(m.TxReceipt.Size()))
		n2, err := m.TxReceipt.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *UnconfirmedTxsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnconfirmedTxsRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.MaxTxs != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTransaction(dAtA, i, uint64(m.MaxTxs))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *UnconfirmTxsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnconfirmTxsResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Count != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTransaction(dAtA, i, uint64(m.Count))
	}
	if len(m.TxEnvelopes) > 0 {
		for _, msg := range m.TxEnvelopes {
			dAtA[i] = 0x12
			i++
			i = encodeVarintTransaction(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintTransaction(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TransactRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TxEnvelope != nil {
		l = m.TxEnvelope.Size()
		n += 1 + l + sovTransaction(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ReceiptResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TxReceipt != nil {
		l = m.TxReceipt.Size()
		n += 1 + l + sovTransaction(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *UnconfirmedTxsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxTxs != 0 {
		n += 1 + sovTransaction(uint64(m.MaxTxs))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *UnconfirmTxsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Count != 0 {
		n += 1 + sovTransaction(uint64(m.Count))
	}
	if len(m.TxEnvelopes) > 0 {
		for _, e := range m.TxEnvelopes {
			l = e.Size()
			n += 1 + l + sovTransaction(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTransaction(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTransaction(x uint64) (n int) {
	return sovTransaction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TransactRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransaction
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
			return fmt.Errorf("proto: TransactRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransactRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxEnvelope", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TxEnvelope == nil {
				m.TxEnvelope = &github_com_gallactic_gallactic_txs.Envelope{}
			}
			if err := m.TxEnvelope.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTransaction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTransaction
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
func (m *ReceiptResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransaction
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
			return fmt.Errorf("proto: ReceiptResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReceiptResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxReceipt", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_gallactic_gallactic_txs.Receipt
			m.TxReceipt = &v
			if err := m.TxReceipt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTransaction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTransaction
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
func (m *UnconfirmedTxsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransaction
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
			return fmt.Errorf("proto: UnconfirmedTxsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnconfirmedTxsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxTxs", wireType)
			}
			m.MaxTxs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxTxs |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTransaction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTransaction
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
func (m *UnconfirmTxsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransaction
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
			return fmt.Errorf("proto: UnconfirmTxsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnconfirmTxsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxEnvelopes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxEnvelopes = append(m.TxEnvelopes, github_com_gallactic_gallactic_txs.Envelope{})
			if err := m.TxEnvelopes[len(m.TxEnvelopes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTransaction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTransaction
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
func skipTransaction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTransaction
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
					return 0, ErrIntOverflowTransaction
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
					return 0, ErrIntOverflowTransaction
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
				return 0, ErrInvalidLengthTransaction
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTransaction
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
				next, err := skipTransaction(dAtA[start:])
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
	ErrInvalidLengthTransaction = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTransaction   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("rpc/grpc/proto3/transaction.proto", fileDescriptor_transaction_260589f6c138e39f)
}
func init() {
	golang_proto.RegisterFile("rpc/grpc/proto3/transaction.proto", fileDescriptor_transaction_260589f6c138e39f)
}

var fileDescriptor_transaction_260589f6c138e39f = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcb, 0x6a, 0xdb, 0x40,
	0x14, 0xf5, 0xb8, 0xd8, 0xd0, 0x51, 0xc1, 0x74, 0x70, 0x5b, 0x21, 0x5c, 0xd9, 0xd5, 0xca, 0xb4,
	0x54, 0x03, 0xf6, 0x07, 0x14, 0x54, 0x42, 0x36, 0x59, 0x09, 0x65, 0x1d, 0xc6, 0xf2, 0x58, 0x16,
	0xc8, 0x33, 0x8a, 0x66, 0x14, 0x04, 0x26, 0x9b, 0xfc, 0x42, 0x7e, 0x20, 0xf9, 0x93, 0x90, 0x95,
	0x97, 0x81, 0xec, 0xbc, 0x30, 0xc1, 0xce, 0x87, 0x04, 0xeb, 0x61, 0xcb, 0xce, 0x83, 0x64, 0x23,
	0xe6, 0xce, 0xbd, 0xe7, 0x9c, 0x7b, 0x8e, 0x06, 0xfe, 0x8a, 0x42, 0x17, 0x7b, 0xeb, 0x4f, 0x18,
	0x71, 0xc9, 0xfb, 0x58, 0x46, 0x84, 0x09, 0xe2, 0x4a, 0x9f, 0x33, 0x33, 0xbd, 0x42, 0xf5, 0xac,
	0xa3, 0xfd, 0xf5, 0x7c, 0x39, 0x8e, 0x07, 0xa6, 0xcb, 0x27, 0xd8, 0xe3, 0x1e, 0xcf, 0x10, 0x83,
	0x78, 0x94, 0x56, 0x69, 0x91, 0x9e, 0x32, 0x98, 0xd6, 0xf2, 0x38, 0xf7, 0x02, 0x8a, 0x49, 0xe8,
	0x63, 0xc2, 0x18, 0x97, 0x64, 0xcd, 0x29, 0xb2, 0xae, 0x31, 0x85, 0x0d, 0x27, 0x57, 0xb2, 0xe9,
	0x69, 0x4c, 0x85, 0x44, 0x63, 0x08, 0x65, 0x72, 0xc0, 0xce, 0x68, 0xc0, 0x43, 0xaa, 0x82, 0x0e,
	0xe8, 0x2a, 0xbd, 0x1f, 0xd9, 0x78, 0xdf, 0xdc, 0x1b, 0xb6, 0xf0, 0x7c, 0xd1, 0xfe, 0x53, 0x5e,
	0x88, 0x04, 0xc1, 0x7a, 0x6b, 0xb7, 0x74, 0x92, 0x89, 0x30, 0x0b, 0x3e, 0xbb, 0xc4, 0x6d, 0x9c,
	0xc0, 0x86, 0x4d, 0x5d, 0xea, 0x87, 0xd2, 0xa6, 0x22, 0xe4, 0x4c, 0x50, 0x74, 0x04, 0x3f, 0x3b,
	0x49, 0x7e, 0x99, 0x6a, 0x7f, 0xb1, 0xcc, 0xf9, 0xa2, 0xfd, 0xfb, 0x1d, 0x12, 0x05, 0xd5, 0x96,
	0xc0, 0xc0, 0xf0, 0xdb, 0x31, 0x73, 0x39, 0x1b, 0xf9, 0xd1, 0x84, 0x0e, 0x9d, 0x44, 0x14, 0x1e,
	0xbf, 0xc3, 0xfa, 0x84, 0x24, 0x4e, 0x22, 0x52, 0x8d, 0x9a, 0x9d, 0x57, 0xc6, 0x35, 0x80, 0xcd,
	0x0d, 0x22, 0x9d, 0xcf, 0xf7, 0x6a, 0xc2, 0xda, 0x7f, 0x1e, 0x33, 0x99, 0xcf, 0x67, 0x05, 0x12,
	0x50, 0xd9, 0xda, 0x11, 0x6a, 0xb5, 0xf3, 0xa9, 0xab, 0xf4, 0x5a, 0x45, 0x56, 0x2f, 0x11, 0x59,
	0xfd, 0xd9, 0xa2, 0x5d, 0xf9, 0x68, 0x68, 0x65, 0x95, 0xde, 0x2d, 0x80, 0x8a, 0xb3, 0x7d, 0x1d,
	0xe8, 0x1f, 0x54, 0xac, 0x88, 0x93, 0xa1, 0x4b, 0x84, 0x74, 0x12, 0xf4, 0xda, 0xaf, 0xd2, 0x36,
	0x8d, 0xfd, 0xcc, 0x23, 0xf8, 0xf5, 0x90, 0xca, 0xdd, 0xa0, 0xd0, 0xcf, 0x67, 0x2e, 0xca, 0x01,
	0x6a, 0x6f, 0x9a, 0x34, 0x3a, 0x17, 0xf7, 0x8f, 0x97, 0x55, 0x0d, 0xa9, 0x78, 0x17, 0x8d, 0xa7,
	0x59, 0xce, 0xe7, 0x96, 0x3a, 0x5b, 0xea, 0xe0, 0x6e, 0xa9, 0x83, 0x87, 0xa5, 0x0e, 0xae, 0x56,
	0x7a, 0xe5, 0x66, 0xa5, 0x83, 0xd9, 0x4a, 0x07, 0x83, 0xfc, 0x99, 0x3f, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x7d, 0xc4, 0x81, 0xfa, 0x12, 0x03, 0x00, 0x00,
}
