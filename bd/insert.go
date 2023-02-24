package bd

import (
	"context"
	"time"

	"github.com/andrelensanro/building_twitter_copy/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertRegister(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() /*se ejecuta si o si y al  fina; de la funcon*/

	db := MongoCONN.Database("twitter")
	collection := db.Collection("users")

	u.Password, _ = EncryptPassword(u.Password)
	result, err := collection.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}
