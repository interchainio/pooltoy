package pooltoy

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
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

	// does this creator have permission to create this new user?

	// if yes
	k.CreateUser(ctx, user)

	// if no
	// throw error

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
