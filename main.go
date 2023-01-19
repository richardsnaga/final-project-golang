package main

import (
	"database/sql"
	"final-project-golang/controllers"
	"final-project-golang/database"
	"fmt"
	"os"
	"strconv"

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
		fmt.Println("Failed load file environment")
	} else {
		fmt.Println("Success read file environment")
	}

	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	DB, _ = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("Database Connection Failed")
		panic(err)
	} else {
		fmt.Println("Database Connected")
	}

	database.DbMigrate(DB)

	defer DB.Close()
	router := gin.Default()
	// comic
	router.GET("/comic", controllers.GetAllComic)
	router.POST("/comic", controllers.CreateComic)
	router.PUT("/comic/:id", controllers.UpdateComic)
	router.DELETE("/comic/:id", controllers.DeleteComic)
	router.GET("/filter-comic", controllers.FilterComic)


	// rating
	router.POST("/rating", controllers.CreateRating)
	router.GET("/rating/:id/comic", controllers.GetRatingByComicId)
	router.PUT("/rating/:id", controllers.UpdateRating)

	// chapter
	router.GET("/chapters", controllers.GetAllChapter)
	router.POST("/chapters", controllers.CreateChapter)
	router.PUT("/chapters/:id", controllers.UpdateChapter)
	router.DELETE("/chapters/:id", controllers.DeleteChapter)

	// comment
	router.GET("/comments", controllers.GetAllComment)
	router.POST("/comments", controllers.CreateComment)
	router.PUT("/comments/:id", controllers.UpdateComment)
	router.DELETE("/comments/:id", controllers.DeleteComment)
	router.GET("/comments/:id/comic", controllers.GetCommentByComicId)

	router.Run("localhost:8090")
}
