package main

import (
	"database/sql"
	"final-project-golang/controllers"
	"final-project-golang/database"
	"final-project-golang/structs"
	"fmt"
	"net/http"
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

	port, err := strconv.Atoi(os.Getenv("PGPORT"))
	if err != nil {
		panic(err)
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		port,
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"))

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

	authorized := router.Group("/")
	authorized.Use(basicAuth())
	{
		// comic
		authorized.POST("/comic", controllers.CreateComic)
		authorized.PUT("/comic/:id", controllers.UpdateComic)
		authorized.DELETE("/comic/:id", controllers.DeleteComic)

		// rating
		authorized.POST("/rating", controllers.CreateRating)
		authorized.PUT("/rating/:id", controllers.UpdateRating)

		// chapter
		authorized.POST("/chapters", controllers.CreateChapter)
		authorized.PUT("/chapters/:id", controllers.UpdateChapter)
		authorized.DELETE("/chapters/:id", controllers.DeleteChapter)

		// comment
		authorized.POST("/comments", controllers.CreateComment)
		authorized.PUT("/comments/:id", controllers.UpdateComment)
		authorized.DELETE("/comments/:id", controllers.DeleteComment)

		// user
		authorized.PUT("/user/:id", controllers.UpdateUser)
		authorized.DELETE("/user/:id", controllers.DeleteUser)
		authorized.GET("/logout", logout)

	}

	// comic
	router.GET("/comic", controllers.GetAllComic)
	router.GET("/filter-comic", controllers.FilterComic)

	// rating
	router.GET("/rating/:id/comic", controllers.GetRatingByComicId)

	// chapter
	router.GET("/chapters", controllers.GetAllChapter)

	// comment
	router.GET("/comments", controllers.GetAllComment)
	router.GET("/comments/:id/chapter", controllers.GetCommentByChapterId)

	// user
	router.POST("/regis", controllers.Register)
	router.POST("/login", Login)

	router.Run(":" + os.Getenv("PORT"))
}

func Login(c *gin.Context) {
	var json structs.User

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM users WHERE email=$1 AND password=$2", json.Email, json.Password).Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if count == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Login Success"})
}

func basicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Basic Authentication credentials
		username, password, hasAuth := c.Request.BasicAuth()
		if hasAuth && username == "admin" && password == "password" {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
		}
	}
}

func logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
