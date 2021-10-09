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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func mongoConnection() *mongo.Client {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}

func insertDocument(collName string, bsonValue bson.D) {
	client := mongoConnection()
	collection := client.Database("picstagram").Collection(collName)
	res, err := collection.InsertOne(context.TODO(), bsonValue)

	log.SetPrefix("response id: ")
	log.Println(res.InsertedID)

	if err != nil {
		log.Panic(err)
	}
}

func findDocument(collName string, docId primitive.ObjectID) primitive.M {
	client := mongoConnection()
	collection := client.Database("picstagram").Collection(collName)

	findone_result := collection.FindOne(context.TODO(), bson.M{"_id": docId})

	var bson_obj bson.M
	if err2 := findone_result.Decode(&bson_obj); err2 != nil {
		fmt.Println(err2)
	}
	return bson_obj

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

		password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
		// race contd
		if err != nil {
			log.Fatal("internal error!")
		}
		hashedPassword := string(password)

		userData := bson.D{
			{Key: "id", Value: user.Id},
			{Key: "name", Value: user.Name},
			{Key: "email", Value: user.Email},
			{Key: "password", Value: hashedPassword},
		}

		insertDocument("users", userData)
	}

}

// Get a User: GET
func GetUserByID(w http.ResponseWriter, req *http.Request) {
	id := strings.Split(req.URL.Path, "/")[2]

	httpStatus := req.Method

	if httpStatus == "GET" {
		w.Header().Set("Content-Type", "application/json")
		docId, _ := primitive.ObjectIDFromHex(id)
		val := findDocument("users", docId)
		json.NewEncoder(w).Encode(val)
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

		postData := bson.D{
			{Key: "id", Value: posts.Id},
			{Key: "caption", Value: posts.Caption},
			{Key: "imageUrl", Value: posts.Image_URL},
			{Key: "timestamp", Value: posts.Timestamp},
		}

		insertDocument("posts", postData)

	}
}

// Get a post using id : GET
func GetPostById(w http.ResponseWriter, req *http.Request) {

	postId := strings.Split(req.URL.Path, "/")[2]

	httpStatus := req.Method

	if httpStatus == "GET" {
		w.Header().Set("Content-Type", "application/json")
		docId, _ := primitive.ObjectIDFromHex(postId)
		val := findDocument("posts", docId)
		json.NewEncoder(w).Encode(val)
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
