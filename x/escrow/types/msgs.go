package types

import (
	fmt "fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = &OfferRequest{}
)

const (
	TypeOffer = "offer"
)

// NewMsgMint is a constructor function for NewMsgMint
func NewOfferRequest(sender sdk.AccAddress, amount string, request string) *OfferRequest {
	return &OfferRequest{Sender: sender.String(), Amount: amount, Request: request}
}

// Route should return the name of the module
func (msg *OfferRequest) Route() string { return RouterKey }

// Type should return the action
func (msg *OfferRequest) Type() string { return TypeOffer}

// ValidateBasic runs stateless checks on the message
func (msg *OfferRequest) ValidateBasic() error {
	//addr, err := sdk.AccAddressFromBech32(msg.Sender)
	//fmt.Println("validation basic!!!!", addr)
	//if err != nil {
	//	return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	//}
	//// todo add more validation

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *OfferRequest) GetSignBytes() []byte {
	panic("amino support disabled")
}

// GetSigners defines whose signature is required
func (msg *OfferRequest) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}
	return []sdk.AccAddress{sdk.AccAddress(sender)}
}
