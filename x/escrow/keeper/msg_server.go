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

func (k Keeper) Offer(c context.Context, msg *types.Offer) (*types.OfferResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	res, err := k.OfferSend(ctx, msg)
	if err != nil {
		return nil, sdkerrors.Wrap(err, fmt.Sprintf("unable to offer"))
	}


	_, err = k.InsertOffer(ctx, *msg)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (k Keeper) OfferSend(ctx sdk.Context, msg *types.Offer) (*types.OfferResponse, error) {

	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		fmt.Println("addr err!!!!")
		return &types.OfferResponse{}, err
	}
//	coins, err := sdk.ParseCoinsNormalized(msg.Amount)
//	if err != nil {
//		return &types.OfferResponse{}, err
//	}

	//	moduleAcc:= k.AccountKeeper.GetModuleAddress(types.ModuleName)

	err = k.BankKeeper.SendCoinsFromAccountToModule(ctx, addr,types.ModuleName,msg.Amount)
	if err != nil {
		fmt.Println("sending err!!!!")
		return &types.OfferResponse{}, err
	}

	presentIdx := k.index
	*k.index +=1   // some checks this index is not re
	return &types.OfferResponse{Sender: msg.Sender, Index: *presentIdx}, nil
}

