package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/todoapp/server02/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
	loadTheEnv()
	createDBInstance()
}

func loadTheEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading the .env file")
	}

}

func createDBInstance() {
	connectionString := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	collName := os.Getenv("DB_COLLECTION_NAME")

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("error: database is not connected")
	}

	err = client.Ping(context.TODO(), nil) //uses the MongoDB Go driver to ping the MongoDB server and check whether a connection can be established.
	if err != nil {                        //if server is not properly working then it give the error
		log.Fatal("error: client ping error")
	}

	fmt.Println("connected to the mongodb")
	collection = client.Database(dbName).Collection(collName)
	fmt.Println("collection instace created")
}

func Getalltasks(ctx *gin.Context) {
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal("error: to find the collection")
	}
	defer cursor.Close(context.Background())
	var results []models.ToDoList
	for cursor.Next(context.Background()) {
		var result models.ToDoList
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal("error: decode the string")
		}
		results = append(results, result)
	}

	ctx.JSON(http.StatusOK, results)

}

func Createtask(ctx *gin.Context) {

	var user models.ToDoList
	err := ctx.ShouldBind(&user) //It will bind the post request data to the user
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	insertResult, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data"})
		return
	}
	ctx.JSON(200, gin.H{"message": "Data inserted successfully", "id": insertResult.InsertedID})

}

func Complete(ctx *gin.Context) {
	taskId := ctx.Param("id")
	var user models.ToDoList
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "NOt the valid payload"})
	}
	filter := bson.M{"id": taskId}
	update := bson.M{"$set": user}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error in the updating the data"})
	}
	if result.ModifiedCount == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)

}

func Deleteone(ctx *gin.Context) {
	fmt.Println("Delete one course")
	taskID := ctx.Param("id")
	//taskID := "1"
	result, err := collection.DeleteOne(context.Background(), bson.M{"id": taskID})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}
	if result.DeletedCount == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted succesfully"})
}
