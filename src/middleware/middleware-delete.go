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

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Variable section
	collectionUser := util.Client.Database(util.DatabaseName).Collection(util.CollectionName[0])
	var result *mongo.DeleteResult
	var filter primitive.D
	var responseJson []byte
	var err error
	var userId primitive.ObjectID
	var _id string
	var claims sjwt.Claims

	// Parse body
	token := r.Header.Get("Authorization")

	claims, err = sjwt.Parse(token)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error35)
		return
	}
	_id, err = claims.GetStr("_id")
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error36)
		return
	}
	userId, err = primitive.ObjectIDFromHex(_id)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error37)
		return
	}

	// Delete data
	filter = bson.D{{Key: "_id", Value: userId}}
	result, err = collectionUser.DeleteOne(context.TODO(), filter)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error38)
		return
	}

	if result.DeletedCount == 0 {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error39)
		return
	}

	clearMerchants(w, userId)

	// Create response
	modelResponse := model.ModelResponse{
		ResponseMessage: "User deleted",
	}
	responseJson, err = json.Marshal(modelResponse)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error40)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(responseJson)
}

func DeleteMerchant(w http.ResponseWriter, r *http.Request) {
	// Variable section
	collectionMerchant := util.Client.Database(util.DatabaseName).Collection(util.CollectionName[1])
	var modelRequestDelete model.ModelRequestDelete
	var result *mongo.DeleteResult
	var filter primitive.D
	var responseJson []byte
	var err error
	var userId primitive.ObjectID
	var _id string
	var claims sjwt.Claims

	// Parse body
	err = json.NewDecoder(r.Body).Decode(&modelRequestDelete)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error41)
		return
	}
	token := r.Header.Get("Authorization")

	claims, err = sjwt.Parse(token)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error42)
		return
	}
	_id, err = claims.GetStr("_id")
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error43)
		return
	}
	userId, err = primitive.ObjectIDFromHex(_id)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error44)
		return
	}

	// Delete data
	filter = bson.D{{Key: "user_id", Value: userId}, {Key: "_id", Value: modelRequestDelete.Item_id}}
	result, err = collectionMerchant.DeleteOne(context.TODO(), filter)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error45)
		return
	}

	if result.DeletedCount == 0 {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error46)
		return
	}

	clearOutlets(w, userId)

	// Create response
	modelResponse := model.ModelResponse{
		ResponseMessage: "Merchant deleted",
	}
	responseJson, err = json.Marshal(modelResponse)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error47)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(responseJson)
}

func DeleteOutlet(w http.ResponseWriter, r *http.Request) {
	// Variable section
	collectionOutlet := util.Client.Database(util.DatabaseName).Collection(util.CollectionName[2])
	var modelRequestDelete model.ModelRequestDelete
	var result *mongo.DeleteResult
	var filter primitive.D
	var responseJson []byte
	var err error
	var userId primitive.ObjectID
	var _id string
	var claims sjwt.Claims

	// Parse body
	err = json.NewDecoder(r.Body).Decode(&modelRequestDelete)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error48)
		return
	}
	token := r.Header.Get("Authorization")

	claims, err = sjwt.Parse(token)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error49)
		return
	}
	_id, err = claims.GetStr("_id")
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error50)
		return
	}
	userId, err = primitive.ObjectIDFromHex(_id)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error51)
		return
	}

	// Delete data
	filter = bson.D{{Key: "user_id", Value: userId}, {Key: "_id", Value: modelRequestDelete.Item_id}}
	result, err = collectionOutlet.DeleteOne(context.TODO(), filter)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error52)
		return
	}

	if result.DeletedCount == 0 {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error53)
		return
	}

	clearTransactions(w, userId)

	// Create response
	modelResponse := model.ModelResponse{
		ResponseMessage: "Outlet deleted",
	}
	responseJson, err = json.Marshal(modelResponse)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error54)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(responseJson)
}

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	// Variable section
	collectionTransaction := util.Client.Database(util.DatabaseName).Collection(util.CollectionName[3])
	var modelRequestDelete model.ModelRequestDelete
	var result *mongo.DeleteResult
	var filter primitive.D
	var responseJson []byte
	var err error
	var userId primitive.ObjectID
	var _id string
	var claims sjwt.Claims

	// Parse body
	err = json.NewDecoder(r.Body).Decode(&modelRequestDelete)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error55)
		return
	}
	token := r.Header.Get("Authorization")

	claims, err = sjwt.Parse(token)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error56)
		return
	}
	_id, err = claims.GetStr("_id")
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error57)
		return
	}
	userId, err = primitive.ObjectIDFromHex(_id)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error58)
		return
	}

	// Delete data
	filter = bson.D{{Key: "user_id", Value: userId}, {Key: "_id", Value: modelRequestDelete.Item_id}}
	result, err = collectionTransaction.DeleteOne(context.TODO(), filter)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error59)
		return
	}

	if result.DeletedCount == 0 {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error60)
		return
	}

	// Create response
	modelResponse := model.ModelResponse{
		ResponseMessage: "Transaction deleted",
	}
	responseJson, err = json.Marshal(modelResponse)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error61)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(responseJson)
}
