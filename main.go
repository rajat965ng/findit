package main

import (
	"findings/controller"

	_ "findings/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title GuardRails Findings API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {
	server := echo.New()
	defer server.Close()

	//server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	grp := server.Group("/api/v1")

	grp.GET("/swagger/*", echoSwagger.WrapHandler)

	repositoryController := controller.NewRepositoryController()
	grp.GET("/repository", repositoryController.FindAll)
	grp.GET("/repository/:status", repositoryController.FindByStatus)
	grp.POST("/repository", repositoryController.Create)
	grp.PUT("/repository", repositoryController.Update)
	grp.DELETE("/repository/:name", repositoryController.Delete)
	grp.PUT("/repository/scan/:name", repositoryController.InitScan)
	grp.PUT("/repository/scan", repositoryController.ExecuteScanner)

	server.Start(":8080")
}
