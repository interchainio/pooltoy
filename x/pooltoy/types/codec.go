package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	proto "github.com/gogo/protobuf/proto"
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgCreateUser{})
	registry.RegisterInterface(
		"pooltoy.User",
		(*User)(nil),
		&User{},
	)
}

// PackUser constructs a new Any packed with the given User value. It returns
// an error if the user data can't be casted to a protobuf message or if the concrete
// implemention is not registered to the protobuf codec.
func PackUser(user User) (*cdctypes.Any, error) {
	msg, ok := user.(proto.Message)
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrPackAny, "cannot proto marshal %T", user)
	}

	anyUser, err := cdctypes.NewAnyWithValue(msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPackAny, err.Error())
	}

	return anyUser, nil
}

// UnpackUser Unpacks an Any into a User. It returns an error if the
// client state can't be Unpacked into a User.
func UnpackUser(any *cdctypes.Any) (User, error) {
	if any == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnpackAny, "protobuf Any message cannot be nil")
	}

	user, ok := any.GetCachedValue().(User)
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnpackAny, "cannot Unpack Any into User %T", any)
	}

	return user, nil
}
