package middleware

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"src/model"
	"src/util"

	"github.com/brianvoe/sjwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Register(w http.ResponseWriter, r *http.Request) {
	// Variable section
	var modelRequestRegister model.ModelRequestRegister
	var filterUserId primitive.D
	var err error
	collectionUser := util.Client.Database(util.DatabaseName).Collection(util.CollectionName[0])
	var modelDatabaseUser model.ModelDatabaseUser
	var modelResponse model.ModelResponse
	var responseJson []byte

	// Parse section
	err = json.NewDecoder(r.Body).Decode(&modelRequestRegister)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error62)
		return
	}

	// Prevent empty user_name
	if modelRequestRegister.User_name == "" {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error63)
		return
	}

	// Check user if exist
	filterUserId = bson.D{{Key: "user_name", Value: modelRequestRegister.User_name}}
	err = collectionUser.FindOne(
		context.TODO(),
		filterUserId,
	).Decode(&modelDatabaseUser)
	if err == nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error64)
		return

	}

	// Insert to database
	modelDatabaseUser = model.ModelDatabaseUser{
		Name:      modelRequestRegister.Name,
		User_name: modelRequestRegister.User_name,
		Password:  modelRequestRegister.Password,
	}
	_, err = collectionUser.InsertOne(context.TODO(), modelDatabaseUser)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error65)
		return
	}

	// Create response
	modelResponse = model.ModelResponse{
		ResponseMessage: "Register success",
	}
	responseJson, err = json.Marshal(modelResponse)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error66)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(responseJson)

}

func Login(w http.ResponseWriter, r *http.Request) {
	// Variable section
	var modelRequestLogin model.ModelRequestLogin
	var err error
	var claims *sjwt.Claims
	var jwt string
	var responseJson []byte
	var filter primitive.M
	var modelResponseLogin model.ModelResponseLogin
	var modelDatabaseUser model.ModelDatabaseUser
	var modelResponseUser model.ModelResponseUser
	collectionUser := util.Client.Database(util.DatabaseName).Collection(util.CollectionName[0])

	// Parse body
	err = json.NewDecoder(r.Body).Decode(&modelRequestLogin)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error67)
		return
	}

	// begin find
	filter = bson.M{"user_name": modelRequestLogin.User_name, "password": modelRequestLogin.Password}
	err = collectionUser.FindOne(
		context.TODO(),
		filter,
	).Decode(&modelDatabaseUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ErrorHandler(err, w, http.StatusNotFound, util.Error68)
			return
		}
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error69)
		return
	}
	claims = sjwt.New()
	claims.Set("_id", modelDatabaseUser.Id)

	// Generate jwt
	jwt = claims.Generate(util.GetScretKey())

	modelResponseUser = model.ModelResponseUser{
		Name:      modelDatabaseUser.Name,
		User_name: modelDatabaseUser.User_name,
	}

	modelResponseLogin = model.ModelResponseLogin{
		ResponseMessage: "Login Success",
		Token:           jwt,
		Profile:         &modelResponseUser,
	}

	responseJson, err = json.Marshal(&modelResponseLogin)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error70)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(responseJson)

}

func ErrorHandler(err error, w http.ResponseWriter, responseCode int, responseMessage string) {
	// Variable section
	var modelResponse model.ModelResponse
	var responseJson []byte

	log.Printf("logInfo : %v => %v", responseMessage, err)
	modelResponse = model.ModelResponse{ResponseMessage: responseMessage}
	responseJson, err = json.Marshal(modelResponse)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(responseCode)
	w.Write(responseJson)

}

func clearMerchants(w http.ResponseWriter, userId primitive.ObjectID) {
	collectionMerchants := util.Client.Database(util.DatabaseName).Collection(util.CollectionName[1])
	collectionOutlets := util.Client.Database(util.DatabaseName).Collection(util.CollectionName[2])
	collectionTransactions := util.Client.Database(util.DatabaseName).Collection(util.CollectionName[3])
	var err error

	filter := bson.D{{Key: "user_id", Value: userId}}
	_, err = collectionMerchants.DeleteMany(context.TODO(), filter)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error71)
		return
	}

	_, err = collectionOutlets.DeleteMany(context.TODO(), filter)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error72)
		return
	}

	_, err = collectionTransactions.DeleteMany(context.TODO(), filter)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error73)
		return
	}
}

func clearOutlets(w http.ResponseWriter, merchantId primitive.ObjectID) {
	collectionOutlets := util.Client.Database(util.DatabaseName).Collection(util.CollectionName[2])
	collectionTransactions := util.Client.Database(util.DatabaseName).Collection(util.CollectionName[3])
	var err error

	filter := bson.D{{Key: "merchant_id", Value: merchantId}}
	_, err = collectionOutlets.DeleteMany(context.TODO(), filter)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error74)
		return
	}

	_, err = collectionTransactions.DeleteMany(context.TODO(), filter)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error75)
		return
	}

}

func clearTransactions(w http.ResponseWriter, outletId primitive.ObjectID) {
	collectionTransactions := util.Client.Database(util.DatabaseName).Collection(util.CollectionName[3])
	filter := bson.D{{Key: "outlet_id", Value: outletId}}

	_, err := collectionTransactions.DeleteMany(context.TODO(), filter)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error76)
		return
	}

}
