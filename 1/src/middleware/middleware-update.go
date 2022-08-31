package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"src/model"
	"src/util"

	"github.com/brianvoe/sjwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Variable section
	collectionUser := util.Client.Database(util.DatabaseName).Collection(util.CollectionName[0])
	var modelRequestUpdateUser model.ModelRequestUpdateUser
	var modelDatabaseUser model.ModelDatabaseUser
	var modelResponse model.ModelResponse
	var update primitive.D
	var filterUserId primitive.D
	var result *mongo.UpdateResult
	var responseJson []byte
	var err error
	var userId primitive.ObjectID
	var _id string
	var claims sjwt.Claims
	var token string

	// Parse body
	err = json.NewDecoder(r.Body).Decode(&modelRequestUpdateUser)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error129)
		return
	}
	token = r.Header.Get("Authorization")

	claims, err = sjwt.Parse(token)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error130)
		return
	}
	_id, err = claims.GetStr("_id")
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error131)
		return
	}
	userId, err = primitive.ObjectIDFromHex(_id)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error132)
		return
	}

	// Check id
	filterUserId = bson.D{{Key: "_id", Value: userId}}
	err = collectionUser.FindOne(
		context.TODO(),
		filterUserId,
	).Decode(&modelDatabaseUser)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error133)
		return

	}

	// Create structure
	modelDatabaseUser = model.ModelDatabaseUser{
		Name:      modelRequestUpdateUser.Name,
		User_name: modelRequestUpdateUser.User_name,
		Password:  modelRequestUpdateUser.Password,
	}

	update = bson.D{{Key: "$set", Value: modelDatabaseUser}}
	result, err = collectionUser.UpdateByID(context.TODO(), userId, update)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error134)
		return
	}
	if result.ModifiedCount == 0 {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error135)
		return
	}

	// Create response
	modelResponse = model.ModelResponse{
		ResponseMessage: "User profile updated",
	}

	responseJson, err = json.Marshal(modelResponse)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error136)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(responseJson)
}

func UpdateMerchant(w http.ResponseWriter, r *http.Request) {
	// Variable section
	collectionMerchant := util.Client.Database(util.DatabaseName).Collection(util.CollectionName[1])
	var modelRequestUpdateMerchant model.ModelRequestUpdateMerchant
	var modelDatabaseMerchant model.ModelDatabaseMerchant
	var modelResponse model.ModelResponse
	var update primitive.D
	var filterId primitive.D
	var result *mongo.UpdateResult
	var responseJson []byte
	var err error
	var userId primitive.ObjectID
	var _id string
	var claims sjwt.Claims
	var token string

	// Parse body
	err = json.NewDecoder(r.Body).Decode(&modelRequestUpdateMerchant)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error137)
		return
	}
	token = r.Header.Get("Authorization")

	claims, err = sjwt.Parse(token)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error138)
		return
	}
	_id, err = claims.GetStr("_id")
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error139)
		return
	}
	userId, err = primitive.ObjectIDFromHex(_id)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error140)
		return
	}

	// Check id
	filterId = bson.D{{Key: "user_id", Value: userId}, {Key: "_id", Value: modelRequestUpdateMerchant.Merchant_id}}
	err = collectionMerchant.FindOne(
		context.TODO(),
		filterId,
	).Decode(&modelDatabaseMerchant)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error141)
		return

	}

	// Create structure
	modelDatabaseMerchant = model.ModelDatabaseMerchant{
		Merchant_name: modelRequestUpdateMerchant.Merchant_name,
	}

	update = bson.D{{Key: "$set", Value: modelDatabaseMerchant}}
	result, err = collectionMerchant.UpdateByID(context.TODO(), modelRequestUpdateMerchant.Merchant_id, update)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error142)
		return
	}
	if result.ModifiedCount == 0 {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error143)
		return
	}

	// Create response
	modelResponse = model.ModelResponse{
		ResponseMessage: "Merchant updated",
	}

	responseJson, err = json.Marshal(modelResponse)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error144)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(responseJson)
}

func UpdateOutlet(w http.ResponseWriter, r *http.Request) {
	// Variable section
	collectionOutlet := util.Client.Database(util.DatabaseName).Collection(util.CollectionName[2])
	var modelRequestUpdateOutlet model.ModelRequestUpdateOutlet
	var modelDatabaseOutlet model.ModelDatabaseOutlet
	var modelResponse model.ModelResponse
	var update primitive.D
	var filterId primitive.D
	var result *mongo.UpdateResult
	var responseJson []byte
	var err error
	var userId primitive.ObjectID
	var _id string
	var claims sjwt.Claims
	var token string

	// Parse body
	err = json.NewDecoder(r.Body).Decode(&modelRequestUpdateOutlet)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error145)
		return
	}
	token = r.Header.Get("Authorization")

	claims, err = sjwt.Parse(token)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error146)
		return
	}
	_id, err = claims.GetStr("_id")
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error147)
		return
	}
	userId, err = primitive.ObjectIDFromHex(_id)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error148)
		return
	}

	// Check id
	filterId = bson.D{{Key: "user_id", Value: userId}, {Key: "merchant_id", Value: modelRequestUpdateOutlet.Merchant_id}, {Key: "_id", Value: modelRequestUpdateOutlet.Outlet_id}}
	err = collectionOutlet.FindOne(
		context.TODO(),
		filterId,
	).Decode(&modelDatabaseOutlet)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error149)
		return

	}

	// Create structure
	modelDatabaseOutlet = model.ModelDatabaseOutlet{
		Outlet_name: modelRequestUpdateOutlet.Outlet_name,
	}

	update = bson.D{{Key: "$set", Value: modelDatabaseOutlet}}
	result, err = collectionOutlet.UpdateByID(context.TODO(), modelRequestUpdateOutlet.Outlet_id, update)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error150)
		return
	}
	if result.ModifiedCount == 0 {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error151)
		return
	}

	// Create response
	modelResponse = model.ModelResponse{
		ResponseMessage: "Outlet updated",
	}

	responseJson, err = json.Marshal(modelResponse)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error152)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(responseJson)
}

func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	// Variable section
	collectionTransaction := util.Client.Database(util.DatabaseName).Collection(util.CollectionName[3])
	var modelRequestUpdateTransaction model.ModelRequestUpdateTransaction
	var modelDatabaseTransaction model.ModelDatabaseTransaction
	var modelResponse model.ModelResponse
	var update primitive.D
	var filterId primitive.D
	var result *mongo.UpdateResult
	var responseJson []byte
	var err error
	var userId primitive.ObjectID
	var _id string
	var claims sjwt.Claims
	var token string

	// Parse body
	err = json.NewDecoder(r.Body).Decode(&modelRequestUpdateTransaction)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error153)
		return
	}
	token = r.Header.Get("Authorization")

	claims, err = sjwt.Parse(token)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error154)
		return
	}
	_id, err = claims.GetStr("_id")
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error155)
		return
	}
	userId, err = primitive.ObjectIDFromHex(_id)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error156)
		return
	}

	// Check id
	filterId = bson.D{{Key: "user_id", Value: userId}, {Key: "merchant_id", Value: modelRequestUpdateTransaction.Merchant_id}, {Key: "outlet_id", Value: modelRequestUpdateTransaction.Outlet_id}, {Key: "_id", Value: modelRequestUpdateTransaction.Transaction_id}}
	err = collectionTransaction.FindOne(
		context.TODO(),
		filterId,
	).Decode(&modelDatabaseTransaction)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error157)
		return

	}

	// Create structure
	modelDatabaseTransaction = model.ModelDatabaseTransaction{
		Bill_total: modelRequestUpdateTransaction.Bill_total,
	}

	update = bson.D{{Key: "$set", Value: modelDatabaseTransaction}}
	result, err = collectionTransaction.UpdateByID(context.TODO(), modelRequestUpdateTransaction.Transaction_id, update)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error158)
		return
	}
	if result.ModifiedCount == 0 {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error159)
		return
	}

	// Create response
	modelResponse = model.ModelResponse{
		ResponseMessage: "Transaction updated",
	}

	responseJson, err = json.Marshal(modelResponse)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error160)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(responseJson)
}
