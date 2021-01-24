package secretdb

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/shunail2029/SecretDB-master/x/secretdb/keeper"
	"github.com/shunail2029/SecretDB-master/x/secretdb/types"
)

// Handle a message to delete item
func handleMsgDeleteItem(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteItem) (*sdk.Result, error) {
	var filter bson.M
	err := bson.UnmarshalExtJSON([]byte(msg.Filter), true, &filter)
	if err != nil {
		return nil, err
	}

	iFil := types.ItemFilter{
		Owner:  msg.Owner,
		Filter: filter,
	}
	res, err := k.DeleteItem(iFil)
	if err != nil {
		return nil, err
	}

	log, _ := json.Marshal(res)
	return &sdk.Result{
		Log:    fmt.Sprintf("%s", string(log)),
		Events: ctx.EventManager().Events(),
	}, nil
}
