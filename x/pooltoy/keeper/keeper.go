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

func (k Keeper) ListAccounts(ctx sdk.Context) []authtypes.AccountI {
	return k.AccountKeeper.GetAllAccounts(ctx)
}

func (k Keeper) InsertUser(ctx sdk.Context, user types.User) error {
	// TODO: use account address as key
	// then we can remove account address from U
	key := []byte(types.UserPrefix + user.Id)

	u, err := k.cdc.MarshalBinaryBare(&user)
	if err != nil {
		return err
	}
	p := []byte(types.UserPrefix)
	store := ctx.KVStore(k.storeKey)
	store.Set(append(p, key...), u)

	// validation already done in msg server
	a, err := sdk.AccAddressFromBech32(user.UserAccount)
	if err != nil {
		return err
	}

	acc := k.AccountKeeper.GetAccount(ctx, a)
	if acc == nil {
		acc = k.AccountKeeper.NewAccountWithAddress(ctx, a)
		k.AccountKeeper.SetAccount(ctx, acc)
	}
	return nil
}

func (k Keeper) GetUserByAccAddress(ctx sdk.Context, queriedUserAccAddress sdk.AccAddress) types.User {
	store := ctx.KVStore(k.storeKey)

	var queriedUser types.User

	iterator := sdk.KVStorePrefixIterator(store, []byte(types.UserPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var user types.User
		k.cdc.MustUnmarshalBinaryBare(store.Get(iterator.Key()), &user)
		if user.UserAccount == queriedUserAccAddress.String() {
			queriedUser = user
		}
	}
	return queriedUser
}

func (k Keeper) ListUsers(ctx sdk.Context) []*types.User {
	var userList []*types.User
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.UserPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var user types.User
		k.cdc.MustUnmarshalBinaryBare(store.Get(iterator.Key()), &user)
		userList = append(userList, &user)
	}
	return userList
}
