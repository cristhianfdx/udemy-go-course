package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/cristhianforerod/udemy-go-course/middlewares"
	"github.com/cristhianforerod/udemy-go-course/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Handler dispatch to assign server port */
func Dispatcher() {
	router := mux.NewRouter()
	var urlBase = "/api"

	router.HandleFunc(urlBase+"/sign_up", middlewares.CheckDataBase(routes.Register)).Methods("POST")
	router.HandleFunc(urlBase+"/login", middlewares.CheckDataBase(routes.Login)).Methods("POST")
	router.HandleFunc(urlBase+"/profile", middlewares.ValidateJWT(routes.ViewProfile)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
