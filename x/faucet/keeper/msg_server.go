package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/interchainberlin/pooltoy/x/faucet/types"
)

var _ types.MsgServer = Keeper{}

func (k Keeper) Mint(c context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if err := k.MintAndSend(ctx, msg); err != nil {
		// TODO: does this error still make sense?
		return nil, sdkerrors.Wrap(err, fmt.Sprintf(" in [%v] hours", k.Limit.Hours()))
	}
	return &types.MsgMintResponse{}, nil
}
