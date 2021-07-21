package keeper

import (
	"fmt"
	"github.com/interchainberlin/pooltoy/x/faucet/utils"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/interchainberlin/pooltoy/x/faucet/types"
	"github.com/tendermint/tendermint/libs/log"
)

const FaucetStoreKey = "DefaultFaucetStoreKey"

// Keeper maintains the link to storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	BankKeeper    bankkeeper.Keeper
	StakingKeeper stakingkeeper.Keeper
	AccountKeeper authkeeper.AccountKeeper
	amount        int64                 // set default amount for each mint.
	Limit         time.Duration         // rate limiting for mint, etc 24 * time.Hours
	storeKey      sdk.StoreKey          // Unexposed key to access store from sdk.Context
	cdc           codec.BinaryMarshaler //
}

// NewKeeper creates new instances of the Faucet Keeper
func NewKeeper(
	bankKeeper bankkeeper.Keeper,
	stakingKeeper stakingkeeper.Keeper,
	accountKeeper authkeeper.AccountKeeper,
	amount int64,
	rateLimit time.Duration,
	storeKey sdk.StoreKey,
	cdc codec.BinaryMarshaler) Keeper {
	return Keeper{
		BankKeeper:    bankKeeper,
		StakingKeeper: stakingKeeper,
		AccountKeeper: accountKeeper,
		amount:        amount,
		Limit:         rateLimit,
		storeKey:      storeKey,
		cdc:           cdc,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// MintAndSend mint coins and send to receiver.
func (k Keeper) MintAndSend(ctx sdk.Context, msg *types.MsgMint) error {
	// TODO: should most of this logic be in the msg_server?

	mintTime := ctx.BlockTime().Unix()
	if msg.Denom == k.StakingKeeper.BondDenom(ctx) {
		return types.ErrCantWithdrawStake
	}

	// refuse mint in 24 hours
	a, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return err
	}
	m := k.getMintHistory(ctx, a)
	if k.isPresent(ctx, a) &&
		time.Unix(int64(m.Lasttime), 0).Add(k.Limit).UTC().After(time.Unix(mintTime, 0)) {
		return types.ErrWithdrawTooOften
	}

	newCoin := sdk.NewCoin(msg.Denom, sdk.NewInt(k.amount))
	m.Tally = m.Tally + k.amount
	m.Lasttime = mintTime
	k.Logger(ctx).Info("*********** mintHistory *********", m)
	k.setMintHistory(ctx, a, m)
	k.Logger(ctx).Info("Mint coin: %s", newCoin)
	newCoins := sdk.NewCoins(newCoin)
	if err := k.BankKeeper.MintCoins(ctx, types.ModuleName, newCoins); err != nil {
		return err
	}

	r, err := sdk.AccAddressFromBech32(msg.Minter)
	if err != nil {
		return err
	}

	receiverAccount := k.AccountKeeper.GetAccount(ctx, r)

	if receiverAccount == nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s does not exist and is not allowed to receive tokens", msg.Minter)
	}

	if err := k.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, r, sdk.NewCoins(newCoin)); err != nil {
		return err
	}

	return nil
}

func (k Keeper) getMintHistory(ctx sdk.Context, minter sdk.AccAddress) types.MintHistory {
	store := ctx.KVStore(k.storeKey)
	if !k.isPresent(ctx, minter) {
		return *types.NewMintHistory(minter, 0)
	}

	bz := store.Get([]byte(minter))
	var history types.MintHistory
	k.cdc.MustUnmarshalBinaryBare(bz, &history)
	return history
}

func (k Keeper) setMintHistory(ctx sdk.Context, minter sdk.AccAddress, history types.MintHistory) {
	if history.Minter == "" {
		k.Logger(ctx).Info("history.Minter is empty")
		return
	}

	if history.Tally == 0 {
		k.Logger(ctx).Info("history.Tally is empty")
		return
	}

	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(minter), k.cdc.MustMarshalBinaryBare(&history))
}

// IsPresent check if the name is present in the store or not
func (k Keeper) isPresent(ctx sdk.Context, minter sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(minter))
}
