package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*User Model*/
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name,omitempty"`
	Lastname  string             `bson:"lastname" json:"lastname,omitempty"`
	Birthdate time.Time          `bson:"birthdate" json:"birthdate,omitempty"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar,omitempty"`
	Banner    string             `bson:"banner" json:"banner,omitempty"`
	Address   string             `bson:"address" json:"address,omitempty"`
	Biography string             `bson:"biography" json:"biography,omitempty"`
	WebSite   string             `bson:"website" json:"website,omitempty"`
}
