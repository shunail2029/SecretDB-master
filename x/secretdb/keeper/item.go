package keeper

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/shunail2029/SecretDB-master/x/secretdb/types"
	"go.mongodb.org/mongo-driver/bson"
)

// StoreItem stores a item
func (k Keeper) StoreItem(item types.Item) {}

// UpdateItem sets a item
func (k Keeper) UpdateItem(iFil types.ItemFilter, update bson.M) {}

// UpdateItems sets some items
func (k Keeper) UpdateItems(iFil types.ItemFilter, update bson.M) {}

// DeleteItem deletes a item
func (k Keeper) DeleteItem(iFil types.ItemFilter) {}

// DeleteItems deletes some items
func (k Keeper) DeleteItems(iFil types.ItemFilter) {}

//
// Functions used by querier
//

// getItem returns the item information
func getItem(path []string, k Keeper) ([]byte, error) {
	msg, pubkey, sigBytes, err := pathUnescape(path, k)
	if err != nil {
		return nil, err
	}
	if !pubkey.VerifyBytes(msg, sigBytes) {
		return nil, errors.New("signature verification failed")
	}

	// insert "_owner" to filter
	var filter bson.M
	err = bson.UnmarshalExtJSON(msg, true, &filter)
	if err != nil {
		return nil, err
	}
	owner := pubkey.Address()
	filter = insertOwner(sdk.AccAddress(owner), filter)

	// dbRes, err := mongodb.GetItem(filter)
	// if err != nil {
	// 	return nil, err
	// }
	// if dbRes.GotItemCount == 0 {
	// 	res, _ := bson.MarshalExtJSON(bson.M{}, true, false)
	// 	return res, nil
	// }

	// var res []byte
	// res, err = bson.MarshalExtJSON(dbRes.Data[0], true, false)
	// if err != nil {
	// 	return nil, err
	// }
	// return res, nil
	return nil, nil
}

// GetItems returns the item information
func getItems(path []string, k Keeper) ([]byte, error) {
	msg, pubkey, sigBytes, err := pathUnescape(path, k)
	if err != nil {
		return nil, err
	}
	if !pubkey.VerifyBytes(msg, sigBytes) {
		return nil, errors.New("signature verification failed")
	}

	// insert "_owner" to filter
	var filter bson.M
	err = bson.UnmarshalExtJSON(msg, true, &filter)
	if err != nil {
		return nil, err
	}
	owner := pubkey.Address()
	filter = insertOwner(sdk.AccAddress(owner), filter)

	// dbRes, err := mongodb.GetItems(filter)
	// if err != nil {
	// 	return nil, err
	// }
	// if dbRes.GotItemCount == 0 {
	// 	res, _ := bson.MarshalExtJSON(bson.M{}, true, false)
	// 	return res, nil
	// }

	// var res []byte
	// for _, data := range dbRes.Data {
	// 	res, err = bson.MarshalExtJSONAppend(res, data, true, false)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }
	// return res, nil
	return nil, nil
}

// GetItemOwner gets owner of the item
func (k Keeper) GetItemOwner(filter bson.M) sdk.AccAddress {
	// res, err := mongodb.GetItem(filter)
	// if err != nil || res.GotItemCount != 1 {
	// 	return nil
	// }
	// switch addr := res.Data[0]["_owner"].(type) {
	// case sdk.AccAddress:
	// 	return addr
	// default:
	// 	return nil
	// }
	return nil
}

// GetItemsOwner gets owner of the items
// If one owner owns all items, return address of the owner
func (k Keeper) GetItemsOwner(filter bson.M) sdk.AccAddress {
	// res, err := mongodb.GetItem(filter)
	// if err != nil || res.GotItemCount == 0 {
	// 	return nil
	// }
	// switch addr := res.Data[0]["_owner"].(type) { // type assertion of res.Data[0]["_owner"]
	// case sdk.AccAddress:
	// 	for _, data := range res.Data {
	// 		switch dataAddr := data["_owner"].(type) { // type assertion of data["_owner"]
	// 		case sdk.AccAddress:
	// 			if !bytes.Equal(dataAddr, addr) {
	// 				return nil
	// 			}
	// 		default:
	// 			return nil
	// 		}
	// 	}
	// 	return addr
	// default:
	// 	return nil
	// }
	return nil
}

// ItemExists checks if the key exists in the store
func (k Keeper) ItemExists(iFil types.ItemFilter) bool {
	// filter := insertOwner(iFil.Owner, iFil.Filter)
	// res, err := mongodb.GetItem(filter)
	// return err == nil && res.GotItemCount > 0
	return true
}
