package raft

// The request sent to a server to vote for a candidate to become a leader.
type RequestVoteRequest struct {
	peer          *Peer
	Term          uint64
	LastLogIndex  uint64
	LastLogTerm   uint64
	CandidateName string
}

// The response returned from a server after a vote for a candidate to become a leader.
type RequestVoteResponse struct {
	peer        *Peer
	Term        uint64
	VoteGranted bool
}

// Creates a new RequestVote request.
func newRequestVoteRequest(term uint64, candidateName string, lastLogIndex uint64, lastLogTerm uint64) *RequestVoteRequest {
	return &RequestVoteRequest{
		Term:          term,
		LastLogIndex:  lastLogIndex,
		LastLogTerm:   lastLogTerm,
		CandidateName: candidateName,
	}
}

// Creates a new RequestVote response.
func newRequestVoteResponse(term uint64, voteGranted bool) *RequestVoteResponse {
	return &RequestVoteResponse{
		Term:        term,
		VoteGranted: voteGranted,
	}
}
