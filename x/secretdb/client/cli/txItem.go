package cli

import (
	"bufio"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/shunail2029/SecretDB-master/x/secretdb/types"
)

// GetCmdStoreItem ...
func GetCmdStoreItem(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "store-item [data]",
		Short: "Stores a new item",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// check whether args[0] can be decoded as bson.M
			err := bson.UnmarshalExtJSON([]byte(args[0]), false, bson.M{})
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			plainData := []byte(args[0])
			kb := txBldr.Keybase()
			cipherData, err := encryptMsg(plainData, cliCtx, kb, cdc)
			if err != nil {
				return err
			}

			fmt.Println(cliCtx.GetFromName())
			keyBaseInfo, err := kb.Get(cliCtx.GetFromName())
			if err != nil {
				return err
			}
			var pubkey secp256k1.PubKeySecp256k1
			pk := keyBaseInfo.GetPubKey()
			err = cdc.UnmarshalBinaryBare(pk.Bytes(), &pubkey)

			msg := types.NewMsgStoreItem(cliCtx.GetFromAddress(), pubkey, cipherData)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdUpdateItem ...
func GetCmdUpdateItem(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "update-item [filter] [update]",
		Short: "Update a new item",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			// check whether args[0] and args[1] can be decoded as bson.M
			err := bson.UnmarshalExtJSON([]byte(args[0]), true, bson.M{})
			if err != nil {
				return err
			}
			err = bson.UnmarshalExtJSON([]byte(args[1]), true, bson.M{})
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			plainFilter := []byte(args[0])
			plainUpdate := []byte(args[1])
			kb := txBldr.Keybase()
			cipherFilter, err := encryptMsg(plainFilter, cliCtx, kb, cdc)
			if err != nil {
				return err
			}
			cipherUpdate, err := encryptMsg(plainUpdate, cliCtx, kb, cdc)
			if err != nil {
				return err
			}

			keyBaseInfo, err := kb.Get(cliCtx.GetFromName())
			if err != nil {
				return err
			}
			var pubkey secp256k1.PubKeySecp256k1
			pk := keyBaseInfo.GetPubKey()
			err = cdc.UnmarshalBinaryBare(pk.Bytes(), &pubkey)

			msg := types.NewMsgUpdateItem(cliCtx.GetFromAddress(), pubkey, cipherFilter, cipherUpdate)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdUpdateItems ...
func GetCmdUpdateItems(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "update-items [filter] [update]",
		Short: "Update some new items",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			// check whether args[0] and args[1] can be decoded as bson.M
			err := bson.UnmarshalExtJSON([]byte(args[0]), true, bson.M{})
			if err != nil {
				return err
			}
			err = bson.UnmarshalExtJSON([]byte(args[1]), true, bson.M{})
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			plainFilter := []byte(args[0])
			plainUpdate := []byte(args[1])
			kb := txBldr.Keybase()
			cipherFilter, err := encryptMsg(plainFilter, cliCtx, kb, cdc)
			if err != nil {
				return err
			}
			cipherUpdate, err := encryptMsg(plainUpdate, cliCtx, kb, cdc)
			if err != nil {
				return err
			}

			keyBaseInfo, err := kb.Get(cliCtx.GetFromName())
			if err != nil {
				return err
			}
			var pubkey secp256k1.PubKeySecp256k1
			pk := keyBaseInfo.GetPubKey()
			err = cdc.UnmarshalBinaryBare(pk.Bytes(), &pubkey)

			msg := types.NewMsgUpdateItems(cliCtx.GetFromAddress(), pubkey, cipherFilter, cipherUpdate)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdDeleteItem ...
func GetCmdDeleteItem(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-item [filter]",
		Short: "Delete a item by Filter",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// check whether args[0] can be decoded as bson.M
			err := bson.UnmarshalExtJSON([]byte(args[0]), true, bson.M{})
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			plainFilter := []byte(args[0])
			kb := txBldr.Keybase()
			cipherFilter, err := encryptMsg(plainFilter, cliCtx, kb, cdc)
			if err != nil {
				return err
			}

			keyBaseInfo, err := kb.Get(cliCtx.GetFromName())
			if err != nil {
				return err
			}
			var pubkey secp256k1.PubKeySecp256k1
			pk := keyBaseInfo.GetPubKey()
			err = cdc.UnmarshalBinaryBare(pk.Bytes(), &pubkey)

			msg := types.NewMsgDeleteItem(cliCtx.GetFromAddress(), pubkey, cipherFilter)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdDeleteItems ...
func GetCmdDeleteItems(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-items [filter]",
		Short: "Delete some items by Filter",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// check whether args[0] can be decoded as bson.M
			err := bson.UnmarshalExtJSON([]byte(args[0]), true, bson.M{})
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			plainFilter := []byte(args[0])
			kb := txBldr.Keybase()
			cipherFilter, err := encryptMsg(plainFilter, cliCtx, kb, cdc)
			if err != nil {
				return err
			}

			keyBaseInfo, err := kb.Get(cliCtx.GetFromName())
			if err != nil {
				return err
			}
			var pubkey secp256k1.PubKeySecp256k1
			pk := keyBaseInfo.GetPubKey()
			err = cdc.UnmarshalBinaryBare(pk.Bytes(), &pubkey)

			msg := types.NewMsgDeleteItems(cliCtx.GetFromAddress(), pubkey, cipherFilter)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
