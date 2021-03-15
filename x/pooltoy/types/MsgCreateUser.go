package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// msg types
const (
	TypeMsgCreateUser = "create-user"
)

var _ sdk.Msg = &MsgCreateUser{}

func NewMsgCreateUser(
	usr User,
	creator string,
) *MsgCreateUser {
	return &MsgCreateUser{
		User:    &usr,
		Creator: creator,
	}
}

// Route implements sdk.Msg
func (MsgCreateUser) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgCreateUser) Type() string {
	return TypeMsgCreateUser
}

func (msg MsgCreateUser) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateUser) GetSignBytes() []byte {
	panic("pooltoy has deprecated amino")
}

func (msg MsgCreateUser) ValidateBasic() error {
	if msg.Creator == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	if (User{} == *msg.User) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "UserAccount can't be empty")
	}
	return nil
}
