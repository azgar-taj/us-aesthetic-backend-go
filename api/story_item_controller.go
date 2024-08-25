package api

import (
	"us-aesthetic-backend-go/internal/model"
	"us-aesthetic-backend-go/internal/connectors/databases"
	"net/http"
	"github.com/gin-gonic/gin"
)

type StoryItemController struct {
	storyDB *databases.StoryDB
}

func NewStoryItemController(clusterName string, username string, password string) (*StoryItemController, error) {
	storyDB, err := databases.NewStoryDB(clusterName, username, password)
	if err != nil {
		return nil, err
	}
	return &StoryItemController{storyDB: storyDB}, nil
}

func (s *StoryItemController) Close() {
	s.storyDB.Close()
}

func (s *StoryItemController) GetAllStoryItems(c *gin.Context) {
	items, err := s.storyDB.GetAllStoryItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (s *StoryItemController) GetStoryItem(c *gin.Context) {
	id := c.Param("id")
	item, err := s.storyDB.GetStoryItem(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

func (s *StoryItemController) CreateStoryItem(c *gin.Context) {
	var item model.StoryItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}
	id, err := s.storyDB.InsertStoryItem(item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, id)
}

func (s *StoryItemController) UpdateStoryItem(c *gin.Context) {
	var item model.StoryItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}
	id, err := s.storyDB.UpdateStoryItem(item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, id)
}

func (s *StoryItemController) DeleteStoryItem(c *gin.Context) {
	id := c.Param("id")
	err := s.storyDB.DeleteStoryItem(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Story item deleted successfully!")
}
