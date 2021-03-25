package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMintHistory = "mint-history"
)

// NewMining returns a new Mining Message
func NewMintHistory(minter sdk.AccAddress, tally int64) *MintHistory {
	return &MintHistory{
		Minter:   minter.String(),
		Lasttime: 0,
		Tally:    tally,
	}
}

// Type should return the action
func (m MintHistory) Type() string { return TypeMintHistory }

// ValidateBasic runs stateless checks on the message
// TODO: more complex validation?
func (m MintHistory) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Minter)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, m.Minter)
	}
	return nil
}

// GetSigners defines whose signature is required
func (m MintHistory) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Minter)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

//Marshal
//Unmarshal
