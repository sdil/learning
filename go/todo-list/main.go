package main

import (
	"io"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"gopkg.in/mgo.v2"
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

var session, _ = mgo.Dial("127.0.0.1")
var connection = session.DB("TutDb").C("ToDo")

type TodoItem struct {
	ID	bson.ObjectId `bson:"_id,omitempty"`
	Date time.Time
	Description string
	Done bool
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("adding a todo")
	_ = connection.Insert(TodoItem{
		bson.NewObjectId(),
		time.Now(),
		r.FormValue("description"),
		false,
	})

	result := TodoItem{}
	_ = connection.Find(bson.M{"description": r.FormValue("description")}).One(&result)
	json.NewEncoder(w).Encode(result)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Todo")
	var res []TodoItem

	vars := mux.Vars(r)
	id := vars["id"]
	if id != "" {
		res = GetByID(id)
	} else {
		fmt.Println("Get All")
		_ = connection.Find(nil).All(&res)
	}

	json.NewEncoder(w).Encode(res)
}

func GetByID(id string) []TodoItem {
	fmt.Println("Get todo by ID")
	var result TodoItem
	var res []TodoItem
	_ = connection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)
	res = append(res, result)
	return res
}

func MarkDone(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Mark as done")
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
	err := connection.Update(bson.M{"_id": id}, bson.M{"$set": bson.M{"done": true}})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `"updated": false, "error": ` + err.Error() + `}` )
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `"updated": true`)
	}
}

func DeleteItem (w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Item")

	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
	err := connection.RemoveId(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `"removed": false, "error": ` + err.Error() + `}` )
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `"removed": true`)
	}
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func main() {
	session.SetMode(mgo.Monotonic, true)
	defer session.Close()
	router := mux.NewRouter()
	router.HandleFunc("/todo", AddTodo).Methods("POST")
	router.HandleFunc("/todo", GetTodo).Methods("GET")
	router.HandleFunc("/todo/{id}", GetTodo).Methods("GET")
	router.HandleFunc("/todo/{id}", MarkDone).Methods("PATCH")
	router.HandleFunc("/todo/{id}", DeleteItem).Methods("DELETE")
	router.HandleFunc("/health", Health).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
