package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Album struct {
	Id     primitive.ObjectID `bson:"_id,omitempty"`
	Title  string             `bson:"title,omitempty"`
	Artist string             `bson:"artist,omitempty"`
	Price  float64            `bson:"price,omitempty"`
}

func GetAlbums() ([]Album, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	albumsCollection := client.Database("Golang").Collection("Albums")
	cursor, err := albumsCollection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var albums []Album
	for cursor.Next(context.Background()) {
		var album Album
		err := cursor.Decode(&album)
		if err != nil {
			return nil, err
		}
		fmt.Println(album)
		albums = append(albums, album)
	}
	return albums, nil
}

func GetAlbumByID(id string) (*Album, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	albumsCollection := client.Database("Golang").Collection("Albums")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{"_id", objectID}}
	result := albumsCollection.FindOne(context.Background(), filter)
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, result.Err()
	}

	var album Album
	err = result.Decode(&album)
	if err != nil {
		return nil, err
	}
	return &album, nil
}

func PostAlbum() {

}
