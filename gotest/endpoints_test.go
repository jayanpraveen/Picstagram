package gotest

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetAllUserPost(t *testing.T) {

	req, err := http.NewRequest("GET", "posts/users/616168fcda1796c0d9ecee99", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllPostsOfUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"_id":"616168fcda1796c0d9ecee99","caption":"cap is good","id":"uip","imageUrl":"http://google.com/post/url","timestamp":"10-09-2021 15:33:40"}]`

	evaluated := strings.TrimSpace(rr.Body.String())

	if evaluated != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			evaluated, expected)
	}

}

func TestPostRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/posts/616168fcda1796c0d9ecee99", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetPostById)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"_id":"616168fcda1796c0d9ecee99","caption":"capisgood","id":"uip","imageUrl":"http://google.com/post/url","timestamp":"10-09-202115:33:40"}`

	evaluated := strings.Join(strings.Fields(rr.Body.String()), "")

	if evaluated != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			evaluated, expected)
	}

}

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}

func TestGetRequest(t *testing.T) {

	req, err := http.NewRequest("GET", "localhost/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUserByID)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"_id":"6161682c8ca584daf56a3a66","email":"jay@gmail.com","id":"id7234","name":"jay","password":"$2a$04$RjCeJdYpys9G5McBFcqX0OUs0b.O7dDxsIROXbrX6JXM6yPIK0s6u"}`

	evaluated := strings.Join(strings.Fields(rr.Body.String()), "")

	if evaluated != expected {
		t.Errorf("handler returned unexpected body: got %s wanted %s",
			rr.Body.String(), expected)
	}
}
