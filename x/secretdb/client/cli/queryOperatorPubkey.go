package cli

import (
	"fmt"
	"net/url"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/shunail2029/SecretDB-master/x/secretdb/types"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

// GetCmdGetOperatorPubkey ...
func GetCmdGetOperatorPubkey(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-operatorpubkey",
		Short: "Query operator's pubkey",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryGetOperatorPubkey), nil)
			if err != nil {
				fmt.Printf("could not resolve OperatorPubkey\n%s\n", err.Error())

				return nil
			}

			var out string
			cdc.MustUnmarshalJSON(res, &out)
			out, err = url.PathUnescape(out)
			if err != nil {
				fmt.Println(err.Error())

				return nil
			}
			var pubkey secp256k1.PubKeySecp256k1
			cdc.UnmarshalBinaryBare([]byte(out), &pubkey)
			return cliCtx.PrintOutput(pubkey)
		},
	}
}
