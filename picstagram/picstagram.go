package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

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

func main() {
	HandleUserRequests()
}
