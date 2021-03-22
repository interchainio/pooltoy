package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/interchainberlin/pooltoy/x/pooltoy/types"
)

func (k Keeper) InsertUser(ctx sdk.Context, user types.User) {
	key := []byte(types.UserPrefix + user.Id)
	k.Set(ctx, key, []byte(types.UserPrefix), user, k.MarshalUser)
	// validation already done in msg server
	a, _ := sdk.AccAddressFromBech32(user.UserAccount)
	acc := k.AccountKeeper.GetAccount(ctx, a)
	if acc == nil {
		acc = k.AccountKeeper.NewAccountWithAddress(ctx, a)
		k.AccountKeeper.SetAccount(ctx, acc)
	}
}

func (k Keeper) GetUserByAccAddress(ctx sdk.Context, queriedUserAccAddress sdk.AccAddress) types.User {
	store := ctx.KVStore(k.storeKey)

	var queriedUser types.User

	iterator := sdk.KVStorePrefixIterator(store, []byte(types.UserPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var user types.User
		// TODO: check unmarshaler, MustUnmarshalBinaryBare?
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

func (k Keeper) MarshalUser(value interface{}) []byte {
	var identifier types.User
	bytes, _ := k.cdc.MarshalBinaryBare(&identifier)
	return bytes
}
