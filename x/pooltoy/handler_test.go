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
	handler := NewHandler(keeper)

	//module github.com/cosmos/cosmos-sdk@latest found (v0.39.2), but does not contain package github.com/cosmos/cosmos-sdk/testutil/testdata
	//Not sure what I am missing here
	//msg := testdata.NewTestMsg()

	//Test for first user creation when creator is not an admin
	TestCreatorAddress := GenerateTestAddress(1)
	TestUserAddress1 := GenerateTestAddress(2)
	testMsg := types.NewMsgCreateUser(TestCreatorAddress, TestUserAddress1, false, "test_one", "test1@test.com")
	_, err := handler(ctx, testMsg)
	require.Error(t, err)

	creator := keeper.GetUserByAccAddress(ctx, TestCreatorAddress)
	testErr := sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("user %s (%s) is not an admin", creator.Name, TestCreatorAddress))
	require.True(t, errors.Is(err, testErr))

	//Test for first  user creation
	testMsg1 := types.NewMsgCreateUser(TestCreatorAddress, TestUserAddress1, true, "test_one", "test1@test.com")
	_, err1 := handler(ctx, testMsg1)
	require.NoError(t, err1)

	//Test for Creator who isn't an admin
	TestUserAddress2 := GenerateTestAddress(3)
	testMsg2 := types.NewMsgCreateUser(TestCreatorAddress, TestUserAddress2, false, "test_two", "test2@test.com")
	_, err2 := handler(ctx, testMsg2)
	require.Error(t, err2)

	creator = keeper.GetUserByAccAddress(ctx, TestCreatorAddress)
	testErr1 := sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("user %s (%s) is not an admin", creator.Name, TestCreatorAddress))
	require.True(t, errors.Is(err2, testErr1))

	//Test for non-existent Creator
	testMsg3 := types.NewMsgCreateUser(TestCreatorAddress, TestUserAddress1, true, "test_three", "test3@test.com")
	_, err3 := handler(ctx, testMsg3)
	require.Error(t, err3)

	testErr2 := sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("user %s does not exist", TestCreatorAddress))
	require.True(t, errors.Is(err3, testErr2))

	//Test for duplicate user TestUserAddress1 where creator is the user itself
	testMsg4 := types.NewMsgCreateUser(TestUserAddress1, TestUserAddress1, true, "test_one", "test1@test.com")
	_, err4 := handler(ctx, testMsg4)
	require.Error(t, err4)

	testErr3 := sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("user %s already exists", TestUserAddress1))
	require.True(t, errors.Is(err4, testErr3))

}
