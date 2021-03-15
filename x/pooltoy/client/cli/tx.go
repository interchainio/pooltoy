package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/interchainberlin/pooltoy/x/pooltoy/types"
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
		
	)

	pooltoyTxCmd.AddCommand(flags.PostCommands(
		// this line is used by starport scaffolding
		GetCmdCreateUser(cdc),
	)...)

	return pooltoyTxCmd
}
