package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome")
}

func TodoList(w http.ResponseWriter, r *http.Request) {
	todos := Todos {
		Todo {Name: "Eat"},
		Todo {Name: "Work"},
		Todo {Name: "Sleep"},
	}

	json.NewEncoder(w).Encode(todos)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todoId := params["todoId"]

    fmt.Fprintf(w, "Show Todo Item: %q", todoId)
}