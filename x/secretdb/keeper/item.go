package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/shunail2029/SecretDB-master/x/secretdb/client/cli"
	"github.com/shunail2029/SecretDB-master/x/secretdb/types"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"go.mongodb.org/mongo-driver/bson"
)

// StoreItem stores a item
func (k Keeper) StoreItem(item types.Item, pubkey secp256k1.PubKeySecp256k1) (sdk.TxResponse, error) {
	data := insertOwner(item.Owner, item.Data)
	dataBytes, err := bson.MarshalExtJSON(data, true, false)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	// encrypt
	key, err := cli.GenerateSharedKey(pubkey, nil, types.OperatorName, types.KeyringPassword, k.cdc)
	if err != nil {
		return sdk.TxResponse{}, err
	}
	cipherData, err := cli.EncryptWithKey(dataBytes, key)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	msg := types.NewMsgStoreItem(types.OperatorAddress, pubkey, cipherData)
	err = msg.ValidateBasic()
	if err != nil {
		return sdk.TxResponse{}, err
	}

	distChild := calcRemainder(item.Owner, types.ChildCount)
	return sendTxToChild(distChild, []sdk.Msg{msg}, k.cdc)
}

// UpdateItem sets a item
func (k Keeper) UpdateItem(iFil types.ItemFilter, update bson.M, pubkey secp256k1.PubKeySecp256k1) (sdk.TxResponse, error) {
	filter := insertOwner(iFil.Owner, iFil.Filter)
	filterBytes, err := bson.MarshalExtJSON(filter, true, false)
	if err != nil {
		return sdk.TxResponse{}, err
	}
	updateBytes, err := bson.MarshalExtJSON(update, true, false)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	// encrypt
	key, err := cli.GenerateSharedKey(pubkey, nil, types.OperatorName, types.KeyringPassword, k.cdc)
	if err != nil {
		return sdk.TxResponse{}, err
	}
	cipherFilter, err := cli.EncryptWithKey(filterBytes, key)
	if err != nil {
		return sdk.TxResponse{}, err
	}
	cipherUpdate, err := cli.EncryptWithKey(updateBytes, key)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	msg := types.NewMsgUpdateItem(types.OperatorAddress, pubkey, cipherFilter, cipherUpdate)
	err = msg.ValidateBasic()
	if err != nil {
		return sdk.TxResponse{}, err
	}

	distChild := calcRemainder(iFil.Owner, types.ChildCount)
	return sendTxToChild(distChild, []sdk.Msg{msg}, k.cdc)
}

// UpdateItems sets some items
func (k Keeper) UpdateItems(iFil types.ItemFilter, update bson.M, pubkey secp256k1.PubKeySecp256k1) (sdk.TxResponse, error) {
	filter := insertOwner(iFil.Owner, iFil.Filter)
	filterBytes, err := bson.MarshalExtJSON(filter, true, false)
	if err != nil {
		return sdk.TxResponse{}, err
	}
	updateBytes, err := bson.MarshalExtJSON(update, true, false)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	// encrypt
	key, err := cli.GenerateSharedKey(pubkey, nil, types.OperatorName, types.KeyringPassword, k.cdc)
	if err != nil {
		return sdk.TxResponse{}, err
	}
	cipherFilter, err := cli.EncryptWithKey(filterBytes, key)
	if err != nil {
		return sdk.TxResponse{}, err
	}
	cipherUpdate, err := cli.EncryptWithKey(updateBytes, key)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	msg := types.NewMsgUpdateItems(types.OperatorAddress, pubkey, cipherFilter, cipherUpdate)
	err = msg.ValidateBasic()
	if err != nil {
		return sdk.TxResponse{}, err
	}

	distChild := calcRemainder(iFil.Owner, types.ChildCount)
	return sendTxToChild(distChild, []sdk.Msg{msg}, k.cdc)
}

// DeleteItem deletes a item
func (k Keeper) DeleteItem(iFil types.ItemFilter, pubkey secp256k1.PubKeySecp256k1) (sdk.TxResponse, error) {
	filter := insertOwner(iFil.Owner, iFil.Filter)
	filterBytes, err := bson.MarshalExtJSON(filter, true, false)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	// encrypt
	key, err := cli.GenerateSharedKey(pubkey, nil, types.OperatorName, types.KeyringPassword, k.cdc)
	if err != nil {
		return sdk.TxResponse{}, err
	}
	cipherFilter, err := cli.EncryptWithKey(filterBytes, key)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	msg := types.NewMsgDeleteItem(types.OperatorAddress, pubkey, cipherFilter)
	err = msg.ValidateBasic()
	if err != nil {
		return sdk.TxResponse{}, err
	}

	distChild := calcRemainder(iFil.Owner, types.ChildCount)
	return sendTxToChild(distChild, []sdk.Msg{msg}, k.cdc)
}

// DeleteItems deletes some items
func (k Keeper) DeleteItems(iFil types.ItemFilter, pubkey secp256k1.PubKeySecp256k1) (sdk.TxResponse, error) {
	filter := insertOwner(iFil.Owner, iFil.Filter)
	filterBytes, err := bson.MarshalExtJSON(filter, true, false)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	// encrypt
	key, err := cli.GenerateSharedKey(pubkey, nil, types.OperatorName, types.KeyringPassword, k.cdc)
	if err != nil {
		return sdk.TxResponse{}, err
	}
	cipherFilter, err := cli.EncryptWithKey(filterBytes, key)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	msg := types.NewMsgDeleteItems(types.OperatorAddress, pubkey, cipherFilter)
	err = msg.ValidateBasic()
	if err != nil {
		return sdk.TxResponse{}, err
	}

	distChild := calcRemainder(iFil.Owner, types.ChildCount)
	return sendTxToChild(distChild, []sdk.Msg{msg}, k.cdc)
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
