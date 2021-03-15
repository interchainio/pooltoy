package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/interchainberlin/pooltoy/x/pooltoy/types"
)

func (k Keeper) CreateUser(ctx sdk.Context, user types.User) {
	key := []byte(types.UserPrefix + user.Id)
	k.Set(ctx, key, []byte(types.UserPrefix), user, k.MarshalUser)
	a, err := sdk.AccAddressFromBech32(user.UserAccount)
	//TODO: error handling
	if err != nil {
		panic(err)
	}
	acc := k.AccountKeeper.GetAccount(ctx, a)
	if acc == nil {
		acc = k.AccountKeeper.NewAccountWithAddress(ctx, a)
		k.AccountKeeper.SetAccount(ctx, acc)
	}
}

func (k Keeper) GetUserByAccAddress(ctx sdk.Context, queriedUserAccAddress sdk.AccAddress) types.User {
	// store := ctx.KVStore(k.storeKey)

	// var queriedUser types.User

	// iterator := sdk.KVStorePrefixIterator(store, []byte(types.UserPrefix))
	// for ; iterator.Valid(); iterator.Next() {
	// 	var user types.User
	// 	k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &user)
	// 	if user.UserAccount.Equals(queriedUserAccAddress) {
	// 		queriedUser = user
	// 	}
	// }
	// return queriedUser

	// TODO
	return types.User{}
}

func (k Keeper) ListUsers(ctx sdk.Context) ([]byte, error) {
	// var userList []types.User
	// store := ctx.KVStore(k.storeKey)
	// iterator := sdk.KVStorePrefixIterator(store, []byte(types.UserPrefix))
	// for ; iterator.Valid(); iterator.Next() {
	// 	var user types.User
	// 	k.Cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &user)
	// 	userList = append(userList, user)
	// }
	// res := codec.MustMarshalJSONIndent(k.Cdc, userList)
	// return res, nil
	// TODO
	return []byte{}, nil
}

func (k Keeper) MarshalUser(value interface{}) []byte {
	identifier := value.(types.User)
	bytes, _ := k.Cdc.MarshalBinaryBare(&identifier)
	return bytes
}
