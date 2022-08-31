package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"src/model"
	"src/util"
	"time"

	"github.com/brianvoe/sjwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateMerchant(w http.ResponseWriter, r *http.Request) {
	// Variable section
	var modelRequestCreateMerchant model.ModelRequestCreateMerchant
	var filterId primitive.D
	var err error
	collectionUser := util.Client.Database(util.DatabaseName).Collection(util.CollectionName[0])
	collectionMerchant := util.Client.Database(util.DatabaseName).Collection(util.CollectionName[1])
	var modelDatabaseUser model.ModelDatabaseUser
	var modelDatabaseMerchant model.ModelDatabaseMerchant
	var modelResponse model.ModelResponse
	var responseJson []byte
	var token string
	var _id string
	var userId primitive.ObjectID
	var claims sjwt.Claims

	// Parse section
	err = json.NewDecoder(r.Body).Decode(&modelRequestCreateMerchant)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error0)
		return
	}
	token = r.Header.Get("Authorization")
	claims, err = sjwt.Parse(token)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error1)
		return
	}

	_id, err = claims.GetStr("_id")
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error2)
		return
	}

	userId, err = primitive.ObjectIDFromHex(_id)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error3)
		return
	}

	// Prevent empty merchant_name
	if modelRequestCreateMerchant.Merchant_name == "" {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error4)
		return
	}

	// Check user_id
	filterId = bson.D{{Key: "_id", Value: userId}}
	err = collectionUser.FindOne(
		context.TODO(),
		filterId,
	).Decode(&modelDatabaseUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ErrorHandler(err, w, http.StatusNotFound, util.Error5)
			return
		}
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error6)
		return
	}

	// Check merchant if exist
	filterId = bson.D{{Key: "merchant_name", Value: modelRequestCreateMerchant.Merchant_name}}
	err = collectionMerchant.FindOne(
		context.TODO(),
		filterId,
	).Decode(&modelDatabaseMerchant)
	if err == nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error7)
		return

	}

	// Insert to database
	modelDatabaseMerchant = model.ModelDatabaseMerchant{
		User_id:       userId,
		Merchant_name: modelRequestCreateMerchant.Merchant_name,
	}
	_, err = collectionMerchant.InsertOne(context.TODO(), modelDatabaseMerchant)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error8)
		return
	}

	// Create response
	modelResponse = model.ModelResponse{
		ResponseMessage: "Create merchant success",
	}
	responseJson, err = json.Marshal(modelResponse)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error9)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(responseJson)

}

func CreateOutlet(w http.ResponseWriter, r *http.Request) {
	// Variable section
	var modelRequestCreateOutlet model.ModelRequestCreateOutlet
	var filterId primitive.D
	var err error
	collectionUser := util.Client.Database(util.DatabaseName).Collection(util.CollectionName[0])
	collectionMerchant := util.Client.Database(util.DatabaseName).Collection(util.CollectionName[1])
	collectionOutlet := util.Client.Database(util.DatabaseName).Collection(util.CollectionName[2])
	var modelDatabaseUser model.ModelDatabaseUser
	var modelDatabaseMerchant model.ModelDatabaseMerchant
	var modelDatabaseOutlet model.ModelDatabaseOutlet
	var modelResponse model.ModelResponse
	var responseJson []byte
	var token string
	var _id string
	var userId primitive.ObjectID
	var claims sjwt.Claims

	// Parse section
	err = json.NewDecoder(r.Body).Decode(&modelRequestCreateOutlet)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error10)
		return
	}
	token = r.Header.Get("Authorization")
	claims, err = sjwt.Parse(token)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error11)
		return
	}

	_id, err = claims.GetStr("_id")
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error12)
		return
	}

	userId, err = primitive.ObjectIDFromHex(_id)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error13)
		return
	}

	// Prevent empty outlet_name
	if modelRequestCreateOutlet.Outlet_name == "" {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error14)
		return
	}

	// Check user_id
	filterId = bson.D{{Key: "_id", Value: userId}}
	err = collectionUser.FindOne(
		context.TODO(),
		filterId,
	).Decode(&modelDatabaseUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ErrorHandler(err, w, http.StatusNotFound, util.Error15)
			return
		}
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error16)
		return
	}

	// Check merchant_id
	filterId = bson.D{{Key: "_id", Value: modelRequestCreateOutlet.Merchant_id}}
	err = collectionMerchant.FindOne(
		context.TODO(),
		filterId,
	).Decode(&modelDatabaseMerchant)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ErrorHandler(err, w, http.StatusNotFound, util.Error17)
			return
		}
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error18)
		return
	}

	// Check outlet if exist
	filterId = bson.D{{Key: "outlet_name", Value: modelRequestCreateOutlet.Outlet_name}}
	err = collectionOutlet.FindOne(
		context.TODO(),
		filterId,
	).Decode(&modelDatabaseOutlet)
	if err == nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error19)
		return

	}

	// Insert to database
	modelDatabaseOutlet = model.ModelDatabaseOutlet{
		User_id:     userId,
		Merchant_id: modelRequestCreateOutlet.Merchant_id,
		Outlet_name: modelRequestCreateOutlet.Outlet_name,
	}
	_, err = collectionOutlet.InsertOne(context.TODO(), modelDatabaseOutlet)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error20)
		return
	}

	// Create response
	modelResponse = model.ModelResponse{
		ResponseMessage: "Create outlet success",
	}
	responseJson, err = json.Marshal(modelResponse)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error21)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(responseJson)

}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	// Variable section
	var modelRequestCreateTransaction model.ModelRequestCreateTransaction
	var filterId primitive.D
	var err error
	collectionUser := util.Client.Database(util.DatabaseName).Collection(util.CollectionName[0])
	collectionMerchant := util.Client.Database(util.DatabaseName).Collection(util.CollectionName[1])
	collectionOutlet := util.Client.Database(util.DatabaseName).Collection(util.CollectionName[2])
	collectionTransaction := util.Client.Database(util.DatabaseName).Collection(util.CollectionName[3])
	var modelDatabaseUser model.ModelDatabaseUser
	var modelDatabaseMerchant model.ModelDatabaseMerchant
	var modelDatabaseOutlet model.ModelDatabaseOutlet
	var modelDatabaseTransaction model.ModelDatabaseTransaction
	var modelResponse model.ModelResponse
	var responseJson []byte
	var token string
	var _id string
	var userId primitive.ObjectID
	var claims sjwt.Claims

	// Parse section
	err = json.NewDecoder(r.Body).Decode(&modelRequestCreateTransaction)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error22)
		return
	}
	token = r.Header.Get("Authorization")
	claims, err = sjwt.Parse(token)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error23)
		return
	}

	_id, err = claims.GetStr("_id")
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error24)
		return
	}

	userId, err = primitive.ObjectIDFromHex(_id)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error25)
		return
	}

	// Prevent zero bill_total
	if modelRequestCreateTransaction.Bill_total == 0 {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error26)
		return
	}

	// Check user_id
	filterId = bson.D{{Key: "_id", Value: userId}}
	err = collectionUser.FindOne(
		context.TODO(),
		filterId,
	).Decode(&modelDatabaseUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ErrorHandler(err, w, http.StatusNotFound, util.Error27)
			return
		}
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error28)
		return
	}

	// Check merchant_id
	filterId = bson.D{{Key: "_id", Value: modelRequestCreateTransaction.Merchant_id}}
	err = collectionMerchant.FindOne(
		context.TODO(),
		filterId,
	).Decode(&modelDatabaseMerchant)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ErrorHandler(err, w, http.StatusNotFound, util.Error29)
			return
		}
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error30)
		return
	}

	// Check outlet_id
	filterId = bson.D{{Key: "_id", Value: modelRequestCreateTransaction.Outlet_id}}
	err = collectionOutlet.FindOne(
		context.TODO(),
		filterId,
	).Decode(&modelDatabaseOutlet)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ErrorHandler(err, w, http.StatusNotFound, util.Error31)
			return
		}
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error32)
		return
	}

	// Insert to database
	modelDatabaseTransaction = model.ModelDatabaseTransaction{
		User_id:     userId,
		Merchant_id: modelRequestCreateTransaction.Merchant_id,
		Outlet_id:   modelRequestCreateTransaction.Outlet_id,
		Bill_total:  modelRequestCreateTransaction.Bill_total,
		Created_at:  primitive.Timestamp{T: uint32(time.Now().Unix())},
	}
	_, err = collectionTransaction.InsertOne(context.TODO(), modelDatabaseTransaction)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error33)
		return
	}

	// Create response
	modelResponse = model.ModelResponse{
		ResponseMessage: "Create transaction success",
	}
	responseJson, err = json.Marshal(modelResponse)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error34)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(responseJson)

}
