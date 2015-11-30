package main

import (
    "log"
    "os"
    "net/http"
)

// http://thenewstack.io/make-a-restful-json-api-go/

func main() {
	router := InitRouters()

    log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}