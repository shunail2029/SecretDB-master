package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/shunail2029/SecretDB-master/x/secretdb/types"
)

// CreateBlockHash creates a BlockHash
func (k Keeper) CreateBlockHash(ctx sdk.Context, BlockHash types.BlockHash) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.BlockHashPrefix + BlockHash.ChainID + "@" + fmt.Sprint(BlockHash.Height))
	value := k.cdc.MustMarshalBinaryLengthPrefixed(BlockHash)
	store.Set(key, value)
}

// GetBlockHash returns the BlockHash information
func (k Keeper) GetBlockHash(ctx sdk.Context, key string) (types.BlockHash, error) {
	store := ctx.KVStore(k.storeKey)
	var BlockHash types.BlockHash
	byteKey := []byte(types.BlockHashPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &BlockHash)
	if err != nil {
		return BlockHash, err
	}
	return BlockHash, nil
}

// SetBlockHash sets a BlockHash
func (k Keeper) SetBlockHash(ctx sdk.Context, BlockHash types.BlockHash) {
	BlockHashKey := BlockHash.ChainID + "@" + fmt.Sprint(BlockHash.Height)
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(BlockHash)
	key := []byte(types.BlockHashPrefix + BlockHashKey)
	store.Set(key, bz)
}

// DeleteBlockHash deletes a BlockHash
func (k Keeper) DeleteBlockHash(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.BlockHashPrefix + key))
}

//
// Functions used by querier
//

func getBlockHash(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	BlockHash, err := k.GetBlockHash(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, BlockHash)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// GetBlockHashOwner gets creator of the item
func (k Keeper) GetBlockHashOwner(ctx sdk.Context, key string) sdk.AccAddress {
	BlockHash, err := k.GetBlockHash(ctx, key)
	if err != nil {
		return nil
	}
	return BlockHash.Creator
}

// BlockHashExists checks if the key exists in the store
func (k Keeper) BlockHashExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.BlockHashPrefix + key))
}
