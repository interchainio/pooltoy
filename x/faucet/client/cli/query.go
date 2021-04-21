package cli

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/interchainberlin/pooltoy/x/faucet/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	// Group pooltoy queries under a subcommand
	pooltoyQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	pooltoyQueryCmd.AddCommand([]*cobra.Command{
		queryWhenBrrr(),
	}...)

	return pooltoyQueryCmd
}

func queryWhenBrrr() *cobra.Command {
	return &cobra.Command{
		Use:   "when-brrr [userAccount]",
		Short: "how many seconds until this user can brrr again",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(ctx)
			req := &types.QueryWhenBrrRequest{}
			res, err := queryClient.QueryWhenBrr(context.Background(), req)
			if err != nil {
				return err
			}
			return ctx.PrintProto(res)
		},
	}
}
