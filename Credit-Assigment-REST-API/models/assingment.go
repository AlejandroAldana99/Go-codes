package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AssigmentData struct {
	ID                 MyObjectID         `json:"id,omitempty" bson:"_id,omitempty"`
	OriginalInvestment float64            `json:"original_investment"`
	Owner              string             `json:"owner,omitempty"`
	OwnerID            string             `json:"ownerid,omitempty"`
	Combinations       []CombinationsData `json:"combinations,omitempty"`
	NonCombinated      bool               `json:"noncombinated" default:"false"`
}

type CombinationsData struct {
	CreditType300 int `json:"credit_type_300"`
	CreditType500 int `json:"credit_type_500"`
	CreditType700 int `json:"credit_type_700"`
}

type MyObjectID string

func (id MyObjectID) MarshalBSONValue() (bsontype.Type, []byte, error) {
	p, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		return bsontype.Null, nil, err
	}

	return bson.MarshalValue(p)
}
