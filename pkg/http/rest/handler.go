package rest

import "github.com/gorilla/mux"

// InitHandlers función para inicializar el enrutador
func InitHandlers() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/", welcomeHandler()).Methods("GET")

	return router
}
