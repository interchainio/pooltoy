package pooltoy

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/interchainberlin/pooltoy/x/pooltoy/types"
)

func handleMsgCreateUser(ctx sdk.Context, k Keeper, msg MsgCreateUser) (*sdk.Result, error) {
	var user = types.User{
		Creator:     msg.Creator,
		UserAccount: msg.UserAccount,
		IsAdmin:     msg.IsAdmin,
		ID:          msg.ID,
		Name:        msg.Name,
		Email:       msg.Email,
	}

	allUsersRaw, err := k.ListUsers(ctx)
	if err != nil {
		return nil, err
	}
	var allUsers []types.User
	k.Cdc.MustUnmarshalJSON(allUsersRaw, &allUsers)

	// does this creator have permission to create this new user?
	// bare in mind special case allows create as initialization when there are no users yet
	creator := k.GetUserByAccAddress(ctx, msg.Creator)
	if creator.UserAccount.Empty() && len(allUsers) != 0 {
		errMsg := fmt.Sprintf("user %s does not exist", msg.Creator)
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, errMsg)
	}

	//  check that new user doesn't exist already
	if existingUser := k.GetUserByAccAddress(ctx, msg.UserAccount); existingUser.UserAccount.Equals(msg.UserAccount) {
		errMsg := fmt.Sprintf("user %s already exists", msg.UserAccount)
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, errMsg)
	}

	// special case allow create as initilization when there are no users yet
	if creator.IsAdmin || len(allUsers) == 0 {
		// if yes
		k.CreateUser(ctx, user)
	} else {
		// if no
		// throw error
		errMsg := fmt.Sprintf("user %s (%s) is not an admin", creator.Name, msg.Creator)
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, errMsg)
	}

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
