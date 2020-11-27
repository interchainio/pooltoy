package pooltoy

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	//"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/interchainberlin/pooltoy/x/pooltoy/keeper"
	"github.com/interchainberlin/pooltoy/x/pooltoy/types"
	"github.com/stretchr/testify/require"
)

func GenerateTestAddress(n byte) sdk.AccAddress {
	return bytes.Repeat([]byte{n}, sdk.AddrLen)
}

func TestBasicMsg(t *testing.T) {
	keeper, ctx, _, _ := keeper.CreateTestKeepers(t)
	//msg := testdata.NewTestMsg()

	TestCreatorAddress := GenerateTestAddress(1)
	TestUserAddress := GenerateTestAddress(2)
	testMsg1 := types.NewMsgCreateUser(TestCreatorAddress, TestUserAddress, true, "test", "test@test.com")
	handler1 := NewHandler(keeper)
	_, err1 := handler1(ctx, testMsg1)
	require.NoError(t, err1)

	//TestCreator doesn't exist
	testMsg2 := types.NewMsgCreateUser(TestCreatorAddress, TestUserAddress, true, "test", "test@test.com")
	handler2 := NewHandler(keeper)
	_, err2 := handler2(ctx, testMsg2)
	errMsg := fmt.Sprintf("user %s does not exist", TestCreatorAddress)
	testErr2 := sdkerrors.Wrap(sdkerrors.ErrUnauthorized, errMsg)
	require.Error(t, err2)
	require.True(t, errors.Is(err2, testErr2))
}
