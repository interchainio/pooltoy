package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenesisState - all pooltoy state that must be provided at genesis
type GenesisState struct {
	User User
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState(user User) GenesisState {
	return GenesisState{
		User: user,
	}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() GenesisState {
	creator, _ := sdk.AccAddressFromBech32("cosmos1qd4gsa4mlnpzmv4zsf9ghdrsgkt5avs8zte65u")
	userAccount, _ := sdk.AccAddressFromBech32("cosmos1qd4gsa4mlnpzmv4zsf9ghdrsgkt5avs8zte65u")
	return GenesisState{
		User: User{
			Creator:     creator,
			UserAccount: userAccount,
			IsAdmin:     true,
			Name:        "Alice",
		},
	}
}

// ValidateGenesis validates the pooltoy genesis parameters
func ValidateGenesis(data GenesisState) error {
	if data.User.Creator.Empty() {
		return fmt.Errorf("invalid creator: %s", data.User.Creator.String())
	}

	if data.User.UserAccount.Empty() {
		return fmt.Errorf("invalid user: %s", data.User.UserAccount.String())
	}

	if !data.User.IsAdmin {
		return fmt.Errorf("user: %s not an admin", data.User.UserAccount.String())
	}
	return nil
}
