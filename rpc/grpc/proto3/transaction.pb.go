// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rpc/grpc/proto3/transaction.proto

package proto3

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/golang/protobuf/ptypes/struct"

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
	return fileDescriptor_transaction_00d342c9de1b70f5, []int{0}
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
	return fileDescriptor_transaction_00d342c9de1b70f5, []int{1}
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
	return fileDescriptor_transaction_00d342c9de1b70f5, []int{2}
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
	return fileDescriptor_transaction_00d342c9de1b70f5, []int{3}
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
	proto.RegisterFile("rpc/grpc/proto3/transaction.proto", fileDescriptor_transaction_00d342c9de1b70f5)
}
func init() {
	golang_proto.RegisterFile("rpc/grpc/proto3/transaction.proto", fileDescriptor_transaction_00d342c9de1b70f5)
}

var fileDescriptor_transaction_00d342c9de1b70f5 = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0xed, 0x7c, 0x1f, 0x2d, 0x7c, 0x93, 0x0f, 0x8a, 0xa1, 0x6a, 0x09, 0x35, 0xad, 0x59, 0x15,
	0xc5, 0x09, 0xb4, 0x0f, 0x20, 0x44, 0xc4, 0x4d, 0x17, 0x12, 0xa2, 0x5b, 0x49, 0xa7, 0xd3, 0x34,
	0x90, 0x66, 0xd2, 0xcc, 0x8d, 0x04, 0x7c, 0x18, 0xfb, 0x28, 0x2e, 0xbb, 0x74, 0xdd, 0x45, 0x91,
	0xf6, 0x45, 0xa4, 0xf9, 0x69, 0x62, 0x55, 0xd0, 0x4d, 0x98, 0x3b, 0xe7, 0xdc, 0x73, 0xee, 0x9c,
	0x5c, 0x7c, 0x1a, 0x06, 0x54, 0x77, 0xb6, 0x9f, 0x20, 0xe4, 0xc0, 0xfb, 0x3a, 0x84, 0xb6, 0x2f,
	0x6c, 0x0a, 0x2e, 0xf7, 0x49, 0x72, 0x25, 0xd7, 0x52, 0x44, 0xb9, 0x70, 0x5c, 0x98, 0x44, 0x43,
	0x42, 0xf9, 0x54, 0x77, 0xb8, 0xc3, 0xd3, 0x8e, 0x61, 0x34, 0x4e, 0xaa, 0xa4, 0x48, 0x4e, 0x69,
	0x9b, 0xd2, 0x72, 0x38, 0x77, 0x3c, 0x56, 0xb0, 0x04, 0x84, 0x11, 0x85, 0x14, 0xd5, 0x9e, 0x70,
	0xdd, 0xca, 0x9c, 0x4c, 0x36, 0x8b, 0x98, 0x00, 0x79, 0x82, 0x31, 0xc4, 0xd7, 0xfe, 0x23, 0xf3,
	0x78, 0xc0, 0x9a, 0xa8, 0x83, 0xba, 0x52, 0xef, 0x38, 0xa5, 0xf7, 0xc9, 0x1e, 0xd9, 0xd0, 0x97,
	0xab, 0xf6, 0x79, 0x79, 0x20, 0xdb, 0xf3, 0xb6, 0x53, 0xd3, 0xd2, 0x09, 0x62, 0x41, 0x72, 0x3d,
	0xb3, 0xa4, 0xad, 0x3d, 0xe0, 0xba, 0xc9, 0x28, 0x73, 0x03, 0x30, 0x99, 0x08, 0xb8, 0x2f, 0x98,
	0x3c, 0xc0, 0xff, 0xac, 0x38, 0xbb, 0x4c, 0xbc, 0xff, 0x1b, 0x64, 0xb9, 0x6a, 0x9f, 0xfd, 0xc0,
	0x22, 0x97, 0x2a, 0x04, 0x34, 0x1d, 0x1f, 0xde, 0xf9, 0x94, 0xfb, 0x63, 0x37, 0x9c, 0xb2, 0x91,
	0x15, 0x8b, 0xfc, 0x8d, 0x47, 0xb8, 0x36, 0xb5, 0x63, 0x2b, 0x16, 0x89, 0x47, 0xd5, 0xcc, 0x2a,
	0xed, 0x19, 0xe1, 0xc6, 0xae, 0x23, 0xe1, 0x67, 0x73, 0x35, 0x70, 0xf5, 0x8a, 0x47, 0x3e, 0x64,
	0xfc, 0xb4, 0x90, 0x67, 0x58, 0x2a, 0x9e, 0x23, 0x9a, 0x7f, 0x3a, 0x7f, 0xbb, 0x52, 0x4f, 0x21,
	0x69, 0xe2, 0x24, 0x4f, 0x9c, 0x0c, 0x5c, 0x01, 0xf7, 0xb6, 0x17, 0x31, 0xa3, 0xbf, 0x58, 0xb5,
	0x2b, 0xbf, 0x8d, 0xac, 0xec, 0xd1, 0x9b, 0x23, 0x2c, 0x59, 0xc5, 0x6e, 0xc8, 0x97, 0x58, 0x32,
	0x42, 0x6e, 0x8f, 0xa8, 0x2d, 0xc0, 0x8a, 0xe5, 0xef, 0x7e, 0x94, 0xb2, 0x03, 0xf6, 0x13, 0xbf,
	0xc5, 0x07, 0x37, 0x0c, 0x3e, 0xc6, 0x24, 0x9f, 0xe4, 0xec, 0x2f, 0xe3, 0x53, 0x5a, 0x9f, 0xe0,
	0x52, 0x56, 0x46, 0x73, 0xb1, 0x56, 0xd1, 0xeb, 0x5a, 0x45, 0x6f, 0x6b, 0x15, 0xcd, 0x37, 0x6a,
	0xe5, 0x65, 0xa3, 0xa2, 0xc5, 0x46, 0x45, 0xc3, 0x6c, 0x85, 0xdf, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x94, 0x7b, 0x8b, 0x94, 0xee, 0x02, 0x00, 0x00,
}