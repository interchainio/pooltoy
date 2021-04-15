package pooltoy

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"
	"github.com/interchainberlin/pooltoy/x/pooltoy/keeper"
	"github.com/interchainberlin/pooltoy/x/pooltoy/types"
)

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) {
	var oneAdmin = false
	for _, u := range data.Users {
		if u.IsAdmin {
			oneAdmin = true
		}
		k.InsertUser(ctx, u)
	}
	if !oneAdmin {
		a := types.MakeAdmin()
		k.InsertUser(ctx, *a)
	}
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	users := k.ListUsers(ctx)

	// chain got into a state such that there were accounts with no corresponding users.
	// ideally this would never happen but as it is currently the case these users need to be manually created.
	allAccounts := k.ListAccounts(ctx)
	if len(allAccounts) != len(users) {
		for _, acct := range allAccounts {
			found := false
			for _, u := range users {
				a := acct.GetAddress()
				if u.UserAccount == a.String() {
					found = true
				}
			}
			if !found {
				n := types.User{
					UserAccount: acct.GetAddress().String(),
					IsAdmin:     false,
					Id:          uuid.New().String(),
					Name:        "",
					Email:       "",
				}
				users = append(users, n)
			}
		}
	}

	return types.NewGenesisState(users)
}
