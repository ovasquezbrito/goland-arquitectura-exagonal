package rest

import (
	"github.com/ovasquezbrito/candyshop/pkg/reading"
	"encoding/json"
	"net/http"
)

func welcomeHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Bienbenido a nuestra tienda de dulces")
	}
}

func getAllCandiesHandler(rs reading.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cs, err := rs.GetAllCandyNames()
		if err != nil {
			http.Error(w, "Cannot Process your", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(cs)
	}
}
