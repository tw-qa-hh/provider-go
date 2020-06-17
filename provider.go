package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const fileStore = "addresses.json"

type address struct {
	ID     string `json:"ID"`
	Zip    string `json:"Zip"`
	Street string `json:"Street"`
}

type allAddresses []address

var addresses allAddresses

func createAddress(w http.ResponseWriter, r *http.Request) {
	var newAddress address
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newAddress)
	addresses = append(addresses, newAddress)
	toFile, _ := json.Marshal(addresses)
	ioutil.WriteFile(fileStore, toFile, 0600)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newAddress)
}

func getSavedAddresses() {
	f, _ := ioutil.ReadFile(fileStore)
	json.Unmarshal([]byte(f), &addresses)
}

func getOneAddress(w http.ResponseWriter, r *http.Request) {
	getSavedAddresses()
	addressID := mux.Vars(r)["id"]

	for _, singleAddress := range addresses {
		if singleAddress.ID == addressID {
			json.NewEncoder(w).Encode(singleAddress)
		}
	}
}

func getAllAddresses(w http.ResponseWriter, r *http.Request) {
	getSavedAddresses()
	json.NewEncoder(w).Encode(addresses)
}

func getPort() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		return "8080"
	}
	return port
}

func main() {
	fmt.Println("Server start...")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", getAllAddresses).Methods(http.MethodGet)
	router.HandleFunc("/", createAddress).Methods(http.MethodPost)
	router.HandleFunc("/address/{id}", getOneAddress).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":"+getPort(), router))
}
