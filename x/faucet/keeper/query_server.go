package keeper

import (
	"context"
	"sort"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/interchainberlin/pooltoy/regex"
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

	// temporarily revert back to the normal regex that doesn't allow emojis
	sdk.SetCoinDenomRegex(sdk.DefaultCoinDenomRegex)
	for _, account := range accounts {
		// filter out module accounts
		if _, ok := account.(authtypes.ModuleAccountI); ok {
			continue
		}

		var amount = int64(0)
		addr = account.GetAddress()
		balances = k.BankKeeper.GetAllBalances(ctx, addr)
		for _, emoji := range balances {

			_, err := sdk.ParseCoinsNormalized("1" + emoji.Denom)
			if err != nil {
				amount += emoji.Amount.Int64()
			}
		}

		ranks = append(ranks, &types.Amount{Address: addr.String(), Total: amount})
	}

	// reinstate the new regex rules in case sdk is used again afterwards
	sdk.SetCoinDenomRegex(func() string {
		return regex.NewDnmRegex
	})

	if int(req.ShowNum) > len(ranks) {
		req.ShowNum = int64(len(ranks))
	}

	sort.Slice(ranks, func(i, j int) bool {
		return ranks[i].Total > ranks[j].Total
	})

	return &types.QueryEmojiRankResponse{Rank: ranks[:req.ShowNum]}, nil
}
