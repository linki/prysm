package protoarray

import "sync"

// ForkChoice defines the overall fork choice store which includes all block nodes, validator's latest votes and balances.
type ForkChoice struct {
	store    *Store
	votes    []Vote   // tracks individual validator's last vote.
	balances []uint64 // tracks individual validator's last justified balances.
}

// Store defines the fork choice store which includes block nodes and the last view of checkpoint information.
type Store struct {
	pruneThreshold  uint64              // do not prune tree unless threshold is reached.
	justifiedEpoch  uint64              // latest justified epoch in store.
	finalizedEpoch  uint64              // latest finalized epoch in store.
	finalizedRoot   [32]byte            // latest finalized root in store.
	nodes           []*Node             // list of block nodes, each node is a representation of one block.
	nodeIndices     map[[32]byte]uint64 // the root of block node and the nodes index in the list.
	nodeIndicesLock sync.RWMutex
}

// Node defines the individual block which includes its block parent, ancestor and how much weight accounted for it.
// This is used as an array based stateful DAG for efficient fork choice look up.
type Node struct {
	Slot           uint64   // slot of the block converted to the node.
	Root           [32]byte // Root of the block converted to the node.
	Parent         uint64   // the parent index of this node.
	justifiedEpoch uint64   // justified epoch of this node.
	finalizedEpoch uint64   // finalized epoch of this node.
	Weight         uint64   // weight of this node.
	bestChild      uint64   // best child index of this node.
	BestDescendent uint64   // head index of this node.
	Graffiti       [32]byte // graffati of the block node.
}

// Vote defines an individual validator's vote.
type Vote struct {
	currentRoot [32]byte // current voting root.
	nextRoot    [32]byte // next voting root.
	nextEpoch   uint64   // epoch of next voting period.
}

// NonExistentNode defines an unknown node which is used for the array based stateful DAG.
const NonExistentNode = ^uint64(0)
