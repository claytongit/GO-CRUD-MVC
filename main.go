package main

import(
	"fmt"
	"net/http"
	"connection"
	"routes"

	"github.com/gorilla/mux"
)

func main()  {

	connection.Db()

	route := mux.NewRouter().StrictSlash(true)

	routes.Route(route)

	fmt.Println("Server run port 8080")

	http.ListenAndServe(":8080", route)
	
}