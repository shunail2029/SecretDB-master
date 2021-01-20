package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	itemCollection = "items"
)

// StoreItem stores one item
func StoreItem(document interface{}) (StoreItemResult, error) {
	// create connection to database
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	c := newConnection(ctx)

	_, err := c.collection(itemCollection).InsertOne(context.Background(), document)
	if err != nil {
		return StoreItemResult{}, err
	}
	return StoreItemResult{
		StoredItemCount: 1,
	}, nil
}

// StoreItems stores some items
func StoreItems(documents []interface{}) (StoreItemResult, error) {
	// create connection to database
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	c := newConnection(ctx)

	res, err := c.collection(itemCollection).InsertMany(context.Background(), documents)
	if err != nil {
		return StoreItemResult{}, err
	}
	return StoreItemResult{
		StoredItemCount: int64(len(res.InsertedIDs)),
	}, nil
}

// GetItem gets one item
func GetItem(filter interface{}) (GetItemResult, error) {
	// create connection to database
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	c := newConnection(ctx)

	var res bson.M
	err := c.collection(itemCollection).FindOne(context.Background(), filter).Decode(&res)
	if err == mongo.ErrNoDocuments {
		return GetItemResult{
			GotItemCount: 0,
			Data:         nil,
		}, nil
	} else if err != nil {
		return GetItemResult{}, err
	}

	return GetItemResult{
		GotItemCount: 1,
		Data:         []bson.M{res},
	}, nil
}

// GetItems gets some items
func GetItems(filter interface{}) (GetItemResult, error) {
	// create connection to database
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	c := newConnection(ctx)

	cursor, err := c.collection(itemCollection).Find(context.Background(), filter)
	if err == mongo.ErrNoDocuments {
		return GetItemResult{
			GotItemCount: 0,
			Data:         nil,
		}, nil
	} else if err != nil {
		return GetItemResult{}, err
	}

	var res []bson.M
	if err = cursor.All(context.Background(), &res); err != nil {
		return GetItemResult{}, err
	}
	return GetItemResult{
		GotItemCount: int64(len(res)),
		Data:         res,
	}, nil
}

// UpdateItem updates one item
func UpdateItem(filter interface{}, update interface{}) (UpdateItemResult, error) {
	// create connection to database
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	c := newConnection(ctx)

	res, err := c.collection(itemCollection).UpdateOne(context.Background(), filter, update)
	if err != nil {
		return UpdateItemResult{}, err
	}
	return UpdateItemResult{
		MatchedCount:  res.MatchedCount,
		ModifiedCount: res.ModifiedCount,
		UpsertedCount: res.UpsertedCount,
	}, nil
}

// UpdateItems updates some items
func UpdateItems(filter interface{}, update interface{}) (UpdateItemResult, error) {
	// create connection to database
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	c := newConnection(ctx)

	res, err := c.collection(itemCollection).UpdateMany(context.Background(), filter, update)
	if err != nil {
		return UpdateItemResult{}, err
	}
	return UpdateItemResult{
		MatchedCount:  res.MatchedCount,
		ModifiedCount: res.ModifiedCount,
		UpsertedCount: res.UpsertedCount,
	}, nil
}

// DeleteItem deletes one item
func DeleteItem(filter interface{}) (DeleteItemResult, error) {
	// create connection to database
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	c := newConnection(ctx)

	res, err := c.collection(itemCollection).DeleteOne(context.Background(), filter)
	if err != nil {
		return DeleteItemResult{}, err
	}
	return DeleteItemResult{
		DeletedCount: res.DeletedCount,
	}, nil
}

// DeleteItems deletes some items
func DeleteItems(filter interface{}) (DeleteItemResult, error) {
	// create connection to database
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	c := newConnection(ctx)

	res, err := c.collection(itemCollection).DeleteMany(context.Background(), filter)
	if err != nil {
		return DeleteItemResult{}, err
	}
	return DeleteItemResult{
		DeletedCount: res.DeletedCount,
	}, nil
}
