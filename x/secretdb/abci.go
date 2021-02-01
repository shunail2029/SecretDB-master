package secretdb

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/shunail2029/SecretDB-master/x/secretdb/keeper"
	"github.com/shunail2029/SecretDB-master/x/secretdb/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// BeginBlocker check for infraction evidence or downtime of validators
// on every begin block
func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	for i := 0; i < types.ChildCount; i++ {
		if types.AccNum[i] != 0 {
			continue
		}

		chainID := types.ChildChainIDs[i]
		nodeURI := types.ChildURIs[i]

		ctx := context.CLIContext{
			FromAddress: types.OperatorAddress,
			ChainID:     chainID,
		}.WithNodeURI(nodeURI)

		var err error
		types.AccNum[i], types.AccSeq[i], err = authtypes.NewAccountRetriever(ctx).GetAccountNumberSequence(types.OperatorAddress)
		if err != nil {
			fmt.Printf("failed to get account info from %dth slave", i+1)
		}
	}
}

// EndBlocker called every block, process inflation, update validator set.
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {}
