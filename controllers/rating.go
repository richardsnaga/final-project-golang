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

func CreateRating(c *gin.Context) {
	var rating structs.Rating
	var errorRating []library.Error

	err := c.ShouldBindJSON(&rating)
	if err != nil {
		panic(err)
	}

	if rating.Rate < 1 {
		result := library.Error{"Rating minimal 1"}.Validate()
		errorRating = append(errorRating, result)
	} else if rating.Rate > 5 {
		result := library.Error{"Rating Maximal 5"}.Validate()
		errorRating = append(errorRating, result)
	}

	if len(errorRating) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Error":  errorRating,
		})
		return
	}

	err = repository.CreateRating(database.DbConnection, rating)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert rating",
	})
}

func GetRatingByComicId(c *gin.Context) {
	var (
		result gin.H
	)

	id, _ := strconv.Atoi(c.Param("id"))
	rating, err := repository.GetRatingByComicId(database.DbConnection, id)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": rating,
		}
	}

	c.JSON(http.StatusOK, result)
}

func UpdateRating(c *gin.Context) {
	var rating structs.Rating
	var errorRating []library.Error

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&rating)
	if err != nil {
		panic(err)
	}

	rating.Id = int(id)

	if rating.Rate < 1 {
		result := library.Error{"Rating minimal 1"}.Validate()
		errorRating = append(errorRating, result)
	} else if rating.Rate > 5 {
		result := library.Error{"Rating Maximal 5"}.Validate()
		errorRating = append(errorRating, result)
	}

	err = repository.UpdateRating(database.DbConnection, rating)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update comic",
	})
}
