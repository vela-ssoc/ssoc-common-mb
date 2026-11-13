package model

import "go.mongodb.org/mongo-driver/v2/bson"

type Operator struct {
	ID       bson.ObjectID `bson:"id"                json:"id"`
	Company  string        `bson:"company"           json:"company"`
	Username string        `bson:"username"          json:"username"`
	Nickname string        `bson:"nickname"          json:"nickname"`
	Service  bool          `bson:"service,omitempty" json:"service,omitzero"`
}
