package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"
)

func NewUser(
	userAccount string,
	isAdmin bool,
	name string,
	email string,
) User {
	return User{
		Id:          uuid.New().String(),
		UserAccount: userAccount,
		IsAdmin:     isAdmin,
		Name:        name,
		Email:       email,
	}
}

// GetBytes is a helper for serialising
func (usr User) GetBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&usr))
}
