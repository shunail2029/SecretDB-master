package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/shunail2029/SecretDB-master/x/secretdb/types"
	"go.mongodb.org/mongo-driver/bson"
)

// StoreItem stores a item
func (k Keeper) StoreItem(item types.Item) (sdk.TxResponse, error) {
	data := insertOwner(item.Owner, item.Data)
	dataBytes, err := bson.MarshalExtJSON(data, true, false)
	if err != nil {
		return sdk.TxResponse{}, err
	}
	msg := types.NewMsgStoreItem(types.ValidatorAddress, string(dataBytes))
	err = msg.ValidateBasic()
	if err != nil {
		return sdk.TxResponse{}, err
	}

	distChild := calcRemainder(item.Owner, types.ChildCount)
	return sendTxToChild(distChild, []sdk.Msg{msg})
}

// UpdateItem sets a item
func (k Keeper) UpdateItem(iFil types.ItemFilter, update bson.M) (sdk.TxResponse, error) {
	filter := insertOwner(iFil.Owner, iFil.Filter)
	filterBytes, err := bson.MarshalExtJSON(filter, true, false)
	if err != nil {
		return sdk.TxResponse{}, err
	}
	updateBytes, err := bson.MarshalExtJSON(update, true, false)
	if err != nil {
		return sdk.TxResponse{}, err
	}
	msg := types.NewMsgUpdateItem(types.ValidatorAddress, string(filterBytes), string(updateBytes))
	err = msg.ValidateBasic()
	if err != nil {
		return sdk.TxResponse{}, err
	}

	distChild := calcRemainder(iFil.Owner, types.ChildCount)
	return sendTxToChild(distChild, []sdk.Msg{msg})
}

// UpdateItems sets some items
func (k Keeper) UpdateItems(iFil types.ItemFilter, update bson.M) (sdk.TxResponse, error) {
	filter := insertOwner(iFil.Owner, iFil.Filter)
	filterBytes, err := bson.MarshalExtJSON(filter, true, false)
	if err != nil {
		return sdk.TxResponse{}, err
	}
	updateBytes, err := bson.MarshalExtJSON(update, true, false)
	if err != nil {
		return sdk.TxResponse{}, err
	}
	msg := types.NewMsgUpdateItems(types.ValidatorAddress, string(filterBytes), string(updateBytes))
	err = msg.ValidateBasic()
	if err != nil {
		return sdk.TxResponse{}, err
	}

	distChild := calcRemainder(iFil.Owner, types.ChildCount)
	return sendTxToChild(distChild, []sdk.Msg{msg})
}

// DeleteItem deletes a item
func (k Keeper) DeleteItem(iFil types.ItemFilter) (sdk.TxResponse, error) {
	filter := insertOwner(iFil.Owner, iFil.Filter)
	filterBytes, err := bson.MarshalExtJSON(filter, true, false)
	if err != nil {
		return sdk.TxResponse{}, err
	}
	msg := types.NewMsgDeleteItem(types.ValidatorAddress, string(filterBytes))
	err = msg.ValidateBasic()
	if err != nil {
		return sdk.TxResponse{}, err
	}

	distChild := calcRemainder(iFil.Owner, types.ChildCount)
	return sendTxToChild(distChild, []sdk.Msg{msg})
}

// DeleteItems deletes some items
func (k Keeper) DeleteItems(iFil types.ItemFilter) (sdk.TxResponse, error) {
	filter := insertOwner(iFil.Owner, iFil.Filter)
	filterBytes, err := bson.MarshalExtJSON(filter, true, false)
	if err != nil {
		return sdk.TxResponse{}, err
	}
	msg := types.NewMsgDeleteItems(types.ValidatorAddress, string(filterBytes))
	err = msg.ValidateBasic()
	if err != nil {
		return sdk.TxResponse{}, err
	}

	distChild := calcRemainder(iFil.Owner, types.ChildCount)
	return sendTxToChild(distChild, []sdk.Msg{msg})
}

//
// Functions used by querier
//

// getItem returns the item information
func getItem(path []string, k Keeper) ([]byte, error) {
	query := fmt.Sprintf("custom/%s/%s/%s/%s/%s", types.StoreKey, types.QueryGetItem, path[0], path[1], path[2])
	_, pubkey, _, _ := pathUnescape(path, k)
	owner := pubkey.Address()

	distChild := calcRemainder(owner, types.ChildCount)
	return sendQueryToChild(distChild, query)
}

// GetItems returns the item information
func getItems(path []string, k Keeper) ([]byte, error) {
	query := fmt.Sprintf("custom/%s/%s/%s/%s/%s", types.StoreKey, types.QueryGetItems, path[0], path[1], path[2])
	_, pubkey, _, _ := pathUnescape(path, k)
	owner := pubkey.Address()

	distChild := calcRemainder(owner, types.ChildCount)
	return sendQueryToChild(distChild, query)
}
