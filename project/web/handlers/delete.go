package handlers

import (
	"fmt"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {

	// nm=url name
	nm := r.URL.Path[len("/users/"):]
	var ck int
	ck = 1
	for i, u := range users {

		if u.Name == nm {
			ck = 0
		}
		if ck == 0 && i+1 != len(users) {
			users[i].Name = users[i+1].Name
			users[i].Password = users[i+1].Password
		}
	}
	if ck == 0 {
		users = users[:len(users)-1]
		fmt.Fprintf(w, "Deleted Successfully: %s", nm)
		return
	}
	fmt.Fprintf(w, "No user found named : %s", nm)

}
