package routes

import (
	"errors"
	"os"
	"strings"

	"github.com/cristhianforerod/udemy-go-course/db"
	"github.com/cristhianforerod/udemy-go-course/models"
	jwt "github.com/dgrijalva/jwt-go"
)

var Email string
var UserID string

/*Token Process route*/
func TokenProcess(token string) (*models.Claim, bool, string, error) {
	key := []byte(os.Getenv("MY_KEY"))
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Token format invalid")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err == nil {
		_, found, _ := db.CheckUser(claims.Email)
		if found {
			Email = claims.Email
			UserID = claims.ID.Hex()
		}
		return claims, found, UserID, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("Invalid token")
	}

	return claims, false, string(""), err
}
