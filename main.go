package main

import (
    "log"
    "net/http"
)

// http://thenewstack.io/make-a-restful-json-api-go/

func main() {
	router := InitRouters()

    log.Fatal(http.ListenAndServe(":8080", router))
}