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

// todo: merge OfferSend and Offer into 1
func (k Keeper) Offer(c context.Context, msg *types.OfferRequest) (*types.OfferResponse, error) {
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

func (k Keeper) OfferSend(ctx sdk.Context, msg *types.OfferRequest) (*types.OfferResponse, error) {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return &types.OfferResponse{}, err
	}

	err = k.BankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, msg.Amount)
	if err != nil {
		return &types.OfferResponse{}, err
	}

	return &types.OfferResponse{Sender: msg.Sender, Index: k.GetLatestID(ctx)}, nil
}

func (k Keeper) Response(c context.Context, msg *types.ResponseRequest) (*types.ResponseResult, error) {
	responser, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	ctx := sdk.UnwrapSDKContext(c)
	offer, err := k.QueryOfferByID(c, &types.QueryOfferByIDRequest{msg.Id})
	if err != nil {
		return nil, err
	}
	offerer, err := sdk.AccAddressFromBech32(offer.Sender)
	if err != nil {
		return nil, err
	}

	escrowAcc :=k.AccountKeeper.GetModuleAccount(ctx, types.ModuleName)
	err = k.BankKeeper.SendCoins(ctx, escrowAcc.GetAddress(), responser, offer.Amount)
	if err != nil {
		return nil, err
	}

	err = k.BankKeeper.SendCoins(ctx, responser, offerer, offer.Request)

	err = k.DeleteOffer(ctx, msg.Id)
	if err != nil {
		return nil, err
	}

	return &types.ResponseResult{}, nil
}


func (k Keeper) CancelOffer(c context.Context, msg *types.CancelOfferRequest) (*types.CancelOfferResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	offer, err := k.QueryOfferByID(c, &types.QueryOfferByIDRequest{msg.Id})
	if err != nil {
		return nil, err
	}
	if offer.Sender != msg.Sender{
		return nil, sdkerrors.Wrap(err, fmt.Sprintf("unauthorized to cancel this offer"))
	}

	offerer, err := sdk.AccAddressFromBech32(offer.Sender)
	if err != nil {
		return nil, err
	}
	escrowAcc :=k.AccountKeeper.GetModuleAccount(ctx, types.ModuleName)
	err = k.BankKeeper.SendCoins(ctx, escrowAcc.GetAddress(), offerer, offer.Amount)
	if err != nil {
		return nil, err
	}

	err = k.DeleteOffer(ctx, msg.Id)
	if err != nil {
		return nil, err
	}

	return &types.CancelOfferResponse{}, nil
}
