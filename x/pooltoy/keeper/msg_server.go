package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/interchainio/pooltoy/x/pooltoy/types"
)

var _ types.MsgServer = Keeper{}

func (k Keeper) CreateUser(c context.Context, msg *types.MsgCreateUser) (*types.MsgCreateUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	u := *msg.User
	allUsers := k.ListUsers(ctx)
	ca, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		// TODO: unify error handling
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("msg.Creator has an invalid address format %s", msg.Creator))
	}

	fmt.Println("allUsers", len(allUsers), allUsers)
	if len(allUsers) == 0 {
		if err := k.InsertUser(ctx, u); err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, fmt.Sprintf("error marshaling u.UserAccount %s", u.UserAccount))
		}
		return &types.MsgCreateUserResponse{}, nil
	}

	creator := k.GetUserByAccAddress(ctx, ca)
	fmt.Println("creator", creator)
	if creator.UserAccount == "" && len(allUsers) != 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("invalid address format for creator.UserAccount %s", creator.UserAccount))
	}

	//  validate address format
	a, err := sdk.AccAddressFromBech32(u.UserAccount)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address format for u.UserAccount %s", u.UserAccount))
	}
	//  check that new user doesn't exist already
	if existingUser := k.GetUserByAccAddress(ctx, a); existingUser.UserAccount == u.UserAccount {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("u.UserAccount %s already exists", u.UserAccount))
	}

	// creator must be an admin
	if creator.IsAdmin {
		// if yes
		if err := k.InsertUser(ctx, u); err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, fmt.Sprintf("error marshaling u.UserAccount %s", u.UserAccount))
		}
	} else {
		// if no
		// throw error
		errMsg := fmt.Sprintf("user %s (%s) is not an admin", creator.Name, msg.Creator)
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, errMsg)
	}

	return &types.MsgCreateUserResponse{}, nil
}
