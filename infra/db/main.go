package db

import (
	"context"
	"fiap-fast-food-ms-producao/adapter/context_manager"
	"fiap-fast-food-ms-producao/adapter/database"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type databaseManager struct {
	client *mongo.Client
}

func (d *databaseManager) Create(collection string, data map[string]interface{}) (any, error) {
	c := d.client.Database("fiap-tech-challenge").Collection(collection)
	insertOne, err := c.InsertOne(context.TODO(), data)
	if err != nil {
		return nil, err
	}
	return insertOne, nil
}

func (d *databaseManager) ReadOne(collection string, query map[string]interface{}) any {
	c := d.client.Database("fiap-tech-challenge").Collection(collection)
	findOne := c.FindOne(context.TODO(), query)
	return findOne
}

func (d *databaseManager) UpdateOne(collection string, query any, data map[string]interface{}) (any, error) {
	c := d.client.Database("fiap-tech-challenge").Collection(collection)
	updateOne, err := c.UpdateOne(context.TODO(), query, data)
	if err != nil {
		return nil, err
	}
	return updateOne, nil
}

func (d *databaseManager) Disconnect() error {
	return nil
}

func NewDatabaseManager(ctx context_manager.ContextManager) (database.DatabaseManger, error) {
	uri := ctx.Get("mongo_url")
	clientOptions := options.Client().ApplyURI(uri.(string))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	if err := client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}
	fmt.Println("Connected to MongoDB")
	return &databaseManager{client: client}, nil
}
