package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"src/model"
	"src/util"

	"github.com/brianvoe/sjwt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func ReadUser(w http.ResponseWriter, r *http.Request) {
	// Variable user
	collectionUser := util.Client.Database(util.DatabaseName).Collection(util.CollectionName[0])
	var modelResponseUser model.ModelResponseUser
	var modelDatabaseUser model.ModelDatabaseUser
	var userId primitive.ObjectID
	var responseJson []byte
	var claims sjwt.Claims
	var filter primitive.D
	var token string
	var _id string
	var err error

	// Parse body
	token = r.Header.Get("Authorization")
	claims, err = sjwt.Parse(token)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error77)
		return
	}
	_id, err = claims.GetStr("_id")
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error78)
		return
	}

	userId, err = primitive.ObjectIDFromHex(_id)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error79)
		return
	}

	// begin find
	filter = bson.D{{Key: "_id", Value: userId}}
	err = collectionUser.FindOne(
		context.TODO(),
		filter,
	).Decode(&modelDatabaseUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ErrorHandler(err, w, http.StatusNotFound, util.Error80)
			return
		}
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error81)
		return
	}

	// Create response
	modelResponseUser = model.ModelResponseUser{
		Name:      modelDatabaseUser.Name,
		User_name: modelDatabaseUser.User_name,
	}
	responseJson, err = json.Marshal(modelResponseUser)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error82)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(responseJson)

}

func ReadMerchants(w http.ResponseWriter, r *http.Request) {
	// Variable section
	collectionMerchants := util.Client.Database(util.DatabaseName).Collection(util.CollectionName[1])
	var modelResponseMerchants []model.ModelResponseMerchant
	var modelDatabaseMerchants []model.ModelDatabaseMerchant
	var modelResponseMerchant model.ModelResponseMerchant
	var modelDatabaseMerchant model.ModelDatabaseMerchant
	var responseObject interface{}
	var userId primitive.ObjectID
	var cursor *mongo.Cursor
	var responseJson []byte
	var filter primitive.D
	var claims sjwt.Claims
	var token string
	var _id string
	var err error

	// Parse body
	token = r.Header.Get("Authorization")
	claims, err = sjwt.Parse(token)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error83)
		return
	}
	_id, err = claims.GetStr("_id")
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error84)
		return
	}

	userId, err = primitive.ObjectIDFromHex(_id)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error85)
		return
	}

	// begin find
	filter = bson.D{{Key: "user_id", Value: userId}}
	cursor, err = collectionMerchants.Find(context.TODO(), filter)
	if err != nil {
		ErrorHandler(err, w, http.StatusNotFound, util.Error86)
		return
	}
	err = cursor.All(context.TODO(), &modelDatabaseMerchants)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error87)
		return
	}

	// Read data
	for _, modelDatabaseMerchant = range modelDatabaseMerchants {
		modelResponseMerchant = model.ModelResponseMerchant{
			Id:            modelDatabaseMerchant.Id,
			Merchant_name: modelDatabaseMerchant.Merchant_name,
			Created_at:    modelDatabaseMerchant.Created_at,
			Created_by:    modelDatabaseMerchant.Created_by,
			Updated_at:    modelDatabaseMerchant.Updated_at,
			Update_by:     modelDatabaseMerchant.Update_by,
		}
		modelResponseMerchants = append(modelResponseMerchants, modelResponseMerchant)

	}

	if modelDatabaseMerchants != nil {
		responseObject = &modelResponseMerchants

	} else {
		responseObject = bson.A{}

	}

	responseJson, err = json.Marshal(responseObject)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error88)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(responseJson)

}

func ReadOutlets(w http.ResponseWriter, r *http.Request) {
	// Variable section
	collectionOutlets := util.Client.Database(util.DatabaseName).Collection(util.CollectionName[2])
	var modelResponseOutlets []model.ModelResponseOutlet
	var modelDatabaseOutlets []model.ModelDatabaseOutlet
	var modelResponseOutlet model.ModelResponseOutlet
	var modelDatabaseOutlet model.ModelDatabaseOutlet
	var merchantId primitive.ObjectID
	var responseObject interface{}
	var userId primitive.ObjectID
	var queryParameter url.Values
	var requestParameter []string
	var merchantIdString string
	var cursor *mongo.Cursor
	var responseJson []byte
	var filter primitive.D
	var claims sjwt.Claims
	var token string
	var _id string
	var err error

	// Parse body
	queryParameter = r.URL.Query()
	requestParameter = queryParameter["merchant_id"]
	if requestParameter == nil {
		ErrorHandler(err, w, http.StatusNotFound, util.Error89)
		return
	}
	merchantIdString = requestParameter[0]

	token = r.Header.Get("Authorization")
	claims, err = sjwt.Parse(token)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error90)
		return
	}
	_id, err = claims.GetStr("_id")
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error91)
		return
	}

	userId, err = primitive.ObjectIDFromHex(_id)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error92)
		return
	}
	merchantId, err = primitive.ObjectIDFromHex(merchantIdString)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error93)
		return
	}

	// begin find
	filter = bson.D{{Key: "user_id", Value: userId}, {Key: "merchant_id", Value: merchantId}}
	cursor, err = collectionOutlets.Find(context.TODO(), filter)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error94)
		return
	}
	err = cursor.All(context.TODO(), &modelDatabaseOutlets)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error95)
		return
	}

	// Read data
	for _, modelDatabaseOutlet = range modelDatabaseOutlets {
		modelResponseOutlet = model.ModelResponseOutlet{
			Id:          modelDatabaseOutlet.Id,
			Outlet_name: modelDatabaseOutlet.Outlet_name,
			Created_at:  modelDatabaseOutlet.Created_at,
			Created_by:  modelDatabaseOutlet.Created_by,
			Updated_at:  modelDatabaseOutlet.Updated_at,
			Update_by:   modelDatabaseOutlet.Update_by,
		}
		modelResponseOutlets = append(modelResponseOutlets, modelResponseOutlet)

	}

	if modelResponseOutlets != nil {
		responseObject = &modelResponseOutlets

	} else {
		responseObject = bson.A{}

	}

	responseJson, err = json.Marshal(responseObject)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error96)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(responseJson)

}

func ReadTransactions(w http.ResponseWriter, r *http.Request) {
	// Variable section
	collectionTransactions := util.Client.Database(util.DatabaseName).Collection(util.CollectionName[3])
	var modelResponseTransactions []model.ModelResponseTransaction
	var modelDatabaseTransactions []model.ModelDatabaseTransaction
	var modelResponseTransaction model.ModelResponseTransaction
	var modelDatabaseTransaction model.ModelDatabaseTransaction
	var merchantId primitive.ObjectID
	var outletId primitive.ObjectID
	var responseObject interface{}
	var userId primitive.ObjectID
	var queryParameter url.Values
	var requestParameter []string
	var merchantIdString string
	var outletIdString string
	var cursor *mongo.Cursor
	var responseJson []byte
	var filter primitive.D
	var claims sjwt.Claims
	var token string
	var _id string
	var err error

	// Parse body
	queryParameter = r.URL.Query()
	requestParameter = queryParameter["merchant_id"]
	if requestParameter == nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error97)
		return
	}
	merchantIdString = requestParameter[0]
	requestParameter = queryParameter["outlet_id"]
	if requestParameter == nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error98)
		return
	}
	outletIdString = requestParameter[0]

	token = r.Header.Get("Authorization")
	claims, err = sjwt.Parse(token)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error99)
		return
	}
	_id, err = claims.GetStr("_id")
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error100)
		return
	}

	userId, err = primitive.ObjectIDFromHex(_id)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error101)
		return
	}
	merchantId, err = primitive.ObjectIDFromHex(merchantIdString)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error102)
		return
	}
	outletId, err = primitive.ObjectIDFromHex(outletIdString)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error103)
		return
	}

	// begin find
	filter = bson.D{{Key: "user_id", Value: userId}, {Key: "merchant_id", Value: merchantId}, {Key: "outlet_id", Value: outletId}}
	cursor, err = collectionTransactions.Find(context.TODO(), filter)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error104)
		return
	}
	err = cursor.All(context.TODO(), &modelDatabaseTransactions)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error105)
		return
	}

	// Read data
	for _, modelDatabaseTransaction = range modelDatabaseTransactions {
		modelResponseTransaction = model.ModelResponseTransaction{
			Id:         modelDatabaseTransaction.Id,
			Bill_total: modelDatabaseTransaction.Bill_total,
			Created_at: modelDatabaseTransaction.Created_at,
			Created_by: modelDatabaseTransaction.Created_by,
			Updated_at: modelDatabaseTransaction.Updated_at,
			Update_by:  modelDatabaseTransaction.Update_by,
		}
		modelResponseTransactions = append(modelResponseTransactions, modelResponseTransaction)

	}

	if modelResponseTransactions != nil {
		responseObject = &modelResponseTransactions

	} else {
		responseObject = bson.A{}

	}

	responseJson, err = json.Marshal(responseObject)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error106)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(responseJson)

}

func ReadTransactionsSimple(w http.ResponseWriter, r *http.Request) {
	// Variable section
	collectionTransactions := util.Client.Database(util.DatabaseName).Collection(util.CollectionName[3])
	var modelResponseTransactionCompletes []model.ModelResponseTransactionComplete
	var timestampFromTime time.Time
	var offsetId primitive.ObjectID
	var responseObject interface{}
	var userId primitive.ObjectID
	var queryParameter url.Values
	var timestampToTime time.Time
	var pipeline []primitive.D
	var cursor *mongo.Cursor
	var timestampFrom string
	var responseJson []byte
	var timestampTo string
	var claims sjwt.Claims
	var limitNumber int
	var offset string
	var limit string
	var token string
	var _id string
	var err error

	// Parse body
	queryParameter = r.URL.Query()
	timestampFrom = queryParameter.Get("timestamp_from")
	timestampTo = queryParameter.Get("timestamp_to")
	offset = queryParameter.Get("offset")
	limit = queryParameter.Get("limit")
	token = r.Header.Get("Authorization")

	claims, err = sjwt.Parse(token)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error107)
		return
	}
	_id, err = claims.GetStr("_id")
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error108)
		return
	}

	if limit == "" {
		ErrorHandler(nil, w, http.StatusBadRequest, util.Error109)
		return
	}

	limitNumber, err = strconv.Atoi(limit)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error110)
		return
	}

	if offset != "" {
		offsetId, err = primitive.ObjectIDFromHex(offset)
		if err != nil {
			ErrorHandler(err, w, http.StatusInternalServerError, util.Error111)
			return
		}
	} else {
		offsetId = primitive.NilObjectID
	}

	userId, err = primitive.ObjectIDFromHex(_id)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error112)
		return
	}

	timestampFromTime, err = time.Parse(time.RFC3339, timestampFrom)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error113)
		return
	}
	timestampToTime, err = time.Parse(time.RFC3339, timestampTo)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error114)
		return
	}

	pipeline = []bson.D{
		{
			{
				Key: "$match",
				Value: bson.M{
					"_id": bson.M{
						"$gte": offsetId,
					},
					"created_at": bson.M{
						"$gte": primitive.Timestamp{
							T: uint32(timestampFromTime.Unix()),
						},
						"$lte": primitive.Timestamp{
							T: uint32(timestampToTime.Unix()),
						},
					},
					"user_id": userId,
				},
			},
		},
		{
			{
				Key: "$lookup",
				Value: bson.M{
					"from":         "merchants",
					"localField":   "merchant_id",
					"foreignField": "_id",
					"as":           "merchant_info",
				},
			},
		},
		{
			{
				Key:   "$unwind",
				Value: "$merchant_info",
			},
		},
		{
			{
				Key: "$group",
				Value: bson.M{
					"_id": bson.M{
						"year": bson.M{
							"$year": "$created_at",
						},
						"month": bson.M{
							"$month": "$created_at",
						},
						"day": bson.M{
							"$dayOfMonth": "$created_at",
						},
					},
					"bill_total": bson.M{
						"$sum": "$bill_total",
					},
					"merchant_name": bson.M{
						"$first": "$merchant_info.merchant_name",
					},
				},
			},
		},
		{
			{
				Key:   "$limit",
				Value: limitNumber,
			},
		},
		{
			{
				Key: "$project",
				Value: bson.M{
					"outlet_name": 0,
				},
			},
		},
	}

	cursor, err = collectionTransactions.Aggregate(context.TODO(), pipeline)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error115)
		return
	}

	err = cursor.All(context.TODO(), &modelResponseTransactionCompletes)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error116)
		return
	}

	if modelResponseTransactionCompletes != nil {
		responseObject = &modelResponseTransactionCompletes

	} else {
		responseObject = bson.A{}

	}

	responseJson, err = json.Marshal(&responseObject)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error117)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(responseJson)

}

func ReadTransactionsComplete(w http.ResponseWriter, r *http.Request) {
	// Variable section
	collectionTransactions := util.Client.Database(util.DatabaseName).Collection(util.CollectionName[3])
	var modelResponseTransactionCompletes []model.ModelResponseTransactionComplete
	var offsetId primitive.ObjectID
	var timestampFromTime time.Time
	var responseObject interface{}
	var queryParameter url.Values
	var userId primitive.ObjectID
	var timestampToTime time.Time
	var pipeline []primitive.D
	var cursor *mongo.Cursor
	var timestampFrom string
	var responseJson []byte
	var timestampTo string
	var claims sjwt.Claims
	var limitNumber int
	var offset string
	var limit string
	var token string
	var _id string
	var err error

	// Parse body
	queryParameter = r.URL.Query()
	timestampFrom = queryParameter.Get("timestamp_from")
	timestampTo = queryParameter.Get("timestamp_to")
	offset = queryParameter.Get("offset")
	limit = queryParameter.Get("limit")
	token = r.Header.Get("Authorization")

	claims, err = sjwt.Parse(token)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error118)
		return
	}
	_id, err = claims.GetStr("_id")
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error119)
		return
	}

	if limit == "" {
		ErrorHandler(nil, w, http.StatusBadRequest, util.Error120)
		return
	}

	limitNumber, err = strconv.Atoi(limit)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error121)
		return
	}

	if offset != "" {
		offsetId, err = primitive.ObjectIDFromHex(offset)
		if err != nil {
			ErrorHandler(err, w, http.StatusInternalServerError, util.Error122)
			return
		}
	} else {
		offsetId = primitive.NilObjectID
	}

	userId, err = primitive.ObjectIDFromHex(_id)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error123)
		return
	}

	timestampFromTime, err = time.Parse(time.RFC3339, timestampFrom)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error124)
		return
	}
	timestampToTime, err = time.Parse(time.RFC3339, timestampTo)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error125)
		return
	}

	pipeline = []bson.D{
		{
			{
				Key: "$match",
				Value: bson.M{
					"_id": bson.M{
						"$gte": offsetId,
					},
					"created_at": bson.M{
						"$gte": primitive.Timestamp{
							T: uint32(timestampFromTime.Unix()),
						},
						"$lte": primitive.Timestamp{
							T: uint32(timestampToTime.Unix()),
						},
					},
					"user_id": userId,
				},
			},
		},
		{
			{
				Key: "$lookup",
				Value: bson.M{
					"from":         "merchants",
					"localField":   "merchant_id",
					"foreignField": "_id",
					"as":           "merchant_info",
				},
			},
		},
		{
			{
				Key:   "$unwind",
				Value: "$merchant_info",
			},
		},
		{
			{
				Key: "$lookup",
				Value: bson.M{
					"from":         "outlets",
					"localField":   "outlet_id",
					"foreignField": "_id",
					"as":           "outlet_info",
				},
			},
		},
		{
			{
				Key:   "$unwind",
				Value: "$outlet_info",
			},
		},
		{
			{
				Key: "$group",
				Value: bson.M{
					"_id": bson.M{
						"year": bson.M{
							"$year": "$created_at",
						},
						"month": bson.M{
							"$month": "$created_at",
						},
						"day": bson.M{
							"$dayOfMonth": "$created_at",
						},
					},
					"bill_total": bson.M{
						"$sum": "$bill_total",
					},
					"merchant_name": bson.M{
						"$first": "$merchant_info.merchant_name",
					},
					"outlet_name": bson.M{
						"$first": "$outlet_info.outlet_name",
					},
				},
			},
		},
		{
			{
				Key:   "$limit",
				Value: limitNumber,
			},
		},
	}

	cursor, err = collectionTransactions.Aggregate(context.TODO(), pipeline)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error126)
		return
	}

	err = cursor.All(context.TODO(), &modelResponseTransactionCompletes)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error127)
		return
	}

	if modelResponseTransactionCompletes != nil {
		responseObject = &modelResponseTransactionCompletes

	} else {
		responseObject = bson.A{}

	}

	responseJson, err = json.Marshal(&responseObject)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error128)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(responseJson)
}
