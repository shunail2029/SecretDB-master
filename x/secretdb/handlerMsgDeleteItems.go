package secretdb

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/shunail2029/SecretDB-master/x/secretdb/client/cli"
	"github.com/shunail2029/SecretDB-master/x/secretdb/keeper"
	"github.com/shunail2029/SecretDB-master/x/secretdb/types"
)

// Handle a message to delete some items
func handleMsgDeleteItems(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteItems) (*sdk.Result, error) {
	// decrypt msg
	key, err := cli.GenerateSharedKey(msg.Pubkey, nil, types.OperatorName, types.KeyringPassword, k.Codec())
	if err != nil {
		return nil, err
	}
	plainFilter, err := cli.DecryptWithKey(msg.Filter, key)
	if err != nil {
		return nil, err
	}

	var filter bson.M
	err = bson.UnmarshalExtJSON(plainFilter, true, &filter)
	if err != nil {
		return nil, err
	}

	iFil := types.ItemFilter{
		Owner:  msg.Owner,
		Filter: filter,
	}
	res, err := k.DeleteItems(iFil, msg.Pubkey)
	if err != nil {
		return nil, err
	}

	log, _ := json.Marshal(res)
	return &sdk.Result{
		Log:    fmt.Sprintf("%s", string(log)),
		Events: ctx.EventManager().Events(),
	}, nil
}
