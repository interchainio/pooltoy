package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/interchainio/pooltoy/x/faucet/types"
	"github.com/spf13/cobra"
)

// GetTxCmd returns the root tx module
func GetTxCmd() *cobra.Command {
	faucetTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "faucet transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	faucetTxCmd.AddCommand(
		txMint(),
		txMintFor(),
	)

	return faucetTxCmd
}

// txMint is the CLI command for minting coins to the sender address
func txMint() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint",
		Short: "mint coin to sender address",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			denom := args[0]
			minter := ctx.GetFromAddress()
			msg := types.NewMsgMint(minter, minter, denom)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// txMint is the CLI command for minting coins to a specified address
func txMintFor() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mintfor [address] [denom]",
		Short: "mint coin for new address",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			minter, _ := sdk.AccAddressFromBech32(args[0])
			sender := ctx.GetFromAddress()
			_, err = sdk.ParseCoinsNormalized("1" + args[1])
			if err != nil {
				return err
			}
			msg := types.NewMsgMint(sender, minter, args[1])
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
