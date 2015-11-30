package main

import (
    "fmt"
	"io"
	"io/ioutil"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/jinzhu/gorm"
    "github.com/davecgh/go-spew/spew"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome")
}

func TodoList(w http.ResponseWriter, r *http.Request) {
	db := InitDB()
	var todos Todos
	response := TodosResponse{Error: true}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	db.Limit(10).Find(&todos)

	response.Error = false
	response.Data = &todos

	json.NewEncoder(w).Encode(response)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db := InitDB()
	response := TodoResponse{Error: true}
	var todo Todo

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	query := db.First(&todo, params["todoId"])

	if query.Error != nil && query.Error == gorm.RecordNotFound {
		w.WriteHeader(404)
		response.Data = nil
	} else {
		response.Error = false
		response.Data = &todo
	}

    json.NewEncoder(w).Encode(response)
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	db := InitDB()
	var todo Todo
	response := TodoResponse{Error: true}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	spew.Dump(body)

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err := json.Unmarshal(body, &todo); err != nil {
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
	} else {
        db.Create(&todo)
        response.Error = false;
        response.Data = &todo
	}

    json.NewEncoder(w).Encode(response)
}

func TodoRemove(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db := InitDB()
	response := TodoResponse{Error: true}
	var todo Todo

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	query := db.First(&todo, params["todoId"])

	if query.Error != nil && query.Error == gorm.RecordNotFound {
		w.WriteHeader(404)
		response.Data = nil
	} else {
		db.Delete(&todo)
		response.Error = false
	}

    json.NewEncoder(w).Encode(response)
}