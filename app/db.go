package app

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Database *mongo.Database
var DBclient *mongo.Client

func DBConnect() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(" "))
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
	defer client.Disconnect(ctx)
	DBclient = client
	Database = client.Database("userdb")
}
