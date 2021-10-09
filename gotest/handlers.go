// handlers.go
package gotest

import (
	"encoding/json"
	"io"
	"net/http"

	mongoDB "example.com/mongoDB"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"alive": true}`)
}

func GetUserByID(w http.ResponseWriter, req *http.Request) {

	httpStatus := req.Method

	if httpStatus == "GET" {
		w.Header().Set("Content-Type", "application/json")
		docId, _ := primitive.ObjectIDFromHex("6161682c8ca584daf56a3a66")
		val := mongoDB.FindDocument("users", docId)
		json.NewEncoder(w).Encode(val)
	}

}

func GetPostById(w http.ResponseWriter, req *http.Request) {

	httpStatus := req.Method

	if httpStatus == "GET" {
		w.Header().Set("Content-Type", "application/json")
		docId, _ := primitive.ObjectIDFromHex("616168fcda1796c0d9ecee99")
		val := mongoDB.FindDocument("posts", docId)
		json.NewEncoder(w).Encode(val)
	}
}

func GetAllPostsOfUser(w http.ResponseWriter, req *http.Request) {

	httpStatus := req.Method

	if httpStatus == "GET" {
		w.Header().Set("Content-Type", "application/json")
		posts := mongoDB.GetAllUserPosts("uip")
		json.NewEncoder(w).Encode(posts)
	}
}
