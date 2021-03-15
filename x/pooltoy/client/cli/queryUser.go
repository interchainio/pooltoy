package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/interchainberlin/pooltoy/x/pooltoy/types"
	"github.com/spf13/cobra"
)

func GetCmdListUsers(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-users",
		Short: "list all users",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(cliCtx)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryListUsers, queryRoute), nil)
			if err != nil {
				fmt.Printf("could not list User\n%s\n", err.Error())
				return nil
			}
			return cliCtx.PrintProto(res)
		},
	}
	return cmd
}
