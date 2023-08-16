package service

import (
	"context"
	"fmt"
	"mongogo/config"
	"mongogo/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ProductContext() *mongo.Collection {
	client, _ := config.ConnectDataBase()
	return config.GetCollection(client, "inventory", "products")
} /*
func InsertProduct() {
	var product models.Product
	product.Name = "Samsung"
	product.Price = 1500000
	product.Description = "Awesome"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := config.ConnectDataBase()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	productCollection := config.GetCollection(client, "sample_restaurants", "restaurants")

	result, err := productCollection.InsertOne(ctx, product)
	if err != nil {
		fmt.Println("Error inserting product:", err)
		return
	}

	fmt.Println("Inserted product ID:", result.InsertedID)
}
*/

func FindRestaurants() ([]*models.RestaurantWithSum, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{"cuisine", "Bakery"}}

	client, err := config.ConnectDataBase()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return nil, err
	}
	defer client.Disconnect(ctx)

	collection := config.GetCollection(client, "sample_restaurants", "restaurants")

	options := options.Find() // You can configure options here if needed
	result, err := collection.Find(ctx, filter, options)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer result.Close(ctx)

	var restaurantsWithSum []*models.RestaurantWithSum

	for result.Next(ctx) {
		restaurant := &models.Restaurant{}
		err := result.Decode(restaurant)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		sum := 0
		for _, grade := range restaurant.Grades {
			sum += grade.Score
		}

		restaurantWithSum := &models.RestaurantWithSum{
			Restaurant:  *restaurant,
			SumOfScores: sum,
		}

		restaurantsWithSum = append(restaurantsWithSum, restaurantWithSum)
	}

	if err := result.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return restaurantsWithSum, nil
}
