package server

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yeahyeahcore/KinoLab-Api/storage"
)

func createAuthor(c *gin.Context) {
	//Ловит форму и зависывает данные
	lname := c.PostForm("lname")
	fname := c.PostForm("fname")
	fmt.Println(lname)
	fmt.Println(fname)

	storage.Exec(fmt.Sprintf(`insert into author(fname, lname) values('%s', '%s')`, fname, lname))
}

func createActor(c *gin.Context) {
	//Ловит форму и зависывает данные
	lname := c.PostForm("lname")
	fname := c.PostForm("fname")
	fmt.Println(lname)
	fmt.Println(fname)

	storage.Exec(fmt.Sprintf(`insert into actor(fname, lname) values('%s', '%s')`, fname, lname))
}

func createFilm(c *gin.Context) {
	//Ловит форму и зависывает данные
	filmName := c.PostForm("filmName")
	description := c.PostForm("description")
	filmYear := c.PostForm("filmYear")
	budget, _ := strconv.ParseFloat(c.PostForm("budget"), 32)
	fmt.Println(filmName)
	fmt.Println(description)
	fmt.Println(filmYear)
	fmt.Println(budget)

	//Ловит форму с файлом
	filmFile, err := c.FormFile("filmFile")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(filmFile.Filename)

	posterFile, err := c.FormFile("posterFile")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(posterFile.Filename)

	bannerFile, err := c.FormFile("bannerFile")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bannerFile.Filename)

	//Задаём файлам их местоположение
	filmFileDst := (fmt.Sprintf("film-store/%s/%s.mp4", filmName, filmName))
	fmt.Println(filmFileDst)

	posterFileDst := (fmt.Sprintf("film-store/%s/%s-poster.jpg", filmName, filmName))
	fmt.Println(posterFileDst)

	bannerFileDst := (fmt.Sprintf("film-store/%s/%s-banner.jpg", filmName, filmName))
	fmt.Println(bannerFileDst)

	//Загружает файлы в их местоположения
	err = c.SaveUploadedFile(filmFile, filmFileDst)
	if err != nil {
		log.Fatal(err)
	}

	err = c.SaveUploadedFile(posterFile, posterFileDst)
	if err != nil {
		log.Fatal(err)
	}

	err = c.SaveUploadedFile(bannerFile, bannerFileDst)
	if err != nil {
		log.Fatal(err)
	}

	//Создаёт переменную url к файлу в БД
	filmFileURL := (fmt.Sprintf("http://%s:%s/media/%s/%s.mp4", Host, Port, filmName, filmName))
	posterFileURL := (fmt.Sprintf("http://%s:%s/media/%s/%s-poster.jpg", Host, Port, filmName, filmName))
	bannerFileURL := (fmt.Sprintf("http://%s:%s/media/%s/%s-banner.jpg", Host, Port, filmName, filmName))
	fmt.Println(filmFileURL)
	fmt.Println(posterFileURL)
	fmt.Println(bannerFileURL)

	sql := fmt.Sprintf(`
		insert into Film(Film_name, Description, Film_year, Budget, File_URL, Poster_URL, Banner_URL)
		values('%s', '%s', '%s', %f, '%s', '%s', '%s')
	`, filmName, description, filmYear, budget, filmFileURL, posterFileURL, bannerFileURL)

	storage.Exec(sql)
}

func createGenre(c *gin.Context) {
	//Ловит форму и зависывает данные
	genreName := c.PostForm("genreName")
	fmt.Println(genreName)

	storage.Exec(fmt.Sprintf(`insert into Genre(Genre_name) values('%s')`, genreName))
}

func createFilmGenre(c *gin.Context) {
	//Ловит форму и зависывает данные
	filmID, _ := strconv.Atoi(c.PostForm("filmID"))
	genreID, _ := strconv.Atoi(c.PostForm("genreID"))
	fmt.Println(filmID)
	fmt.Println(genreID)

	storage.Exec(fmt.Sprintf(`insert into Film_Genre(Film_ID, Genre_ID) values(%d, %d)`, filmID, genreID))
}

func createFilmAuthor(c *gin.Context) {
	//Ловит форму и зависывает данные
	filmID, _ := strconv.Atoi(c.PostForm("filmID"))
	authorID, _ := strconv.Atoi(c.PostForm("authorID"))
	fmt.Println(filmID)
	fmt.Println(authorID)

	storage.Exec(fmt.Sprintf(`insert into Film_Author(Film_ID, Author_ID) values(%d, %d)`, filmID, authorID))
}

func createFilmActor(c *gin.Context) {
	//Ловит форму и зависывает данные
	filmID, _ := strconv.Atoi(c.PostForm("filmID"))
	actorID, _ := strconv.Atoi(c.PostForm("actorID"))
	fmt.Println(filmID)
	fmt.Println(actorID)

	storage.Exec(fmt.Sprintf(`insert into Film_Actor(Film_ID, Actor_ID) values(%d, %d)`, filmID, actorID))
}

func createFilmComment(c *gin.Context) {
	//Ловит форму и зависывает данные
	filmID, _ := strconv.Atoi(c.PostForm("filmID"))
	name := c.PostForm("name")
	text := c.PostForm("text")
	fmt.Println(filmID)
	fmt.Println(name)
	fmt.Println(text)

	storage.Exec(fmt.Sprintf(`
		insert into Film_Comments(Film_ID, commentator_name, comment_text)
		values (%d, '%s', '%s')
	`, filmID, name, text))
}
