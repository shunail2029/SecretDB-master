package secretdb

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/shunail2029/SecretDB-master/x/secretdb/keeper"
	abci "github.com/tendermint/tendermint/abci/types"
	// abci "github.com/tendermint/tendermint/abci/types"
)

// BeginBlocker check for infraction evidence or downtime of validators
// on every begin block
func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	// 	FIXME: fill out if your application requires beginblock, if not you can delete this function
}

// EndBlocker called every block, process inflation, update validator set.
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	// 	FIXME: fill out if your application requires endblock, if not you can delete this function
}
