package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/interchainberlin/pooltoy/x/faucet/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sort"
	"time"
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
		// has never minted if not present in the keeper
		return &types.QueryWhenBrrResponse{
			TimeLeft: 0,
		}, nil
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

func (k Keeper) QueryEmojiRank(c context.Context, req *types.QueryEmojiRankRequest) (*types.QueryEmojiRankResponse, error) {

	var addr sdk.AccAddress
	var balances sdk.Coins
	var ranks = []*types.Amount{}

	ctx := sdk.UnwrapSDKContext(c)
	accounts := k.AccountKeeper.GetAllAccounts(ctx)
	for _, account := range accounts {
		// filter out module accounts
		if _, ok := account.(authtypes.ModuleAccountI); ok{
			continue
		}

		var amount = int64(0)
		addr = account.GetAddress()
		balances = k.BankKeeper.GetAllBalances(ctx, addr)
		for _, emoji := range balances {
			amount += emoji.Amount.Int64()
		}

		ranks = append(ranks, &types.Amount{Address: addr.String(), Total: amount})
	}

	if int(req.ShowNum) > len(ranks){
		return nil,status.Error(codes.InvalidArgument, "the rank list is shorter than the requested")
	}

	sort.Slice(ranks, func(i, j int) bool {
		return ranks[i].Total > ranks[j].Total
	})

	return &types.QueryEmojiRankResponse{Rank: ranks[:req.ShowNum]}, nil
}
