package keeper

import (
	"context"
	_ "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktype "github.com/cosmos/cosmos-sdk/x/bank/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/interchainberlin/pooltoy/x/pooltoy/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) QueryListUsers(c context.Context, req *types.QueryListUsersRequest) (*types.QueryListUsersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	users := k.ListUsers(ctx)
	for _, u := range users {
		_, err := sdk.AccAddressFromBech32(u.UserAccount)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	return &types.QueryListUsersResponse{
		Users: users,
	}, nil
}

func (k Keeper) QueryMostEmojiOwner(c context.Context, req *types.QueryMostEmojiOwnerRequest) (*types.QueryMostEmojiOwnerResponse, error) {
	var mostMojiAdrr string
	var allBalance int64
	var maxAllBalance = int64(0)
	var allBalanceReq *banktype.QueryAllBalancesRequest
	//if req != nil {
	//	return nil, status.Error(codes.InvalidArgument, "non-empty request")
	//}
	ctx := sdk.UnwrapSDKContext(c)
	users := k.ListUsers(ctx)
	if len(users) == 0{
		return nil, status.Error(codes.InvalidArgument, "no users")
	}
	for _, u := range users {
		addr, err := sdk.AccAddressFromBech32(u.UserAccount)
		if err!= nil{
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		allBalanceReq =  &banktype.QueryAllBalancesRequest{Address:addr.String()}
		allBalanceResp, err := k.BankKeeper.AllBalances(c, allBalanceReq)
		if allBalanceResp == nil{
			continue
		}
		for _, emoji := range allBalanceResp.GetBalances(){
			allBalance += emoji.Amount.Int64()
		}

		if maxAllBalance < allBalance{
			maxAllBalance = allBalance
			mostMojiAdrr = addr.String()
		}
	}

	return &types.QueryMostEmojiOwnerResponse{Address: mostMojiAdrr, Total: maxAllBalance}, nil
}
