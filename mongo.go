package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Contact struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Email string             `bson:"email"`
	Tags  []string           `bson:"tags"`
}

type Post struct {
	Title string `bson:"title,omitempty"`
	Body  string `bson:"body,omitempty"`
}

var url = "mongodb+srv://zparinthornk:es7HBYTR@cluster0.nymtjxi.mongodb.net/?retryWrites=true&w=majority"
var dbName = "Bogdanov"

func main() {

	var ctct = ContactGetById("63012a164e5aae33ca50a579")

	fmt.Println(ContactDelete(ctct))
}

func CreateDbConnection(collName string) (context.Context, *mongo.Client, *mongo.Collection) {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	/*databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)*/
	collection := client.Database(dbName).Collection(collName)
	return ctx, client, collection
}

func ContactList() []Contact {
	ctx, client, collection := CreateDbConnection("Contact")
	cur, currErr := collection.Find(ctx, bson.D{})
	if currErr != nil {
		panic(currErr)
	}
	var ctcts []Contact
	if err := cur.All(ctx, &ctcts); err != nil {
		panic(err)
	}
	cur.Close(ctx)
	client.Disconnect(ctx)
	return ctcts
}

func ContactGetById(id string) Contact {
	ctx, client, collection := CreateDbConnection("Contact")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}
	cur, currErr := collection.Find(ctx, bson.M{"_id": objectId})
	if currErr != nil {
		panic(currErr)
	}
	var ctcts []Contact
	if err := cur.All(ctx, &ctcts); err != nil {
		panic(err)
	}
	cur.Close(ctx)
	client.Disconnect(ctx)
	return ctcts[0]
}

func ContactInsert(c Contact) string {
	ctx, client, collection := CreateDbConnection("Contact")
	res, _ := collection.InsertOne(ctx, c)
	client.Disconnect(ctx)
	dfs := fmt.Sprintf("%s", res.InsertedID)
	return dfs[9:34]
}

func ContactSave(c Contact) bool {
	ctx, client, collection := CreateDbConnection("Contact")
	_, er := collection.UpdateOne(ctx, bson.M{"_id": c.ID}, bson.D{{"$set", c}})
	client.Disconnect(ctx)
	return er == nil
}

func ContactDelete(c Contact) bool {
	ctx, client, collection := CreateDbConnection("Contact")
	_, er := collection.DeleteOne(ctx, bson.M{"_id": c.ID})
	client.Disconnect(ctx)
	return er == nil
}
