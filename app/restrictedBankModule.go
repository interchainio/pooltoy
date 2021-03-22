package app

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

// RestrictedBankModule overrides the NFT module for custom handlers
type RestrictedBankModule struct {
	bank.AppModule
	keeper        bankkeeper.Keeper
	accountKeeper authkeeper.AccountKeeper
}

// NewRestrictedBankModule generates a new NFT Module
func NewRestrictedBankModule(appModule bank.AppModule, keeper bankkeeper.Keeper, accountKeeper authkeeper.AccountKeeper) RestrictedBankModule {
	return RestrictedBankModule{
		AppModule:     appModule,
		keeper:        keeper,
		accountKeeper: accountKeeper,
	}
}

// NewHandler module handler for the OerrideNFTModule
func (am RestrictedBankModule) NewHandler() sdk.Handler {
	return RestrictedBankHandler(am.keeper, am.accountKeeper)
}

// RestrictedBankHandler routes the messages to the handlers
func RestrictedBankHandler(k bankkeeper.Keeper, ak authkeeper.AccountKeeper) sdk.Handler {
	oldHandler := bank.NewHandler(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		fmt.Println("new restricted bank handler getting used")

		switch msg := msg.(type) {
		case bank.MsgSend:
			return handleMsgSend(ctx, k, msg, ak, oldHandler)

		// NOTE: disabling multisend for now rather than re-writing it with account checking

		// case types.MsgMultiSend:
		// 	return handleMsgMultiSend(ctx, k, msg)

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized bank message type: %T", msg)
		}
	}
}

// Handle MsgSend.
func handleMsgSend(ctx sdk.Context, k bankkeeper.Keeper, msg bank.MsgSend, ak authkeeper.AccountKeeper, oldHandler sdk.Handler) (*sdk.Result, error) {
	fmt.Println("new restricted bank handler handleMsgSend getting used - updated")

	acc := ak.GetAccount(ctx, msg.ToAddress)
	fmt.Println("this is the account", acc, msg.ToAddress)
	if acc == nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to receive transactions", msg.ToAddress)
	}

	return oldHandler(ctx, msg)
	// if !k.GetSendEnabled(ctx) {
	// 	return nil, types.ErrSendDisabled
	// }

	// if k.BlacklistedAddr(msg.ToAddress) {
	// 	return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to receive transactions", msg.ToAddress)
	// }

	// err := k.SendCoins(ctx, msg.FromAddress, msg.ToAddress, msg.Amount)
	// if err != nil {
	// 	return nil, err
	// }

	// ctx.EventManager().EmitEvent(
	// 	sdk.NewEvent(
	// 		sdk.EventTypeMessage,
	// 		sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
	// 	),
	// )

	// return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
