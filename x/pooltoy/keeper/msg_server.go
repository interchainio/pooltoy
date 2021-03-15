package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/interchainberlin/pooltoy/x/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// CreateUser creates a new user
func (k msgServer) CreateUser(
	goCtx context.Context,
	msg *types.MsgCreateUser,
) (
	, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.Keeper.GetVerifiableCredential(ctx, []byte(msg.VerifiableCredential.Id))
	if found {
		return nil, sdkerrors.Wrapf(
			types.ErrVerifiableCredentialFound,
			"vc already exists",
		)

	}

	k.Keeper.SetVerifiableCredential(
		ctx,
		[]byte(msg.VerifiableCredential.Id),
		*msg.VerifiableCredential,
	)

	return &types.MsgCreateVerifiableCredentialResponse{}, nil
}
