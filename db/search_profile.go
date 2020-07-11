package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/cristhianforerod/udemy-go-course/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Search user profile*/
func SearchProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(os.Getenv("DB_NAME"))
	col := db.Collection("users")

	var profile models.User

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""

	if err != nil {
		fmt.Println("Profile not found" + err.Error())
		return profile, err
	}

	return profile, nil

}
