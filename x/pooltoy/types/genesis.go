package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenesisState - all pooltoy state that must be provided at genesis
type GenesisState struct {
	Users []User
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState(users []User) GenesisState {
	return GenesisState{
		Users: users,
	}
}

func MakeAdmin() User {
	creator, _ := sdk.AccAddressFromBech32("cosmos1qd4gsa4mlnpzmv4zsf9ghdrsgkt5avs8zte65u")
	userAccount, _ := sdk.AccAddressFromBech32("cosmos1qd4gsa4mlnpzmv4zsf9ghdrsgkt5avs8zte65u")
	return User{
		Creator:     creator,
		UserAccount: userAccount,
		IsAdmin:     true,
		Name:        "Alice",
	}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() GenesisState {
	return GenesisState{
		Users: []User{
			MakeAdmin(),
		},
	}
}

// ValidateGenesis validates the pooltoy genesis parameters
func ValidateGenesis(data GenesisState) error {
	for _, user := range data.Users {
		if user.Creator.Empty() {
			return fmt.Errorf("invalid creator: %s", user.Creator.String())
		}
		if user.UserAccount.Empty() {
			return fmt.Errorf("invalid user: %s", user.UserAccount.String())
		}
	}
	return nil
}
