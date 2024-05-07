package handlers

import (
	"net/http"
	"project/web/utils"
)

func List(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		utils.SendData(w, users)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
