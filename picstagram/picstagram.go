package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	mongoOps "example.com/mongoDB"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	HandleUserRequests()
}

func notFoundHandler(w http.ResponseWriter) {
	w.Write([]byte("404 - The page does not exist or has been moved."))
}

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var user User

type Post struct {
	Id        string `json:"id"`
	Caption   string `json:"caption"`
	Image_URL string `json:"imageUrl"`
	Timestamp string `json:"timestamp"`
}

var posts Post

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

		mongoOps.InsertDocument("users", userData)
	}

}

// Get a User: GET
func GetUserByID(w http.ResponseWriter, req *http.Request) {
	id := strings.Split(req.URL.Path, "/")[2]

	httpStatus := req.Method

	if httpStatus == "GET" {
		w.Header().Set("Content-Type", "application/json")
		docId, _ := primitive.ObjectIDFromHex(id)
		user := mongoOps.FindDocument("users", docId)
		if user == nil {
			notFoundHandler(w)
		}
		json.NewEncoder(w).Encode(user)

	}
}

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

		mongoOps.InsertDocument("posts", postData)

	}
}

// Get a post using id : GET
func GetPostById(w http.ResponseWriter, req *http.Request) {

	postId := strings.Split(req.URL.Path, "/")[2]

	httpStatus := req.Method

	if httpStatus == "GET" {
		w.Header().Set("Content-Type", "application/json")
		docId, _ := primitive.ObjectIDFromHex(postId)
		userPost := mongoOps.FindDocument("posts", docId)
		if userPost == nil {
			notFoundHandler(w)
		}
		json.NewEncoder(w).Encode(userPost)

	}
}

// List all posts of a user : GET
func GetAllPostsOfUser(w http.ResponseWriter, req *http.Request) {

	userId := strings.Split(req.URL.Path, "/")[3]

	httpStatus := req.Method

	if httpStatus == "GET" {
		w.Header().Set("Content-Type", "application/json")
		posts := mongoOps.GetAllUserPosts(userId)
		if posts == nil {
			notFoundHandler(w)
		}
		json.NewEncoder(w).Encode(posts)

	}
}
