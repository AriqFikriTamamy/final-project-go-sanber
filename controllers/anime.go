package controllers

import (
	"final-project-go-sanber/database"
	"final-project-go-sanber/repository"
	"final-project-go-sanber/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllAnime(c *gin.Context) {
	var (
		result gin.H
	)

	anime, err := repository.GetAllAnime(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": anime,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertAnime(c *gin.Context) {
	var anime structs.Anime

	err := c.BindJSON(&anime)
	if err != nil {
		panic(err)
	}

	err = repository.InsertAnime(database.DbConnection, anime)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, anime)
}

func UpdateAnime(c *gin.Context) {
	var anime structs.Anime
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&anime)
	if err != nil {
		panic(err)
	}

	anime.ID = id

	err = repository.UpdateAnime(database.DbConnection, anime)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, anime)
}

func DeleteAnime(c *gin.Context) {
	var anime structs.Anime
	id, _ := strconv.Atoi(c.Param("id"))

	anime.ID = id
	err := repository.DeleteAnime(database.DbConnection, anime)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, anime)
}
