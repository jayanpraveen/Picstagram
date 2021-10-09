package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func makeConn() *mongo.Client {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	collection := client.Database("picstagram").Collection("collectionName")

	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})

	id := res.InsertedID
	log.Println(id)

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	return client
}

func insertOneDocument(collName string, bsonValue bson.M) {
	val := makeConn()
	col := val.Database("picstagram").Collection(collName)
	res, err := col.InsertOne(context.TODO(), bsonValue)
	id := res.InsertedID
	log.Println(id)
	log.Panic(err)
}

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var user User

func HandleUserRequests() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", CreateUser)
	mux.HandleFunc("/users/", GetUserByID)
	mux.HandleFunc("/posts", CreatePost)
	mux.HandleFunc("/posts/", GetPostById)
	mux.HandleFunc("/posts/users/", GetAllPostsOfUser)
	log.Fatal(http.ListenAndServe(":80", mux))
}

// Create User: POST
func CreateUser(w http.ResponseWriter, req *http.Request) {

	httpStatus := req.Method

	if httpStatus == "POST" {
		jsonDecoder := json.NewDecoder(req.Body)
		err := jsonDecoder.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(user)
		fmt.Println(user.Id)
	}

}

// Get a User: GET
func GetUserByID(w http.ResponseWriter, req *http.Request) {
	id := strings.Split(req.URL.Path, "/")[2]

	httpStatus := req.Method

	if httpStatus == "GET" {
		// search
		fmt.Println(id)
	}

}

type Post struct {
	Id        string `json:"id"`
	Caption   string `json:"caption"`
	Image_URL string `json:"imageUrl"`
	Timestamp string `json:"timestamp"`
}

var posts Post

// Create a Post : POST
func CreatePost(w http.ResponseWriter, req *http.Request) {

	httpStatus := req.Method

	if httpStatus == "POST" {
		jsonDecoder := json.NewDecoder(req.Body)
		err := jsonDecoder.Decode(&posts)
		if err != nil {
			log.Fatal(err)
		}

		posts.Timestamp = time.Now().Format("01-02-2006 15:04:05")

		fmt.Println(posts)
		fmt.Println(posts.Timestamp)
		fmt.Println(posts.Caption)
	}
}

// Get a post using id : GET
func GetPostById(w http.ResponseWriter, req *http.Request) {

	postId := strings.Split(req.URL.Path, "/")[2]

	httpStatus := req.Method

	if httpStatus == "GET" {
		fmt.Println(postId)
	}

}

// List all posts of a user : GET
func GetAllPostsOfUser(w http.ResponseWriter, req *http.Request) {

	userId := strings.Split(req.URL.Path, "/")[3]

	httpStatus := req.Method

	if httpStatus == "GET" {
		fmt.Println(userId)
		fmt.Println("print all users")
	}
}

func main() {
	HandleUserRequests()
}
