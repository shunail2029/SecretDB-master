package secretdb

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/shunail2029/SecretDB-master/x/secretdb/keeper"
	"github.com/shunail2029/SecretDB-master/x/secretdb/types"
)

func handleMsgSetBlockHash(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetBlockHash) (*sdk.Result, error) {
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
	if !msg.Creator.Equals(k.GetBlockHashOwner(ctx, msg.ChainID+"@"+fmt.Sprint(msg.Height))) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetBlockHash(ctx, BlockHash)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
