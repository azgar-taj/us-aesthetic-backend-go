package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StoryItem struct {
	Id 	  	 primitive.ObjectID 	`json:"id" bson:"_id"`
	Date  	 string  	`json:"date"`
	Title 	 string 	`json:"title"`
	ImageUrl string 	`json:"imageUrl"`
}
