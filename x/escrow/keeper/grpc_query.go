package keeper

import (
	"github.com/interchainberlin/pooltoy/x/escrow/types"
)

var _ types.QueryServer = Keeper{}
