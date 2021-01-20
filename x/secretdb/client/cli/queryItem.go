package cli

import (
	"bufio"
	"fmt"
	"net/url"

	"github.com/spf13/viper"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/codec"
	crypto "github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/shunail2029/SecretDB-master/x/secretdb/types"
	"github.com/spf13/cobra"
)

// GetCmdGetItem ...
func GetCmdGetItem(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-item [filter]",
		Short: "Query a item by filter",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			var filter bson.M
			err := bson.UnmarshalExtJSON([]byte(args[0]), true, &filter)
			if err != nil {
				return err
			}

			owner := cliCtx.GetFromAddress()
			if owner.Empty() {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "owner can't be empty")
			}

			// create keybase
			keybase, err := crypto.NewKeyring(
				sdk.KeyringServiceName(),
				viper.GetString(flags.FlagKeyringBackend),
				viper.GetString(flags.FlagHome),
				bufio.NewReader(cmd.InOrStdin()),
			)
			if err != nil {
				fmt.Printf("failed to create keybase\n%s\n", err.Error())
				return nil
			}

			// get owner from cli and generate signature
			fromName := cliCtx.GetFromName()
			pubkey, sig, err := makeSignature(keybase, fromName, keys.DefaultKeyPass, []byte(args[0]))
			if err != nil {
				fmt.Printf("failed to generate signature\n%s\n", err.Error())
				return nil
			}

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf(
					"custom/%s/%s/%s/%s/%s",
					queryRoute,
					types.QueryGetItem,
					url.PathEscape(args[0]),
					url.PathEscape(string(pubkey.Bytes())),
					url.PathEscape(string(sig)),
				),
				nil,
			)
			if err != nil {
				fmt.Printf("could not resolve item %s \n%s\n", args[0], err.Error())
				return nil
			}

			var out bson.M
			err = bson.UnmarshalExtJSON(res, true, &out)
			if err != nil {
				return err
			}
			return printOutput(out, cliCtx.OutputFormat, cliCtx.Indent)
		},
	}
	// to generate signature in get-item and get-items command
	cmd.Flags().String(flags.FlagFrom, "", "Name or address of private key with which to sign")
	cmd.Flags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend, "Select keyring's backend (os|file|test)")
	viper.BindPFlag(flags.FlagKeyringBackend, cmd.Flags().Lookup(flags.FlagKeyringBackend))
	return cmd
}

// GetCmdGetItems ...
func GetCmdGetItems(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-items [filter]",
		Short: "Query some items by filter",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			var filter bson.M
			err := bson.UnmarshalExtJSON([]byte(args[0]), true, &filter)
			if err != nil {
				return err
			}

			owner := cliCtx.GetFromAddress()
			if owner.Empty() {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "owner can't be empty")
			}

			// create keybase
			keybase, err := crypto.NewKeyring(
				sdk.KeyringServiceName(),
				viper.GetString(flags.FlagKeyringBackend),
				viper.GetString(flags.FlagHome),
				bufio.NewReader(cmd.InOrStdin()),
			)
			if err != nil {
				fmt.Printf("failed to create keybase\n%s\n", err.Error())
				return nil
			}

			// get owner from cli and generate signature
			fromName := cliCtx.GetFromName()
			pubkey, sig, err := makeSignature(keybase, fromName, keys.DefaultKeyPass, []byte(args[0]))
			if err != nil {
				fmt.Printf("failed to generate signature\n%s\n", err.Error())
				return nil
			}

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf(
					"custom/%s/%s/%s/%s/%s",
					queryRoute,
					types.QueryGetItems,
					url.PathEscape(args[0]),
					url.PathEscape(string(pubkey.Bytes())),
					url.PathEscape(string(sig)),
				),
				nil,
			)
			if err != nil {
				fmt.Printf("could not resolve item %s\n%s\n", args[0], err.Error())
				return nil
			}

			idxs := findItemBeginnings(res)
			for _, idx := range idxs {
				var out bson.M
				err = bson.UnmarshalExtJSON(res[idx:], true, &out)
				if err != nil {
					return err
				}
				err = printOutput(out, cliCtx.OutputFormat, cliCtx.Indent)
				if err != nil {
					return err
				}
			}

			return nil
		},
	}
	// to use "from" flag in get-item and get-items command
	cmd.Flags().String(flags.FlagFrom, "", "Name or address of private key with which to sign")
	cmd.Flags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend, "Select keyring's backend (os|file|test)")
	viper.BindPFlag(flags.FlagKeyringBackend, cmd.Flags().Lookup(flags.FlagKeyringBackend))
	return cmd
}
