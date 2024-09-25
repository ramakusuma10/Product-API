package checklistcontroller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ramakusuma10/ginproject/models"
)

var validate = validator.New()

func CreateChecklist(c *gin.Context) {
	var checklist models.Checklist

	if err := c.ShouldBindJSON(&checklist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(&checklist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Create(&checklist)
	c.JSON(http.StatusOK, gin.H{"message": "Checklist created successfully", "checklist": checklist})
}

func GetChecklists(c *gin.Context) {
	var checklists []models.Checklist
	models.DB.Preload("ChecklistItems").Find(&checklists)
	c.JSON(http.StatusOK, gin.H{"checklists": checklists})
}

func GetChecklist(c *gin.Context) {
	checklistID, _ := strconv.Atoi(c.Param("checklistId"))
	var checklist models.Checklist
	if err := models.DB.Preload("ChecklistItems").First(&checklist, checklistID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Checklist not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"checklist": checklist})
}

func DeleteChecklist(c *gin.Context) {
	checklistID, _ := strconv.Atoi(c.Param("checklistId"))
	models.DB.Delete(&models.Checklist{}, checklistID)
	c.JSON(http.StatusOK, gin.H{"message": "Checklist deleted successfully"})
}
