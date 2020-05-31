package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	handleRequest()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the HomePage!")
}
func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/getAddress", getAddressBookAll)
	http.ListenAndServe(getPort(), nil)
}

type addressBook struct {
	Firstname string
	Lastname  string
	Code      int
	Phone     string
}

func getAddressBookAll(w http.ResponseWriter, r *http.Request) {
	addBook := addressBook{
		Firstname: "Chaiyarin",
		Lastname:  "Niamsuwan",
		Code:      1993,
		Phone:     "0870940955",
	}
	json.NewEncoder(w).Encode(addBook) // (1)
}

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Println("No Port In Heroku " + port)
	}
	return ":" + port
}
