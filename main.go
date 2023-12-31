package main

import (
	"fmt"
	"os"
	"sistem-pengelolaan-bank-sampah/database"
	"sistem-pengelolaan-bank-sampah/pkg/middleware"
	"sistem-pengelolaan-bank-sampah/pkg/postgres"
	"sistem-pengelolaan-bank-sampah/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// load environment variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading environment variables file, the apps will read global environtment variabels on this system")
	}

	// database initialization
	postgres.DatabaseInit()
	// redis.RedisInit()

	// database migration & seeder
	database.DropMigration()
	database.RunMigration()
	database.RunSeeder()

	// gin Mode
	gin.SetMode(os.Getenv("GIN_MODE"))

	// create new router
	router := gin.Default()

	// call logger middleware before route to any routes
	router.Use(middleware.Logger())

	//	set up CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Replace with your allowed origins
	config.AllowMethods = []string{"HEAD", "OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE"}
	config.AllowHeaders = []string{"Origin", "X-Requested-With", "Content-Type", "Authorization"}

	// Add cors middleware on all route
	router.Use(cors.New(config))

	// call routerinit with pathprefix
	routes.RouterInit(router.Group("/api/v1"))

	// file server endpoint
	router.Static("/static", "./uploads")

	// Running services
	fmt.Println("server running on localhost:" + os.Getenv("PORT"))
	router.Run(":" + os.Getenv("PORT"))
}
