// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package consts

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/hypersdk/chain"
	"github.com/ava-labs/hypersdk/codec"
	"github.com/ava-labs/hypersdk/consts"
)

const (
	// Choose a human-readable part (HRP) for your hyperchain
	// For example, "hyp" for "Hyperchain".
	HRP = "Khyati_chain"

	// Choose a name for your hyperchain
	// For example, "HyperNet".
	Name = "Khyati"

	// Choose a token symbol for your hyperchain
	// For example, "HYP" for the native token.
	Symbol = "KHYT"
)

var ID ids.ID

func init() {
	// Ensure the name is appropriately padded for the ID generation.
	b := make([]byte, consts.IDLen)
	copy(b, []byte(Name))

	// Generate the ID based on the Name and handle errors.
	vmID, err := ids.ToID(b)
	if err != nil {
		panic(err)
	}
	ID = vmID
}

// Instantiate registry here so it can be imported by any package.
// We set these values in [controller/registry].
var (
	ActionRegistry *codec.TypeParser[chain.Action, *warp.Message, bool]
	AuthRegistry   *codec.TypeParser[chain.Auth, *warp.Message, bool]
)
