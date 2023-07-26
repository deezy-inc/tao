package model

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	Transfer struct {
		ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
		To      string        `json:"to" bson:"to"`
		From    string        `json:"from" bson:"from"`
		Message string        `json:"message" bson:"message"`
		Sats    int           `json:"sats" bson:"sats"`
		Invoice string        `json:"invoice" bson:"invoice"`
		Paid    bool          `json:"paid" bson:"paid"`
		Read    bool          `json:"read" bson:"read"`
		Public  string        `json:"public" bson:"public"`

		// LN fields
		RHash          byte   `json:"r_hash"`
		PaymentRequest string `json:"payment_request"`
		AddIndex       string `json:"add_index"`
	}
)
