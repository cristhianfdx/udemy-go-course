package jwt

import (
	"os"
	"time"

	"github.com/cristhianforerod/udemy-go-course/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*Generate JSON Web Token*/
func GenerateJWT(user models.User) (string, error) {
	payload := jwt.MapClaims{
		"email":     user.Email,
		"name":      user.Name,
		"lastname":  user.Lastname,
		"birthdate": user.Birthdate,
		"biography": user.Biography,
		"website":   user.WebSite,
		"_id":       user.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString([]byte(os.Getenv("MY_KEY")))
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
