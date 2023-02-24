package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/andrelensanro/building_twitter_copy/middlew"
	"github.com/andrelensanro/building_twitter_copy/routers"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/cors"
)

/*Pongo a escuchar el servidor*/
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.CheckDB(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT")
	log.Println("Valor de PORT" + PORT)
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
