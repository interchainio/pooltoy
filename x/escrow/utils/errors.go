package utils

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var ErrParseEmojiandNum  = sdkerrors.Register("escrow", 106, "parse emoji string into numbers and emojis failed")
var ErrParseEmojiToCoins  = sdkerrors.Register("escrow", 107, "parse emoji string into coins failed")

var ErrEmojiCoinsCheck  = sdkerrors.Register("escrow", 108, "emoji coins check failed")

var ErrEmojiStr  = sdkerrors.Register("escrow", 109, "parse emoji string into emoji failed")
