package utils

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var ErrParsEscrowEmoji  = sdkerrors.Register("escrow", 106, "parse emoji failed")
