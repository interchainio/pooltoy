package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
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

	pooltoyQueryCmd.AddCommand(queryWhenBrrr(), queryEmojiRank())

	return pooltoyQueryCmd
}

func queryWhenBrrr() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "when-brrr [Address]",
		Short: "how many seconds until this user can brrr again",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(ctx)
			req := &types.QueryWhenBrrRequest{
				Address: args[0],
			}
			res, err := queryClient.QueryWhenBrr(context.Background(), req)
			if err != nil {
				return err
			}
			return ctx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func queryEmojiRank() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "emoji-rank [show number]",
		Short: "emoji rank",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(ctx)
			num, err := strconv.Atoi(args[0])
			if err != nil {
				panic(err)
			}

			req := &types.QueryEmojiRankRequest{
				ShowNum: int64(num),
			}
			res, err := queryClient.QueryEmojiRank(context.Background(), req)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
