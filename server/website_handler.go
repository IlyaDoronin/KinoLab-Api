package server

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yeahyeahcore/KinoLab-Api/storage"
)

func getFilmsWeb(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(404, gin.H{"error": "неверный формат"})
		return
	}
	films := DBFilmHandler.SelectRangeWeb(page)
	c.JSON(200, gin.H{"films": films})
}

func getWebSiteFilm(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(404, gin.H{"error": "неверный формат"})
		return
	}
	film := DBWebSiteFilmHandler.Select(id)
	c.JSON(200, gin.H{"film": film})
}

func getBanners(c *gin.Context) {
	banners := DBBannersHandler.SelectRange()
	c.JSON(200, gin.H{"banners": banners})
}

func getFilteredFilm(c *gin.Context) {

	//Объявляем объект структуры
	var filterJSON storage.Filter

	//Ловим json из post запроса
	c.ShouldBindJSON(&filterJSON)

	//Генерируем sql запрос и записываем в переменную
	resultSQL := DBFilterHandler.FilmFilterQueryGeneration(filterJSON)
	fmt.Println(resultSQL)

	//Получаем список отфильтрованных фильмов
	films := DBFilterHandler.SelectFilter(resultSQL)
	fmt.Println(films)

	c.JSON(200, films)
}

func getAllActors(c *gin.Context) {
	actors := DBActorHandler.SelectAll()
	c.JSON(200, gin.H{"actors": actors})
}

func getAllAuthors(c *gin.Context) {
	authors := DBAuthorHandler.SelectAll()
	c.JSON(200, gin.H{"authors": authors})
}

func getAllGenres(c *gin.Context) {
	genres := DBGenreHandler.SelectAll()
	c.JSON(200, gin.H{"genres": genres})
}

func getAllFilmYears(c *gin.Context) {
	filmYears := DBFilmHandler.SelectAllYears()
	c.JSON(200, gin.H{"film_years": filmYears})
}
