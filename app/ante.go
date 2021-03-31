package app

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

type AccountExistsCheckDecorator struct {
	ak         authkeeper.AccountKeeper
	bankKeeper types.BankKeeper
}

func NewAccountExistsCheckDecorator(ak authkeeper.AccountKeeper, bk types.BankKeeper) AccountExistsCheckDecorator {
	return AccountExistsCheckDecorator{
		ak:         ak,
		bankKeeper: bk,
	}
}

func (ad AccountExistsCheckDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	msgs := tx.GetMsgs()
	for _, m := range msgs {
		// won't catch ibc messages which is fine bc it doesn't matter if dest account is in another chain
		msg, ok := m.(*banktypes.MsgSend)
		// msg.
		if !ok {
			continue
		}
		a, err := sdk.AccAddressFromBech32(msg.ToAddress)
		if err != nil {
			return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "account address %s is not a valid address", a)
		}

		acc := ad.ak.GetAccount(ctx, a)
		if acc == nil {
			return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "account address %s is not allowed to receive transactions", msg.ToAddress)
		}

	}

	return next(ctx, tx, simulate)
}

// NewAnteHandler returns an AnteHandler that checks and increments sequence
// numbers, checks signatures & account numbers, and deducts fees from the first
// signer
func NewAnteHandler(
	ak authkeeper.AccountKeeper,
	bankKeeper types.BankKeeper,
	sigGasConsumer ante.SignatureVerificationGasConsumer,
	signModeHandler signing.SignModeHandler,
) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		NewAccountExistsCheckDecorator(ak, bankKeeper),
		ante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		ante.NewRejectExtensionOptionsDecorator(),
		ante.NewMempoolFeeDecorator(),
		ante.NewValidateBasicDecorator(),
		ante.TxTimeoutHeightDecorator{},
		ante.NewValidateMemoDecorator(ak),
		ante.NewConsumeGasForTxSizeDecorator(ak),
		ante.NewRejectFeeGranterDecorator(),
		ante.NewSetPubKeyDecorator(ak), // SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewValidateSigCountDecorator(ak),
		ante.NewDeductFeeDecorator(ak, bankKeeper),
		ante.NewSigGasConsumeDecorator(ak, sigGasConsumer),
		ante.NewSigVerificationDecorator(ak, signModeHandler),
		ante.NewIncrementSequenceDecorator(ak),
	)
}
