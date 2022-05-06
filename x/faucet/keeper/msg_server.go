package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/interchainio/pooltoy/x/faucet/types"
)

var _ types.MsgServer = Keeper{}

func (k Keeper) Mint(c context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if err := k.MintAndSend(ctx, msg); err != nil {
		return nil, err
	}
	return &types.MsgMintResponse{}, nil
}
