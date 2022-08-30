package middleware

import (
	"context"
	"encoding/json"
	"log"
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

func Login(w http.ResponseWriter, r *http.Request) {
	var modelRequestLogin model.ModelRequestLogin
	var err error
	var claims *sjwt.Claims
	var jwt string
	var responseJson []byte
	var modelResponseLogin model.ModelResponseLogin

	// Parse body
	err = json.NewDecoder(r.Body).Decode(&modelRequestLogin)
	if err != nil {
		ErrorHandler(err, w, http.StatusBadRequest, util.Error0)
		return
	}
	log.Printf("logInfo : Email => %v\n", modelRequestLogin.Email)
	log.Printf("logInfo : Password => %v\n", modelRequestLogin.Password)

	if (modelRequestLogin.Email == "mantap@gmail.com") && (modelRequestLogin.Password == "wokeh") {
		// Add Claims
		claims = sjwt.New()
		claims.Set("id", "630c0ec7a335ddfa8f96e7bc")

		// Generate jwt
		jwt = claims.Generate(util.GetScretKey())
		log.Println(jwt)

		modelResponseLogin = model.ModelResponseLogin{
			ResponseMessage: "Login Success",
			Token:           jwt,
		}

		responseJson, err = json.Marshal(&modelResponseLogin)
		if err != nil {
			ErrorHandler(err, w, http.StatusInternalServerError, util.Error1)
			return
		}

		log.Printf("logInfo : => %v\n", string(responseJson))

		w.Header().Set("content-type", "application/json")
		w.Write(responseJson)
	} else {
		ErrorHandler(err, w, http.StatusForbidden, util.Error2)
	}

}

func ReadReport(w http.ResponseWriter, r *http.Request) {
	var queryParameter url.Values
	var offset string
	var limit string
	var limitNumber int
	var timestampFrom string
	var timestampFromTime time.Time
	var timestampTo string
	var timestampToTime time.Time
	var offsetId primitive.ObjectID
	var userId primitive.ObjectID
	var id string
	var claims sjwt.Claims
	var token string
	var pipeline []primitive.D
	var cursor *mongo.Cursor
	var err error
	var responseJson []byte
	var results []model.ModelResponseReport
	collectionTransactions := util.Client.Database(util.DatabaseName).Collection(util.CollectionName[3])

	// Parse body
	queryParameter = r.URL.Query()
	timestampFrom = queryParameter.Get("timestamp_from")
	timestampTo = queryParameter.Get("timestamp_to")
	offset = queryParameter.Get("offset")
	limit = queryParameter.Get("limit")
	token = r.Header.Get("Authorization")

	claims, err = sjwt.Parse(token)
	if err != nil {
		ErrorHandler(err, w, http.StatusForbidden, util.Error3)
		return
	}

	// Get claims
	id, err = claims.GetStr("id")
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error4)
		return
	}

	if limit == "" {
		ErrorHandler(nil, w, http.StatusBadRequest, util.Error5)
		return
	}
	log.Printf("logInfo : timestampFrom => %v\n", timestampFrom)
	log.Printf("logInfo : timestampTo => %v\n", timestampTo)
	log.Printf("logInfo : offset => %v\n", offset)
	log.Printf("logInfo : limit => %v\n", limit)

	limitNumber, err = strconv.Atoi(limit)
	if err != nil {
		ErrorHandler(err, w, http.StatusForbidden, util.Error6)
		return
	}

	if offset != "" {
		offsetId, err = primitive.ObjectIDFromHex(offset)
		if err != nil {
			ErrorHandler(err, w, http.StatusInternalServerError, util.Error7)
			return
		}
	} else {
		offsetId = primitive.NilObjectID
	}

	userId, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error8)
		return
	}

	timestampFromTime, err = time.Parse(time.RFC1123Z, timestampFrom)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error9)
		return
	}
	timestampToTime, err = time.Parse(time.RFC1123Z, timestampTo)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error10)
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
					"updated_at": bson.M{
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
	}

	cursor, err = collectionTransactions.Aggregate(context.TODO(), pipeline)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error11)
		return
	}

	err = cursor.All(context.TODO(), &results)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error12)
		return
	}

	responseJson, err = json.Marshal(&results)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error13)
		return
	}

	log.Printf("logInfo : => %v\n", string(responseJson))

	w.Header().Set("content-type", "application/json")
	w.Write(responseJson)

}

func ReadReportAdvanced(w http.ResponseWriter, r *http.Request) {
	var queryParameter url.Values
	var offset string
	var limit string
	var limitNumber int
	var timestampFrom string
	var timestampFromTime time.Time
	var timestampTo string
	var timestampToTime time.Time
	var offsetId primitive.ObjectID
	var userId primitive.ObjectID
	var id string
	var claims sjwt.Claims
	var token string
	var pipeline []primitive.D
	var cursor *mongo.Cursor
	var err error
	var responseJson []byte
	var results []model.ModelResponseReport
	collectionTransactions := util.Client.Database(util.DatabaseName).Collection(util.CollectionName[3])

	// Parse body
	queryParameter = r.URL.Query()
	timestampFrom = queryParameter.Get("timestamp_from")
	timestampTo = queryParameter.Get("timestamp_to")
	offset = queryParameter.Get("offset")
	limit = queryParameter.Get("limit")
	token = r.Header.Get("Authorization")

	claims, err = sjwt.Parse(token)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error14)
		return
	}

	// Get claims
	id, err = claims.GetStr("id")
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error15)
		return
	}

	if limit == "" {
		ErrorHandler(nil, w, http.StatusBadRequest, util.Error16)
		return
	}
	log.Printf("logInfo : timestampFrom => %v\n", timestampFrom)
	log.Printf("logInfo : timestampTo => %v\n", timestampTo)
	log.Printf("logInfo : offset => %v\n", offset)
	log.Printf("logInfo : limit => %v\n", limit)

	limitNumber, err = strconv.Atoi(limit)
	if err != nil {
		ErrorHandler(err, w, http.StatusForbidden, util.Error17)
		return
	}

	if offset != "" {
		offsetId, err = primitive.ObjectIDFromHex(offset)
		if err != nil {
			ErrorHandler(err, w, http.StatusInternalServerError, util.Error18)
			return
		}
	} else {
		offsetId = primitive.NilObjectID
	}

	userId, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error19)
		return
	}

	timestampFromTime, err = time.Parse(time.RFC1123Z, timestampFrom)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error20)
		return
	}
	timestampToTime, err = time.Parse(time.RFC1123Z, timestampTo)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error21)
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
					"updated_at": bson.M{
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
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error22)
		return
	}

	err = cursor.All(context.TODO(), &results)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error23)
		return
	}

	responseJson, err = json.Marshal(&results)
	if err != nil {
		ErrorHandler(err, w, http.StatusInternalServerError, util.Error24)
		return
	}

	log.Printf("logInfo : => %v\n", string(responseJson))

	w.Header().Set("content-type", "application/json")
	w.Write(responseJson)
}
