package controllers

import (
	"api-rest/database"
	"api-rest/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func FetchPersons(w http.ResponseWriter, r *http.Request) {
	var p []models.Person
	database.DB.Find(&p)

	json.NewEncoder(w).Encode(p)
}

func FetchPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Person

	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var p models.Person
	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Create(&p)
	json.NewEncoder(w).Encode(p)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Person

	database.DB.Delete(&p, id)
	json.NewEncoder(w).Encode(p)
}

func EditPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Person

	database.DB.First(&p, id)
	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Save(&p)
	json.NewEncoder(w).Encode(p)
}
