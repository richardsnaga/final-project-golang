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

func GetAllComic(c *gin.Context) {
	var (
		result gin.H
	)

	comic, err := repository.GetAllComic(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": comic,
		}
	}

	c.JSON(http.StatusOK, result)
}

func CreateComic(c *gin.Context) {
	var comic structs.Comic
	var errorComic []library.Error

	err := c.ShouldBindJSON(&comic)
	if err != nil {
		panic(err)
	}

	if _, err := url.ParseRequestURI(comic.ImageURL); err != nil {
		result := library.Error{"Gambar Harus Berupa URL"}.Validate()
		errorComic = append(errorComic, result)
	}

	if len(errorComic) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Error":  errorComic,
		})
		return
	}

	err = repository.CreateComic(database.DbConnection, comic)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert comic",
	})
}

func UpdateComic(c *gin.Context) {
	var comic structs.Comic
	var errorComic []library.Error

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&comic)
	if err != nil {
		panic(err)
	}

	comic.Id = int(id)

	if _, err := url.ParseRequestURI(comic.ImageURL); err != nil {
		result := library.Error{"Gambar Harus Berupa URL"}.Validate()
		errorComic = append(errorComic, result)
	}

	err = repository.UpdateComic(database.DbConnection, comic)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update comic",
	})
}

func DeleteComic(c *gin.Context) {
	var comic structs.Comic

	id, err := strconv.Atoi(c.Param("id"))

	comic.Id = int(id)

	err = repository.DeleteComic(database.DbConnection, comic)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete comic",
	})
}

func FilterComic(c *gin.Context) {
	var (
		result gin.H
	)

	genre := c.Query("genre")
	tipe := c.Query("type")
	status, _ := strconv.ParseBool(c.Query("status"))
	comic, err := repository.FilterComic(database.DbConnection,genre,tipe,status)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": comic,

		}
	}

	c.JSON(http.StatusOK, result)
}

