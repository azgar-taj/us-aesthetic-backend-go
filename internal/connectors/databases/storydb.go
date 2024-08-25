package databases

import (
	"context"
	"fmt"
	"log"
	"us-aesthetic-backend-go/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StoryDB struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewStoryDB(clusterName string, username string, password string) (*StoryDB, error) {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@%s.uo3wozi.mongodb.net/?retryWrites=true&w=majority&appName=SmrutajUs", username, password, clusterName))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
		return nil, err
	}
	collection := client.Database("stories").Collection("stories")
	return &StoryDB{client: client, collection: collection}, nil
}

func (s *StoryDB) Close() {
	err := s.client.Disconnect(context.Background())
	if err != nil {
		log.Printf("Failed to disconnect from MongoDB: %v", err)
	}
}

func (s *StoryDB) InsertStoryItem(item model.StoryItem) (primitive.ObjectID, error) {
	item.Id = primitive.NewObjectID()
	result, err := s.collection.InsertOne(context.Background(), item)
	if err != nil {
		log.Printf("Failed to insert story item: %v", err)
		return primitive.NilObjectID, err
	}
	log.Println("Inserted story item:", result.InsertedID)
	return result.InsertedID.(primitive.ObjectID), nil
}

func (s *StoryDB) DeleteStoryItem(id string) error {
	formatted_id, _ := primitive.ObjectIDFromHex(id)
	_, err := s.collection.DeleteOne(context.Background(), bson.D{{"_id", formatted_id}})
	if err != nil {
		log.Printf("Failed to delete story item: %v", err)
		return err
	}
	log.Println("Deleted story item with ID:", id)
	return nil
}

func (s *StoryDB) UpdateStoryItem(item model.StoryItem) (primitive.ObjectID, error) {
	_, err := s.collection.UpdateOne(context.Background(), bson.D{{"id", item.Id}}, bson.D{{"$set", item}})
	if err != nil {
		log.Printf("Failed to update story item: %v", err)
		return primitive.NilObjectID, err
	}
	log.Println("Updated story item:", item.Id)
	return item.Id, nil
}

func (s *StoryDB) GetStoryItem(id string) (model.StoryItem, error) {
	var item model.StoryItem
	formatted_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Failed to convert ID to ObjectID: %v", err)
		return model.StoryItem{}, err
	}
	err = s.collection.FindOne(context.Background(), bson.D{{"_id", formatted_id}}).Decode(&item)
	if err != nil {
		log.Printf("Failed to get story item: %v", err)
		return model.StoryItem{}, err
	}
	log.Println("Retrieved story item:", item)
	return item, nil
}

func (s *StoryDB) GetAllStoryItems() ([]model.StoryItem, error) {
	cursor, err := s.collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Printf("Failed to get all story items: %v", err)
		return nil, err
	}
	defer cursor.Close(context.Background())
	var items []model.StoryItem
	for cursor.Next(context.Background()) {
		var item model.StoryItem
		cursor.Decode(&item)
		items = append(items, item)
	}
	log.Println("Retrieved all story items:", items)
	return items, nil
}