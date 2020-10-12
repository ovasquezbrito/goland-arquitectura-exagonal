package rest

import (
	"encoding/json"
	"net/http"
)

func welcomeHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Bienbenido a nuestra tienda de dulces")
	}
}
