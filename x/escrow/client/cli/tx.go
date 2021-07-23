package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	//"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
	escrowTxCmd.AddCommand(escrowOffer())
	return escrowTxCmd
}

func escrowOffer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "offer [address] [offer] [request] --from [username]",
		Short: "send coins to escrow",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			offer := types.NewOfferRequest(addr, args[1], args[2])

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), offer)


			// method:2
			//	msgClient := types.NewMsgClient(ctx)
			//	offer := &types.OfferRequest{Sender: args[0], Amount: args[1], Request: args[2]}
			//res, err := msgClient.Offer(context.Background(), offer)
			// if err != nil{
		    //	return err
		    //}
			//	return ctx.PrintProto(res)

		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

