package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Restaurant struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Address Address            `json:"address" bson:"address"`
	Borough string             `json:"borough" bson:"borough"`
	Cuisine string             `json:"cuisine" bson:"cuisine"`
	Grades  []Grade            `json:"grades" bson:"grades"`
}

type Address struct {
	Building string    `json:"building" bson:"building"`
	Coord    []float64 `json:"coord" bson:"coord"`
	Street   string    `json:"street" bson:"street"`
	Zipcode  string    `json:"zipcode" bson:"zipcode"`
}

type Grade struct {
	Date  time.Time `json:"date" bson:"date"`
	Grade string    `json:"grade" bson:"grade"`
	Score int       `json:"score" bson:"score"`
}
type RestaurantWithSum struct {
	Restaurant  Restaurant
	SumOfScores int
}
