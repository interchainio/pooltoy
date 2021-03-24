package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	proto "github.com/gogo/protobuf/proto"
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgMint{},
	)

	registry.RegisterInterface(
		"faucet.MintHistory",
		(*MintHistory)(nil),
		&MintHistory{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// PackMintHistory constructs a new Any packed with the given History value. It returns
// an error if the history data can't be casted to a protobuf message or if the concrete
// implemention is not registered to the protobuf codec.
func PackMintHistory(h MintHistory) (*cdctypes.Any, error) {
	msg, ok := h.(proto.Message)
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrPackAny, "cannot proto marshal %T", h)
	}

	anyHistory, err := cdctypes.NewAnyWithValue(msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPackAny, err.Error())
	}

	return anyHistory, nil
}

// UnpackHistory Unpacks an Any into a History. It returns an error if the
// client state can't be Unpacked into a History.
func UnpackMintHistory(any *cdctypes.Any) (MintHistory, error) {
	if any == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnpackAny, "protobuf Any message cannot be nil")
	}

	h, ok := any.GetCachedValue().(MintHistory)
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnpackAny, "cannot Unpack Any into MintHistory %T", any)
	}

	return h, nil
}
