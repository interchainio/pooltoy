package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/interchainberlin/pooltoy/x/pooltoy/types"
	"github.com/tendermint/tendermint/libs/log"
)

// Keeper of the pooltoy store
type Keeper struct {
	Cdc      codec.BinaryMarshaler
	StoreKey sdk.StoreKey

	CoinKeeper    types.BankKeeper
	AccountKeeper types.AccountKeeper
	// paramspace types.ParamSubspace
}

// NewKeeper creates a pooltoy keeper
func NewKeeper(
	cdc codec.BinaryMarshaler,
	key sdk.StoreKey,

	coinKeeper types.BankKeeper,
	accountKeeper types.AccountKeeper,
) Keeper {
	return Keeper{
		Cdc:      cdc,
		StoreKey: key,

		CoinKeeper:    coinKeeper,
		AccountKeeper: accountKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// TODO: check out where accounts are now
// func (k Keeper) ListAccounts(ctx sdk.Context) []exported.Account {
// 	return k.AccountKeeper.GetAllAccounts(ctx)
// }

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
