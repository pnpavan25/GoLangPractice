package main

import (
	"log"
	"net/http"
	"pavan/a3http/controller"
	"pavan/a3http/mapstore"
	crouter "pavan/a3http/router"
)

//Entry point of the program
func main() {
	ctl := &controller.CustomerController{
		Store : mapstore.NewMapStore(), // Injecting dependency
	}
	ctl.InitializeCustomerController()

	r := crouter.InitializeRoutes(ctl) // configure routes

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe() // Run the http server
}
