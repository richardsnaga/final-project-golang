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

func GetAllComment(c *gin.Context) {
	var (
		result gin.H
	)

	comment, err := repository.GetAllComment(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": comment,
		}
	}

	c.JSON(http.StatusOK, result)
}

func CreateComment(c *gin.Context) {
	var comment structs.Comment
	var errorComment []library.Error

	err := c.ShouldBindJSON(&comment)
	if err != nil {
		panic(err)
	}

	
	if len(errorComment) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Error":  errorComment,
		})
		return
	}

	err = repository.CreateComment(database.DbConnection, comment)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert comment",
	})
}

func UpdateComment(c *gin.Context) {
	var comment structs.Comment
	var errorComment []library.Error

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&comment)
	if err != nil {
		panic(err)
	}

	comment.Id = int(id)

	if len(errorComment) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Error":  errorComment,
		})
		return
	}
	err = repository.UpdateComment(database.DbConnection, comment)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update comment",
	})
}

func DeleteComment(c *gin.Context) {
	var comment structs.Comment

	id, err := strconv.Atoi(c.Param("id"))

	comment.Id = int(id)

	err = repository.DeleteComment(database.DbConnection, comment)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete comment",
	})
}

func GetCommentByComicId(c *gin.Context) {
	var (
		result gin.H
	)

	id, _ := strconv.Atoi(c.Param("id"))
	comment, err := repository.GetCommentByComicId(database.DbConnection,id)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": comment,
		}
	}

	c.JSON(http.StatusOK, result)
}