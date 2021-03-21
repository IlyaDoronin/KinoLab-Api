package server

import (
	"github.com/gin-gonic/gin"
	"github.com/yeahyeahcore/KinoLab-Api/storage"
)

var (
	DBAuthorHandler      storage.Author
	DBActorHandler       storage.Actor
	DBGenreHandler       storage.Genre
	DBFilmHandler        storage.Film
	DBFilmAuthorHandler  storage.FilmAuthor
	DBFilmActorHandler   storage.FilmActor
	DBFilmGenreHandler   storage.FilmGenre
	DBFilmCommentHandler storage.FilmComments
	DBWebSiteFilmHandler storage.WebSiteFilm
	DBBannersHandler     storage.Banner
	DBFilterHandler      storage.Filter
)

func getAllTables(c *gin.Context) {
	sql := "SELECT table_name FROM information_schema.tables  where table_schema='public' ORDER BY table_name;"
	tables := storage.FetchAll(sql)
	c.JSON(200, gin.H{"tables": tables})
}

func getPageCount(c *gin.Context) {
	tableName := c.Query("table")
	pageCount := storage.GetPageCount(tableName)
	c.JSON(200, gin.H{"page_count": pageCount})
}
