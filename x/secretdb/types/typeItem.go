package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"go.mongodb.org/mongo-driver/bson"
)

// Item is a type of data stored in MongoDB
type Item struct {
	Owner sdk.AccAddress `json:"owner" yaml:"owner"`
	Data  bson.M         `json:"data" yaml:"data"`
}

// ItemFilter is a type of filter of data stored in MongoDB
type ItemFilter struct {
	Owner  sdk.AccAddress `json:"owner" yaml:"owner"`
	Filter bson.M         `json:"filter" yaml:"filter"`
}
