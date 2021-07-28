package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/keeper"
	// this line is used by starport scaffolding # ibc/keeper/import
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/interchainberlin/pooltoy/x/escrow/types"
	"github.com/tendermint/tendermint/libs/log"

)

type Keeper struct {
	BankKeeper bankkeeper.Keeper
	AccountKeeper keeper.AccountKeeper
	cdc        codec.Marshaler
	storeKey   sdk.StoreKey
//	memKey     sdk.StoreKey
	// this line is used by starport scaffolding # ibc/keeper/attribute
	index *int64
}

func NewKeeper(
	bankKeeper bankkeeper.Keeper,
	accountKeeper keeper.AccountKeeper,
	cdc codec.Marshaler,
	storeKey sdk.StoreKey,
	//memKey sdk.StoreKey,
	index *int64,
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
