package main

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ToString(x any) (y string) {
	return fmt.Sprintf("%s", x)
}

func GetType(x any) (y reflect.Type) {
	return reflect.TypeOf(x)
}

func ExtractObjectId(x any) (y string) {
	var s1 = ToString(x)
	res3 := strings.Split(s1, "\"")
	return res3[1]
}

func MongoGetCollection(collName string) (context.Context, *mongo.Client, *mongo.Collection) {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database(dbName).Collection(collName)
	return ctx, client, collection
}

func MhInsert(x any) (y string) {
	ctx, client, collection := MongoGetCollection(ToString(GetType(x)))
	res, _ := collection.InsertOne(ctx, x)
	client.Disconnect(ctx)
	return ExtractObjectId(res.InsertedID)
}

func MhList[T any]() (y []T) {
	t_str := ToString(GetType(y))
	t_str = t_str[2:len(t_str)]
	ctx, client, collection := MongoGetCollection(t_str)
	cur, _ := collection.Find(ctx, bson.D{})
	var ctcts []T
	if err := cur.All(ctx, &ctcts); err != nil {
		panic(err)
	}
	cur.Close(ctx)
	client.Disconnect(ctx)
	return ctcts
}

func MhUpdate(x any) (y bool) {
	var id = ExtractObjectId(reflect.ValueOf(&x).Elem().Elem().FieldByName("ID"))
	objectId, _ := primitive.ObjectIDFromHex(id)
	ctx, client, collection := MongoGetCollection(ToString(GetType(x)))
	_, er := collection.UpdateOne(ctx, bson.M{"_id": objectId}, bson.D{{"$set", x}})
	client.Disconnect(ctx)
	return er == nil
}

func MhDelete(x any) (y bool) {
	var id = ExtractObjectId(reflect.ValueOf(&x).Elem().Elem().FieldByName("ID"))
	objectId, _ := primitive.ObjectIDFromHex(id)
	ctx, client, collection := MongoGetCollection(ToString(GetType(x)))
	_, er := collection.DeleteOne(ctx, bson.M{"_id": objectId})
	client.Disconnect(ctx)
	return er == nil
}

// ==================================================================================================== //

var url = "mongodb+srv://zparinthornk:es7HBYTR@cluster0.nymtjxi.mongodb.net/?retryWrites=true&w=majority"
var dbName = "InterfaceDemo4"

var POSITION_OUTSOUIRCE = "OUTSOUIRCE"
var POSITION_JUNIOR = "JUNIOR"
var POSITION_SENIOR1 = "SENIOR1"
var POSITION_SENIOR2 = "SENIOR2"
var POSITION_MANAGER = "MANAGER"

var STATUS_PURPOSED = "PURPOSED"
var STATUS_DESIGN = "DESIGN"
var STATUS_DEV = "DEV"
var STATUS_TESTING = "TESTING"
var STATUS_GOLIVE = "GOLIVE"

// ==================================================================================================== //

type Member struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Position    string             `bson:"position"`
	Phone       string             `bson:"phone"`
}

type Machine struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	IpAddrIn    string             `bson:"ipAddrIn"`
	IpAddrEx    string             `bson:"ipAddrEx"`
}

type Authentication struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
}

// ==================================================================================================== //

type ServiceProvider struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Status      string             `bson:"status"`
	Machine     Machine            `bson:"machine"`
	Owner       Member             `bson:"owner"`
	HostName    string             `bson:"hostName"`
	Port        int                `bson:"port"`
	Protocol    string             `bson:"protocol"`
	URL         string             `bson:"url"`
}

type Subscription struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Name           string             `bson:"name"`
	Description    string             `bson:"description"`
	Status         string             `bson:"status"`
	Provider       ServiceProvider    `bson:"provider"`
	Authentication Authentication     `bson:"authentication"`
}

type ESubscription struct {
	SubscriptionDEV []Subscription `bson:"subscriptionDEV"`
	SubscriptionSIT []Subscription `bson:"subscriptionSIT"`
	SubscriptionUAT []Subscription `bson:"subscriptionUAT"`
	SubscriptionPRD []Subscription `bson:"subscriptionPRD"`
}

type EServiceProvider struct {
	ProviderDEV ServiceProvider `bson:"providerDEV"`
	ProviderSIT ServiceProvider `bson:"providerSIT"`
	ProviderUAT ServiceProvider `bson:"providerUAT"`
	ProviderPRD ServiceProvider `bson:"providerPRD"`
}

type Interface struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Status      string             `bson:"status"`
	PIC         Member             `bson:"pic"`
	Backends    ESubscription      `bson:"backends"`
	Exposed     EServiceProvider   `bson:"exposed"`
}

type Pipeline struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Name              string             `bson:"name"`
	Description       string             `bson:"description"`
	InterfacesPlanned []Interface        `bson:"interfacesPlanned"`
	InterfacesActual  []Interface        `bson:"interfacesActual"`
}

// ==================================================================================================== //

func main() {

}

// ==================================================================================================== //
