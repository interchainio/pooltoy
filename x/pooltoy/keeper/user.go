package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/interchainberlin/pooltoy/x/pooltoy/types"
)

func (k Keeper) InsertUser(ctx sdk.Context, user types.User) error {
	key := []byte(types.UserPrefix + user.Id)
	u, err := k.cdc.MarshalBinaryBare(&user)
	if err != nil {
		return err
	}

	k.SetUser(ctx, key, []byte(types.UserPrefix), u)
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
		var user *types.User
		k.cdc.MustUnmarshalBinaryBare(store.Get(iterator.Key()), user)
		userList = append(userList, user)
	}
	return userList
}
