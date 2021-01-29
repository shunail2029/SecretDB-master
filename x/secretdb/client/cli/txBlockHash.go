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
	"github.com/shunail2029/SecretDB-master/x/secretdb/types"
)

// GetCmdCreateBlockHash ...
func GetCmdCreateBlockHash(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-BlockHash [chain-id] [height] [hash]",
		Short: "Creates a new BlockHash",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsChainID := string(args[0])
			argsHeight, err := strconv.ParseInt(args[1], 10, 64)
			argsHash := []byte(args[2])
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateBlockHash(cliCtx.GetFromAddress(), argsChainID, argsHeight, argsHash)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdSetBlockHash ...
func GetCmdSetBlockHash(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-BlockHash [chainId] [height] [hash]",
		Short: "Set a new BlockHash",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsChainID := string(args[0])
			argsHeight, err := strconv.ParseInt(args[1], 10, 64)
			argsHash := []byte(args[2])
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgSetBlockHash(cliCtx.GetFromAddress(), argsChainID, argsHeight, argsHash)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdDeleteBlockHash ...
func GetCmdDeleteBlockHash(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-BlockHash [chain-id] [height]",
		Short: "Delete a new BlockHash by ID",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsChainID := string(args[0])
			argsHeight, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteBlockHash(cliCtx.GetFromAddress(), argsChainID, argsHeight)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
