package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ramakusuma10/ginproject/controllers/authcontroller"
	"github.com/ramakusuma10/ginproject/controllers/checklistcontroller"
	"github.com/ramakusuma10/ginproject/controllers/checklistitemcontroller"
	"github.com/ramakusuma10/ginproject/middleware"

	"github.com/ramakusuma10/ginproject/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.POST("/api/register", authcontroller.Register)
	r.POST("/api/login", authcontroller.Login)

	// Group routes that require authentication
	authRoutes := r.Group("/api", middleware.JWTAuthMiddleware()) // Middleware untuk verifikasi JWT
	{
		authRoutes.POST("/checklist", checklistcontroller.CreateChecklist)
		authRoutes.GET("/checklist", checklistcontroller.GetChecklists)
		authRoutes.GET("/checklist/:checklistId", checklistcontroller.GetChecklist)
		authRoutes.DELETE("/checklist/:checklistId", checklistcontroller.DeleteChecklist)

		authRoutes.POST("/checklist/:checklistId/item", checklistitemcontroller.CreateChecklistItem)
		authRoutes.GET("/checklist/:checklistId/item/:itemId", checklistitemcontroller.GetChecklist)
		authRoutes.PUT("/checklist/:checklistId/item/:itemId", checklistitemcontroller.UpdateChecklistItem)
		authRoutes.PUT("/checklist/:checklistId/item/:itemId/status", checklistitemcontroller.UpdateItemStatus)
		authRoutes.DELETE("/checklist/:checklistId/item/:itemId", checklistitemcontroller.DeleteChecklistItem)
	}

	r.Run()
}
