package server

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yeahyeahcore/KinoLab-Api/conf"
)

var (
	Host string
	Port string
)

//Start Запускает роутинг
func Start() {

	r := gin.New()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.Static("/media", "film-store/")

	r.GET("/tables", getAllTables)
	r.GET("/pageCount", getPageCount)

	tableElementsGroup := r.Group("/get/table")
	{
		tableElementsGroup.GET("/authors", getAuthors)
		tableElementsGroup.GET("/actors", getActors)
		tableElementsGroup.GET("/films", getFilms)
		tableElementsGroup.GET("/genres", getGenres)
		tableElementsGroup.GET("/films_genres", getFilmsGenres)
		tableElementsGroup.GET("/films_authors", getFilmsAuthors)
		tableElementsGroup.GET("/films_actors", getFilmsActors)
		tableElementsGroup.GET("/film_comments", getFilmComments)

		tableElementsGroup.GET("/authors/all", getAllAuthors)
		tableElementsGroup.GET("/actors/all", getAllActors)
		tableElementsGroup.GET("/films/all", getAllFilms)
		tableElementsGroup.GET("/genres/all", getAllGenres)
		tableElementsGroup.GET("/films_genres/all", getAllFilmsGenres)
		tableElementsGroup.GET("/films_authors/all", getAllFilmsAuthors)
		tableElementsGroup.GET("/films_actors/all", getAllFilmsActors)
		tableElementsGroup.GET("/film_comments/all", getAllFilmComments)
	}

	editGroup := r.Group("/get/edit")
	{
		editGroup.GET("/author", getAuthorEditPage)
		editGroup.GET("/actor", getActorEditPage)
		editGroup.GET("/film", getFilmEditPage)
		editGroup.GET("/genre", getGenreEditPage)
		editGroup.GET("/film_genre", getFilmGenreEditPage)
		editGroup.GET("/film_author", getFilmAuthorEditPage)
		editGroup.GET("/film_actor", getFilmActorEditPage)
		editGroup.GET("/film_comment", getFilmCommentEditPage)
	}

	createGroup := r.Group("/post/create")
	{
		createGroup.POST("/author", createAuthor)
		createGroup.POST("/actor", createActor)
		createGroup.POST("/film", createFilm)
		createGroup.POST("/genre", createGenre)
		createGroup.POST("/film_genre", createFilmGenre)
		createGroup.POST("/film_author", createFilmAuthor)
		createGroup.POST("/film_actor", createFilmActor)
		createGroup.POST("/film_comment", createFilmComment)
	}

	updateGroup := r.Group("/post/update")
	{
		updateGroup.PATCH("/author", updateAuthor)
		updateGroup.PATCH("/actor", updateActor)
		updateGroup.PATCH("/film", updateFilm)
		updateGroup.PATCH("/genre", updateGenre)
		updateGroup.PATCH("/film_genre", updateFilmGenre)
		updateGroup.PATCH("/film_author", updateFilmAuthor)
		updateGroup.PATCH("/film_actor", updateFilmActor)
		updateGroup.PATCH("/film_comment", updateFilmComment)
	}

	deleteGroup := r.Group("/post/delete")
	{
		deleteGroup.DELETE("/author", deleteAuthor)
		deleteGroup.DELETE("/actor", deleteActor)
		deleteGroup.DELETE("/film", deleteFilm)
		deleteGroup.DELETE("/genre", deleteGenre)
		deleteGroup.DELETE("/film_genre", deleteFilmGenre)
		deleteGroup.DELETE("/film_author", deleteFilmAuthor)
		deleteGroup.DELETE("/film_actor", deleteFilmActor)
		deleteGroup.DELETE("/film_comment", deleteFilmComment)
	}

	websiteGroup := r.Group("/website")
	{
		websiteGroup.GET("/film", getWebSiteFilm)
		websiteGroup.GET("/banners", getBanners)
		websiteGroup.GET("/comments", getFilmCommentsForFilm)
		websiteGroup.GET("/films", getFilmsWeb)
		websiteGroup.GET("/actors", getAllActors)
		websiteGroup.GET("/authors", getAllAuthors)
		websiteGroup.GET("/genres", getAllGenres)
		websiteGroup.GET("/years", getAllFilmYears)
		websiteGroup.POST("/filter", getFilteredFilm)
		websiteGroup.POST("/postComment", createFilmComment)
	}

	Host = conf.Server.Host
	Port = conf.Server.Port

	//Server run
	r.Run(fmt.Sprintf("%s:%s", Host, Port)) // listen and serve
}
