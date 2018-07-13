// Code generated by protoc-gen-go. DO NOT EDIT.
// source: eos.proto

/*
Package proto is a generated protocol buffer package.

based on https://github.com/Appscrunch/Multy-back/

It is generated from these files:
	eos.proto

It has these top-level messages:
	OkErrResponse
	Empty
	ChainInfo
	BalanceReq
	Accounts
	Account
	Balances
	Asset
	AccountInfo
	Transaction
	PushTransactionResp
	TransactionID
	TransactionInfo
	AccountCreateReq
	Exist
	RAMPrice
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Compression int32

const (
	Compression_NONE Compression = 0
	Compression_ZLIB Compression = 1
)

var Compression_name = map[int32]string{
	0: "NONE",
	1: "ZLIB",
}
var Compression_value = map[string]int32{
	"NONE": 0,
	"ZLIB": 1,
}

func (x Compression) String() string {
	return proto1.EnumName(Compression_name, int32(x))
}
func (Compression) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type OkErrResponse struct {
	Ok    bool   `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *OkErrResponse) Reset()                    { *m = OkErrResponse{} }
func (m *OkErrResponse) String() string            { return proto1.CompactTextString(m) }
func (*OkErrResponse) ProtoMessage()               {}
func (*OkErrResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *OkErrResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *OkErrResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto1.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ChainInfo struct {
	HeadBlockNum             uint32 `protobuf:"varint,1,opt,name=head_block_num,json=headBlockNum" json:"head_block_num,omitempty"`
	HeadBlockId              string `protobuf:"bytes,2,opt,name=head_block_id,json=headBlockId" json:"head_block_id,omitempty"`
	HeadBlockTime            int64  `protobuf:"varint,3,opt,name=head_block_time,json=headBlockTime" json:"head_block_time,omitempty"`
	LastIrreversibleBlockNum uint32 `protobuf:"varint,4,opt,name=last_irreversible_block_num,json=lastIrreversibleBlockNum" json:"last_irreversible_block_num,omitempty"`
}

func (m *ChainInfo) Reset()                    { *m = ChainInfo{} }
func (m *ChainInfo) String() string            { return proto1.CompactTextString(m) }
func (*ChainInfo) ProtoMessage()               {}
func (*ChainInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ChainInfo) GetHeadBlockNum() uint32 {
	if m != nil {
		return m.HeadBlockNum
	}
	return 0
}

func (m *ChainInfo) GetHeadBlockId() string {
	if m != nil {
		return m.HeadBlockId
	}
	return ""
}

func (m *ChainInfo) GetHeadBlockTime() int64 {
	if m != nil {
		return m.HeadBlockTime
	}
	return 0
}

func (m *ChainInfo) GetLastIrreversibleBlockNum() uint32 {
	if m != nil {
		return m.LastIrreversibleBlockNum
	}
	return 0
}

type BalanceReq struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Symbol string `protobuf:"bytes,2,opt,name=symbol" json:"symbol,omitempty"`
	Code   string `protobuf:"bytes,3,opt,name=code" json:"code,omitempty"`
}

func (m *BalanceReq) Reset()                    { *m = BalanceReq{} }
func (m *BalanceReq) String() string            { return proto1.CompactTextString(m) }
func (*BalanceReq) ProtoMessage()               {}
func (*BalanceReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BalanceReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BalanceReq) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *BalanceReq) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type Accounts struct {
	Accounts []*Account `protobuf:"bytes,1,rep,name=accounts" json:"accounts,omitempty"`
}

func (m *Accounts) Reset()                    { *m = Accounts{} }
func (m *Accounts) String() string            { return proto1.CompactTextString(m) }
func (*Accounts) ProtoMessage()               {}
func (*Accounts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Accounts) GetAccounts() []*Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

type Account struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto1.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Account) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Balances struct {
	Account string   `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
	Assets  []*Asset `protobuf:"bytes,2,rep,name=assets" json:"assets,omitempty"`
}

func (m *Balances) Reset()                    { *m = Balances{} }
func (m *Balances) String() string            { return proto1.CompactTextString(m) }
func (*Balances) ProtoMessage()               {}
func (*Balances) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Balances) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *Balances) GetAssets() []*Asset {
	if m != nil {
		return m.Assets
	}
	return nil
}

type Asset struct {
	Ammount   int64  `protobuf:"varint,1,opt,name=ammount" json:"ammount,omitempty"`
	Precision uint32 `protobuf:"varint,2,opt,name=precision" json:"precision,omitempty"`
	Symbol    string `protobuf:"bytes,3,opt,name=symbol" json:"symbol,omitempty"`
}

func (m *Asset) Reset()                    { *m = Asset{} }
func (m *Asset) String() string            { return proto1.CompactTextString(m) }
func (*Asset) ProtoMessage()               {}
func (*Asset) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Asset) GetAmmount() int64 {
	if m != nil {
		return m.Ammount
	}
	return 0
}

func (m *Asset) GetPrecision() uint32 {
	if m != nil {
		return m.Precision
	}
	return 0
}

func (m *Asset) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

type AccountInfo struct {
	Name              string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	RamAvailable      int64  `protobuf:"varint,2,opt,name=ram_available,json=ramAvailable" json:"ram_available,omitempty"`
	CoreLiquidBalance *Asset `protobuf:"bytes,3,opt,name=core_liquid_balance,json=coreLiquidBalance" json:"core_liquid_balance,omitempty"`
}

func (m *AccountInfo) Reset()                    { *m = AccountInfo{} }
func (m *AccountInfo) String() string            { return proto1.CompactTextString(m) }
func (*AccountInfo) ProtoMessage()               {}
func (*AccountInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *AccountInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AccountInfo) GetRamAvailable() int64 {
	if m != nil {
		return m.RamAvailable
	}
	return 0
}

func (m *AccountInfo) GetCoreLiquidBalance() *Asset {
	if m != nil {
		return m.CoreLiquidBalance
	}
	return nil
}

type Transaction struct {
	Signatures            []string    `protobuf:"bytes,1,rep,name=signatures" json:"signatures,omitempty"`
	Compression           Compression `protobuf:"varint,2,opt,name=compression,enum=proto.Compression" json:"compression,omitempty"`
	PackedContextFreeData []byte      `protobuf:"bytes,3,opt,name=packed_context_free_data,json=packedContextFreeData,proto3" json:"packed_context_free_data,omitempty"`
	PackedTrx             []byte      `protobuf:"bytes,4,opt,name=packed_trx,json=packedTrx,proto3" json:"packed_trx,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto1.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Transaction) GetSignatures() []string {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func (m *Transaction) GetCompression() Compression {
	if m != nil {
		return m.Compression
	}
	return Compression_NONE
}

func (m *Transaction) GetPackedContextFreeData() []byte {
	if m != nil {
		return m.PackedContextFreeData
	}
	return nil
}

func (m *Transaction) GetPackedTrx() []byte {
	if m != nil {
		return m.PackedTrx
	}
	return nil
}

type PushTransactionResp struct {
	Id         string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	StatusCode string `protobuf:"bytes,3,opt,name=status_code,json=statusCode" json:"status_code,omitempty"`
}

func (m *PushTransactionResp) Reset()                    { *m = PushTransactionResp{} }
func (m *PushTransactionResp) String() string            { return proto1.CompactTextString(m) }
func (*PushTransactionResp) ProtoMessage()               {}
func (*PushTransactionResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *PushTransactionResp) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PushTransactionResp) GetStatusCode() string {
	if m != nil {
		return m.StatusCode
	}
	return ""
}

type TransactionID struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *TransactionID) Reset()                    { *m = TransactionID{} }
func (m *TransactionID) String() string            { return proto1.CompactTextString(m) }
func (*TransactionID) ProtoMessage()               {}
func (*TransactionID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *TransactionID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type TransactionInfo struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	BlockNum uint32 `protobuf:"varint,2,opt,name=block_num,json=blockNum" json:"block_num,omitempty"`
}

func (m *TransactionInfo) Reset()                    { *m = TransactionInfo{} }
func (m *TransactionInfo) String() string            { return proto1.CompactTextString(m) }
func (*TransactionInfo) ProtoMessage()               {}
func (*TransactionInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *TransactionInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TransactionInfo) GetBlockNum() uint32 {
	if m != nil {
		return m.BlockNum
	}
	return 0
}

type AccountCreateReq struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	ActiveKey string `protobuf:"bytes,2,opt,name=active_key,json=activeKey" json:"active_key,omitempty"`
	OwnerKey  string `protobuf:"bytes,3,opt,name=owner_key,json=ownerKey" json:"owner_key,omitempty"`
	RamCost   uint64 `protobuf:"varint,4,opt,name=ram_cost,json=ramCost" json:"ram_cost,omitempty"`
}

func (m *AccountCreateReq) Reset()                    { *m = AccountCreateReq{} }
func (m *AccountCreateReq) String() string            { return proto1.CompactTextString(m) }
func (*AccountCreateReq) ProtoMessage()               {}
func (*AccountCreateReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *AccountCreateReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AccountCreateReq) GetActiveKey() string {
	if m != nil {
		return m.ActiveKey
	}
	return ""
}

func (m *AccountCreateReq) GetOwnerKey() string {
	if m != nil {
		return m.OwnerKey
	}
	return ""
}

func (m *AccountCreateReq) GetRamCost() uint64 {
	if m != nil {
		return m.RamCost
	}
	return 0
}

type Exist struct {
	Exist bool `protobuf:"varint,1,opt,name=exist" json:"exist,omitempty"`
}

func (m *Exist) Reset()                    { *m = Exist{} }
func (m *Exist) String() string            { return proto1.CompactTextString(m) }
func (*Exist) ProtoMessage()               {}
func (*Exist) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *Exist) GetExist() bool {
	if m != nil {
		return m.Exist
	}
	return false
}

type RAMPrice struct {
	Price float64 `protobuf:"fixed64,1,opt,name=price" json:"price,omitempty"`
}

func (m *RAMPrice) Reset()                    { *m = RAMPrice{} }
func (m *RAMPrice) String() string            { return proto1.CompactTextString(m) }
func (*RAMPrice) ProtoMessage()               {}
func (*RAMPrice) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *RAMPrice) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func init() {
	proto1.RegisterType((*OkErrResponse)(nil), "proto.OkErrResponse")
	proto1.RegisterType((*Empty)(nil), "proto.Empty")
	proto1.RegisterType((*ChainInfo)(nil), "proto.ChainInfo")
	proto1.RegisterType((*BalanceReq)(nil), "proto.BalanceReq")
	proto1.RegisterType((*Accounts)(nil), "proto.Accounts")
	proto1.RegisterType((*Account)(nil), "proto.Account")
	proto1.RegisterType((*Balances)(nil), "proto.Balances")
	proto1.RegisterType((*Asset)(nil), "proto.Asset")
	proto1.RegisterType((*AccountInfo)(nil), "proto.AccountInfo")
	proto1.RegisterType((*Transaction)(nil), "proto.Transaction")
	proto1.RegisterType((*PushTransactionResp)(nil), "proto.PushTransactionResp")
	proto1.RegisterType((*TransactionID)(nil), "proto.TransactionID")
	proto1.RegisterType((*TransactionInfo)(nil), "proto.TransactionInfo")
	proto1.RegisterType((*AccountCreateReq)(nil), "proto.AccountCreateReq")
	proto1.RegisterType((*Exist)(nil), "proto.Exist")
	proto1.RegisterType((*RAMPrice)(nil), "proto.RAMPrice")
	proto1.RegisterEnum("proto.Compression", Compression_name, Compression_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NodeCommunications service

type NodeCommunicationsClient interface {
	// EventGetChainInfo gets current state info
	EventGetChainInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ChainInfo, error)
	// EventGetBalance get balance
	// from standard eosio.token smart-contract by default
	// for all the tokens by default
	EventGetBalance(ctx context.Context, in *BalanceReq, opts ...grpc.CallOption) (*Balances, error)
	// EventGetAccount gets account data such as unused ram or core liquid balance
	EventGetAccount(ctx context.Context, in *Account, opts ...grpc.CallOption) (*AccountInfo, error)
	// EventPushTransaction pushes transaction to chain
	EventPushTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*PushTransactionResp, error)
	// EventGetTransactionInfo get transaction block number
	EventGetTransactionInfo(ctx context.Context, in *TransactionID, opts ...grpc.CallOption) (*TransactionInfo, error)
	// Event account create creates account if cost is right
	EventAccountCreate(ctx context.Context, in *AccountCreateReq, opts ...grpc.CallOption) (*OkErrResponse, error)
	// EventAccountCheck checks if account exists.
	EventAccountCheck(ctx context.Context, in *Account, opts ...grpc.CallOption) (*Exist, error)
	// BalanceChanged streams changed balances
	BalanceChanged(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeCommunications_BalanceChangedClient, error)
	// EventTrackAccount marks account as tracked
	EventTrackAccount(ctx context.Context, in *Account, opts ...grpc.CallOption) (*Empty, error)
	// EventSetTrackedAccount resets tracked accounts
	EventSetTrackedAccounts(ctx context.Context, in *Accounts, opts ...grpc.CallOption) (*Empty, error)
	// EventGetRAMPrice get actual RAM price using rammarket
	EventGetRAMPrice(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RAMPrice, error)
}

type nodeCommunicationsClient struct {
	cc *grpc.ClientConn
}

func NewNodeCommunicationsClient(cc *grpc.ClientConn) NodeCommunicationsClient {
	return &nodeCommunicationsClient{cc}
}

func (c *nodeCommunicationsClient) EventGetChainInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ChainInfo, error) {
	out := new(ChainInfo)
	err := grpc.Invoke(ctx, "/proto.NodeCommunications/EventGetChainInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeCommunicationsClient) EventGetBalance(ctx context.Context, in *BalanceReq, opts ...grpc.CallOption) (*Balances, error) {
	out := new(Balances)
	err := grpc.Invoke(ctx, "/proto.NodeCommunications/EventGetBalance", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeCommunicationsClient) EventGetAccount(ctx context.Context, in *Account, opts ...grpc.CallOption) (*AccountInfo, error) {
	out := new(AccountInfo)
	err := grpc.Invoke(ctx, "/proto.NodeCommunications/EventGetAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeCommunicationsClient) EventPushTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*PushTransactionResp, error) {
	out := new(PushTransactionResp)
	err := grpc.Invoke(ctx, "/proto.NodeCommunications/EventPushTransaction", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeCommunicationsClient) EventGetTransactionInfo(ctx context.Context, in *TransactionID, opts ...grpc.CallOption) (*TransactionInfo, error) {
	out := new(TransactionInfo)
	err := grpc.Invoke(ctx, "/proto.NodeCommunications/EventGetTransactionInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeCommunicationsClient) EventAccountCreate(ctx context.Context, in *AccountCreateReq, opts ...grpc.CallOption) (*OkErrResponse, error) {
	out := new(OkErrResponse)
	err := grpc.Invoke(ctx, "/proto.NodeCommunications/EventAccountCreate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeCommunicationsClient) EventAccountCheck(ctx context.Context, in *Account, opts ...grpc.CallOption) (*Exist, error) {
	out := new(Exist)
	err := grpc.Invoke(ctx, "/proto.NodeCommunications/EventAccountCheck", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeCommunicationsClient) BalanceChanged(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeCommunications_BalanceChangedClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_NodeCommunications_serviceDesc.Streams[0], c.cc, "/proto.NodeCommunications/BalanceChanged", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeCommunicationsBalanceChangedClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeCommunications_BalanceChangedClient interface {
	Recv() (*Balances, error)
	grpc.ClientStream
}

type nodeCommunicationsBalanceChangedClient struct {
	grpc.ClientStream
}

func (x *nodeCommunicationsBalanceChangedClient) Recv() (*Balances, error) {
	m := new(Balances)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeCommunicationsClient) EventTrackAccount(ctx context.Context, in *Account, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/proto.NodeCommunications/EventTrackAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeCommunicationsClient) EventSetTrackedAccounts(ctx context.Context, in *Accounts, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/proto.NodeCommunications/EventSetTrackedAccounts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeCommunicationsClient) EventGetRAMPrice(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RAMPrice, error) {
	out := new(RAMPrice)
	err := grpc.Invoke(ctx, "/proto.NodeCommunications/EventGetRAMPrice", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NodeCommunications service

type NodeCommunicationsServer interface {
	// EventGetChainInfo gets current state info
	EventGetChainInfo(context.Context, *Empty) (*ChainInfo, error)
	// EventGetBalance get balance
	// from standard eosio.token smart-contract by default
	// for all the tokens by default
	EventGetBalance(context.Context, *BalanceReq) (*Balances, error)
	// EventGetAccount gets account data such as unused ram or core liquid balance
	EventGetAccount(context.Context, *Account) (*AccountInfo, error)
	// EventPushTransaction pushes transaction to chain
	EventPushTransaction(context.Context, *Transaction) (*PushTransactionResp, error)
	// EventGetTransactionInfo get transaction block number
	EventGetTransactionInfo(context.Context, *TransactionID) (*TransactionInfo, error)
	// Event account create creates account if cost is right
	EventAccountCreate(context.Context, *AccountCreateReq) (*OkErrResponse, error)
	// EventAccountCheck checks if account exists.
	EventAccountCheck(context.Context, *Account) (*Exist, error)
	// BalanceChanged streams changed balances
	BalanceChanged(*Empty, NodeCommunications_BalanceChangedServer) error
	// EventTrackAccount marks account as tracked
	EventTrackAccount(context.Context, *Account) (*Empty, error)
	// EventSetTrackedAccount resets tracked accounts
	EventSetTrackedAccounts(context.Context, *Accounts) (*Empty, error)
	// EventGetRAMPrice get actual RAM price using rammarket
	EventGetRAMPrice(context.Context, *Empty) (*RAMPrice, error)
}

func RegisterNodeCommunicationsServer(s *grpc.Server, srv NodeCommunicationsServer) {
	s.RegisterService(&_NodeCommunications_serviceDesc, srv)
}

func _NodeCommunications_EventGetChainInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeCommunicationsServer).EventGetChainInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeCommunications/EventGetChainInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeCommunicationsServer).EventGetChainInfo(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeCommunications_EventGetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeCommunicationsServer).EventGetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeCommunications/EventGetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeCommunicationsServer).EventGetBalance(ctx, req.(*BalanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeCommunications_EventGetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Account)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeCommunicationsServer).EventGetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeCommunications/EventGetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeCommunicationsServer).EventGetAccount(ctx, req.(*Account))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeCommunications_EventPushTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeCommunicationsServer).EventPushTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeCommunications/EventPushTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeCommunicationsServer).EventPushTransaction(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeCommunications_EventGetTransactionInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeCommunicationsServer).EventGetTransactionInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeCommunications/EventGetTransactionInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeCommunicationsServer).EventGetTransactionInfo(ctx, req.(*TransactionID))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeCommunications_EventAccountCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeCommunicationsServer).EventAccountCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeCommunications/EventAccountCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeCommunicationsServer).EventAccountCreate(ctx, req.(*AccountCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeCommunications_EventAccountCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Account)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeCommunicationsServer).EventAccountCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeCommunications/EventAccountCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeCommunicationsServer).EventAccountCheck(ctx, req.(*Account))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeCommunications_BalanceChanged_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeCommunicationsServer).BalanceChanged(m, &nodeCommunicationsBalanceChangedServer{stream})
}

type NodeCommunications_BalanceChangedServer interface {
	Send(*Balances) error
	grpc.ServerStream
}

type nodeCommunicationsBalanceChangedServer struct {
	grpc.ServerStream
}

func (x *nodeCommunicationsBalanceChangedServer) Send(m *Balances) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeCommunications_EventTrackAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Account)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeCommunicationsServer).EventTrackAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeCommunications/EventTrackAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeCommunicationsServer).EventTrackAccount(ctx, req.(*Account))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeCommunications_EventSetTrackedAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Accounts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeCommunicationsServer).EventSetTrackedAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeCommunications/EventSetTrackedAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeCommunicationsServer).EventSetTrackedAccounts(ctx, req.(*Accounts))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeCommunications_EventGetRAMPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeCommunicationsServer).EventGetRAMPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeCommunications/EventGetRAMPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeCommunicationsServer).EventGetRAMPrice(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeCommunications_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.NodeCommunications",
	HandlerType: (*NodeCommunicationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EventGetChainInfo",
			Handler:    _NodeCommunications_EventGetChainInfo_Handler,
		},
		{
			MethodName: "EventGetBalance",
			Handler:    _NodeCommunications_EventGetBalance_Handler,
		},
		{
			MethodName: "EventGetAccount",
			Handler:    _NodeCommunications_EventGetAccount_Handler,
		},
		{
			MethodName: "EventPushTransaction",
			Handler:    _NodeCommunications_EventPushTransaction_Handler,
		},
		{
			MethodName: "EventGetTransactionInfo",
			Handler:    _NodeCommunications_EventGetTransactionInfo_Handler,
		},
		{
			MethodName: "EventAccountCreate",
			Handler:    _NodeCommunications_EventAccountCreate_Handler,
		},
		{
			MethodName: "EventAccountCheck",
			Handler:    _NodeCommunications_EventAccountCheck_Handler,
		},
		{
			MethodName: "EventTrackAccount",
			Handler:    _NodeCommunications_EventTrackAccount_Handler,
		},
		{
			MethodName: "EventSetTrackedAccounts",
			Handler:    _NodeCommunications_EventSetTrackedAccounts_Handler,
		},
		{
			MethodName: "EventGetRAMPrice",
			Handler:    _NodeCommunications_EventGetRAMPrice_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BalanceChanged",
			Handler:       _NodeCommunications_BalanceChanged_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "eos.proto",
}

func init() { proto1.RegisterFile("eos.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 904 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x55, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0x9e, 0xec, 0x38, 0xb1, 0x8e, 0x7f, 0xc3, 0x66, 0xad, 0xe7, 0x2e, 0xab, 0xa7, 0x15, 0x43,
	0xd0, 0x8b, 0x62, 0x49, 0xd7, 0x6d, 0x17, 0xdb, 0x80, 0x44, 0x71, 0x0a, 0x6f, 0x59, 0x5a, 0x70,
	0x01, 0x06, 0xec, 0x46, 0xa0, 0xa5, 0xd3, 0x5a, 0xb0, 0x25, 0xba, 0x24, 0x9d, 0x25, 0x17, 0xbb,
	0xdb, 0x1b, 0xec, 0x71, 0xb6, 0x87, 0x2b, 0x48, 0x51, 0xb2, 0xac, 0x3a, 0x57, 0xe2, 0xf9, 0xce,
	0x77, 0x7e, 0x78, 0xf8, 0x91, 0x02, 0x17, 0xb9, 0x7c, 0xbe, 0x14, 0x5c, 0x71, 0xd2, 0x30, 0x1f,
	0xef, 0x25, 0x74, 0x5e, 0xcf, 0xc7, 0x42, 0x50, 0x94, 0x4b, 0x9e, 0x4a, 0x24, 0x5d, 0xa8, 0xf1,
	0xf9, 0xc0, 0x19, 0x39, 0x47, 0x4d, 0x5a, 0xe3, 0x73, 0x72, 0x00, 0x0d, 0x14, 0x82, 0x8b, 0x41,
	0x6d, 0xe4, 0x1c, 0xb9, 0x34, 0x33, 0xbc, 0x3d, 0x68, 0x8c, 0x93, 0xa5, 0xba, 0xf3, 0xfe, 0x77,
	0xc0, 0xf5, 0x67, 0x2c, 0x4e, 0x27, 0xe9, 0x5b, 0x4e, 0x9e, 0x42, 0x77, 0x86, 0x2c, 0x0a, 0xa6,
	0x0b, 0x1e, 0xce, 0x83, 0x74, 0x95, 0x98, 0x44, 0x1d, 0xda, 0xd6, 0xe8, 0x99, 0x06, 0xaf, 0x56,
	0x09, 0xf1, 0xa0, 0x53, 0x62, 0xc5, 0x91, 0x4d, 0xdd, 0x2a, 0x48, 0x93, 0x88, 0x7c, 0x0d, 0xbd,
	0x12, 0x47, 0xc5, 0x09, 0x0e, 0xea, 0x23, 0xe7, 0xa8, 0x4e, 0x3b, 0x05, 0xeb, 0x3a, 0x4e, 0x90,
	0xfc, 0x04, 0x8f, 0x17, 0x4c, 0xaa, 0x20, 0x16, 0x02, 0x6f, 0x50, 0xc8, 0x78, 0xba, 0xc0, 0x52,
	0xf9, 0x1d, 0x53, 0x7e, 0xa0, 0x29, 0x93, 0x12, 0x23, 0x6f, 0xc5, 0xbb, 0x04, 0x38, 0x63, 0x0b,
	0x96, 0x86, 0x48, 0xf1, 0x3d, 0x21, 0xb0, 0x93, 0xb2, 0x04, 0x4d, 0xd3, 0x2e, 0x35, 0x6b, 0xf2,
	0x10, 0x76, 0xe5, 0x5d, 0x32, 0xe5, 0x0b, 0xdb, 0xa5, 0xb5, 0x34, 0x37, 0xe4, 0x51, 0xd6, 0x95,
	0x4b, 0xcd, 0xda, 0xfb, 0x0e, 0x9a, 0xa7, 0x61, 0xc8, 0x57, 0xa9, 0x92, 0xe4, 0x19, 0x34, 0x99,
	0x5d, 0x0f, 0x9c, 0x51, 0xfd, 0xa8, 0x75, 0xd2, 0xcd, 0x26, 0xff, 0xdc, 0x52, 0x68, 0xe1, 0xf7,
	0x0e, 0x61, 0xcf, 0x82, 0xdb, 0x5a, 0xf0, 0x7e, 0x81, 0xa6, 0x6d, 0x52, 0x92, 0x01, 0xec, 0xd9,
	0x30, 0x4b, 0xc9, 0x4d, 0xf2, 0x14, 0x76, 0x99, 0x94, 0xa8, 0xe4, 0xa0, 0x66, 0xca, 0xb5, 0xf3,
	0x72, 0x1a, 0xa4, 0xd6, 0xe7, 0xfd, 0x01, 0x0d, 0x03, 0x98, 0x44, 0x49, 0x52, 0x24, 0xaa, 0xd3,
	0xdc, 0x24, 0x9f, 0x83, 0xbb, 0x14, 0x18, 0xc6, 0x32, 0xe6, 0xa9, 0xd9, 0x74, 0x87, 0xae, 0x81,
	0xd2, 0x3c, 0xea, 0xe5, 0x79, 0x78, 0xff, 0x38, 0xd0, 0xb2, 0x9b, 0x30, 0x52, 0xd8, 0x36, 0xcb,
	0xaf, 0xa0, 0x23, 0x58, 0x12, 0xb0, 0x1b, 0x16, 0x2f, 0xd8, 0x74, 0x81, 0x26, 0x7b, 0x9d, 0xb6,
	0x05, 0x4b, 0x4e, 0x73, 0x8c, 0xfc, 0x08, 0x0f, 0x42, 0x2e, 0x30, 0x58, 0xc4, 0xef, 0x57, 0x71,
	0x14, 0x4c, 0xb3, 0x9d, 0x9b, 0x6a, 0xd5, 0x4d, 0xed, 0x6b, 0xe2, 0xa5, 0xe1, 0xd9, 0x01, 0x79,
	0xff, 0x39, 0xd0, 0xba, 0x16, 0x2c, 0x95, 0x2c, 0x54, 0xba, 0xdd, 0x2f, 0x00, 0x64, 0xfc, 0x2e,
	0x65, 0x6a, 0x25, 0x30, 0x3b, 0x08, 0x97, 0x96, 0x10, 0xf2, 0x2d, 0xb4, 0x42, 0x9e, 0x2c, 0x05,
	0xca, 0x62, 0xbb, 0xdd, 0x13, 0x62, 0xab, 0xf8, 0x6b, 0x0f, 0x2d, 0xd3, 0xc8, 0xf7, 0x30, 0x58,
	0xb2, 0x70, 0x8e, 0x51, 0x10, 0xf2, 0x54, 0xe1, 0xad, 0x0a, 0xde, 0x0a, 0xc4, 0x20, 0x62, 0x8a,
	0x99, 0x46, 0xdb, 0xf4, 0xd3, 0xcc, 0xef, 0x67, 0xee, 0x0b, 0x81, 0x78, 0xce, 0x14, 0x23, 0x87,
	0x00, 0x36, 0x50, 0x89, 0x5b, 0xa3, 0xce, 0x36, 0x75, 0x33, 0xe4, 0x5a, 0xdc, 0x7a, 0x17, 0xf0,
	0xe0, 0xcd, 0x4a, 0xce, 0x4a, 0x1b, 0xd0, 0xf7, 0x52, 0xdf, 0xc9, 0x38, 0xb2, 0x93, 0xac, 0xc5,
	0x11, 0x79, 0x02, 0x2d, 0xa9, 0x98, 0x5a, 0xc9, 0xa0, 0x24, 0x41, 0xc8, 0x20, 0x5f, 0x0b, 0xf1,
	0x09, 0x74, 0x4a, 0x39, 0x26, 0xe7, 0xd5, 0x0c, 0xde, 0xcf, 0xd0, 0x2b, 0x13, 0xf4, 0x81, 0x55,
	0x8b, 0x3c, 0x06, 0x77, 0x7d, 0x8f, 0x32, 0x19, 0x34, 0xa7, 0xf9, 0xbd, 0xf9, 0x1b, 0xfa, 0xf6,
	0xb0, 0x7d, 0x81, 0x4c, 0xdd, 0x7b, 0x7b, 0x0e, 0x01, 0x74, 0x89, 0x1b, 0x0c, 0xe6, 0x78, 0x67,
	0x6f, 0x90, 0x9b, 0x21, 0xbf, 0xe2, 0x9d, 0xae, 0xc1, 0xff, 0x4a, 0x51, 0x18, 0x6f, 0xb6, 0x8d,
	0xa6, 0x01, 0xb4, 0xf3, 0x33, 0x68, 0x6a, 0xb5, 0x84, 0x5c, 0x2a, 0x33, 0xa9, 0x1d, 0xba, 0x27,
	0x58, 0xe2, 0x73, 0xa9, 0xbc, 0x43, 0x68, 0x8c, 0x6f, 0x63, 0xa9, 0xcc, 0xeb, 0xa4, 0x17, 0xf6,
	0xc1, 0xca, 0x0c, 0x6f, 0x04, 0x4d, 0x7a, 0xfa, 0xdb, 0x1b, 0x11, 0x87, 0xa8, 0x19, 0x4b, 0xbd,
	0x30, 0x0c, 0x87, 0x66, 0xc6, 0xb3, 0x2f, 0xa1, 0x55, 0x3a, 0x5c, 0xd2, 0x84, 0x9d, 0xab, 0xd7,
	0x57, 0xe3, 0xfe, 0x27, 0x7a, 0xf5, 0xe7, 0xe5, 0xe4, 0xac, 0xef, 0x9c, 0xfc, 0xdb, 0x00, 0x72,
	0xc5, 0x23, 0xf4, 0x79, 0x92, 0xac, 0xd2, 0x38, 0x64, 0x7a, 0x52, 0x92, 0xbc, 0x80, 0xfd, 0xf1,
	0x0d, 0xa6, 0xea, 0x15, 0xaa, 0xf5, 0xbb, 0x97, 0xcb, 0xd2, 0xbc, 0x89, 0xc3, 0x7e, 0x2e, 0x9f,
	0xc2, 0xff, 0x12, 0x7a, 0x79, 0x90, 0x15, 0x2a, 0xd9, 0xb7, 0xa4, 0xf5, 0xf3, 0x33, 0xec, 0x6d,
	0x42, 0xb2, 0x1c, 0x96, 0xbf, 0x0f, 0x95, 0x47, 0x64, 0x48, 0x36, 0x6d, 0x53, 0xed, 0x02, 0x0e,
	0x4c, 0x58, 0x45, 0x4a, 0x24, 0xe7, 0x96, 0xb0, 0xe1, 0xd0, 0x62, 0xdb, 0x64, 0xf7, 0x0a, 0x1e,
	0xe5, 0xe5, 0xab, 0x62, 0x39, 0xf8, 0x38, 0xd5, 0xe4, 0x7c, 0xf8, 0x70, 0x0b, 0xaa, 0xd9, 0x3e,
	0x10, 0x93, 0x68, 0x43, 0x32, 0xe4, 0xd1, 0x66, 0xeb, 0x85, 0x90, 0x86, 0x79, 0xf2, 0xcd, 0x1f,
	0xd3, 0xb1, 0x1d, 0x7c, 0x4e, 0x9f, 0x61, 0x38, 0xff, 0x68, 0x1c, 0xc5, 0x41, 0x18, 0x75, 0x1c,
	0x43, 0xd7, 0xce, 0xd2, 0x9f, 0xb1, 0xf4, 0x1d, 0x46, 0x95, 0x83, 0xaa, 0x0e, 0xfc, 0x1b, 0xa7,
	0xa8, 0x72, 0x2d, 0x58, 0x38, 0xbf, 0x6f, 0xe8, 0x1b, 0x59, 0xc8, 0x0f, 0x76, 0x4c, 0xbf, 0x63,
	0x16, 0x85, 0x51, 0xf1, 0x13, 0xe8, 0x6d, 0x06, 0xca, 0x4a, 0xe4, 0x31, 0xf4, 0xf3, 0x01, 0x17,
	0x7a, 0xdd, 0xde, 0x61, 0xee, 0x9e, 0xee, 0x1a, 0xfb, 0xc5, 0x87, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xf8, 0x7b, 0x33, 0xdd, 0xca, 0x07, 0x00, 0x00,
}
