package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Database *mongo.Database
var DBclient *mongo.Client

func DBConnect() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading .env file")
	}
	value := os.Getenv("MONGOURI")
	fmt.Println(value)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(value))
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
	defer client.Disconnect(ctx)
	DBclient = client
	Database = client.Database("userdb")
}
