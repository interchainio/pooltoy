package cli

import (
	"bufio"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/interchainberlin/pooltoy/x/pooltoy/types"
)

func GetCmdCreateUser(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-user [userAccount] [isAdmin] [name] [email]",
		Short: "Creates a new user",
		Args:  cobra.MinimumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsUserAccount := string(args[0])
			userAccount, err := sdk.AccAddressFromBech32(argsUserAccount)
			if err != nil {
				return err
			}
			var isAdmin bool
			isAdmin, err = strconv.ParseBool(args[1])
			if err != nil {
				return err
			}
			argsName := string(args[2])
			argsEmail := string(args[3])

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateUser(cliCtx.GetFromAddress(), userAccount, isAdmin, argsName, argsEmail)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
