package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateUser{}

type MsgCreateUser struct {
	ID          string
	Creator     sdk.AccAddress `json:"creator" yaml:"creator"`
	UserAccount sdk.AccAddress `json:"userAccount" yaml:"userAccount"`
	IsAdmin     bool           `json:"isAdmin" yaml:"isAdmin"`
	Name        string         `json:"name" yaml:"name"`
	Email       string         `json:"email" yaml:"email"`
}

func NewMsgCreateUser(creator sdk.AccAddress, userAccount sdk.AccAddress, isAdmin bool, name string, email string) MsgCreateUser {
	return MsgCreateUser{
		ID:          uuid.New().String(),
		Creator:     creator,
		UserAccount: userAccount,
		IsAdmin:     isAdmin,
		Name:        name,
		Email:       email,
	}
}

func (msg MsgCreateUser) Route() string {
	return RouterKey
}

func (msg MsgCreateUser) Type() string {
	return "CreateUser"
}

func (msg MsgCreateUser) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateUser) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreateUser) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	if msg.UserAccount.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "UserAccount can't be empty")
	}
	if msg.Name == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Name can't be empty")
	}
	return nil
}
