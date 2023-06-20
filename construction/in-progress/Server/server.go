package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/users", handleUsers)

	fmt.Println("Listening on port 8080..")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func handleUsers(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		fmt.Println("Not implemented")
	case "POST":
		fmt.Println("Not implemented")
	case "DELETE":
		fmt.Println("Not implemented")

	default:
		w.Header().Set("Allow", "GET, POST, DELETE")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}
