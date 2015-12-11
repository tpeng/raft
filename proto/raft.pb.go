// Code generated by protoc-gen-go.
// source: proto/raft.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto/raft.proto

It has these top-level messages:
	RequestVoteRequest
	RequestVoteResponse
	LogEntry
	AppendEntriesRequest
	AppendEntriesResponse
	SnapshotRequest
	SnapshotResponse
	SnapshotRecoveryRequest
	SnapshotRecoveryResponse
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

// the request/response defintion for RequestVote
type RequestVoteRequest struct {
	Term          uint64 `protobuf:"varint,1,opt,name=Term" json:"Term,omitempty"`
	LastLogIndex  uint64 `protobuf:"varint,2,opt,name=LastLogIndex" json:"LastLogIndex,omitempty"`
	LastLogTerm   uint64 `protobuf:"varint,3,opt,name=LastLogTerm" json:"LastLogTerm,omitempty"`
	CandidateName string `protobuf:"bytes,4,opt,name=CandidateName" json:"CandidateName,omitempty"`
}

func (m *RequestVoteRequest) Reset()                    { *m = RequestVoteRequest{} }
func (m *RequestVoteRequest) String() string            { return proto1.CompactTextString(m) }
func (*RequestVoteRequest) ProtoMessage()               {}
func (*RequestVoteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RequestVoteResponse struct {
	Term        uint64 `protobuf:"varint,1,opt,name=Term" json:"Term,omitempty"`
	VoteGranted bool   `protobuf:"varint,2,opt,name=VoteGranted" json:"VoteGranted,omitempty"`
}

func (m *RequestVoteResponse) Reset()                    { *m = RequestVoteResponse{} }
func (m *RequestVoteResponse) String() string            { return proto1.CompactTextString(m) }
func (*RequestVoteResponse) ProtoMessage()               {}
func (*RequestVoteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// the request/response defintion for AppendEntries
type LogEntry struct {
	Index       uint64 `protobuf:"varint,1,opt,name=Index" json:"Index,omitempty"`
	Term        uint64 `protobuf:"varint,2,opt,name=Term" json:"Term,omitempty"`
	CommandName string `protobuf:"bytes,3,opt,name=CommandName" json:"CommandName,omitempty"`
	Command     []byte `protobuf:"bytes,4,opt,name=Command,proto3" json:"Command,omitempty"`
}

func (m *LogEntry) Reset()                    { *m = LogEntry{} }
func (m *LogEntry) String() string            { return proto1.CompactTextString(m) }
func (*LogEntry) ProtoMessage()               {}
func (*LogEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type AppendEntriesRequest struct {
	Term         uint64      `protobuf:"varint,1,opt,name=Term" json:"Term,omitempty"`
	PrevLogIndex uint64      `protobuf:"varint,2,opt,name=PrevLogIndex" json:"PrevLogIndex,omitempty"`
	PrevLogTerm  uint64      `protobuf:"varint,3,opt,name=PrevLogTerm" json:"PrevLogTerm,omitempty"`
	CommitIndex  uint64      `protobuf:"varint,4,opt,name=CommitIndex" json:"CommitIndex,omitempty"`
	LeaderName   string      `protobuf:"bytes,5,opt,name=LeaderName" json:"LeaderName,omitempty"`
	Entries      []*LogEntry `protobuf:"bytes,6,rep,name=Entries" json:"Entries,omitempty"`
}

func (m *AppendEntriesRequest) Reset()                    { *m = AppendEntriesRequest{} }
func (m *AppendEntriesRequest) String() string            { return proto1.CompactTextString(m) }
func (*AppendEntriesRequest) ProtoMessage()               {}
func (*AppendEntriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AppendEntriesRequest) GetEntries() []*LogEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type AppendEntriesResponse struct {
	Term        uint64 `protobuf:"varint,1,opt,name=Term" json:"Term,omitempty"`
	Index       uint64 `protobuf:"varint,2,opt,name=Index" json:"Index,omitempty"`
	CommitIndex uint64 `protobuf:"varint,3,opt,name=CommitIndex" json:"CommitIndex,omitempty"`
	Success     bool   `protobuf:"varint,4,opt,name=Success" json:"Success,omitempty"`
}

func (m *AppendEntriesResponse) Reset()                    { *m = AppendEntriesResponse{} }
func (m *AppendEntriesResponse) String() string            { return proto1.CompactTextString(m) }
func (*AppendEntriesResponse) ProtoMessage()               {}
func (*AppendEntriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type SnapshotRequest struct {
	LeaderName string `protobuf:"bytes,1,opt,name=LeaderName" json:"LeaderName,omitempty"`
	LastIndex  uint64 `protobuf:"varint,2,opt,name=LastIndex" json:"LastIndex,omitempty"`
	LastTerm   uint64 `protobuf:"varint,3,opt,name=LastTerm" json:"LastTerm,omitempty"`
}

func (m *SnapshotRequest) Reset()                    { *m = SnapshotRequest{} }
func (m *SnapshotRequest) String() string            { return proto1.CompactTextString(m) }
func (*SnapshotRequest) ProtoMessage()               {}
func (*SnapshotRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type SnapshotResponse struct {
	Success bool `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
}

func (m *SnapshotResponse) Reset()                    { *m = SnapshotResponse{} }
func (m *SnapshotResponse) String() string            { return proto1.CompactTextString(m) }
func (*SnapshotResponse) ProtoMessage()               {}
func (*SnapshotResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type SnapshotRecoveryRequest struct {
	LeaderName string                          `protobuf:"bytes,1,opt,name=LeaderName" json:"LeaderName,omitempty"`
	LastIndex  uint64                          `protobuf:"varint,2,opt,name=LastIndex" json:"LastIndex,omitempty"`
	LastTerm   uint64                          `protobuf:"varint,3,opt,name=LastTerm" json:"LastTerm,omitempty"`
	Peers      []*SnapshotRecoveryRequest_Peer `protobuf:"bytes,4,rep,name=Peers" json:"Peers,omitempty"`
	State      []byte                          `protobuf:"bytes,5,opt,name=State,proto3" json:"State,omitempty"`
}

func (m *SnapshotRecoveryRequest) Reset()                    { *m = SnapshotRecoveryRequest{} }
func (m *SnapshotRecoveryRequest) String() string            { return proto1.CompactTextString(m) }
func (*SnapshotRecoveryRequest) ProtoMessage()               {}
func (*SnapshotRecoveryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *SnapshotRecoveryRequest) GetPeers() []*SnapshotRecoveryRequest_Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

type SnapshotRecoveryRequest_Peer struct {
	Name             string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	ConnectionString string `protobuf:"bytes,2,opt,name=ConnectionString" json:"ConnectionString,omitempty"`
}

func (m *SnapshotRecoveryRequest_Peer) Reset()                    { *m = SnapshotRecoveryRequest_Peer{} }
func (m *SnapshotRecoveryRequest_Peer) String() string            { return proto1.CompactTextString(m) }
func (*SnapshotRecoveryRequest_Peer) ProtoMessage()               {}
func (*SnapshotRecoveryRequest_Peer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7, 0} }

type SnapshotRecoveryResponse struct {
	Term        uint64 `protobuf:"varint,1,opt,name=Term" json:"Term,omitempty"`
	Success     bool   `protobuf:"varint,2,opt,name=Success" json:"Success,omitempty"`
	CommitIndex uint64 `protobuf:"varint,3,opt,name=CommitIndex" json:"CommitIndex,omitempty"`
}

func (m *SnapshotRecoveryResponse) Reset()                    { *m = SnapshotRecoveryResponse{} }
func (m *SnapshotRecoveryResponse) String() string            { return proto1.CompactTextString(m) }
func (*SnapshotRecoveryResponse) ProtoMessage()               {}
func (*SnapshotRecoveryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func init() {
	proto1.RegisterType((*RequestVoteRequest)(nil), "proto.RequestVoteRequest")
	proto1.RegisterType((*RequestVoteResponse)(nil), "proto.RequestVoteResponse")
	proto1.RegisterType((*LogEntry)(nil), "proto.LogEntry")
	proto1.RegisterType((*AppendEntriesRequest)(nil), "proto.AppendEntriesRequest")
	proto1.RegisterType((*AppendEntriesResponse)(nil), "proto.AppendEntriesResponse")
	proto1.RegisterType((*SnapshotRequest)(nil), "proto.SnapshotRequest")
	proto1.RegisterType((*SnapshotResponse)(nil), "proto.SnapshotResponse")
	proto1.RegisterType((*SnapshotRecoveryRequest)(nil), "proto.SnapshotRecoveryRequest")
	proto1.RegisterType((*SnapshotRecoveryRequest_Peer)(nil), "proto.SnapshotRecoveryRequest.Peer")
	proto1.RegisterType((*SnapshotRecoveryResponse)(nil), "proto.SnapshotRecoveryResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Raft service

type RaftClient interface {
	RequestVote(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteResponse, error)
	AppendEntries(ctx context.Context, in *AppendEntriesRequest, opts ...grpc.CallOption) (*AppendEntriesResponse, error)
	RequestSnapshot(ctx context.Context, in *SnapshotRequest, opts ...grpc.CallOption) (*SnapshotResponse, error)
	RequestSnapshotRecovery(ctx context.Context, in *SnapshotRecoveryRequest, opts ...grpc.CallOption) (*SnapshotRecoveryResponse, error)
}

type raftClient struct {
	cc *grpc.ClientConn
}

func NewRaftClient(cc *grpc.ClientConn) RaftClient {
	return &raftClient{cc}
}

func (c *raftClient) RequestVote(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteResponse, error) {
	out := new(RequestVoteResponse)
	err := grpc.Invoke(ctx, "/proto.Raft/RequestVote", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftClient) AppendEntries(ctx context.Context, in *AppendEntriesRequest, opts ...grpc.CallOption) (*AppendEntriesResponse, error) {
	out := new(AppendEntriesResponse)
	err := grpc.Invoke(ctx, "/proto.Raft/AppendEntries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftClient) RequestSnapshot(ctx context.Context, in *SnapshotRequest, opts ...grpc.CallOption) (*SnapshotResponse, error) {
	out := new(SnapshotResponse)
	err := grpc.Invoke(ctx, "/proto.Raft/RequestSnapshot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftClient) RequestSnapshotRecovery(ctx context.Context, in *SnapshotRecoveryRequest, opts ...grpc.CallOption) (*SnapshotRecoveryResponse, error) {
	out := new(SnapshotRecoveryResponse)
	err := grpc.Invoke(ctx, "/proto.Raft/RequestSnapshotRecovery", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Raft service

type RaftServer interface {
	RequestVote(context.Context, *RequestVoteRequest) (*RequestVoteResponse, error)
	AppendEntries(context.Context, *AppendEntriesRequest) (*AppendEntriesResponse, error)
	RequestSnapshot(context.Context, *SnapshotRequest) (*SnapshotResponse, error)
	RequestSnapshotRecovery(context.Context, *SnapshotRecoveryRequest) (*SnapshotRecoveryResponse, error)
}

func RegisterRaftServer(s *grpc.Server, srv RaftServer) {
	s.RegisterService(&_Raft_serviceDesc, srv)
}

func _Raft_RequestVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(RequestVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RaftServer).RequestVote(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Raft_AppendEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(AppendEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RaftServer).AppendEntries(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Raft_RequestSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RaftServer).RequestSnapshot(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Raft_RequestSnapshotRecovery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SnapshotRecoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RaftServer).RequestSnapshotRecovery(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Raft_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Raft",
	HandlerType: (*RaftServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestVote",
			Handler:    _Raft_RequestVote_Handler,
		},
		{
			MethodName: "AppendEntries",
			Handler:    _Raft_AppendEntries_Handler,
		},
		{
			MethodName: "RequestSnapshot",
			Handler:    _Raft_RequestSnapshot_Handler,
		},
		{
			MethodName: "RequestSnapshotRecovery",
			Handler:    _Raft_RequestSnapshotRecovery_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x53, 0x4f, 0x6f, 0x12, 0x41,
	0x14, 0x77, 0xe9, 0xd2, 0xc2, 0x03, 0x02, 0x4e, 0x5b, 0x59, 0x57, 0xa3, 0x64, 0x7a, 0xe9, 0x09,
	0x13, 0xbc, 0x78, 0x35, 0xf8, 0x27, 0x1a, 0x6c, 0x9a, 0x62, 0x3c, 0x18, 0x2f, 0x2b, 0xfb, 0xac,
	0x1c, 0x98, 0x59, 0x67, 0xa6, 0x8d, 0xfd, 0x2a, 0x7e, 0x1a, 0x2f, 0x7e, 0x2f, 0x67, 0xde, 0xce,
	0xc2, 0xc2, 0x2e, 0x5e, 0x3c, 0xb1, 0xf9, 0xcd, 0xe3, 0xf7, 0xef, 0xcd, 0xc0, 0x20, 0x53, 0xd2,
	0xc8, 0x67, 0x2a, 0xf9, 0x66, 0xc6, 0xf4, 0xc9, 0x9a, 0xf4, 0xc3, 0x53, 0x60, 0x57, 0xf8, 0xe3,
	0x06, 0xb5, 0xf9, 0x24, 0x0d, 0xfa, 0x4f, 0xd6, 0x85, 0xf0, 0x23, 0xaa, 0x55, 0x14, 0x8c, 0x82,
	0xf3, 0x90, 0x9d, 0x40, 0x77, 0x96, 0x68, 0x33, 0x93, 0xd7, 0xef, 0x44, 0x8a, 0x3f, 0xa3, 0x06,
	0xa1, 0xc7, 0xd0, 0xf1, 0x28, 0x8d, 0x1e, 0x10, 0x78, 0x0a, 0xbd, 0x69, 0x22, 0xd2, 0x65, 0x9a,
	0x18, 0xbc, 0x48, 0x56, 0x18, 0x85, 0x16, 0x6e, 0xf3, 0x17, 0x70, 0xbc, 0xa5, 0xa2, 0x33, 0x29,
	0x34, 0xee, 0xc8, 0x58, 0x42, 0x77, 0xfa, 0x56, 0x25, 0xc2, 0x60, 0x4a, 0x2a, 0x2d, 0xfe, 0x01,
	0x5a, 0x56, 0xe1, 0xb5, 0x30, 0xea, 0x8e, 0xf5, 0xa0, 0x99, 0x1b, 0xc8, 0xe7, 0x8b, 0x7f, 0xaf,
	0xed, 0x4c, 0xe5, 0x6a, 0x65, 0xc5, 0x49, 0xd7, 0xd9, 0x69, 0xb3, 0x3e, 0x1c, 0x79, 0x90, 0x8c,
	0x74, 0xf9, 0xaf, 0x00, 0x4e, 0x5e, 0x66, 0x19, 0x8a, 0xd4, 0x51, 0x2e, 0x51, 0xef, 0x4d, 0x7c,
	0xa9, 0xf0, 0xb6, 0x9a, 0xd8, 0xa3, 0xa5, 0xc4, 0x5e, 0x77, 0x69, 0xf2, 0xc9, 0x90, 0x40, 0x06,
	0x30, 0xc3, 0x24, 0x45, 0x45, 0x5e, 0x9a, 0xe4, 0x65, 0x04, 0x47, 0x5e, 0x33, 0x3a, 0x1c, 0x1d,
	0x9c, 0x77, 0x26, 0xfd, 0x7c, 0x13, 0xe3, 0x22, 0x1f, 0xff, 0x0c, 0xa7, 0x3b, 0xde, 0x6a, 0x7b,
	0x5a, 0xd7, 0xd0, 0xa8, 0x33, 0x90, 0xbb, 0xb2, 0xc1, 0xe7, 0x37, 0x8b, 0x05, 0x6a, 0x4d, 0x8e,
	0x5a, 0xfc, 0x3d, 0xf4, 0xe7, 0x22, 0xc9, 0xf4, 0x77, 0x69, 0x8a, 0xc8, 0xdb, 0x26, 0x03, 0x32,
	0x79, 0x1f, 0xda, 0x6e, 0xa9, 0x65, 0xfe, 0x81, 0xdd, 0x80, 0x85, 0x36, 0x91, 0xf9, 0x19, 0x0c,
	0x36, 0x5c, 0xde, 0x62, 0x49, 0x30, 0x20, 0xc1, 0x3f, 0x01, 0x0c, 0x37, 0x53, 0x0b, 0x79, 0x8b,
	0xea, 0xee, 0x7f, 0x95, 0xd9, 0x04, 0x9a, 0x97, 0x88, 0xca, 0x85, 0x72, 0x0d, 0x9e, 0xf9, 0x06,
	0xf7, 0xe8, 0x8c, 0xdd, 0xac, 0xab, 0x6b, 0x6e, 0xec, 0x75, 0xa4, 0x35, 0x74, 0xe3, 0x31, 0x84,
	0x04, 0xdb, 0x4e, 0x4b, 0xea, 0x11, 0x0c, 0xa6, 0x52, 0x08, 0x5c, 0x98, 0xa5, 0x14, 0x73, 0xdb,
	0xbe, 0xb8, 0x26, 0x13, 0x6d, 0x7e, 0x01, 0x51, 0x95, 0xbe, 0x76, 0x2f, 0xa5, 0x0a, 0xe8, 0xee,
	0xd6, 0x6e, 0x66, 0xf2, 0xbb, 0x01, 0xe1, 0x95, 0x7d, 0x86, 0xec, 0x0d, 0x74, 0x4a, 0x6f, 0x82,
	0x3d, 0xf4, 0x59, 0xaa, 0xaf, 0x31, 0x8e, 0xeb, 0x8e, 0x72, 0x0b, 0xfc, 0x1e, 0x9b, 0x41, 0x6f,
	0xeb, 0xd6, 0xb0, 0x47, 0x7e, 0xbc, 0xee, 0x9e, 0xc7, 0x8f, 0xeb, 0x0f, 0xd7, 0x6c, 0xaf, 0xa0,
	0xef, 0x47, 0x8b, 0xd4, 0xec, 0x41, 0xa5, 0xe5, 0x9c, 0x6a, 0x58, 0xc1, 0xd7, 0x2c, 0x5f, 0x60,
	0xb8, 0xc3, 0x52, 0x74, 0xc7, 0x9e, 0xfc, 0x7b, 0x67, 0xf1, 0xd3, 0xbd, 0xe7, 0x05, 0xfb, 0xd7,
	0x43, 0x9a, 0x78, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x78, 0xee, 0x20, 0xd5, 0x04, 0x00,
	0x00,
}
