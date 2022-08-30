package model

type ModelResponseError struct {
	ResponseMessage string `json:"responseMessage,omitempty"`
}

type ModelResponseId struct {
	Year  int `bson:"year,omitempty" json:"year,omitempty"`
	Month int `bson:"month,omitempty" json:"month,omitempty"`
	Day   int `bson:"day,omitempty" json:"day,omitempty"`
}

type ModelResponseReport struct {
	Date          *ModelResponseId `bson:"_id,omitempty" json:"date,omitempty"`
	Bill_total    int              `bson:"bill_total,omitempty" json:"bill_total,omitempty"`
	Merchant_name string           `bson:"merchant_name,omitempty" json:"merchant_name,omitempty"`
	Outlet_name   string           `bson:"outlet_name,omitempty" json:"outlet_name,omitempty"`
}

type ModelResponseLogin struct {
	ResponseMessage string `json:"responseMessage,omitempty"`
	Token           string `json:"token,omitempty"`
}
