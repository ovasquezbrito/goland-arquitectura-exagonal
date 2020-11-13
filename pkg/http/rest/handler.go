package rest

import (
	"github.com/ovasquezbrito/candyshop/pkg/reading"
	"github.com/gorilla/mux"
)

// InitHandlers función para inicializar el enrutador
func InitHandlers(rs reading.Service) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/", welcomeHandler()).Methods("GET")
	router.HandleFunc("/api/candies", getAllCandiesHandler(rs)).Methods("GET")

	return router
}
