package secretdb

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/shunail2029/SecretDB-master/x/secretdb/keeper"
	"github.com/shunail2029/SecretDB-master/x/secretdb/types"
)

func handleMsgCreateBlockHash(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateBlockHash) (*sdk.Result, error) {
	// check sender is operator
	if !types.OperatorAddress.Equals(msg.GetSigners()[0]) {
		return nil, errors.New("tx from operator is acceptable")
	}

	var BlockHash = types.BlockHash{
		Creator: msg.Creator,
		ChainID: msg.ChainID,
		Height:  msg.Height,
		Hash:    msg.Hash,
	}
	k.CreateBlockHash(ctx, BlockHash)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
