package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Task struct{
	ID string `json:"id"`
	Item string `json:"item"`
}

var tasks []Task

func getTasks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// Create
func createTask(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var task Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = strconv.Itoa(rand.Intn(100))
	tasks = append(tasks, task)
	json.NewEncoder(w).Encode(task)
}

// Read
func getTask(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range tasks{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

//update
func updateTask(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range tasks{
		if item.ID == params["id"]{
			tasks = append(tasks[:index], tasks[index+1:]...)
			var task Task
			_ = json.NewDecoder(r.Body).Decode(&task)
			task.ID = params["id"]
			tasks = append(tasks, task)
			json.NewEncoder(w).Encode(task)
			return
		}
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range tasks{
		if item.ID == params["id"]{
			tasks = append(tasks[:index], tasks[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode(tasks)
}

func main(){
	r := mux.NewRouter()

	tasks = append(tasks, Task{ID: "1", Item: "Complete Golang Tutorial"})
	tasks = append(tasks, Task{ID: "2", Item: "Fill the bottles"})
	tasks = append(tasks, Task{ID: "3", Item: "Take the dog for a walk"})
	tasks = append(tasks, Task{ID: "4", Item: "Browse Pinterest"})
	tasks = append(tasks, Task{ID: "5", Item: "House Chores"})
	tasks = append(tasks, Task{ID: "6", Item: "Finish reading novels"})


	r.HandleFunc("/tasks/", getTasks).Methods("GET")
	r.HandleFunc("/tasks", createTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	r.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")


	fmt.Println("This Server is Running on Port 8081\n")
	log.Fatal(http.ListenAndServe(":8081", r))
}