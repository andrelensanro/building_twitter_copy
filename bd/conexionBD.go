package bd

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*Sera exportada a otro paquete*/
var MongoCONN = ConnectDB()

/*Sera local*/
var pass = os.Getenv("PASS_DB")
var clientOptions = options.Client().ApplyURI("mongodb+srv://andreassanro:" + pass + "@serviceaws.efglhsh.mongodb.net/?retryWrites=true&w=majority")

/*ConnectDB es la función que me permite conectar la BD*/
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Connected succesfully")
	return client
}

func CheckConnection() int {
	err := MongoCONN.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return 0
	}
	return 1
}
