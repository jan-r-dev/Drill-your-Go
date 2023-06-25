package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Identifier struct {
	ID string `json:"id"`
}

type User struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	FullName string `json:"fullName,omitempty"`
}

type UserHandler struct {
	Users map[string]*User
}

func (h UserHandler) respContentTypeIncorrect(w http.ResponseWriter, req *http.Request) {
	http.Error(w,
		fmt.Sprintf("Incorrect content type. Required: \"application/json\" Received: %q", req.Header.Get("Content-Type")),
		http.StatusBadRequest,
	)
}

func (h UserHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(h.Users)

	case "POST":
		if req.Header.Get("Content-Type") != "application/json" {
			h.respContentTypeIncorrect(w, req)
			return
		}

		bs, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "Bad request body", http.StatusBadRequest)
			return
		}

		user := User{}
		json.Unmarshal(bs, &user)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		if user.ID == "" || user.Email == "" {
			http.Error(w, "User's ID and email cannot be empty", http.StatusBadRequest)
			return
		}

		h.Users[user.ID] = &User{
			ID:       user.ID,
			Email:    user.Email,
			FullName: user.FullName,
		}

		if val, ok := h.Users[user.ID]; ok {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf("User with ID %q created/overwritten successfully.\n", val.ID)))
			return
		} else {
			http.Error(w, "Failed to create user.", http.StatusInternalServerError)
		}

	case "DELETE":
		if req.Header.Get("Content-Type") != "application/json" {
			h.respContentTypeIncorrect(w, req)
			return
		}

		bs, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "Bad request body", http.StatusBadRequest)
			return
		}

		id := &Identifier{}
		err = json.Unmarshal(bs, id)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		delete(h.Users, id.ID)

		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte(fmt.Sprintf("User with ID %q deleted successfully.\n", id.ID)))

	default:
		w.Header().Set("Allow", "GET, POST, DELETE")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	handler := UserHandler{
		Users: make(map[string]*User),
	}

	mux := http.NewServeMux()
	mux.Handle("/api/v1/users", handler)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Listening on port 8080..")
	log.Fatal(s.ListenAndServe())
}
