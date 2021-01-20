package mongodb

import "go.mongodb.org/mongo-driver/bson"

// StoreItemResult contains result of StoreItem/StoreItems
type StoreItemResult struct {
	StoredItemCount int64 `json:"storedItemCount" yaml:"storedItemCount"`
}

// GetItemResult contains result of GetItem/GetItems
type GetItemResult struct {
	GotItemCount int64    `json:"gotItemCount" yaml:"gotItemCount"`
	Data         []bson.M `json:"data" yaml:"data"`
}

// UpdateItemResult contains result of UpdateItem/UpdateItems
type UpdateItemResult struct {
	MatchedCount  int64 `json:"matchedCount" yaml:"matchedCount"`
	ModifiedCount int64 `json:"modifiedCount" yaml:"modifiedCount"`
	UpsertedCount int64 `json:"upsertedCount" yaml:"upsertedCount"`
}

// DeleteItemResult contains result of DeleteItem/DeleteItems
type DeleteItemResult struct {
	DeletedCount int64 `json:"deletedCount" yaml:"deletedCount"`
}
