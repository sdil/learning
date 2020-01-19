package main

import (
	"io"
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"gopkg.in/mgo.v2"
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"github.com/rs/cors"
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

        w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func GetIncompleteTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Incomplete Todo")
	var ItemsIncomplete []TodoItem
	_ = connection.Find(bson.M{"done": false}).All(&ItemsIncomplete)
        w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(ItemsIncomplete)
}


func GetCompletedTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Completed Todo")
	var ItemsCompleted []TodoItem
	_ = connection.Find(bson.M{"done": true}).All(&ItemsCompleted)
        w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(ItemsCompleted)
}

func GetTodoByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get todo by ID")
	var result TodoItem

	vars := mux.Vars(r)
	id := vars["id"]
	_ = connection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)
	json.NewEncoder(w).Encode(result)
}

func MarkDone(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Mark as done")
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
	fmt.Println(id)
	err := connection.Update(bson.M{"_id": id}, bson.M{"$set": bson.M{"done": true}})
	if err != nil {
		io.WriteString(w, `"updated": false, "error": ` + err.Error() + `}` )
	} else {
		io.WriteString(w, `"updated": true`)
	}
}

func DeleteItem (w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Item")

	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
	err := connection.RemoveId(id)
	if err != nil {
		io.WriteString(w, `"removed": false, "error": ` + err.Error() + `}` )
	} else {
		io.WriteString(w, `"removed": true`)
	}
}

func Health(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `{"alive": true}`)
}

func main() {
	fmt.Println("Starting todolist backend")
	fmt.Println("Creating mongo session")
	session.SetMode(mgo.Monotonic, true)
	defer session.Close()
	router := mux.NewRouter()

	router.HandleFunc("/todo", AddTodo).Methods("POST")
	router.HandleFunc("/todo-completed", GetCompletedTodo).Methods("GET")
	router.HandleFunc("/todo-incomplete", GetIncompleteTodo).Methods("GET")
	router.HandleFunc("/todo/{id}", GetTodoByID).Methods("GET")
	router.HandleFunc("/todo/{id}", MarkDone).Methods("PATCH")
	router.HandleFunc("/todo/{id}", DeleteItem).Methods("DELETE")
	router.HandleFunc("/health", Health).Methods("GET")

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS"},
	}).Handler(router)

	http.ListenAndServe(":8000", handler)
}
