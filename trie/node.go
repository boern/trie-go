// Copyright 2021 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package trie

import (
	"github.com/octopus-network/trie-go/substrate/node"
)

// Node is a node in the trie and can be a leaf or a branch.
type Node = node.Node