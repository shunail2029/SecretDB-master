package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/shunail2029/SecretDB-master/x/secretdb/types"
	"github.com/spf13/cobra"
)

// GetCmdGetBlockHash ...
func GetCmdGetBlockHash(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-BlockHash [chain-id] [height]",
		Short: "Query a BlockHash by chain-id and heignt",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			key := args[0] + args[1]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryGetBlockHash, key), nil)
			if err != nil {
				fmt.Printf("could not resolve BlockHash %s \n%s\n", key, err.Error())

				return nil
			}

			var out types.BlockHash
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
