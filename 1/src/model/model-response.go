package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ModelResponse struct {
	ResponseMessage string `json:"responseMessage,omitempty"`
}

type ModelResponseLogin struct {
	ResponseMessage string             `json:"responseMessage,omitempty"`
	Token           string             `json:"token,omitempty"`
	Profile         *ModelResponseUser `json:"profile,omitempty"`
}

type ModelResponseUser struct {
	Name      string `json:"name,omitempty"`
	User_name string `json:"user_name,omitempty"`
}

type ModelResponseMerchant struct {
	Id            primitive.ObjectID  `json:"_id,omitempty"`
	Merchant_name string              `json:"merchant_name,omitempty"`
	Created_at    primitive.Timestamp `json:"created_at,omitempty"`
	Created_by    int64               `json:"created_by,omitempty"`
	Updated_at    primitive.Timestamp `json:"updated_at,omitempty"`
	Update_by     int64               `json:"updated_by,omitempty"`
}

type ModelResponseOutlet struct {
	Id          primitive.ObjectID  `json:"_id,omitempty"`
	Outlet_name string              `json:"outlet_name,omitempty"`
	Created_at  primitive.Timestamp `json:"created_at,omitempty"`
	Created_by  int64               `json:"created_by,omitempty"`
	Updated_at  primitive.Timestamp `json:"updated_at,omitempty"`
	Update_by   int64               `json:"updated_by,omitempty"`
}

type ModelResponseTransaction struct {
	Id         primitive.ObjectID  `json:"_id,omitempty"`
	Bill_total int64               `json:"bill_total,omitempty"`
	Created_at primitive.Timestamp `json:"created_at,omitempty"`
	Created_by int64               `json:"created_by,omitempty"`
	Updated_at primitive.Timestamp `json:"updated_at,omitempty"`
	Update_by  int64               `json:"updated_by,omitempty"`
}
