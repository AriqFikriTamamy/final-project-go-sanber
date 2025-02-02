package main

import (
	"database/sql"
	"final-project-go-sanber/controllers"
	"final-project-go-sanber/database"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")

	if err != nil {
		panic("Error loading .env file")
	}
	psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`, os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"))

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
	router.GET("/anime", controllers.GetAllAnime)
	router.POST("/anime", controllers.InsertAnime)
	router.PUT("/anime:id", controllers.UpdateAnime)
	router.DELETE("/anime:id", controllers.DeleteAnime)

	router.Run(":" + os.Getenv("PGPORT"))
	fmt.Println("Successfully Connected")
}
