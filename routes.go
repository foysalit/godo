package main

import (
    "net/http"
    "github.com/gorilla/mux"
)

type Route struct {
    Name string
    Method string
    Pattern string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes {
    Route {
        Name: "home.index",
        Method: "GET",
        Pattern: "/", 
        HandlerFunc: Index,
    },
    Route {
        Name: "todos.list",
        Method: "GET",
        Pattern: "/todos", 
        HandlerFunc: TodoList,
    },
    Route {
        Name: "todos.create",
        Method: "POST",
        Pattern: "/todos", 
        HandlerFunc: TodoCreate,
    },
    Route {
        Name: "todos.single",
        Method: "GET",
        Pattern: "/todos/{todoId}", 
        HandlerFunc: TodoShow,
    },
    Route {
        Name: "todos.remove",
        Method: "DELETE",
        Pattern: "/todos/{todoId}", 
        HandlerFunc: TodoRemove,
    },
}

func InitRouters() *mux.Router{
	router := mux.NewRouter().StrictSlash(true)

    for _, route := range routes {
        router.
            // Name(route.Name).
            Methods(route.Method).
            Path(route.Pattern).
            HandlerFunc(route.HandlerFunc)
    }

    return router
}