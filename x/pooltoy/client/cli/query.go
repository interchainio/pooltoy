package cli

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/interchainberlin/pooltoy/x/pooltoy/types"
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

	pooltoyQueryCmd.AddCommand(queryListUsers())

	return pooltoyQueryCmd
}

func queryListUsers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-users",
		Short: "list all users",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(ctx)
			req := &types.QueryListUsersRequest{}
			res, err := queryClient.QueryListUsers(context.Background(), req)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
