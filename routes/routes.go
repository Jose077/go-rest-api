package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jose077/go-rest-api/controllers"
	"github.com/jose077/go-rest-api/middleware"
)

func HandleRequests() {

	// istancia do roteador do mux
	r := mux.NewRouter()

	// middleware
	r.Use(middleware.ContentTypeMiddleware)

	// Routes
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("GET")
	r.HandleFunc("/api/personalidades", controllers.CriaUmaNovaPersonalidade).Methods("POST")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeleteUmaPersonalidade).Methods("DELETE")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaPersonalidade).Methods("PUT")

	// Sobe servidor com CORS habilitado
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
