package faucet

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/interchainberlin/pooltoy/x/faucet/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto"
	emoji "github.com/tmdvs/Go-Emoji-Utils"
)

// TODO: rewrite test
func TestEmoji(t *testing.T) {
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte("foo")))
	moduleAcct2 := sdk.AccAddress(crypto.AddressHash([]byte("bar")))
	denom := "🥵"
	msg := types.NewMsgMint(moduleAcct, moduleAcct2, denom)

	err := msg.ValidateBasic()
	require.NoError(t, err)

	results := emoji.FindAll(msg.Denom)
	if len(results) != 1 {
		fmt.Println("results did not equal 1")
		require.True(t, false)
	}
	emo, ok := results[0].Match.(emoji.Emoji)
	if !ok {
		fmt.Println("Not correct interface for Emoji")
		require.True(t, false)
	}
	msg.Denom = emo.Value

	fmt.Println("final msg.Denom", msg.Denom)
	// require.True(t, false)
}
