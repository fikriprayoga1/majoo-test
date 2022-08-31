package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ModelRequestLogin struct {
	User_name string `json:"user_name,omitempty"`
	Password  string `json:"password,omitempty"`
}

type ModelRequestRegister struct {
	Name      string `json:"name,omitempty"`
	User_name string `json:"user_name,omitempty"`
	Password  string `json:"password,omitempty"`
}

type ModelRequestCreateMerchant struct {
	Merchant_name string `json:"merchant_name,omitempty"`
}

type ModelRequestCreateOutlet struct {
	Merchant_id primitive.ObjectID `json:"merchant_id,omitempty"`
	Outlet_name string             `json:"outlet_name,omitempty"`
}

type ModelRequestCreateTransaction struct {
	Merchant_id primitive.ObjectID `json:"merchant_id,omitempty"`
	Outlet_id   primitive.ObjectID `json:"outlet_id,omitempty"`
	Bill_total  int64              `json:"bill_total,omitempty"`
}

type ModelRequestUpdateUser struct {
	Name      string `json:"name,omitempty"`
	User_name string `json:"user_name,omitempty"`
	Password  string `json:"password,omitempty"`
}

type ModelRequestUpdateMerchant struct {
	Merchant_id   primitive.ObjectID `json:"merchant_id,omitempty"`
	Merchant_name string             `json:"merchant_name,omitempty"`
}

type ModelRequestUpdateOutlet struct {
	Merchant_id primitive.ObjectID `json:"merchant_id,omitempty"`
	Outlet_id   primitive.ObjectID `json:"outlet_id,omitempty"`
	Outlet_name string             `json:"outlet_name,omitempty"`
}

type ModelRequestUpdateTransaction struct {
	Merchant_id    primitive.ObjectID `json:"merchant_id,omitempty"`
	Outlet_id      primitive.ObjectID `json:"outlet_id,omitempty"`
	Transaction_id primitive.ObjectID `json:"transaction_id,omitempty"`
	Bill_total     int64              `json:"bill_total,omitempty"`
}

type ModelRequestDelete struct {
	Item_id primitive.ObjectID `json:"item_id,omitempty"`
}
