package pooltoy

import (
	"bytes"
	"testing"

	//"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/interchainberlin/pooltoy/x/pooltoy/keeper"
	"github.com/interchainberlin/pooltoy/x/pooltoy/types"
	"github.com/stretchr/testify/require"
)

var (
	TestUserAddress    sdk.AccAddress = bytes.Repeat([]byte{1}, sdk.AddrLen)
	TestCreatorAddress sdk.AccAddress = bytes.Repeat([]byte{2}, sdk.AddrLen)
)

func TestBasicMsg(t *testing.T) {
	keeper, ctx, _, _ := keeper.CreateTestKeepers(t)
	//msg := testdata.NewTestMsg()
	msg := types.NewMsgCreateUser(TestCreatorAddress, TestUserAddress, true, "test", "test@test.com")
	handler := NewHandler(keeper)

	_, err := handler(ctx, msg)
	require.NoError(t, err)

}
