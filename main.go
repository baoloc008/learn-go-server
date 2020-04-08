package main

import (
	"fmt"
	"io"
	"net/http"
)

// create a handler struct
type HttpHandler struct{}

// implement `ServeHTTP` method on `HttpHandler` struct
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	//// create response binary data
	//data := []byte("Hello World!") // slice of bytes
	//// write `data` to response
	//res.Write(data)


	// write `Hello` using `io.WriteString` function
	io.WriteString(res, "Hello")
	// write `World` using `fmt.Fprint` function
	fmt.Fprint(res, " World! ")
	// write `❤️` using simple `Write` call
	res.Write([]byte("❤️"))
}

func main() {
	// create a new handler
	handler := HttpHandler{}
	// listen and serve
	http.ListenAndServe(":9620", handler)
}
