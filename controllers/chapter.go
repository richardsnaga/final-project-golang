package controllers

import (
	"final-project-golang/database"
	"final-project-golang/library"
	"final-project-golang/repository"
	"final-project-golang/structs"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllChapter(c *gin.Context) {
	var (
		result gin.H
	)

	chapter, err := repository.GetAllChapter(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": chapter,
		}
	}

	c.JSON(http.StatusOK, result)
}


func CreateChapter(c *gin.Context) {
	var chapter structs.Chapter
	var errorChapter []library.Error

	err := c.ShouldBindJSON(&chapter)
	if err != nil {
		panic(err)
	}

	if _, err := url.ParseRequestURI(chapter.ImageUrl); err != nil {
		result := library.Error{"Gambar Harus Berupa URL"}.Validate()
		errorChapter = append(errorChapter, result)
	}

	if len(errorChapter) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Error":  errorChapter,
		})
		return
	}

	err = repository.CreateChapter(database.DbConnection, chapter)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert chapter",
	})
}

func UpdateChapter(c *gin.Context) {
	var chapter structs.Chapter
	var errorChapter []library.Error

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&chapter)
	if err != nil {
		panic(err)
	}

	chapter.Id = int(id)

	if _, err := url.ParseRequestURI(chapter.ImageUrl); err != nil {
		result := library.Error{"Gambar Harus Berupa URL"}.Validate()
		errorChapter = append(errorChapter, result)
	}

	err = repository.UpdateChapter(database.DbConnection, chapter)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update chapter",
	})
}

func DeleteChapter(c *gin.Context) {
	var chapter structs.Chapter

	id, err := strconv.Atoi(c.Param("id"))

	chapter.Id = int(id)

	err = repository.DeleteChapter(database.DbConnection, chapter)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete chapter",
	})
}
