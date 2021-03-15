package cli

import (
	"bufio"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/interchainberlin/pooltoy/x/pooltoy/types"
)

func GetCmdCreateUser() *cobra.Command {
	return &cobra.Command{
		Use:   "create-user [userAccount] [isAdmin] [name] [email]",
		Short: "Creates a new user",
		Args:  cobra.MinimumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := client.GetClientContextFromCmd(cmd)
			config := sdk.NewConfig()

			// parse arguments from cmd
			u := args[0]
			a, err := strconv.ParseBool(args[1])
			if err != nil {
				return err
			}
			n := string(args[2])
			e := string(args[3])

			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(config.GetTxEncoder())
			msg := types.NewMsgCreateUser(cliCtx.GetFromAddress().String(), u, a, n, e)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return tx.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
