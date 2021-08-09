package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	//"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/spf13/cobra"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/interchainberlin/pooltoy/x/escrow/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	escrowTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s escrow offer subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1
	escrowTxCmd.AddCommand(escrowOffer(), escrowResponse(), escrowCancelOffer())
	return escrowTxCmd
}

func escrowOffer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "offer [offer] [request] --from [username]",
		Short: "send coins to escrow",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			addr := ctx.FromAddress

			offerReq := types.NewOfferRequest(addr, args[1], args[2])

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), offerReq)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func escrowResponse() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "response [id] --from [username]",
		Short: "response to an offer at escrow",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			addr := ctx.FromAddress

			i, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}
			responseReq := types.NewResponseRequest(addr, i)
			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), responseReq)

		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func escrowCancelOffer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel [id] --from [username]",
		Short: "cancel an offer",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			addr := ctx.FromAddress

			i, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}
			cancelReq := types.NewCancelOfferRequest(addr, i)
			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), cancelReq)

		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
