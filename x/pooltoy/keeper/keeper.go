package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/interchainberlin/pooltoy/x/pooltoy/types"
	"github.com/tendermint/tendermint/libs/log"
)

// can just use codec marshal binary

// Keeper of the pooltoy store
type Keeper struct {
	cdc           codec.BinaryMarshaler
	storeKey      sdk.StoreKey
	AccountKeeper authkeeper.AccountKeeper
}

// NewKeeper creates a pooltoy keeper
func NewKeeper(
	cdc codec.BinaryMarshaler,
	storeKey sdk.StoreKey,
	accountKeeper authkeeper.AccountKeeper,
) Keeper {
	return Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		AccountKeeper: accountKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Set sets a value in the db with a prefixed key
func (k Keeper) SetUser(ctx sdk.Context, key []byte, prefix []byte, user []byte) {
	store := ctx.KVStore(k.storeKey)
	store.Set(append(prefix, key...), user)
}

// GetAll values from with a prefix from the store
func (k Keeper) GetAll(ctx sdk.Context, prefix []byte) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, prefix)
}

func (k Keeper) ListAccounts(ctx sdk.Context) []authtypes.AccountI {
	return k.AccountKeeper.GetAllAccounts(ctx)
}
