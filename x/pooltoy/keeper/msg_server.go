package keeper

// import (
// 	"github.com/interchainberlin/pooltoy/x/pooltoy/types"
// )

// type msgServer struct {
// 	Keeper
// }

// // NewMsgServerImpl returns an implementation of the MsgServer interface
// // for the provided Keeper.
// func NewMsgServerImpl(keeper Keeper) types.MsgServer {
// 	return &msgServer{Keeper: keeper}
// }

// var _ types.MsgServer = msgServer{}

// // CreateUser creates a new user
// func (k msgServer) CreateUser(
// 	goCtx context.Context,
// 	msg *types.MsgCreateUser,
// ) (*types.MsgCreateUser, error) {
// 	ctx := sdk.UnwrapSDKContext(goCtx)

// 	_, found := k.Keeper.CreateUser(ctx, []byte(msg.User))
// 	if found {
// 		return nil, err

// 	}

// 	k.Keeper.SetVerifiableCredential(
// 		ctx,
// 		[]byte(msg.VerifiableCredential.Id),
// 		*msg.VerifiableCredential,
// 	)

// 	return &types.MsgCreateUser{}, nil
// }
