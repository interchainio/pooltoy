package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewUser(
	userAccount string,
	isAdmin bool,
	name string,
	email string,
) User {
	return User{
		Id:          "",
		UserAccount: userAccount,
		IsAdmin:     isAdmin,
		Name:        name,
		Email:       email,
	}
}

// ValidateBasic runs stateless checks on the message
// TODO: more complex validation?
func (u User) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(u.Creator)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, u.Creator)
	}

	_, err = sdk.AccAddressFromBech32(u.UserAccount)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, u.UserAccount)
	}

	return nil
}

//Marshal
//Unmarshal
