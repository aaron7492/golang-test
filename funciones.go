package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

//conexion a la bd
func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return session
}

var collection = getSession().DB("udemygo").C("tareas")

//como responder peticiones
func responseTarea(w http.ResponseWriter, status int, results Tarea) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}

func responseTareas(w http.ResponseWriter, status int, results Tareas) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}

//Funciones de la app
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hola")
}

func listartareas(w http.ResponseWriter, r *http.Request) {
	var results []Tarea
	err := collection.Find(nil).All(&results)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("resultados: ", results)
	}
	responseTareas(w, 200, results)
}

func showtask(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	tarea_id := parametros["id"]
	if !bson.IsObjectIdHex(tarea_id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(tarea_id)
	fmt.Println(tarea_id)
	fmt.Println(oid)

	results := Tarea{}
	err := collection.FindId(oid).One(&results)

	fmt.Println(results)

	if err != nil {
		w.WriteHeader(404)
		return
	}
	responseTarea(w, 200, results)

}

func taskadd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var data Tarea
	err := decoder.Decode(&data)

	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	log.Println(data)
	err = collection.Insert(data)

	if err != nil {
		w.WriteHeader(500)
		return
	}
	responseTarea(w, 200, data)
}

func taskupdate(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	tarea_id := parametros["id"]

	if !bson.IsObjectIdHex(tarea_id) {
		w.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(tarea_id)
	decoder := json.NewDecoder(r.Body)
	var data Tarea
	err := decoder.Decode(&data)

	if err != nil {
		panic(err)
		w.WriteHeader(500)
		return
	}
	defer r.Body.Close()

	document := bson.M{"_id": oid}
	change := bson.M{"$set": data}
	err = collection.Update(document, change)
	if err != nil {
		w.WriteHeader(404)
		return
	}
}

func (this *Message) setStatus(data string) {
	this.Status = data
}

func (this *Message) setMessage(data string) {
	this.Mensaje = data
}

func taskremove(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	tarea_id := parametros["id"]
	if !bson.IsObjectIdHex(tarea_id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(tarea_id)
	err := collection.RemoveId(oid)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	//results := Message{"Success", "se elimino la tarea id" + tarea_id}
	carma := new(Message)
	carma.setStatus("Success")
	carma.setMessage("se elimino la tarea id" + tarea_id)
	results := carma
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)

}
