package bd

import (
	"context"
	"time"

	"github.com/andrelensanro/building_twitter_copy/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckExists(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCONN.Database("twitter")
	collection := db.Collection("users")
	condition := bson.M{"email": email}
	var result models.User
	err := collection.FindOne(ctx, condition).Decode(&result)
	ID := result.Id.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
