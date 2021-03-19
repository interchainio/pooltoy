package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/interchainberlin/pooltoy/x/pooltoy/types"
)

var _ types.MsgServer = Keeper{}

func (k Keeper) CreateUser(c context.Context, msg *types.MsgCreateUser) (*types.MsgCreateUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	u := *msg.User
	allUsers := k.ListUsers(ctx)

	// does this creator have permission to create this new user?
	// bear in mind special case allows create as initialization when there are no users yet
	ca, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		// TODO: unify error handling
		errMsg := fmt.Sprintf("invalid address format", msg.Creator)
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, errMsg)
	}
	creator := k.GetUserByAccAddress(ctx, ca)
	if creator.UserAccount == "" && len(allUsers) != 0 {
		errMsg := fmt.Sprintf("user %s does not exist", msg.Creator)
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, errMsg)
	}

	//  validate address format
	a, err := sdk.AccAddressFromBech32(u.UserAccount)
	if err != nil {
		errMsg := fmt.Sprintf("invalid address format", u.UserAccount)
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, errMsg)
	}
	//  check that new user doesn't exist already
	if existingUser := k.GetUserByAccAddress(ctx, a); existingUser.UserAccount == u.UserAccount {
		errMsg := fmt.Sprintf("user %s already exists", u.UserAccount)
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, errMsg)
	}

	// creator must be an admin
	if creator.IsAdmin || (u.IsAdmin && len(allUsers) == 0) {
		// if yes
		k.InsertUser(ctx, u)
	} else {
		// if no
		// throw error
		errMsg := fmt.Sprintf("user %s (%s) is not an admin", creator.Name, msg.Creator)
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, errMsg)
	}

	return &types.MsgCreateUserResponse{}, nil
}
