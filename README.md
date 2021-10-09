# Picstagram

- Run `picstagram.go` to expose rest apis' @ localhost:80, this requires a local mongodb instance running @ port `27017`
- Tests are located at [gotest](./gotest/endpoints_test.go)

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
