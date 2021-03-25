package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgCreateUser{}
)

const (
	TypeCreateUser = "create_user"
)

///////////////////////
// MsgCreateUser //
///////////////////////

// NewMsgCreateUser is a constructor function for MsgCreateUser
func NewMsgCreateUser(creator sdk.AccAddress, user User) *MsgCreateUser {
	return &MsgCreateUser{Creator: creator.String(), User: &user}
}

// Route should return the name of the module
func (msg *MsgCreateUser) Route() string { return RouterKey }

// Type should return the action
func (msg *MsgCreateUser) Type() string { return TypeCreateUser }

// ValidateBasic runs stateless checks on the message
func (msg *MsgCreateUser) ValidateBasic() error {
	_, err := sdk.ValAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgCreateUser) GetSignBytes() []byte {
	panic("amino support disabled")
}

// GetSigners defines whose signature is required
func (msg *MsgCreateUser) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}
