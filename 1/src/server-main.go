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
	// Variable section
	var err error

	// Program section
	initDatabase()
	initServer()

	err = util.Client.Disconnect(context.TODO())
	if err != nil {
		log.Print(err)
	}
}

func initDatabase() {
	// Variable section
	var databaseUrl string
	var ctx context.Context
	var cancel context.CancelFunc
	var err error

	// Init timeout connection
	log.Println("logInfo : Database init start")
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to database
	databaseUrl = "mongodb://" + os.Getenv("MONGO_HOST") + ":27017"
	util.Client, err = mongo.Connect(ctx, options.Client().ApplyURI(databaseUrl))
	if err != nil {
		log.Print(err)
		return
	}

	// Ping test
	err = util.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Print(err)
		return
	}

	log.Println("logInfo : Database connected")
	log.Println("logInfo : Database init finish")

}

func initServer() {
	http.HandleFunc("/register", middleware.Register)
	http.HandleFunc("/login", middleware.Login)
	http.HandleFunc("/create/merchant", middleware.CreateMerchant)
	http.HandleFunc("/create/outlet", middleware.CreateOutlet)
	http.HandleFunc("/create/transaction", middleware.CreateTransaction)
	http.HandleFunc("/read/user", middleware.ReadUser)
	http.HandleFunc("/read/merchants", middleware.ReadMerchants)
	http.HandleFunc("/read/outlets", middleware.ReadOutlets)
	http.HandleFunc("/read/transactions", middleware.ReadTransactions)
	http.HandleFunc("/read/transactions/simple", middleware.ReadTransactionsSimple)
	http.HandleFunc("/read/transactions/complete", middleware.ReadTransactionsComplete)
	http.HandleFunc("/update/user", middleware.UpdateUser)
	http.HandleFunc("/update/merchant", middleware.UpdateMerchant)
	http.HandleFunc("/update/outlet", middleware.UpdateOutlet)
	http.HandleFunc("/update/transaction", middleware.UpdateTransaction)
	http.HandleFunc("/delete/user", middleware.DeleteUser)
	http.HandleFunc("/delete/merchant", middleware.DeleteMerchant)
	http.HandleFunc("/delete/outlet", middleware.DeleteOutlet)
	http.HandleFunc("/delete/transaction", middleware.DeleteTransaction)

	log.Println("logInfo : Server listener start")
	log.Fatal(http.ListenAndServe(":"+util.Port, nil))
}
