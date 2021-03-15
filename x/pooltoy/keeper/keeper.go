package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/interchainberlin/pooltoy/x/pooltoy/types"
	"github.com/tendermint/tendermint/libs/log"
)

// UnmarshalFn is a generic function to unmarshal bytes
type UnmarshalFn func(value []byte) (interface{}, bool)

// UnmarshalFn is a generic function to unmarshal bytes
type MarshalFn func(value interface{}) []byte

// Keeper of the pooltoy store
type Keeper struct {
	Cdc      codec.Marshaler
	storeKey sdk.StoreKey
	memKey   sdk.StoreKey

	CoinKeeper    bankkeeper.Keeper
	AccountKeeper authkeeper.AccountKeeper
	// paramspace types.ParamSubspace
}

// NewKeeper creates a pooltoy keeper
func NewKeeper(
	cdc codec.Marshaler,
	storeKey,
	memKey sdk.StoreKey,

	coinKeeper bankkeeper.Keeper,
	accountKeeper authkeeper.AccountKeeper,
) *Keeper {
	return &Keeper{
		Cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,

		CoinKeeper:    coinKeeper,
		AccountKeeper: accountKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Set sets a value in the db with a prefixed key
func (k Keeper) Set(ctx sdk.Context, key []byte, prefix []byte, i interface{}, marshal MarshalFn) {
	store := ctx.KVStore(k.storeKey)
	store.Set(append(prefix, key...), marshal(i))
}

// Get gets an item from the store by bytes
func (k Keeper) Get(ctx sdk.Context, key []byte, prefix []byte, unmarshal UnmarshalFn) (i interface{}, found bool) {
	store := ctx.KVStore(k.storeKey)
	value := store.Get(append(prefix, key...))

	return unmarshal(value)
}

// GetAll values from with a prefix from the store
func (k Keeper) GetAll(ctx sdk.Context, prefix []byte) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, prefix)
}

// TODO: check out where accounts are now
func (k Keeper) ListAccounts(ctx sdk.Context) []authtypes.AccountI {
	return k.AccountKeeper.GetAllAccounts(ctx)
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
