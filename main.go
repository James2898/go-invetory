package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Item struct {
	UID   string  `json:"UID"`
	Name  string  `json:"Name"`
	Desc  string  `json:"Desc"`
	Price float64 `json:"Price"`
}

var inventory []Item

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint called: homePage")
}

func getInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(inventory)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)

	inventory = append(inventory, item)

	json.NewEncoder(w).Encode(inventory)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)

	params := mux.Vars(r)

	// Delete the item at UID
	_deleteIteamAtUid(params["uid"])

	// Create it with new data
	inventory = append(inventory, item)

	json.NewEncoder(w).Encode(inventory)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	_deleteIteamAtUid(params["uid"])

	json.NewEncoder(w).Encode(inventory)
}

func _deleteIteamAtUid(uid string) {
	for index, item := range inventory {
		if item.UID == uid {
			// Dleete item from slice
			inventory = append(inventory[:index], inventory[index+1:]...)
		}
	}
}

func handlerRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/inventory", getInventory).Methods("GET")
	router.HandleFunc("/inventory", createItem).Methods("POST")
	router.HandleFunc("/inventory/{uid}", deleteItem).Methods("DELETE")
	router.HandleFunc("/inventory/{uid}", updateItem).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", router))

}

func main() {
	inventory = append(inventory, Item{
		UID:   "0",
		Name:  "Cheese",
		Desc:  "cheeze wiz",
		Price: 3.99,
	})

	inventory = append(inventory, Item{
		UID:   "1",
		Name:  "Bacon",
		Desc:  "Sweet Bacon",
		Price: 4.99,
	})

	handlerRequests()
}
