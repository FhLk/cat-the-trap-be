package main

import (
	context "context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

type Album struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func main() {
	//router := gin.Default()
	//router.GET("/albums", getAlbums)
	//router.GET("/albums/:id", getAlbumByID)
	//router.POST("/albums", postAlbum)
	//router.PUT("/albums/:id", updateAlbum)
	//router.DELETE("/albums/:id", deleteAlbum)

	//router.Run("localhost:8080")

}

func getAlbums(c *gin.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	database := client.Database("Golang")
	albumsCollection := database.Collection("Albums")
	cursor, err := albumsCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var albums []bson.M
	if err = cursor.All(ctx, &albums); err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	database := client.Database("Golang")
	albumsCollection := database.Collection("Albums")
	if err != nil {
		log.Fatal(err)
	}

	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	// Create a filter for the _id field matching the given ID
	filter := bson.M{"_id": objID}

	// Find the matching document
	cursor, err := albumsCollection.Find(context.Background(), filter)
	fmt.Println(cursor)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "document not found"})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var album []bson.M
	if err = cursor.All(ctx, &album); err != nil {
		log.Fatal(err)
	}

	// Return a success message
	c.IndentedJSON(http.StatusOK, album)
}

func deleteAlbum(c *gin.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	database := client.Database("Golang")
	albumsCollection := database.Collection("Albums")
	if err != nil {
		log.Fatal(err)
	}

	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	filter := bson.M{"_id": objID}
	result, err := albumsCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "album not found"})
		return
	}

	// Return a success message
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Delete Successfully"})
}

func postAlbum(c *gin.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	database := client.Database("Golang")
	albumsCollection := database.Collection("Albums")
	if err != nil {
		log.Fatal(err)
	}
	var newAlbum Album
	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := albumsCollection.InsertOne(context.Background(), newAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the ID of the newly created person
	c.IndentedJSON(http.StatusOK, gin.H{"id": result.InsertedID})
}

func updateAlbum(c *gin.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	database := client.Database("Golang")
	albumsCollection := database.Collection("Albums")
	if err != nil {
		log.Fatal(err)
	}
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var editAlbum Album
	err = c.ShouldBindJSON(&editAlbum)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filter := bson.M{"_id": objID}

	// Create an update for the fields in the request body
	update := bson.M{
		"$set": bson.M{
			"title": editAlbum.Title,
			"artit": editAlbum.Artist,
			"price": editAlbum.Price,
		},
	}

	result, err := albumsCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Check if a document was updated
	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
		return
	}

	// Return a success message
	c.JSON(http.StatusOK, gin.H{"message": "Update Successfully"})

}
