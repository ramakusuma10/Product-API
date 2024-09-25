package checklistitemcontroller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ramakusuma10/ginproject/models"
)

var validate = validator.New()

func GetChecklist(c *gin.Context) {
	checklistID, _ := strconv.Atoi(c.Param("checklistId"))
	var checklist models.Checklist
	if err := models.DB.Preload("ChecklistItems").First(&checklist, checklistID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Checklist not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"checklist": checklist})
}

func CreateChecklistItem(c *gin.Context) {
	checklistID, _ := strconv.Atoi(c.Param("checklistId"))
	var item models.ChecklistItem

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item.ChecklistID = uint(checklistID)
	models.DB.Create(&item)
	c.JSON(http.StatusOK, gin.H{"message": "Checklist item created successfully", "item": item})
}

func UpdateChecklistItem(c *gin.Context) {
	itemID, _ := strconv.Atoi(c.Param("itemId"))
	var item models.ChecklistItem

	if err := models.DB.First(&item, itemID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Save(&item)
	c.JSON(http.StatusOK, gin.H{"message": "Checklist item updated successfully", "item": item})
}

func UpdateItemStatus(c *gin.Context) {
	itemID, _ := strconv.Atoi(c.Param("itemId"))
	var item models.ChecklistItem
	if err := models.DB.First(&item, itemID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&item).Update("status", item.Status)
	c.JSON(http.StatusOK, gin.H{"message": "Item status updated successfully", "item": item})
}

// DeleteChecklistItem - API to delete a checklist item
func DeleteChecklistItem(c *gin.Context) {
	itemID, _ := strconv.Atoi(c.Param("itemId"))
	models.DB.Delete(&models.ChecklistItem{}, itemID)
	c.JSON(http.StatusOK, gin.H{"message": "Checklist item deleted successfully"})
}
