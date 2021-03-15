package types

// import (
// 	"fmt"

// 	sdk "github.com/cosmos/cosmos-sdk/types"
// )

// // NewGenesisState creates a new GenesisState object
// func NewGenesisState(users []*User) GenesisState {
// 	return GenesisState{
// 		users,
// 	}
// }

// func MakeAdmin() *User {
// 	c, _ := sdk.AccAddressFromBech32("cosmos1qd4gsa4mlnpzmv4zsf9ghdrsgkt5avs8zte65u")
// 	u, _ := sdk.AccAddressFromBech32("cosmos1qd4gsa4mlnpzmv4zsf9ghdrsgkt5avs8zte65u")

// 	return &User{
// 		Creator:     c.String(),
// 		UserAccount: u.String(),
// 		IsAdmin:     true,
// 		Name:        "Alice",
// 	}
// }

// // DefaultGenesisState - default GenesisState used by Cosmos Hub
// func DefaultGenesisState() GenesisState {
// 	a := MakeAdmin()
// 	u := []*User{a}
// 	return GenesisState{u}
// }

// // ValidateGenesis validates the pooltoy genesis parameters
// func ValidateGenesis(data GenesisState) error {
// 	for _, user := range data.User {
// 		if user.Creator == "" {
// 			return fmt.Errorf("invalid creator: %s", user.Creator)
// 		}
// 		if user.UserAccount == "" {
// 			return fmt.Errorf("invalid user: %s", user.UserAccount)
// 		}
// 	}
// 	return nil
// }
