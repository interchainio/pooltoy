package cli

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/flags"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/interchainberlin/pooltoy/x/escrow/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group escrow queries under a subcommand
	escrowQuerycmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1
	escrowQuerycmd.AddCommand(escrowOfferListAll(), escrowOfferByAddr())
	return escrowQuerycmd
}

func escrowOfferListAll() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "offer-list-all",
		Short: "show all the offers",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(ctx)
			offerList, err := queryClient.QueryOfferListAll(context.Background(), &types.OfferListAllRequest{})
			if err != nil {
				return err
			}

			return ctx.PrintProto(offerList)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}


func escrowOfferByAddr() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "offer-by-addr [querier] [offerer]",
		Short: "show all the offers of an address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(ctx)
			offerList, err := queryClient.QueryOfferByAddr(context.Background(), &types.QueryOfferByAddrRequest{Offerer: args[0]})
			if err != nil {
				return err
			}

			return ctx.PrintProto(offerList)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
