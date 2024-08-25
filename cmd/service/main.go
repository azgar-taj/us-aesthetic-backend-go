package main

import (
	"os"
	"fmt"
	"us-aesthetic-backend-go/api"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Creating a server...")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	storyItemController, err := api.NewStoryItemController(
		os.Getenv("MONGO_CLUSTER"),
		os.Getenv("MONGO_USERNAME"),
		os.Getenv("MONGO_PASSWORD"),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer storyItemController.Close()
	router.GET("/story_items", storyItemController.GetAllStoryItems)
	router.GET("/GetAllStoriesData", storyItemController.GetAllStoryItems)
	router.GET("/story_items/:id", storyItemController.GetStoryItem)
	router.POST("/story_items", storyItemController.CreateStoryItem)
	router.PUT("/story_items/:id", storyItemController.UpdateStoryItem)
	router.DELETE("/story_items/:id", storyItemController.DeleteStoryItem)
	fmt.Println("Server is running on localhost:80")
    router.Run(":80")
}