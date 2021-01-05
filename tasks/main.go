/*Frankenstein de :
1 - https://tutorialedge.net/golang/creating-restful-api-with-golang/
2 - https://levelup.gitconnected.com/working-with-mongodb-using-golang-754ead0c10c
*/
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Date Layout
const layoutISO = "2006-01-02"

/* Used to create a singleton object of MongoDB client.
Initialized and exposed through  GetMongoClient().*/
var clientInstance *mongo.Client

//Used during creation of singleton client object in GetMongoClient().
var clientInstanceError error

//Used to execute client creation procedure only once.
var mongoOnce sync.Once

//I have used below constants just to hold required database config's.
const (
	CONNECTIONSTRING = "mongodb://localhost:27017"
	DB               = "db_tasks_manager"
	TASKS            = "col_tasks"
)

//GetMongoClient - Return mongodb connection to work with
func getMongoClient() (*mongo.Client, error) {
	//Perform connection creation operation only once.
	mongoOnce.Do(func() {
		// Set client options
		clientOptions := options.Client().ApplyURI(CONNECTIONSTRING)
		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
		}
		// Check the connection
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}

// Task Structure
type Task struct {
	ID      int    `json:"ID"  bson:"_id,omitempty"`
	Title   string `json:"Title" bson:"Title"`
	Content string `json:"Content" bson:"Content"`
	DueDate string `json:"DueDate" bson:"DueDate"`
}

func dbReturnAllTasks() ([]Task, error) {
	//Define filter query for fetching specific document from collection
	filter := bson.D{{}} //bson.D{{}} specifies 'all documents'
	tasks := []Task{}
	//Get MongoDB connection using connectionhelper.
	client, err := getMongoClient()
	if err != nil {
		return tasks, err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(DB).Collection(TASKS)
	//Perform Find operation & validate against the error.
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return tasks, findError
	}
	//Map result to slice
	for cur.Next(context.TODO()) {
		t := Task{}
		err := cur.Decode(&t)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, t)
	}
	// once exhausted, close the cursor
	cur.Close(context.TODO())
	if len(tasks) == 0 {
		return tasks, mongo.ErrNoDocuments
	}
	return tasks, nil
}

//createTask - Insert a new document in the collection.
func dbCreateNewTask(task Task) error {
	//Get MongoDB connection using connectionhelper.
	client, err := getMongoClient()
	if err != nil {
		return err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(DB).Collection(TASKS)
	//Perform InsertOne operation & validate against the error.
	_, err = collection.InsertOne(context.TODO(), task)
	if err != nil {
		return err
	}
	//Return success without any error.
	return nil
}

func dbDeleteTask(id int) error {
	//Define filter query for fetching specific document from collection
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	//Get MongoDB connection using connectionhelper.
	client, err := getMongoClient()
	if err != nil {
		return err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(DB).Collection(TASKS)
	//Perform DeleteOne operation & validate against the error.
	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	//Return success without any error.
	return nil
}

// MarkCompleted - MarkCompleted
func dbUpdateTask(t Task) error {
	//Define filter query for fetching specific document from collection
	filter := bson.D{primitive.E{Key: "_id", Value: t.ID}}

	//Define updater for to specifiy change to be updated.
	updater := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "Title", Value: t.Title},
		primitive.E{Key: "Content", Value: t.Content},
		primitive.E{Key: "DueDate", Value: t.DueDate},
	}}}

	//Get MongoDB connection using connectionhelper.
	client, err := getMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database(DB).Collection(TASKS)

	//Perform UpdateOne operation & validate against the error.
	_, err = collection.UpdateOne(context.TODO(), filter, updater)
	if err != nil {
		return err
	}
	//Return success without any error.
	return nil
}

func dbReturnTask(id int) (Task, error) {
	result := Task{}
	//Define filter query for fetching specific document from collection
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	//Get MongoDB connection using connectionhelper.
	client, err := getMongoClient()
	if err != nil {
		return result, err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(DB).Collection(TASKS)
	//Perform FindOne operation & validate against the error.
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	//Return result without any error.
	return result, nil
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/tasks", apiReturnAllTasks)
	myRouter.HandleFunc("/task", apiCreateNewTask).Methods("POST")
	myRouter.HandleFunc("/task/{id}", apiDeleteTask).Methods("DELETE")
	myRouter.HandleFunc("/task/{id}", apiUpdateTask).Methods("PUT")
	myRouter.HandleFunc("/task/{id}", apiReturnTask)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func apiReturnAllTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: apiReturnAllTasks")
	tasks, err := dbReturnAllTasks()
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(tasks)

}

func apiCreateNewTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: apiCreateNewTask")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var task Task
	json.Unmarshal(reqBody, &task)

	err := dbCreateNewTask(task)
	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(task)
}

func apiDeleteTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: apiDeleteTask")
	vars := mux.Vars(r)
	id := vars["id"]
	_id, _ := strconv.Atoi(id)
	err := dbDeleteTask(_id)
	if err != nil {
		fmt.Println(err)
	}
}

func apiUpdateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: apiUpdateTask")
	vars := mux.Vars(r)
	id := vars["id"]
	_id, _ := strconv.Atoi(id)

	reqBody, _ := ioutil.ReadAll(r.Body)
	var task Task
	json.Unmarshal(reqBody, &task)

	if _id == task.ID {
		dbUpdateTask(task)
	}
}

func apiReturnTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: apiReturnTask")
	vars := mux.Vars(r)
	id := vars["id"]
	_id, _ := strconv.Atoi(id)

	task, err := dbReturnTask(_id)
	if err == nil {
		json.NewEncoder(w).Encode(task)
	} else {
		fmt.Println(err)
	}
}

func main() {
	fmt.Println("Task API.")

	dbCreateNewTask(Task{ID: 1,
		Title:   "Primeira Tarefa",
		Content: "Conteúdo da Primeira Tarefa",
		DueDate: "2021-01-06"})

	dbCreateNewTask(Task{ID: 2,
		Title:   "Segunda Tarefa",
		Content: "Conteúdo da Segunda Tarefa",
		DueDate: "2021-01-07"})

	handleRequests()
}
