package pooltoy

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	// abci "github.com/tendermint/tendermint/abci/types"
)

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
func InitGenesis(ctx sdk.Context, k Keeper, data GenesisState) {
	k.CreateUser(ctx, data.User)
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	userAddr, _ := sdk.AccAddressFromBech32("cosmos1qd4gsa4mlnpzmv4zsf9ghdrsgkt5avs8zte65u")
	return NewGenesisState(k.GetUserByAccAddress(ctx, userAddr))
}
