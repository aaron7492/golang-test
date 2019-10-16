package main

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandleFunc)
	}
	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Listas de tareas",
		"GET",
		"/tareas",
		listartareas,
	},
	Route{
		"Detalle Tarea",
		"GET",
		"/tareas/{id}",
		showtask,
	},
	Route{
		"taskadd",
		"POST",
		"/tareas",
		taskadd,
	},
	Route{
		"taskupdate",
		"PUT",
		"/tareas/{id}",
		taskupdate,
	},
	Route{
		"taskremove",
		"DELETE",
		"/tareas/{id}",
		taskremove,
	},
}
