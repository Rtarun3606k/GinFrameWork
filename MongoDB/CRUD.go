package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DatabaseOperations defines an interface for database operations
type DatabaseOperations interface {
	Query(client *mongo.Client, ctx context.Context, dataBase, col string, query, field interface{}) (result *mongo.Cursor, err error)
	Insert(client *mongo.Client, ctx context.Context, dataBase, col string, document interface{}) (result *mongo.InsertOneResult, err error)
	Update(client *mongo.Client, ctx context.Context, dataBase, col string, filter, update interface{}, updateOneOrmany bool) (result *mongo.UpdateResult, err error)
	Delete(client *mongo.Client, ctx context.Context, dataBase, col string, filter interface{}, deleteOneOrmany bool) (result *mongo.DeleteResult, err error)
}

// MongoDBOperations implements the DatabaseOperations interface for MongoDB
type MongoDBOperations struct{}

func (m MongoDBOperations) Query(client *mongo.Client, ctx context.Context, dataBase, col string, query, field interface{}) (result *mongo.Cursor, err error) {
	collection := client.Database(dataBase).Collection(col)
	result, err = collection.Find(ctx, query, options.Find().SetProjection(field))
	return
}

func (m MongoDBOperations) Insert(client *mongo.Client, ctx context.Context, dataBase, col string, document interface{}) (result *mongo.InsertOneResult, err error) {
	collection := client.Database(dataBase).Collection(col)
	result, err = collection.InsertOne(ctx, document)
	return
}

func (m MongoDBOperations) Update(client *mongo.Client, ctx context.Context, dataBase, col string, filter, update interface{}, updateOneOrmany bool) (result *mongo.UpdateResult, err error) {
	collection := client.Database(dataBase).Collection(col)
	if updateOneOrmany {

		result, err = collection.UpdateOne(ctx, filter, update)
	} else {
		result, err = collection.UpdateMany(ctx, filter, update)

	}
	return
}

func (m MongoDBOperations) Delete(client *mongo.Client, ctx context.Context, dataBase, col string, filter interface{}, deleteOneOrmany bool) (result *mongo.DeleteResult, err error) {
	collection := client.Database(dataBase).Collection(col)
	if deleteOneOrmany {
		result, err = collection.DeleteOne(ctx, filter)
	} else {
		result, err = collection.DeleteMany(ctx, filter)

	}
	// result, err = collection.DeleteOne(ctx, filter)
	return
}
