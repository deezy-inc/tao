package model

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	User struct {
		ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Username string        `json:"username" form:"username" bson:"username"`
		Email    string        `json:"email" form:"email" bson:"email"`
		Password string        `json:"password,omitempty" form:"password" bson:"password"`
		FeeRate  int           `json:"fee_rate" form:"fee_rate" bson:"fee_rate"`
		Avatar   string        `json:"avatar, omitempty" form:"avatar" bson:"avatar"`
		PgpKey   string        `json:"pgpkey, omitempty" form:"pgpkey" bson:"pgpkey"`
		Twitter  string        `json:"twitter, omitempty" form:"twitter" bson:"twitter"`

		Token     string   `json:"token,omitempty" bson:"-"`
		Followers []string `json:"followers,omitempty" bson:"followers,omitempty"`
	}

	UpdateUser struct {
		FeeRate int    `json:"fee_rate" form:"fee_rate" bson:"fee_rate"`
		Avatar  string `json:"avatar, omitempty" form:"avatar" bson:"avatar"`
		PgpKey  string `json:"pgpkey, omitempty" form:"pgpkey" bson:"pgpkey"`
		Twitter string `json:"twitter, omitempty" form:"twitter" bson:"twitter"`
	}

	PublicUser struct {
		Username string `json:"username" bson:"username"`
		FeeRate  int    `json:"fee_rate" bson:"fee_rate"`
		Avatar   string `json:"avatar, omitempty" bson:"avatar"`
		PgpKey   string `json:"pgpkey, omitempty" bson:"pgpkey"`
		Twitter  string `json:"twitter, omitempty" bson:"twitter"`

		Token string `json:"token,omitempty" bson:"-"`
	}

	Withdrawal struct {
		ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Username string        `json:"username" bson:"username"`
		Sats     int           `json:"sats" bson:"sats"`
		Invoice  string        `json:"invoice, omitempty" bson:"invoice"`
		Paid     bool          `json:"paid, omitempty" bson:"paid"`
	}
)
