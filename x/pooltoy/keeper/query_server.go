package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/interchainio/pooltoy/x/pooltoy/types"
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
