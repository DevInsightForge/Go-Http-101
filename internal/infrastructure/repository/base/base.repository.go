package base_repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"http101/internal/infrastructure/database"
)

type BaseRepository[T any] struct {
	collection *mongo.Collection
}

func NewBaseRepository[T any](collectionName string) (*BaseRepository[T], error) {
	collection := database.GetCollection(collectionName)
	return &BaseRepository[T]{collection: collection}, nil
}

type FindFilterOptions struct {
	Filter     interface{}
	Projection interface{}
	SortBy     string
	SortDir    int
	Page       int
	PageSize   int
}

func (r *BaseRepository[T]) FindAllWithOptions(findOptions FindFilterOptions) ([]T, error) {
	ctx := context.Background()

	filterBson := bson.M{}
	queryOptions := options.Find()

	if findOptions.Filter != nil {
		filterMap, ok := findOptions.Filter.(map[string]interface{})
		if !ok {
			return nil, errors.New("filter is not in the expected format")
		}
		if filterMap["_id"] != nil {
			id, ok := filterMap["_id"].(string)
			if !ok {
				return nil, errors.New("invalid _id type")
			}
			oid, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				return nil, errors.New("invalid _id format")
			}
			filterMap["_id"] = oid
		}
		filterBson = bson.M(filterMap)
	}

	if findOptions.Projection != nil {
		projectionMap, ok := findOptions.Projection.(map[string]interface{})
		if !ok {
			return nil, errors.New("projection is not in the expected format")
		}
		projectionBson := bson.M(projectionMap)
		queryOptions.SetProjection(projectionBson)
	}

	if findOptions.SortBy != "" {
		if findOptions.SortDir != 1 && findOptions.SortDir != -1 {
			findOptions.SortDir = 1
		}
		queryOptions.SetSort(bson.D{{Key: findOptions.SortBy, Value: findOptions.SortDir}})
	}

	if findOptions.Page > 0 && findOptions.PageSize > 0 {
		queryOptions.SetSkip(int64((findOptions.Page - 1) * findOptions.PageSize))
		queryOptions.SetLimit(int64(findOptions.PageSize))
	}

	cur, err := r.collection.Find(ctx, filterBson, queryOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var results []T = make([]T, 0)
	for cur.Next(ctx) {
		var result T
		if err := cur.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

func (r *BaseRepository[T]) FindAll() ([]T, error) {
	var results []T = make([]T, 0)
	ctx := context.Background()

	cur, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result T
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

func (r *BaseRepository[T]) FindByID(id string) (T, error) {
	var result T
	ctx := context.Background()
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, errors.New("invalid ID format")
	}
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return result, errors.New("document not found")
		}
		return result, err
	}
	return result, nil
}

func (r *BaseRepository[T]) FindBy(filter interface{}) ([]T, error) {
	var results []T = make([]T, 0)
	ctx := context.Background()

	filterMap, ok := filter.(map[string]interface{})
	if !ok {
		return nil, errors.New("filter is not in the expected format")
	}
	if filterMap["_id"] != nil {
		id, ok := filterMap["_id"].(string)
		if !ok {
			return nil, errors.New("invalid _id type")
		}
		oid, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, errors.New("invalid _id format")
		}
		filterMap["_id"] = oid
	}
	filterBson := bson.M(filterMap)

	cur, err := r.collection.Find(ctx, filterBson)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result T
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

func (r *BaseRepository[T]) Create(t T) (primitive.ObjectID, error) {
	ctx := context.Background()
	result, err := r.collection.InsertOne(ctx, t)

	insertedID := result.InsertedID.(primitive.ObjectID)
	return insertedID, err
}

func (r *BaseRepository[T]) Update(id string, t T) error {
	ctx := context.Background()
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID format")
	}
	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": t})
	if err != nil {
		return errors.New("update operation failed")
	}
	return nil
}

func (r *BaseRepository[T]) Delete(id string) error {
	ctx := context.Background()
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID format")
	}
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return errors.New("delete operation failed")
	}
	return nil
}

func (r *BaseRepository[T]) Count() (int64, error) {
	ctx := context.Background()
	count, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *BaseRepository[T]) CountWithFilter(filter interface{}) (int64, error) {
	ctx := context.Background()

	filterMap, ok := filter.(map[string]interface{})
	if !ok {
		return 0, errors.New("filter is not in the expected format")
	}
	filterBson := bson.M(filterMap)

	count, err := r.collection.CountDocuments(ctx, filterBson)
	if err != nil {
		return 0, err
	}
	return count, nil
}
