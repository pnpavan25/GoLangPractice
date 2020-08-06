package router

import (
	"github.com/gorilla/mux"
	"pavan/a3http/controller"
)

func InitializeRoutes(ctl *controller.CustomerController) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/customers", ctl.GetAll).Methods("GET")
	r.HandleFunc("/api/customers/{id}", ctl.Get).Methods("GET")
	r.HandleFunc("/api/customers", ctl.Post).Methods("POST")
	r.HandleFunc("/api/customers/{id}", ctl.Put).Methods("PUT")
	r.HandleFunc("/api/customers/{id}", ctl.Delete).Methods("DELETE")

	return r
}
