package pooltoy

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"
	"github.com/interchainberlin/pooltoy/x/pooltoy/keeper"
	"github.com/interchainberlin/pooltoy/x/pooltoy/types"
	// abci "github.com/tendermint/tendermint/abci/types"
)

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) {
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
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	allUsersRaw, err := k.ListUsers(ctx)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	var allUsers []types.User
	k.Cdc.MustUnmarshalJSON(allUsersRaw, &allUsers)

	// chain got into a state such that there were accounts with no corresponding users.
	// ideally this would never happen but as it is currently the case these users need to be manually created.
	allAccounts := k.ListAccounts(ctx)
	if len(allAccounts) != len(allUsers) {
		for _, acct := range allAccounts {
			found := false
			for _, user := range allUsers {
				if user.UserAccount.Equals(acct.GetAddress()) {
					found = true
				}
			}
			if !found {
				var newUser = types.User{
					UserAccount: acct.GetAddress(),
					IsAdmin:     false,
					ID:          uuid.New().String(),
					Name:        "",
					Email:       "",
				}
				allUsers = append(allUsers, newUser)
			}
		}
	}

	return types.NewGenesisState(allUsers)
}

// "gov": {
// 	"starting_proposal_id": 1,
// 	"deposits": [],
// 	"votes": [],
// 	"proposals": [],
// 	"deposit_params": {
// 		"min_deposit": [],
// 		"max_deposit_period": 86400000000000
// 	},
// 	"voting_params": {
// 		"voting_period": 86400000000000
// 	},
// 	"tally_params": {
// 		"quorum": {
// 			"int": ""
// 		},
// 		"threshold": ,
// 		"veto":
// 	}

// },
