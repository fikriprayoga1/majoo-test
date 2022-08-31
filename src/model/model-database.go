package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ModelDatabaseUser struct {
	Id         primitive.ObjectID  `bson:"_id,omitempty"`
	Name       string              `bson:"name,omitempty"`
	User_name  string              `bson:"user_name,omitempty"`
	Password   string              `bson:"password,omitempty"`
	Created_at primitive.Timestamp `bson:"created_at,omitempty"`
	Created_by int64               `bson:"created_by,omitempty"`
	Updated_at primitive.Timestamp `bson:"updated_at,omitempty"`
	Update_by  int64               `bson:"updated_by,omitempty"`
}

type ModelDatabaseMerchant struct {
	Id            primitive.ObjectID  `bson:"_id,omitempty"`
	User_id       primitive.ObjectID  `bson:"user_id,omitempty"`
	Merchant_name string              `bson:"merchant_name,omitempty"`
	Created_at    primitive.Timestamp `bson:"created_at,omitempty"`
	Created_by    int64               `bson:"created_by,omitempty"`
	Updated_at    primitive.Timestamp `bson:"updated_at,omitempty"`
	Update_by     int64               `bson:"updated_by,omitempty"`
}

type ModelDatabaseOutlet struct {
	Id          primitive.ObjectID  `bson:"_id,omitempty"`
	User_id     primitive.ObjectID  `bson:"user_id,omitempty"`
	Merchant_id primitive.ObjectID  `bson:"merchant_id,omitempty"`
	Outlet_name string              `bson:"outlet_name,omitempty"`
	Created_at  primitive.Timestamp `bson:"created_at,omitempty"`
	Created_by  int64               `bson:"created_by,omitempty"`
	Updated_at  primitive.Timestamp `bson:"updated_at,omitempty"`
	Update_by   int64               `bson:"updated_by,omitempty"`
}

type ModelDatabaseTransaction struct {
	Id          primitive.ObjectID  `bson:"_id,omitempty"`
	User_id     primitive.ObjectID  `bson:"user_id,omitempty"`
	Merchant_id primitive.ObjectID  `bson:"merchant_id,omitempty"`
	Outlet_id   primitive.ObjectID  `bson:"outlet_id,omitempty"`
	Bill_total  int64               `bson:"bill_total,omitempty"`
	Created_at  primitive.Timestamp `bson:"created_at,omitempty"`
	Created_by  int64               `bson:"created_by,omitempty"`
	Updated_at  primitive.Timestamp `bson:"updated_at,omitempty"`
	Update_by   int64               `bson:"updated_by,omitempty"`
}

type ModelDatabaseId struct {
	Year  int `bson:"year,omitempty" json:"year,omitempty"`
	Month int `bson:"month,omitempty" json:"month,omitempty"`
	Day   int `bson:"day,omitempty" json:"day,omitempty"`
}

type ModelDatabaseReport struct {
	Date          *ModelDatabaseId `bson:"_id,omitempty" json:"_id,omitempty"`
	Bill_total    int              `bson:"bill_total,omitempty" json:"bill_total,omitempty"`
	Merchant_name string           `bson:"merchant_name,omitempty" json:"merchant_name,omitempty"`
	Outlet_name   string           `bson:"outlet_name,omitempty" json:"outlet_name,omitempty"`
}
