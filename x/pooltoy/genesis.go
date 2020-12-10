package pooltoy

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/interchainberlin/pooltoy/x/pooltoy/types"
	// abci "github.com/tendermint/tendermint/abci/types"
)

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
func InitGenesis(ctx sdk.Context, k Keeper, data GenesisState) {
	var oneAdmin = false
	for _, user := range data.Users {
		if user.IsAdmin {
			oneAdmin = true
		}
		k.CreateUser(ctx, user)
	}
	if !oneAdmin {
		k.CreateUser(ctx, types.MakeAdmin())
	}
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	allUsersRaw, err := k.ListUsers(ctx)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	var allUsers []types.User
	k.Cdc.MustUnmarshalJSON(allUsersRaw, &allUsers)

	return NewGenesisState(allUsers)
}
