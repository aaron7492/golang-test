package main

import (
	"net/http"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

type Tarea struct {
	Autor  string `json:"autor"`
	Titulo string `json:"titulo"`
	Tarea  string `json:"tarea"`
	Fecha  string `json:"fecha"`
}

type Tareas []Tarea

type Message struct {
	Status  string `json:"status:"`
	Mensaje string `json:"mensaje:"`
}
