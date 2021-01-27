package secretdb

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/shunail2029/SecretDB-master/x/secretdb/keeper"
	"github.com/shunail2029/SecretDB-master/x/secretdb/types"
)

// Handle a message to delete name
func handleMsgDeleteBlockHash(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteBlockHash) (*sdk.Result, error) {
	key := msg.ChainID + fmt.Sprint(msg.Height)

	if !k.BlockHashExists(ctx, key) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, key)
	}
	if !msg.Creator.Equals(k.GetBlockHashOwner(ctx, key)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeleteBlockHash(ctx, key)
	return &sdk.Result{}, nil
}
