package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/interchainberlin/pooltoy/x/pooltoy/types"
)

func (k Keeper) CreateUser(ctx sdk.Context, user types.User) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.UserPrefix + user.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(user)
	store.Set(key, value)
}

func listUsers(ctx sdk.Context, k Keeper) ([]byte, error) {
	var userList []types.User
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.UserPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var user types.User
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &user)
		userList = append(userList, user)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, userList)
	return res, nil
}
