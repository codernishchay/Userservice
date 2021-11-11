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

var DATABASE *mongo.Database
var DBclient *mongo.Client
var DBCollection *mongo.Collection

func DBConnect() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading .env file")
	}
	value := os.Getenv("MONGOURI")
	fmt.Println(value)
	clientOptions := options.Client().
		ApplyURI(value)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database")
	DBclient = client
	DATABASE = client.Database("userdb")
}
