package pooltoy

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	pooltoykeeper "github.com/interchainio/pooltoy/x/pooltoy/keeper"
	"github.com/interchainio/pooltoy/x/pooltoy/types"
	pooltoytypes "github.com/interchainio/pooltoy/x/pooltoy/types"
)

func NewHandler(k pooltoykeeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case *pooltoytypes.MsgCreateUser:
			res, err := k.CreateUser(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
