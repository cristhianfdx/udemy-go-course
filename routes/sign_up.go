package routes

import (
	"encoding/json"
	"net/http"

	"github.com/cristhianforerod/udemy-go-course/db"
	"github.com/cristhianforerod/udemy-go-course/models"
)

/*Users Sign Up route */
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Data error: "+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "Email is required", 417)
		return
	}

	if len(user.Password) < 6 {
		http.Error(w, "Password must be at least 6 characters", 417)
		return
	}

	_, found, _ := db.CheckUser(user.Email)
	if found {
		http.Error(w, "User already email exists", 417)
		return
	}

	_, status, err := db.Save(user)
	if err != nil || !status {
		http.Error(w, "Failed to register user"+err.Error(), 417)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
