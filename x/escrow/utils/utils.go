package utils

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/interchainberlin/pooltoy/x/faucet/utils"
	"strconv"
)

func ParseCoins(coins string) (sdk.Coins, error) {
	sdkcoins := sdk.Coins{}
	coin := sdk.Coin{}
	emojiMap := utils.ReverseMapKV(utils.EmojiCodeMap)
	num := ""

	// cannot be all num
	_, err := strconv.Atoi(coins)
	if err == nil {
		return sdk.Coins{}, ErrParsEscrowEmoji
	}

	for _, v := range coins {
		_, err := strconv.Atoi(string(v))
		_, found := emojiMap[string(v)]

		// not emoji, not number
		if err != nil && !found {
			return sdk.Coins{}, ErrParsEscrowEmoji
		}

		// is num
		if err == nil {
			num += string(v)
		}

		// is emoji
		if found {

			var bigAmt sdk.Int
			if num == "" {
				bigAmt = sdk.NewInt(1)
			} else {
				amt, err := strconv.Atoi(num)
				if err == nil && amt != 0 {
					bigAmt = sdk.NewInt(int64(amt))
				} else {
					return sdk.Coins{}, ErrParsEscrowEmoji

				}
			}

			coin.Amount = bigAmt
			coin.Denom = string(v)

			sdkcoins = append(sdkcoins, coin)
			coin = sdk.Coin{}
			num = ""
		}

	}

	return sdkcoins, nil

}
