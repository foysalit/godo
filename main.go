package main

import (
    "log"
    "net/http"
)

func main() {
	router := InitRouters()

    log.Fatal(http.ListenAndServe(":8080", router))
}