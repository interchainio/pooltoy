package utils

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var ErrParseEmoji  = sdkerrors.Register("faucet", 105, "parse emoji failed")

