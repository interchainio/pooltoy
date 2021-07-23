package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/interchainberlin/pooltoy/x/escrow/utils"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/interchainberlin/pooltoy/x/escrow/types"
	// this line is used by starport scaffolding # ibc/keeper/import
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

type Keeper struct {
	BankKeeper bankkeeper.Keeper
	AccountKeeper keeper.AccountKeeper
	cdc        codec.Marshaler
	storeKey   sdk.StoreKey
//	memKey     sdk.StoreKey
	// this line is used by starport scaffolding # ibc/keeper/attribute
	index int64
}

func NewKeeper(
	bankKeeper bankkeeper.Keeper,
	accountKeeper keeper.AccountKeeper,
	cdc codec.Marshaler,
	storeKey sdk.StoreKey,
	//memKey sdk.StoreKey,
	index int64,
// this line is used by starport scaffolding # ibc/keeper/parameter
) Keeper {
	return Keeper{
		BankKeeper: bankKeeper,
		AccountKeeper: accountKeeper,
		cdc:      cdc,
		storeKey: storeKey,
		index:    index,
		// this line is used by starport scaffolding # ibc/keeper/return
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) OfferSend(ctx sdk.Context, msg *types.OfferRequest) (*types.OfferResponse, error) {

	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		fmt.Println("addr err!!!!")
		return &types.OfferResponse{}, err
	}
	amount, err := utils.ParseCoins(msg.Amount)
	if err != nil {
		return &types.OfferResponse{}, err
	}

//	moduleAcc:= k.AccountKeeper.GetModuleAddress(types.ModuleName)

	err = k.BankKeeper.SendCoinsFromAccountToModule(ctx, addr,types.ModuleName, amount)
	if err != nil {
		fmt.Println("sending err!!!!")
		return &types.OfferResponse{}, err
	}

	//presentIdx := k.index
	//k.index +=1   // some checks this index is not re
	return &types.OfferResponse{}, nil
}
