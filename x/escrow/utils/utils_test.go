package utils

import (
	"errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"

)

func TestParseCoins_0(t *testing.T) {
	coins := "1🆗200🍍"
	res, err := ParseCoins(coins)
	if err != nil {
		t.Fatal(err)
	}

	if r :=res[0].Amount.BigInt().Cmp(sdk.NewInt(1).BigInt()); r!=0{
		t.Fatal(res[1].Amount.BigInt())
	}
	if res[0].Denom != "🆗"{
		t.Fatal(res[0].Denom)
	}

	if r :=res[1].Amount.BigInt().Cmp(sdk.NewInt(200).BigInt()); r!=0{
		t.Fatal(res[1].Amount.BigInt())
	}
	if res[1].Denom != "🍍"{
		t.Fatal(res[1].Denom)
	}
}


func TestParseCoins_1(t *testing.T) {
	coins := "🆗200🍍"
	res, err := ParseCoins(coins)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res.String())

	if r :=res[0].Amount.BigInt().Cmp(sdk.NewInt(1).BigInt()); r!=0{
		t.Fatal(res[1].Amount.BigInt())
	}
	if res[0].Denom != "🆗"{
		t.Fatal(res[0].Denom)
	}

	if r :=res[1].Amount.BigInt().Cmp(sdk.NewInt(200).BigInt()); r!=0{
		t.Fatal(res[1].Amount.BigInt())
	}
	if res[1].Denom != "🍍"{
		t.Fatal(res[1].Denom)
	}
}

func TestParseCoins_2(t *testing.T) {
	coins := "🆗🍍"
	res, err := ParseCoins(coins)
	if err != nil {
		t.Fatal(err)
	}

	if r :=res[0].Amount.BigInt().Cmp(sdk.NewInt(1).BigInt()); r!=0{
		t.Fatal(res[1].Amount.BigInt())
	}
	if res[0].Denom != "🆗"{
		t.Fatal(res[0].Denom)
	}

	if r :=res[1].Amount.BigInt().Cmp(sdk.NewInt(1).BigInt()); r!=0{
		t.Fatal(res[1].Amount.BigInt())
	}
	if res[1].Denom != "🍍"{
		t.Fatal(res[1].Denom)
	}
}

// -----------------------------------------------------------------------------
//fail cases
// -----------------------------------------------------------------------------
func TestParseCoins_3(t *testing.T) {
	coins := "1🆗0🍍"
	_, err := ParseCoins(coins)
	if !errors.Is(err, ErrParsEscrowEmoji) {
		t.Fatal(err)
	}
}

func TestParseCoins_4(t *testing.T) {
	coins := "0🆗0🍍"
	_, err := ParseCoins(coins)
	if !errors.Is(err, ErrParsEscrowEmoji) {
		t.Fatal(err)
	}
}

