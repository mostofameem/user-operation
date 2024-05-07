package handlers

import (
	"encoding/json"
	"net/http"
)

func GetUserByName(w http.ResponseWriter, r *http.Request) {

	// nm=url name
	nm := r.URL.Path[len("/users/"):]
	for _, u := range users {

		if u.Name == nm {
			w.Header().Set("content-Type", "application/json")
			json.NewEncoder(w).Encode(u)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("User not found"))
}
