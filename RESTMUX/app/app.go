package app

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter().StrictSlash(true)

func MakeRequest(url string, ch chan<- string) {
	start := time.Now()
	resp, _ := http.Get(url)
	secs := time.Since(start).Seconds()
	body, _ := ioutil.ReadAll(resp.Body)
	ch <- fmt.Sprintf("%.2f elapsed with response length: %d %s", secs, len(body), url)
}

func StartApp() {
	Routes()
	http.ListenAndServe(":8080", router)
}
