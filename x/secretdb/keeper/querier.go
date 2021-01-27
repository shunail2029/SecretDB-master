package keeper

import (
	// this line is used by starport scaffolding # 1
	"github.com/shunail2029/SecretDB-master/x/secretdb/types"

	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for secretdb clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		// this line is used by starport scaffolding # 2
		case types.QueryGetBlockHash:
			return getBlockHash(ctx, path[1:], k)
		case types.QueryGetItem:
			return getItem(path[1:], k)
		case types.QueryGetItems:
			return getItems(path[1:], k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown secretdb query endpoint")
		}
	}
}
