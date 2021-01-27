package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BlockHash is block hash of child chains
type BlockHash struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ChainID string         `json:"chainID" yaml:"chainID"`
	Height  int64          `json:"height" yaml:"height"`
	Hash    []byte         `json:"hash" yaml:"hash"`
}
