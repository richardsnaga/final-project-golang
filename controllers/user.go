package controllers

import (
	"final-project-golang/database"
	"final-project-golang/library"
	"final-project-golang/repository"
	"final-project-golang/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user structs.User
	var errorUser []library.Error

	err := c.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}

	if len(errorUser) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Error":  errorUser,
		})
		return
	}

	err = repository.Register(database.DbConnection, user)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Registration user",
	})
}

// func Login(c *gin.Context) {
// 	var user structs.User
// 	var errorUser []library.Error

// 	err := c.ShouldBindJSON(&user)
// 	if err != nil {
// 		panic(err)
// 	}

// 	if len(errorUser) > 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"Status": false,
// 			"Error":  errorUser,
// 		})
// 		return
// 	}

// 	err = repository.Register(database.DbConnection, user)
// 	if err != nil {
// 		panic(err)
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"result": err,
// 	})
// }

func UpdateUser(c *gin.Context) {
	var user structs.User

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}

	user.Id = int(id)

	err = repository.UpdateUser(database.DbConnection, user)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update user",
	})
}

func DeleteUser(c *gin.Context) {
	var user structs.User

	id, err := strconv.Atoi(c.Param("id"))

	user.Id = int(id)

	err = repository.DeleteUser(database.DbConnection, user)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete user",
	})
}
