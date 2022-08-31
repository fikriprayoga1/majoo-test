package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"src/middleware"
	"src/model"
	"testing"

	"github.com/stretchr/testify/require"
)

var token string
var merchantId string
var outletId string
var transactionId string

func TestMain(m *testing.M) {
	initDatabase()

	m.Run()
}

func TestRegister(t *testing.T) {
	var rr *httptest.ResponseRecorder
	var req *http.Request
	var handler http.HandlerFunc
	var err error

	jsonStr := []byte(`{
		"name":"Joko",
		"user_name": "budi0",
		"password": "mantab"
	}`)
	req, err = http.NewRequest("POST", "/register", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(middleware.Register)
	handler.ServeHTTP(rr, req)

	log.Printf("logInfo : responseBody => %v", rr.Body.String())
	require.Equal(t, http.StatusOK, rr.Code)

}

func TestLogin(t *testing.T) {
	var rr *httptest.ResponseRecorder
	var req *http.Request
	var handler http.HandlerFunc
	var err error
	var modelResponseLogin model.ModelResponseLogin

	jsonStr := []byte(`{
		"user_name": "budi0",
    	"password": "mantab"
	}`)
	req, err = http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(middleware.Login)
	handler.ServeHTTP(rr, req)

	log.Printf("logInfo : responseBody => %v", rr.Body.String())
	require.Equal(t, http.StatusOK, rr.Code)

	err = json.NewDecoder(rr.Body).Decode(&modelResponseLogin)
	if err != nil {
		log.Printf("logInfo : %v", err)
		t.FailNow()
	}
	token = modelResponseLogin.Token

}

func TestCreateMerchant(t *testing.T) {
	var rr *httptest.ResponseRecorder
	var req *http.Request
	var handler http.HandlerFunc
	var err error

	jsonStr := []byte(`{
		"merchant_name": "Toko Kuda"
	}`)
	req, err = http.NewRequest("POST", "/create/merchant", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(middleware.CreateMerchant)
	handler.ServeHTTP(rr, req)

	log.Printf("logInfo : responseBody => %v", rr.Body.String())
	require.Equal(t, http.StatusOK, rr.Code)

}

func TestReadMerchant(t *testing.T) {
	var rr *httptest.ResponseRecorder
	var req *http.Request
	var handler http.HandlerFunc
	var modelResponseMerchants []model.ModelResponseMerchant
	var err error

	req, err = http.NewRequest("GET", "/read/merchants", nil)
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(middleware.ReadMerchants)
	handler.ServeHTTP(rr, req)

	log.Printf("logInfo : responseBody => %v", rr.Body.String())
	require.Equal(t, http.StatusOK, rr.Code)

	err = json.NewDecoder(rr.Body).Decode(&modelResponseMerchants)
	if err != nil {
		log.Printf("logInfo : %v", err)
		t.FailNow()
	}
	merchantId = modelResponseMerchants[0].Id.Hex()

}

func TestCreateOutlets(t *testing.T) {
	var rr *httptest.ResponseRecorder
	var req *http.Request
	var handler http.HandlerFunc
	var err error

	jsonStr := []byte(`{
		"merchant_id": "` + merchantId + `",
    	"outlet_name": "Cabang Sakti 2"
	}`)
	req, err = http.NewRequest("POST", "/create/outlet", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(middleware.CreateOutlet)
	handler.ServeHTTP(rr, req)

	log.Printf("logInfo : responseBody => %v", rr.Body.String())
	require.Equal(t, http.StatusOK, rr.Code)

}

func TestReadOutlets(t *testing.T) {
	var rr *httptest.ResponseRecorder
	var req *http.Request
	var handler http.HandlerFunc
	var modelResponseOutlets []model.ModelResponseOutlet
	var err error

	req, err = http.NewRequest("GET", "/read/outlets", nil)
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	q := req.URL.Query()
	q.Add("merchant_id", merchantId)
	req.URL.RawQuery = q.Encode()

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(middleware.ReadOutlets)
	handler.ServeHTTP(rr, req)

	log.Printf("logInfo : responseBody => %v", rr.Body.String())
	require.Equal(t, http.StatusOK, rr.Code)

	err = json.NewDecoder(rr.Body).Decode(&modelResponseOutlets)
	if err != nil {
		log.Printf("logInfo : %v", err)
		t.FailNow()
	}
	outletId = modelResponseOutlets[0].Id.Hex()

}

func TestCreateTransactions(t *testing.T) {
	var rr *httptest.ResponseRecorder
	var req *http.Request
	var handler http.HandlerFunc
	var err error

	jsonStr := []byte(`{
		"merchant_id": "` + merchantId + `",
    	"outlet_id": "` + outletId + `",
    	"bill_total": 10000
	}`)
	req, err = http.NewRequest("POST", "/create/transaction", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(middleware.CreateTransaction)
	handler.ServeHTTP(rr, req)

	log.Printf("logInfo : responseBody => %v", rr.Body.String())
	require.Equal(t, http.StatusOK, rr.Code)

}

func TestReadTransactions(t *testing.T) {
	var rr *httptest.ResponseRecorder
	var req *http.Request
	var handler http.HandlerFunc
	var modelResponseTransactions []model.ModelResponseTransaction
	var err error

	req, err = http.NewRequest("GET", "/read/transactions", nil)
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	q := req.URL.Query()
	q.Add("merchant_id", merchantId)
	q.Add("outlet_id", outletId)
	req.URL.RawQuery = q.Encode()

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(middleware.ReadTransactions)
	handler.ServeHTTP(rr, req)

	log.Printf("logInfo : responseBody => %v", rr.Body.String())
	require.Equal(t, http.StatusOK, rr.Code)

	err = json.NewDecoder(rr.Body).Decode(&modelResponseTransactions)
	if err != nil {
		log.Printf("logInfo : %v", err)
		t.FailNow()
	}
	transactionId = modelResponseTransactions[0].Id.Hex()

}

func TestReadTransactionsSimple(t *testing.T) {
	var rr *httptest.ResponseRecorder
	var req *http.Request
	var handler http.HandlerFunc
	var err error

	req, err = http.NewRequest("GET", "/read/transactions/simple", nil)
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	q := req.URL.Query()
	q.Add("timestamp_from", "2000-01-01T00:00:00Z")
	q.Add("timestamp_to", "2023-01-01T00:00:00Z")
	q.Add("limit", "10")
	req.URL.RawQuery = q.Encode()

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(middleware.ReadTransactionsSimple)
	handler.ServeHTTP(rr, req)

	log.Printf("logInfo : responseBody => %v", rr.Body.String())
	require.Equal(t, http.StatusOK, rr.Code)

}

func TestReadTransactionsComplete(t *testing.T) {
	var rr *httptest.ResponseRecorder
	var req *http.Request
	var handler http.HandlerFunc
	var err error

	req, err = http.NewRequest("GET", "/read/transactions/complete", nil)
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	q := req.URL.Query()
	q.Add("timestamp_from", "2000-01-01T00:00:00Z")
	q.Add("timestamp_to", "2023-01-01T00:00:00Z")
	q.Add("limit", "10")
	req.URL.RawQuery = q.Encode()

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(middleware.ReadTransactionsComplete)
	handler.ServeHTTP(rr, req)

	log.Printf("logInfo : responseBody => %v", rr.Body.String())
	require.Equal(t, http.StatusOK, rr.Code)

}

func TestUpdateUser(t *testing.T) {
	var rr *httptest.ResponseRecorder
	var req *http.Request
	var handler http.HandlerFunc
	var err error

	jsonStr := []byte(`{
		"name":"Mayap",
    	"user_name": "budi0",
    	"password": "mantab"
	}`)
	req, err = http.NewRequest("POST", "/update/user", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(middleware.UpdateUser)
	handler.ServeHTTP(rr, req)

	log.Printf("logInfo : responseBody => %v", rr.Body.String())
	require.Equal(t, http.StatusOK, rr.Code)

}

func TestUpdateMerchant(t *testing.T) {
	var rr *httptest.ResponseRecorder
	var req *http.Request
	var handler http.HandlerFunc
	var err error

	jsonStr := []byte(`{
		"merchant_id":"` + merchantId + `",
    	"merchant_name": "budi2"
	}`)
	req, err = http.NewRequest("POST", "/update/merchant", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(middleware.UpdateMerchant)
	handler.ServeHTTP(rr, req)

	log.Printf("logInfo : responseBody => %v", rr.Body.String())
	require.Equal(t, http.StatusOK, rr.Code)

}

func TestUpdateOutlet(t *testing.T) {
	var rr *httptest.ResponseRecorder
	var req *http.Request
	var handler http.HandlerFunc
	var err error

	jsonStr := []byte(`{
		"merchant_id":"` + merchantId + `",
		"outlet_id":"` + outletId + `",
    	"outlet_name": "budi2"
	}`)
	req, err = http.NewRequest("POST", "/update/outlet", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(middleware.UpdateOutlet)
	handler.ServeHTTP(rr, req)

	log.Printf("logInfo : responseBody => %v", rr.Body.String())
	require.Equal(t, http.StatusOK, rr.Code)

}

func TestUpdateTransaction(t *testing.T) {
	var rr *httptest.ResponseRecorder
	var req *http.Request
	var handler http.HandlerFunc
	var err error

	jsonStr := []byte(`{
		"merchant_id":"` + merchantId + `",
		"outlet_id":"` + outletId + `",
		"transaction_id":"` + transactionId + `",
    	"bill_total": 3000
	}`)
	req, err = http.NewRequest("POST", "/update/transaction", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(middleware.UpdateTransaction)
	handler.ServeHTTP(rr, req)

	log.Printf("logInfo : responseBody => %v", rr.Body.String())
	require.Equal(t, http.StatusOK, rr.Code)

}

func TestDeleteTransaction(t *testing.T) {
	var rr *httptest.ResponseRecorder
	var req *http.Request
	var handler http.HandlerFunc
	var err error

	jsonStr := []byte(`{		
		"item_id":"` + transactionId + `"    	
	}`)
	req, err = http.NewRequest("POST", "/delete/transaction", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(middleware.DeleteTransaction)
	handler.ServeHTTP(rr, req)

	log.Printf("logInfo : responseBody => %v", rr.Body.String())
	require.Equal(t, http.StatusOK, rr.Code)

}

func TestDeleteOutlet(t *testing.T) {
	var rr *httptest.ResponseRecorder
	var req *http.Request
	var handler http.HandlerFunc
	var err error

	jsonStr := []byte(`{		
		"item_id":"` + outletId + `"    	
	}`)
	req, err = http.NewRequest("POST", "/delete/outlet", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(middleware.DeleteOutlet)
	handler.ServeHTTP(rr, req)

	log.Printf("logInfo : responseBody => %v", rr.Body.String())
	require.Equal(t, http.StatusOK, rr.Code)

}

func TestDeleteMerchant(t *testing.T) {
	var rr *httptest.ResponseRecorder
	var req *http.Request
	var handler http.HandlerFunc
	var err error

	jsonStr := []byte(`{		
		"item_id":"` + merchantId + `"    	
	}`)
	req, err = http.NewRequest("POST", "/delete/merchant", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(middleware.DeleteMerchant)
	handler.ServeHTTP(rr, req)

	log.Printf("logInfo : responseBody => %v", rr.Body.String())
	require.Equal(t, http.StatusOK, rr.Code)

}

func TestDeleteUser(t *testing.T) {
	var rr *httptest.ResponseRecorder
	var req *http.Request
	var handler http.HandlerFunc
	var err error

	req, err = http.NewRequest("POST", "/delete/user", nil)
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(middleware.DeleteUser)
	handler.ServeHTTP(rr, req)

	log.Printf("logInfo : responseBody => %v", rr.Body.String())
	require.Equal(t, http.StatusOK, rr.Code)

}
