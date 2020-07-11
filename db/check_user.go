package db

import (
	"context"
	"os"
	"time"

	"github.com/cristhianforerod/udemy-go-course/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckUser(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(os.Getenv("DB_NAME"))
	col := db.Collection("users")

	condition := bson.M{"email": email}

	var result models.User

	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}
	return result, true, ID

}
