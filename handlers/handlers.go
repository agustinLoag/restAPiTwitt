package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/agustinLoag/restAPiTwitt/middlewares"

	"github.com/agustinLoag/restAPiTwitt/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Manejadores se setea el puerto y se ponen el servidor */
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlewares.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
