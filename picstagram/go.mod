module example.com/picstagram

go 1.17

require (
	example.com/mongoDB v0.0.0-00010101000000-000000000000
	go.mongodb.org/mongo-driver v1.7.3
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
)

replace example.com/mongoDB => ../mongoDb

require (
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.0.2 // indirect
	github.com/xdg-go/stringprep v1.0.2 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e // indirect
	golang.org/x/text v0.3.5 // indirect
)
