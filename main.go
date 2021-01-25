package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Eleve struct {
	Id string `json:"id"`
	Nom    string `json:"nom"`
	Prenom string `json:"prenom"`
	Email  string `json:"email"`
}

var Eleves []Eleve

func AllEleve(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Tous les élèves")
	json.NewEncoder(w).Encode(Eleves)
}

func EleveById(w http.ResponseWriter, r *http.Request)  {

	fmt.Println("Elève by ID")
	vars := mux.Vars(r)
	key := vars["id"]
	for _, eleve := range Eleves {
		if eleve.Id == key {
			json.NewEncoder(w).Encode(eleve)
		}
	}
}

func CreateEleve(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Créer un élève")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var eleve Eleve
	json.Unmarshal(reqBody, &eleve)
	// update our global Articles array to include
	// our new Article
	Eleves = append(Eleves, eleve)

	json.NewEncoder(w).Encode(eleve)
}

func DeleteEleve(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, eleve := range Eleves {
		if eleve.Id == id {
			Eleves = append(Eleves[:index], Eleves[index+1:]...)
		}
	}

}

func UpdateEleve(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var updatedEleve Eleve

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(w, "Update eleve")
	}
	json.Unmarshal(reqBody, &updatedEleve)

	for index, eleve := range Eleves {
		if eleve.Id == id {
			eleve.Nom = updatedEleve.Nom
			eleve.Prenom = updatedEleve.Prenom
			eleve.Email = updatedEleve.Email
			Eleves = append(Eleves[:index], eleve)
			json.NewEncoder(w).Encode(eleve)
		}
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("HomePage")
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/eleves", AllEleve).Methods("GET")
	myRouter.HandleFunc("/eleve/{id}", EleveById).Methods("GET")
	myRouter.HandleFunc("/eleve", CreateEleve).Methods("POST")
	myRouter.HandleFunc("/eleve/{id}", UpdateEleve).Methods("PATCH")
	myRouter.HandleFunc("/eleve/{id}", DeleteEleve).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	Eleves = []Eleve{
		{	Id: "1",
			Nom: "Dubois",
			Prenom: "Jean",
			Email:  "jean@dubois.com"},
		{	Id: "2",
			Nom: "Lenoy",
			Prenom: "Franck",
			Email:  "franck@lenoy.com"},
	}
	handleRequests()
}
