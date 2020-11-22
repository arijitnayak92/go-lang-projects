package app

import (
	"fmt"
	"net/http"

	"github.com/arijitnayak92/taskAfford/RESTTODO/controllers"
)

func response(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Backend Responding !")
}

func Routes() {
	router.HandleFunc("/", response)

	api := router.PathPrefix("/item").Subrouter()
	api.HandleFunc("/getOneItem/{item_id}", controllers.GetOneProduct).Methods("GET")
	api.HandleFunc("/logout", controllers.Logout).Methods("POST")
	api.HandleFunc("/getAllItem", controllers.GetAllProduct).Methods("GET")
	api.HandleFunc("/addItem", controllers.AddProduct).Methods("POST")
	api.HandleFunc("/updateItem/{item_id}", controllers.UpdateOneProduct).Methods("PUT")
	api.HandleFunc("/deleteItem/{item_id}", controllers.DeleteOneProduct).Methods("DELETE")

}
