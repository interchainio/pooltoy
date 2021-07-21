package faucet

import (
	"fmt"
	"github.com/interchainberlin/pooltoy/x/faucet/utils"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/interchainberlin/pooltoy/x/faucet/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto"
)

// TODO: rewrite test
func TestEmoji(t *testing.T) {
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte("foo")))
	moduleAcct2 := sdk.AccAddress(crypto.AddressHash([]byte("bar")))
	denom := "🥵"
	msg := types.NewMsgMint(moduleAcct, moduleAcct2, denom)

	err := msg.ValidateBasic()
	require.NoError(t, err)

	emo, err := utils.ParseEmoji(msg.Denom)
	if err!= nil{
		fmt.Println("Not correct interface for Emoji")
	}
	msg.Denom = emo

	fmt.Println("final msg.Denom", msg.Denom)
	// require.True(t, false)
}
