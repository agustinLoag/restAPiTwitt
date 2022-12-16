package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
var ctx = context.TODO()

/* ConectarBD permite conectarse a la base de datos */
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Connected Sucessful")
	return client
}

/* ChequeoConnection el ping a la base de datos */
func ChequeoConnection() int {
	err := MongoCN.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err.Error())
		return 0
	}
	return 1
}
