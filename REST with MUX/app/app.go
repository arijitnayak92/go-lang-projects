package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter().StrictSlash(true)

func StartApp() {
	Routes()
	http.ListenAndServe(":8080", router)
}
