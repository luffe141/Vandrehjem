package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	Client     *mongo.Client
	Database   *mongo.Database
	Collection *mongo.Collection
	Ctx        context.Context
	Cancel     context.CancelFunc
}

func NewStorage(uri string, database string, collection string, uniqueFields ...string) *Storage {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(database)
	col := db.Collection(collection)

	for _, field := range uniqueFields {
		indexModel := mongo.IndexModel{
			Keys:    bson.M{field: 1},
			Options: options.Index().SetUnique(true),
		}
		_, err = col.Indexes().CreateOne(ctx, indexModel)
		if err != nil {
			log.Fatal(err)
		}
	}

	return &Storage{Client: client, Database: db, Collection: col, Ctx: ctx, Cancel: cancel}
}

func (storage *Storage) CreateData(data any) error {
	_, err := storage.Collection.InsertOne(storage.Ctx, data)
	return err
}

func (storage *Storage) ReadData(filter bson.M) ([]bson.M, error) {
	cursor, err := storage.Collection.Find(storage.Ctx, filter)
	if err != nil {
		return nil, err
	}

	var results []bson.M
	if err = cursor.All(storage.Ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (storage *Storage) UpdateData(filter bson.M, update bson.M) error {
	_, err := storage.Collection.UpdateMany(storage.Ctx, filter, update)
	return err
}

func (storage *Storage) DeleteData(filter bson.M) error {
	_, err := storage.Collection.DeleteMany(storage.Ctx, filter)
	return err
}

func (storage *Storage) Close() {
	storage.Cancel()
	if err := storage.Client.Disconnect(storage.Ctx); err != nil {
		panic(err)
	}
}
