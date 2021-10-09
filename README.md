# Picstagram

- Run [`picstagram.go`](./picstagram/picstagram.go) to expose rest apis' @ localhost:80, this requires a local mongodb instance running @ port `27017`
```go
go run picstagram.go
``` 

- Tests are located at [gotest](./gotest/endpoints_test.go)

## End points
 - Create a user: http://localhost/users 
 - Get user with Id: http://localhost/users/:id 
 - Create a post: http://localhost/posts 
 - Get post with Id: http://localhost/posts/:id 
 - List all posts of user: http://localhost/posts/users/:id

## File structure

```
.
├── README.md
├── gotest
│   ├── endpoints_test.go
│   ├── go.mod
│   ├── go.sum
│   └── handlers.go
├── mongoDB
│   ├── go.mod
│   ├── go.sum
│   └── mongoDB.go
└── picstagram
    ├── go.mod
    ├── go.sum
    └── picstagram.go
```
