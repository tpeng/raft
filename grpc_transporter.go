package raft

import (
	"log"
	"time"

	nc "golang.org/x/net/context"
	"google.golang.org/grpc"
	protobuf "raft/proto"
)

type GRPCTransporter struct {
}

func NewGRPCTransporter(prefix string, timeout time.Duration) *GRPCTransporter {
	t := &GRPCTransporter{}
	return t
}

func (t *GRPCTransporter) Install(server Server, mux interface{}) {
	// nothing to do
}

func (t *GRPCTransporter) SendAppendEntriesRequest(server Server, peer *Peer, req *AppendEntriesRequest) *AppendEntriesResponse {
	conn, err := grpc.Dial(peer.ConnectionString, grpc.WithInsecure())
	if err != nil {
		log.Println("can't connect to", peer.ConnectionString)
		log.Println(err)
		return nil
	}
	defer conn.Close()
	client := protobuf.NewRaftClient(conn)
	resp, err := client.AppendEntries(nc.Background(), &protobuf.AppendEntriesRequest{
		Term:         req.Term,
		PrevLogIndex: req.PrevLogIndex,
		PrevLogTerm:  req.PrevLogTerm,
		CommitIndex:  req.CommitIndex,
		LeaderName:   req.LeaderName,
		Entries:      req.Entries,
	})
	if err != nil {
		log.Println("error when call AppendEntries", err)
		return nil
	}
	// fmt.Println("AppendEntries resp:", resp)
	return newAppendEntriesResponse(resp.Term, resp.Success, resp.Index, resp.CommitIndex)
}

func (t *GRPCTransporter) SendVoteRequest(server Server, peer *Peer, req *RequestVoteRequest) *RequestVoteResponse {
	// log.Println("SendVoteRequest to", peer.ConnectionString)
	conn, err := grpc.Dial(peer.ConnectionString, grpc.WithInsecure())
	if err != nil {
		log.Println("can't connect to", peer.ConnectionString)
		log.Println(err)
		return nil
	}
	defer conn.Close()
	client := protobuf.NewRaftClient(conn)
	resp, err := client.RequestVote(nc.Background(), &protobuf.RequestVoteRequest{
		Term:          req.Term,
		LastLogIndex:  req.LastLogIndex,
		LastLogTerm:   req.LastLogTerm,
		CandidateName: req.CandidateName,
	})
	if err != nil {
		log.Println("error when call RequestVote", err)
		return nil
	}
	// fmt.Println("RequestVote resp:", resp)
	return newRequestVoteResponse(resp.Term, resp.VoteGranted)
}

func (t *GRPCTransporter) SendSnapshotRequest(server Server, peer *Peer, req *SnapshotRequest) *SnapshotResponse {
	conn, err := grpc.Dial(peer.ConnectionString, grpc.WithInsecure())
	if err != nil {
		log.Println("can't connect to", peer.ConnectionString)
		return nil
	}
	defer conn.Close()
	client := protobuf.NewRaftClient(conn)
	resp, err := client.RequestSnapshot(nc.Background(), &protobuf.SnapshotRequest{
		LeaderName: req.LeaderName,
		LastIndex:  req.LastIndex,
		LastTerm:   req.LastTerm,
	})
	if err != nil {
		log.Println("error when call Snapshot", err)
		return nil
	}
	// fmt.Println("Snapshot resp:", resp)
	return newSnapshotResponse(resp.Success)
}

func (t *GRPCTransporter) SendSnapshotRecoveryRequest(server Server, peer *Peer, req *SnapshotRecoveryRequest) *SnapshotRecoveryResponse {
	conn, err := grpc.Dial(peer.ConnectionString, grpc.WithInsecure())
	if err != nil {
		log.Println("can't connect to", peer.ConnectionString)
		return nil
	}
	defer conn.Close()
	client := protobuf.NewRaftClient(conn)
	resp, err := client.RequestSnapshotRecovery(nc.Background(), &protobuf.SnapshotRecoveryRequest{
		LeaderName: req.LeaderName,
		LastIndex:  req.LastIndex,
		LastTerm:   req.LastTerm,
	})

	if err != nil {
		log.Println("error when call SnapshotRecovery", err)
		return nil
	}
	log.Println("SnapshotRecovery resp:", resp)
	return newSnapshotRecoveryResponse(resp.Term, resp.Success, resp.CommitIndex)
}
