package utils

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/interchainberlin/pooltoy/x/faucet/utils"
	"strconv"
	"strings"
	"unicode"
)

// -----------------------------------------------------------------------------
//
// -----------------------------------------------------------------------------

// find the max mactch, e.g. str =  "\U0001f1e6\U0001f1fd\U0001f6e9", \U0001f1e6\U0001f1fd\" is an emoji, "\U0001f1e6" is also an emoji, it should find "\U0001f1e6\U0001f1fd".
// maxlen is the max length of a single emoji's runes, here is 5 according to the emojiCodeMap.
// -----------------------------------------------------------------------------

func GetCoinsWithCheck(emostr string) (coins sdk.Coins, err error) {
	nums, emos, err := ParseEmojiandNum(emostr)
	if err != nil {
		return sdk.Coins{}, err
	}

	coins, err = GetCoins(nums, emos)
	if err != nil {
		return sdk.Coins{}, err
	}
	emptyCoins := sdk.Coins{}
	if coins.IsEqual(emptyCoins) { // todo do we allow empty coins?
		return coins, nil
	}
	if !coins.IsAllPositive() {
		return sdk.Coins{}, ErrEmojiCoinsCheck
	}
	//if err := coins.Validate(); err != nil {
	//	fmt.Println("validation failed", err)
	//	return sdk.Coins{}, err
	//}

	//not allow the same emoji appear twice ⭐⭐, one should write 2⭐, this is for distinguish accidental input
	memo := map[string]bool{}
	for _, coin := range coins {
		_, found := memo[coin.Denom]
		if found {
			return sdk.Coins{}, ErrEmojiCoinsCheck
		} else {
			memo[coin.Denom] = true
		}
	}

	return coins, nil
}

func GetCoins(nums, emos []string) (coins sdk.Coins, err error) {
	if len(nums) != len(emos) {
		return sdk.Coins{}, ErrParseEmojiToCoins
	}
	coin := sdk.Coin{}
	j := 0
	for _, emo := range emos {
		emolist, err := ParseEmo(emo, 5)
		if err != nil {
			return sdk.Coins{}, ErrParseEmojiToCoins
		}

		for i := range emolist {
			if i == 0 {
				n, err := strconv.Atoi(nums[j])
				if err != nil {
					return sdk.Coins{}, ErrParseEmojiToCoins
				}
				coin.Amount = sdk.NewInt(int64(n))
				j++
			} else {
				coin.Amount = sdk.NewInt(1)
			}

			coin.Denom = emolist[i]
			coins = append(coins, coin)
		}
	}

	return coins, nil
}

// parse the string, to get list of nums and emojis
func ParseEmojiandNum(str string) (nums []string, emos []string, err error) {
	numBuf := strings.Builder{}
	emojiBuf := strings.Builder{}
	writebuf := true

	for _, r := range []rune(str) {

		ok := unicode.IsDigit(r)

		switch writebuf {
		case true:
			if ok {
				numBuf.WriteRune(r)
			} else {
				emojiBuf.WriteRune(r)
				writebuf = false
			}

		case false:
			if ok {
				num, emo, ok := FlushBuf(&numBuf, &emojiBuf)
				if !ok {
					return nil, nil, ErrParseEmojiandNum
				}
				nums = append(nums, num)
				emos = append(emos, emo)
				numBuf.WriteRune(r)
				writebuf = true
			} else {
				emojiBuf.WriteRune(r)
			}
		}
	}

	num, emo, ok := FlushBuf(&numBuf, &emojiBuf)
	if !ok {
		return nil, nil, ErrParseEmojiandNum
	}
	nums = append(nums, num)
	emos = append(emos, emo)

	return nums, emos, nil
}

func FlushBuf(numBuf, emojiBuf *strings.Builder) (num string, emo string, ok bool) {

	if numBuf.String() == "" && emojiBuf.String() == "" {
		return "", "", false
	}

	if numBuf.String() != "" && emojiBuf.String() == "" {
		return "", "", false
	}

	if numBuf.String() == "" && emojiBuf.String() != "" {
		emo = emojiBuf.String()
		num = "1"
		numBuf.Reset()
		emojiBuf.Reset()
	}

	if numBuf.String() != "" && emojiBuf.String() != "" {
		num = numBuf.String()
		emo = emojiBuf.String()
		numBuf.Reset()
		emojiBuf.Reset()
	}

	return num, emo, true
}

// parse emojis string which does not contain numbers
func ParseEmo(emo string, emojiMaxLen int) (emolist []string, err error) {
	// use rune rather than string to avoid index jumping
	r := []rune(emo)
	emojimap := utils.ReverseMapKV(utils.EmojiCodeMap)

	i := 0
	for i < len(r) {
		maxLen := emojiMaxLen
		if i+emojiMaxLen >= len(r) {
			maxLen = len(r) - i
		}

		hasEmoji := true

		for j := maxLen; j > 0; j-- {
			_, found := emojimap[string(r[i:i+j])]
			if found {
				emolist = append(emolist, string(r[i:i+j]))
				i += j
				hasEmoji = true
				break
			}
			hasEmoji = false
		}

		if hasEmoji == false {
			return []string{}, ErrEmojiStr
		}
	}

	return emolist, nil
}
