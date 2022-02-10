package types

import (
	 "fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = &OfferRequest{}
	_ sdk.Msg = &ResponseRequest{}
	_ sdk.Msg = &CancelOfferRequest{}
)

const (
	TypeOfferRequest = "offer"
	TypeResponseRequest = "response"
	TypeCancelOfferRequest = "CancelOffer"
)

// NewMsgMint is a constructor function for NewMsgMint
func NewOfferRequest(sender sdk.AccAddress, amount string, request string) *OfferRequest {
	amt,_ := sdk.ParseCoinsNormalized(amount)
	req,_ := sdk.ParseCoinsNormalized(request)
	return &OfferRequest{Sender: sender.String(), Amount: amt, Request: req}
}

// Route should return the name of the module
func (msg *OfferRequest) Route() string { return RouterKey }

// Type should return the action
func (msg *OfferRequest) Type() string { return TypeOfferRequest}

// ValidateBasic runs stateless checks on the message
func (msg *OfferRequest) ValidateBasic() error {
	// todo add more validation

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

func NewResponseRequest(sender sdk.AccAddress, id int64) *ResponseRequest {
	return &ResponseRequest{Sender: sender.String(), Id: id}
}

// Route should return the name of the module
func (msg *ResponseRequest) Route() string { return RouterKey }

// Type should return the action
func (msg *ResponseRequest) Type() string { return TypeResponseRequest}

// ValidateBasic runs stateless checks on the message
func (msg *ResponseRequest) ValidateBasic() error {
	// todo add more validation
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *ResponseRequest) GetSignBytes() []byte {
	panic("amino support disabled")
}

// GetSigners defines whose signature is required
func (msg *ResponseRequest) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}
	return []sdk.AccAddress{sdk.AccAddress(sender)}
}



func NewCancelOfferRequest(addr sdk.Address, id int64) *CancelOfferRequest {
	return &CancelOfferRequest{Sender:addr.String(), Id: id}
}

// Route should return the name of the module
func (msg *CancelOfferRequest) Route() string { return RouterKey }

// Type should return the action
func (msg *CancelOfferRequest) Type() string { return TypeCancelOfferRequest}

// ValidateBasic runs stateless checks on the message
func (msg *CancelOfferRequest) ValidateBasic() error {

	// todo add more validation

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *CancelOfferRequest) GetSignBytes() []byte {
	panic("amino support disabled")
}

// GetSigners defines whose signature is required
func (msg *CancelOfferRequest) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}
	return []sdk.AccAddress{sdk.AccAddress(sender)}
}
