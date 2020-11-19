package app

import (
	"fmt"
	"net/http"

	"github.com/arijitnayak92/taskAfford/RESTMUXJWT/controllers"
)

func response(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Backend Responding !")
}

func Routes() {
	router.HandleFunc("/", response)
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/logout", controllers.Logout).Methods("POST")
	router.HandleFunc("/create", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/refreshToken", controllers.RefreshToken).Methods("POST")
	router.HandleFunc("/getOneItem/{item_id}", controllers.GetOneProduct).Methods("GET")
	router.HandleFunc("/getAllItem", controllers.GetAllProduct).Methods("GET")
	router.HandleFunc("/addItem", controllers.AddProduct).Methods("POST")
	router.HandleFunc("/updateItem/{item_id}", controllers.UpdateOneProduct).Methods("PUT")
	router.HandleFunc("/deleteItem/{item_id}", controllers.DeleteOneProduct).Methods("DELETE")

}
