package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var MyID = os.Getenv("SCOOTER_ID")

func helloScooter(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello scooter, your id = %s\n", MyID)
}

func main() {
	if MyID != "" {
		log.Printf("Application B started for scooter id %s.", MyID)
		http.HandleFunc("/scooter", helloScooter)
		http.ListenAndServe(":8080", nil)
	}
}
