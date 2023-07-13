package routes

import (
	"api-rest/controllers"
	"api-rest/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/persons", controllers.FetchPersons).Methods("Get")
	r.HandleFunc("/api/persons", controllers.CreatePerson).Methods("Post")
	r.HandleFunc("/api/persons/{id}", controllers.FetchPerson).Methods("Get")
	r.HandleFunc("/api/persons/{id}", controllers.DeletePerson).Methods("Delete")
	r.HandleFunc("/api/persons/{id}", controllers.EditPerson).Methods("Put")
	log.Fatal(
		http.ListenAndServe(":8000", handlers.CORS(
			handlers.AllowedOrigins(
				[]string{"http://localhost:3000"},
			),
		)(r)),
	)
}
