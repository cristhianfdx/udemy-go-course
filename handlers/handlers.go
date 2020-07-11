package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/cristhianforerod/udemy-go-course/middlewares"
	"github.com/cristhianforerod/udemy-go-course/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Handler dispatch to assign server port */
func Dispatcher() {
	router := mux.NewRouter()

	router.HandleFunc("/sign_up", middlewares.CheckDataBase(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
