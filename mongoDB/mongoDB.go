package mongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnection() *mongo.Client {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Panic(err)
	}
	return client
}

func InsertDocument(collName string, bsonValue bson.D) {
	client := MongoConnection()
	collection := client.Database("picstagram").Collection(collName)
	res, err := collection.InsertOne(context.TODO(), bsonValue)

	log.SetPrefix("response id: ")
	log.Println(res.InsertedID)

	if err != nil {
		log.Println(err)
	}
}

func FindDocument(collName string, docId primitive.ObjectID) primitive.M {
	client := MongoConnection()
	collection := client.Database("picstagram").Collection(collName)

	findone_result := collection.FindOne(context.TODO(), bson.M{"_id": docId})

	var bson_obj bson.M
	if err2 := findone_result.Decode(&bson_obj); err2 != nil {
		log.Println(err2)
	}
	return bson_obj

}

func GetAllUserPosts(userId string) []primitive.M {

	// here posts.id == users._id
	client := MongoConnection()
	collection := client.Database("picstagram").Collection("posts")

	filterCursor, err := collection.Find(context.TODO(), bson.M{"id": userId})
	if err != nil {
		log.Println(err)
	}
	var bson_objs []bson.M
	if err = filterCursor.All(context.TODO(), &bson_objs); err != nil {
		log.Println(err)
	}
	return bson_objs
}
