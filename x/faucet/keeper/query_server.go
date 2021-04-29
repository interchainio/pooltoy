package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/interchainberlin/pooltoy/x/faucet/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) QueryWhenBrr(c context.Context, req *types.QueryWhenBrrRequest) (*types.QueryWhenBrrResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	a, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	mintTime := ctx.BlockTime().Unix()
	m := k.getMintHistory(ctx, a)
	ma, err := sdk.AccAddressFromBech32(m.Minter)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	isPresent := k.isPresent(ctx, ma)
	var timeLeft int64
	if !isPresent {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	lastTime := time.Unix(m.Lasttime, 0)
	currentTime := time.Unix(mintTime, 0)

	lastTimePlusLimit := lastTime.Add(k.Limit).UTC()
	isAfter := lastTimePlusLimit.After(currentTime)
	if isAfter {
		timeLeft = int64(lastTime.Add(k.Limit).UTC().Sub(currentTime).Seconds())
	} else {
		timeLeft = 0
	}

	return &types.QueryWhenBrrResponse{
		TimeLeft: timeLeft,
	}, nil
}
