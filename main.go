package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/database"
	middleware "github.com/heroku/go-getting-started/middleware"
	routes "github.com/heroku/go-getting-started/routes"
	_ "github.com/heroku/x/hmetrics/onload"
	"go.mongodb.org/mongo-driver/mongo"
)

//get foodCollection
var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)

	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)
}
