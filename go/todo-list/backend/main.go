package main

import (
	"io"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
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

func GetIncompleteTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Incomplete Todo")
	var itemsIncomplete []TodoItem
	_ = connection.Find(bson.M{"done": false}).All(&itemsIncomplete)

	fmt.Println(itemsIncomplete[0].Description)
	resIncomplete := []string{}
	for _, todo := range itemsIncomplete{
		resIncomplete = append(resIncomplete, todo.Description)
	}

	// w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(resIncomplete)
}


func GetCompletedTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Todo")
	var itemsCompleted []TodoItem
	_ = connection.Find(bson.M{"done": true}).All(&itemsCompleted)

	fmt.Println(itemsCompleted[0].Description)
	resCompleted := []string{}
	resCompleted = append(resCompleted, itemsCompleted[0].Description)

	// w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(resCompleted)
}

func GetTodoByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get todo by ID")
	var result TodoItem

	vars := mux.Vars(r)
	id := vars["id"]
	_ = connection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(result)
}

func MarkDone(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Mark as done")
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
	fmt.Println(id)
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
	fmt.Println("Starting todolist backend")
	fmt.Println("Creating mongo session")
	session.SetMode(mgo.Monotonic, true)
	defer session.Close()
	router := mux.NewRouter()
	origins := handlers.AllowedOrigins([]string{"*"})

	router.HandleFunc("/todo", AddTodo).Methods("POST")
	router.HandleFunc("/todo-completed", GetCompletedTodo).Methods("GET")
	router.HandleFunc("/todo-incomplete", GetIncompleteTodo).Methods("GET")
	router.HandleFunc("/todo/{id}", GetTodoByID).Methods("GET")
	router.HandleFunc("/todo/{id}", MarkDone).Methods("PATCH")
	router.HandleFunc("/todo/{id}", DeleteItem).Methods("DELETE")
	router.HandleFunc("/health", Health).Methods("GET")

	http.ListenAndServe(":8000", handlers.CORS(origins)(router))
}
