package main

import (
	"database/sql"
	"encoding/base64"
	"final-project-go-sanber/controllers"
	"final-project-go-sanber/database"
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func Authentication(username, password string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Basic ") {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		auth := strings.TrimPrefix(authHeader, "Basic ")
		decodedAuth, err := base64.StdEncoding.DecodeString(auth)
		if err != nil {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		decodedAuthStr := string(decodedAuth)
		creds := strings.Split(decodedAuthStr, ":")
		if len(creds) != 2 {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		if creds[0] != username || creds[1] != password {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// if authHeader != fmt.Sprintf("Basic %s", username+":"+password) {
		// 	c.JSON(401, gin.H{"error": "Unauthorized"})
		// 	c.Abort()
		// 	return
		// }
		c.Next()
	}
}

func main() {
	err = godotenv.Load("config/.env")

	if err != nil {
		panic("Error loading .env file")
	}
	psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`, os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"))
	// psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`, os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		panic(err)
	}
	defer DB.Close()
	if err != nil {
		panic(err)
	}
	database.DBMigrate(DB)

	router := gin.Default()
	username := "admin"
	password := "admin123"
	router.Use(Authentication(username, password))
	router.GET("/anime", controllers.GetAllAnime)
	router.POST("/anime", controllers.InsertAnime)
	router.PUT("/anime:id", controllers.UpdateAnime)
	router.DELETE("/anime:id", controllers.DeleteAnime)

	// router.Run(":8080")
	router.Run(":" + os.Getenv("PGPORT"))
	fmt.Println("Successfully Connected")
}
