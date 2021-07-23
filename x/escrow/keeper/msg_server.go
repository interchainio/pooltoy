package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/interchainberlin/pooltoy/x/escrow/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = Keeper{}

func (k Keeper) Offer(c context.Context, msg *types.OfferRequest) (*types.OfferResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	res, err := k.OfferSend(ctx, msg)
	if err != nil {
		return nil, sdkerrors.Wrap(err, fmt.Sprintf("unable to offer"))
	}

	return res, nil
}

//func (k Keeper) getIndex(ctx sdk.Context) int64{
//	//store := ctx.KVStore(k.storeKey)
//	//if k.isPresent(ctx, )
//	return 1
//}
