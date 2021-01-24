package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/interchainberlin/pooltoy/x/pooltoy/types"
)

// Keeper of the pooltoy store
type Keeper struct {
	CoinKeeper    bank.Keeper
	accountKeeper auth.AccountKeeper
	storeKey      sdk.StoreKey
	Cdc           *codec.Codec
	// paramspace types.ParamSubspace
}

// NewKeeper creates a pooltoy keeper
func NewKeeper(coinKeeper bank.Keeper, accountKeeper auth.AccountKeeper, cdc *codec.Codec, key sdk.StoreKey) Keeper {
	keeper := Keeper{
		CoinKeeper:    coinKeeper,
		accountKeeper: accountKeeper,
		storeKey:      key,
		Cdc:           cdc,
		// paramspace: paramspace.WithKeyTable(types.ParamKeyTable()),
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) ListAccounts(ctx sdk.Context) []exported.Account {
	return k.accountKeeper.GetAllAccounts(ctx)
}

// Get returns the pubkey from the adddress-pubkey relation
// func (k Keeper) Get(ctx sdk.Context, key string) (/* TODO: Fill out this type */, error) {
// 	store := ctx.KVStore(k.storeKey)
// 	var item /* TODO: Fill out this type */
// 	byteKey := []byte(key)
// 	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &item)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return item, nil
// }

// func (k Keeper) set(ctx sdk.Context, key string, value /* TODO: fill out this type */ ) {
// 	store := ctx.KVStore(k.storeKey)
// 	bz := k.cdc.MustMarshalBinaryLengthPrefixed(value)
// 	store.Set([]byte(key), bz)
// }

// func (k Keeper) delete(ctx sdk.Context, key string) {
// 	store := ctx.KVStore(k.storeKey)
// 	store.Delete([]byte(key))
// }
