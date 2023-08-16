package main

import (
	"fmt"
	"mongogo/service"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoClient *mongo.Client
)

func main() {
	products, _ := service.FindRestaurants()
	for _, product := range products {
		fmt.Println(product)
	}
	fmt.Println("Successfully conneted")
}
