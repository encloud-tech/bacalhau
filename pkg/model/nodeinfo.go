package model

import (
	"context"

	"github.com/libp2p/go-libp2p/core/peer"
)

//go:generate stringer -type=NodeType -trimprefix=NodeType -output=nodeinfo_string.go
type NodeType int

const (
	NodeTypeRequester NodeType = iota
	NodeTypeCompute
)

type NodeInfoProvider interface {
	GetNodeInfo(ctx context.Context) NodeInfo
}

type ComputeNodeInfoProvider interface {
	GetComputeInfo(ctx context.Context) ComputeNodeInfo
}

type NodeInfo struct {
	BacalhauVersion BuildVersionInfo  `json:"BacalhauVersion"`
	PeerInfo        peer.AddrInfo     `json:"PeerInfo"`
	NodeType        NodeType          `json:"NodeType"`
	Labels          map[string]string `json:"Labels"`
	ComputeNodeInfo *ComputeNodeInfo  `json:"ComputeNodeInfo"`
}

// IsComputeNode returns true if the node is a compute node
func (n NodeInfo) IsComputeNode() bool {
	return n.NodeType == NodeTypeCompute
}

type ComputeNodeInfo struct {
	ExecutionEngines   []Engine            `json:"ExecutionEngines"`
	Verifiers          []Verifier          `json:"Verifiers"`
	Publishers         []Publisher         `json:"Publishers"`
	StorageSources     []StorageSourceType `json:"StorageSources"`
	MaxCapacity        ResourceUsageData   `json:"MaxCapacity"`
	AvailableCapacity  ResourceUsageData   `json:"AvailableCapacity"`
	MaxJobRequirements ResourceUsageData   `json:"MaxJobRequirements"`
	RunningExecutions  int                 `json:"RunningExecutions"`
	EnqueuedExecutions int                 `json:"EnqueuedExecutions"`
}
