package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/interchainberlin/pooltoy/x/escrow/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) QueryOfferListAll(c context.Context, req *types.OfferListAllRequest) (*types.OfferListResponse, error){
	//if req != nil{
	//	return nil, status.Error(codes.InvalidArgument, "non-empty request")
	//}

	ctx := sdk.UnwrapSDKContext(c)
	offers, err :=k.ListOffer(ctx, *req)
	if err != nil {
		return nil, err
	}

	return &offers, nil

}

func (k Keeper)QueryOfferByAddr(c context.Context, req *types.QueryOfferByAddrRequest) (*types.OfferListResponse, error){
	ctx := sdk.UnwrapSDKContext(c)
	offers, err :=k.ListOfferByAddr(ctx, *req)
	if err != nil {
		return nil, err
	}

	return &offers, nil
}

func (k Keeper) QueryOfferByID(c context.Context, req *types.QueryOfferByIDRequest) (*types.Offer, error){
	ctx := sdk.UnwrapSDKContext(c)
	offer, err :=k.ListOfferByID(ctx, *req)
	if err != nil {
		return nil, err
	}

	return &offer, nil
}
