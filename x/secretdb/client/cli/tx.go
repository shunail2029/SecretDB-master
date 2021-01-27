package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/shunail2029/SecretDB-master/x/secretdb/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	secretdbTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	secretdbTxCmd.AddCommand(flags.PostCommands(
		// this line is used by starport scaffolding # 1
		GetCmdCreateBlockHash(cdc),
		GetCmdSetBlockHash(cdc),
		GetCmdDeleteBlockHash(cdc),
		GetCmdStoreItem(cdc),
		GetCmdUpdateItem(cdc),
		GetCmdUpdateItems(cdc),
		GetCmdDeleteItem(cdc),
		GetCmdDeleteItems(cdc),
	)...)

	return secretdbTxCmd
}
