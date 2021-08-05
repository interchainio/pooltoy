package cli

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"strconv"

	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/interchainberlin/pooltoy/x/escrow/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	// Group escrow queries under a subcommand
	escrowQuerycmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1
	escrowQuerycmd.AddCommand(escrowOfferListAll(), escrowOfferByAddr(), escrowOfferByID())
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
		Use:   "offer-by-addr [offerer]",
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

func escrowOfferByID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "offer-by-id [ID]",
		Short: "show all the offer of an ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(ctx)

			i, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return err
			}

			offer, err := queryClient.QueryOfferByID(context.Background(), &types.QueryOfferByIDRequest{Id: i})
			if err != nil {
				return err
			}

			return ctx.PrintProto(offer)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
