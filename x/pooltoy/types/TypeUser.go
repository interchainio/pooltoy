package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type User struct {
	Creator     sdk.AccAddress `json:"creator" yaml:"creator"`
	ID          string         `json:"id" yaml:"id"`
	UserAccount sdk.AccAddress `json:"userAccount" yaml:"userAccount"`
	IsAdmin     bool           `json:"isAdmin" yaml:"isAdmin"`
	Name        string         `json:"name" yaml:"name"`
	Email       string         `json:"email" yaml:"email"`
}
