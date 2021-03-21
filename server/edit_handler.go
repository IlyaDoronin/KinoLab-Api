package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

//Группа запросов на получение строк для редактирования

func getAuthorEditPage(c *gin.Context) {
	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(404, gin.H{"error": "неверный формат"})
		return
	}

	author := DBAuthorHandler.Select(ID)

	c.JSON(200, gin.H{"author": author})
}

func getActorEditPage(c *gin.Context) {
	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(404, gin.H{"error": "неверный формат"})
		return
	}

	author := DBActorHandler.Select(ID)

	c.JSON(200, gin.H{"author": author})
}

func getGenreEditPage(c *gin.Context) {
	ID, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		c.JSON(404, gin.H{"error": "неверный формат"})
		return
	}

	genre := DBGenreHandler.Select(ID)

	c.JSON(200, gin.H{"genre": genre})
}

func getFilmCommentEditPage(c *gin.Context) {

	ID, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		c.JSON(404, gin.H{"error": "неверный формат"})
		return
	}

	film_comment := DBFilmCommentHandler.Select(ID)

	c.JSON(200, gin.H{"film_comment": film_comment})
}

func getFilmEditPage(c *gin.Context) {

	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(404, gin.H{"error": "неверный формат"})
		return
	}

	film := DBFilmHandler.Select(ID)

	c.JSON(200, gin.H{"film": film})
}

func getFilmGenreEditPage(c *gin.Context) {

	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(404, gin.H{"error": "неверный формат"})
		return
	}

	film_genre := DBFilmGenreHandler.Select(ID)

	c.JSON(200, gin.H{"film_genre": film_genre})
}

func getFilmAuthorEditPage(c *gin.Context) {

	ID, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		c.JSON(404, gin.H{"error": "неверный формат"})
		return
	}

	film_author := DBFilmAuthorHandler.Select(ID)

	c.JSON(200, gin.H{"film_author": film_author})
}

func getFilmActorEditPage(c *gin.Context) {

	ID, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		c.JSON(404, gin.H{"error": "неверный формат"})
		return
	}

	film_actor := DBFilmActorHandler.Select(ID)

	c.JSON(200, gin.H{"film_actor": film_actor})
}
