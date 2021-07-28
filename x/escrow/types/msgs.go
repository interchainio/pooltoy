package types

import (
	fmt "fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = &Offer{}
)

const (
	TypeOffer = "offer"
)

// NewMsgMint is a constructor function for NewMsgMint
func NewOfferRequest(sender sdk.AccAddress, amount string, request string) *Offer {
	amt,_ := sdk.ParseCoinsNormalized(amount)
	req,_ := sdk.ParseCoinsNormalized(request)
	return &Offer{Sender: sender.String(), Amount: amt, Request: req}
}

// Route should return the name of the module
func (msg *Offer) Route() string { return RouterKey }

// Type should return the action
func (msg *Offer) Type() string { return TypeOffer}

// ValidateBasic runs stateless checks on the message
func (msg *Offer) ValidateBasic() error {
	//addr, err := sdk.AccAddressFromBech32(msg.Sender)
	//fmt.Println("validation basic!!!!", addr)
	//if err != nil {
	//	return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	//}
	//// todo add more validation

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *Offer) GetSignBytes() []byte {
	panic("amino support disabled")
}

// GetSigners defines whose signature is required
func (msg *Offer) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}
	return []sdk.AccAddress{sdk.AccAddress(sender)}
}
