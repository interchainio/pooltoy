package utils

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"
)

func TestParseCoins_Cases_Pass(t *testing.T) {
	type ParseCase struct {
		EmoStr string
		Coins  sdk.Coins
	}

	samples := []ParseCase{
		{
			EmoStr: "1🆗",
			Coins: []sdk.Coin{
				{Denom: "🆗", Amount: sdk.NewInt(1)},
			},
		},
		{
			EmoStr: "🆗🍍",
			Coins: []sdk.Coin{
				{Denom: "🆗", Amount: sdk.NewInt(1)},
				{Denom: "🍍", Amount: sdk.NewInt(1)},
			},
		},
		{
			EmoStr: "12🆗🍍",
			Coins: []sdk.Coin{
				{Denom: "🆗", Amount: sdk.NewInt(12)},
				{Denom: "🍍", Amount: sdk.NewInt(1)},
			},
		},
		{
			EmoStr: "12" + "\U0001f1e6\U0001f1f2",
			Coins: []sdk.Coin{
				{Denom: "\U0001f1e6\U0001f1f2", Amount: sdk.NewInt(12)},
			},
		},

		{
			EmoStr: "10\U0001f9d1\u200d\U0001f3a8" +
				"40\U0001f632" + "\U0001f1e6\U0001f1fa" + "\U0001f3a8" + "2\U0001f9d6\U0001f3fb\u200d\u2640\ufe0f",
			Coins: []sdk.Coin{
				{Denom: "\U0001f9d1\u200d\U0001f3a8", Amount: sdk.NewInt(10)},
				{Denom: "\U0001f632", Amount: sdk.NewInt(40)},
				{Denom: "\U0001f1e6\U0001f1fa", Amount: sdk.NewInt(1)},
				{Denom: "\U0001f3a8", Amount: sdk.NewInt(1)},
				{Denom: "\U0001f9d6\U0001f3fb\u200d\u2640\ufe0f", Amount: sdk.NewInt(2)},
			},
		},
	}

	for i, sample := range samples {
		coins, err := GetCoinsWithCheck(sample.EmoStr)
		if err != nil {
			t.Fatalf("failed case %d", i)
		}

		if !coins.IsEqual(sample.Coins) {
			t.Fatalf("failed case %d: %s", i, coins.String())
		}
	}
}

func TestParseCoins_Cases_Fail(t *testing.T) {
	str := []string{
		"",
		"-123🆗",    // negative
		"0🆗",       // zero
		"1e10🆗",    // scientific notation
		"1.2🆗",     // decimal
		"1,223🆗",   // formatted
		"1,223",     // no emoji
		"1🆗🆗",     // duplicate emoji
		"1🆗2🍍3🆗", // duplicate emoji
		//	"00002🆗⭐", // leading zero  //todo if allow ???
		"01 🆗", //space
	}

	for i, s := range str {
		_, err := GetCoinsWithCheck(s)
		if err == nil {
			t.Fatalf("failed case %d", i)
		}
	}
}
