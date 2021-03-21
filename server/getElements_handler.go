package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func getAuthors(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(404, gin.H{"error": "неверный формат"})
		return
	}
	authors := DBAuthorHandler.SelectRange(page)
	c.JSON(200, gin.H{"authors": authors})
}

func getActors(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(404, gin.H{"error": "неверный формат"})
		return
	}
	actors := DBActorHandler.SelectRange(page)
	c.JSON(200, gin.H{"actors": actors})
}

func getFilms(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(404, gin.H{"error": "неверный формат"})
		return
	}
	films := DBFilmHandler.SelectRange(page)
	c.JSON(200, gin.H{"films": films})
}

func getGenres(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(404, gin.H{"error": "неверный формат"})
		return
	}
	genres := DBGenreHandler.SelectRange(page)
	c.JSON(200, gin.H{"genres": genres})
}

func getFilmsAuthors(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(404, gin.H{"error": "неверный формат"})
		return
	}
	films_authors := DBFilmAuthorHandler.SelectRange(page)
	c.JSON(200, gin.H{"films_authors": films_authors})
}

func getFilmsActors(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(404, gin.H{"error": "неверный формат"})
		return
	}
	films_actors := DBFilmActorHandler.SelectRange(page)
	c.JSON(200, gin.H{"films_actors": films_actors})
}

func getFilmsGenres(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(404, gin.H{"error": "неверный формат"})
		return
	}
	films_genres := DBFilmGenreHandler.SelectRange(page)
	c.JSON(200, gin.H{"films_genres": films_genres})
}

func getFilmComments(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(404, gin.H{"error": "неверный формат"})
		return
	}
	film_comments := DBFilmCommentHandler.SelectRange(page)
	c.JSON(200, gin.H{"film_comments": film_comments})
}

//Для получения ВСЕХ элементов

func getAllFilms(c *gin.Context) {
	films := DBFilmHandler.SelectAll()
	c.JSON(200, gin.H{"films": films})
}

func getAllFilmsAuthors(c *gin.Context) {
	films_authors := DBFilmAuthorHandler.SelectAll()
	c.JSON(200, gin.H{"films_authors": films_authors})
}

func getAllFilmsActors(c *gin.Context) {
	films_actors := DBFilmActorHandler.SelectAll()
	c.JSON(200, gin.H{"films_actors": films_actors})
}

func getAllFilmsGenres(c *gin.Context) {
	films_genres := DBFilmGenreHandler.SelectAll()
	c.JSON(200, gin.H{"films_genres": films_genres})
}

func getAllFilmComments(c *gin.Context) {
	film_comments := DBFilmCommentHandler.SelectAll()
	c.JSON(200, gin.H{"film_comments": film_comments})
}
