package main

import (
	"fmt"
	"log"
	"net/http"
)

func test() {
	/// create a new `ServeMux`
	mux := http.NewServeMux()

	// handle `/` route
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello World!")
	})

	// handle `/hello/golang` route
	mux.HandleFunc("/hello/golang", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello Golang!")
	})

	// listen and serve
	http.ListenAndServe(":9620", mux)
}

func main() {
	// handle `/` route to `http.DefaultServeMux`
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		// get response headers
		header := res.Header()

		// set content type header
		header.Set("Content-Type", "application/json")

		// reset date header (inline call)
		res.Header().Set("Date", "01/01/2020")

		// the WriteHeader call should be followed by the Write call but after all the response headers are set.
		// set status header
		res.WriteHeader(http.StatusBadRequest) // http.StatusBadRequest == 400

		// respond with a JSON string
		fmt.Fprint(res, `{"status":"FAILURE"}`)
	})

	// listen and serve using `http.DefaultServeMux`
	log.Fatal(http.ListenAndServe(":9620", nil))
}
