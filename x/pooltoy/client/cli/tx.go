package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/interchainio/pooltoy/x/pooltoy/types"
	"github.com/spf13/cobra"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	pooltoyTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	pooltoyTxCmd.AddCommand(
		txCreateUser(),
	)

	return pooltoyTxCmd
}

func txCreateUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-user [userAccount] [isAdmin] [name] [email]",
		Short: "Creates a new user",
		Args:  cobra.MinimumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// parse arguments from cmd
			a, err := strconv.ParseBool(args[1])
			if err != nil {
				return err
			}

			// populate user
			user := types.User{}
			user.Creator = ctx.GetFromAddress().String()
			user.UserAccount = args[0]
			user.IsAdmin = a
			user.Name = string(args[2])
			user.Email = string(args[3])

			msg := types.NewMsgCreateUser(ctx.FromAddress, user)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
