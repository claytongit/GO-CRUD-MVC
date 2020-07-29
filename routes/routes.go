package routes

import(
	"controller"
	"github.com/gorilla/mux"
)

func Route(route *mux.Router)  {

	route.HandleFunc("/clients", controller.UserGet).Methods("GET")

	route.HandleFunc("/clients/{userId}", controller.UserGetId).Methods("GET")

	route.HandleFunc("/clients", controller.UserPost).Methods("POST")

	route.HandleFunc("/clients/{userId}", controller.UserUpdata).Methods("PUT")

	route.HandleFunc("/clients/{userId}", controller.UserDelete).Methods("DELETE")
	
}