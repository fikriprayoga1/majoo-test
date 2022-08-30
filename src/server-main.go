package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"src/middleware"
	"time"

	"src/util"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	initDatabase()
	initServer()

	defer func() {
		err := util.Client.Disconnect(context.TODO())
		if err != nil {
			panic(err)
		}
	}()
}

func initDatabase() {
	log.Println("logInfo : Database initializing . . .")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	databaseUrl := "mongodb://" + os.Getenv("MONGO_HOST") + ":27017"
	var err error
	util.Client, err = mongo.Connect(ctx, options.Client().ApplyURI(databaseUrl))
	if err != nil {
		panic(err)
	}

	// Ping the primary
	err2 := util.Client.Ping(ctx, readpref.Primary())
	if err2 != nil {
		panic(err2)
	}

	log.Println("logInfo : Successfully connected and pinged")
	log.Println("logInfo : Database initialized")

}

func initServer() {
	http.HandleFunc("/login", middleware.Login)
	http.HandleFunc("/read", middleware.ReadReport)
	http.HandleFunc("/readAdvanced", middleware.ReadReportAdvanced)

	log.Printf("logInfo : Server listener started\n\n")
	log.Fatal(http.ListenAndServe(":"+util.Port, nil))
}
