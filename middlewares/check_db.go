package middlewares

import (
	"net/http"

	"github.com/cristhianforerod/udemy-go-course/db"
)

/*Middleware that check database connection*/
func CheckDataBase(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == false {
			http.Error(w, "Failed DB connection", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
