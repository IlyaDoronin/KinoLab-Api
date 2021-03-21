package server

import (
	"fmt"
	"mime/multipart"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yeahyeahcore/KinoLab-Api/storage"
)

//Удаляет и загружает необходимые файлы для замены
func removeOldFiles(c *gin.Context, id int, newFilmName string, files []*multipart.FileHeader) {

	filmName := storage.Fetch(fmt.Sprintf("select film_name from film where id = %d", id))
	filmFileName := fmt.Sprintf("%s.mp4", newFilmName)
	posterFileName := fmt.Sprintf("%s-poster.jpg", newFilmName)
	bannerFileName := fmt.Sprintf("%s-banner.jpg", newFilmName)

	if files[0] != nil {
		_ = os.Remove(fmt.Sprintf("film-store/%s/%s.mp4", filmName, filmName))
		_ = c.SaveUploadedFile(files[0], filmFileName)
	}

	if files[1] != nil {
		_ = os.Remove(fmt.Sprintf("film-store/%s/%s-poster.jpg", filmName, filmName))
		_ = c.SaveUploadedFile(files[1], posterFileName)
	}

	if files[2] != nil {
		_ = os.Remove(fmt.Sprintf("film-store/%s/%s-banner.jpg", filmName, filmName))
		_ = c.SaveUploadedFile(files[2], bannerFileName)
	}

}

func updateAuthor(c *gin.Context) {
	//Ловит форму и зависывает данные
	id, _ := strconv.Atoi(c.PostForm("id"))
	lname := c.PostForm("lname")
	fname := c.PostForm("fname")
	fmt.Println(lname)
	fmt.Println(fname)

	storage.Exec(fmt.Sprintf(`update author set FName = '%s', LName = '%s' where id = %d`, fname, lname, id))
}

func updateActor(c *gin.Context) {
	//Ловит форму и зависывает данные
	id, _ := strconv.Atoi(c.PostForm("id"))
	lname := c.PostForm("lname")
	fname := c.PostForm("fname")
	fmt.Println(lname)
	fmt.Println(fname)

	storage.Exec(fmt.Sprintf(`update actor set FName = '%s', LName = '%s' where id = %d`, fname, lname, id))
}

func updateFilm(c *gin.Context) {
	//Ловит форму и зависывает данные
	id, _ := strconv.Atoi(c.PostForm("id"))
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

	//Создаём массив с пришедшими файлами
	files := []*multipart.FileHeader{filmFile, posterFile, bannerFile}

	//Если приходят новые файлы, то мы удаляем старые
	removeOldFiles(c, id, filmName, files)

	//Задаём файлам их местоположение
	filmFileDst := fmt.Sprintf("film-store/%s/%s.mp4", filmName, filmName)
	posterFileDst := fmt.Sprintf("film-store/%s/%s-poster.jpg", filmName, filmName)
	bannerFileDst := fmt.Sprintf("film-store/%s/%s-banner.jpg", filmName, filmName)
	fmt.Println(filmFileDst)
	fmt.Println(posterFileDst)
	fmt.Println(bannerFileDst)

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

func updateGenre(c *gin.Context) {
	//Ловит форму и зависывает данные
	id, _ := strconv.Atoi(c.PostForm("id"))
	ganreName := c.PostForm("ganreName")
	fmt.Println(ganreName)

	storage.Exec(fmt.Sprintf(`update genre set Genre_name = '%s' where id = %d`, ganreName, id))
}

func updateFilmGenre(c *gin.Context) {
	//Ловит форму и зависывает данные
	id, _ := strconv.Atoi(c.PostForm("id"))
	filmID, _ := strconv.Atoi(c.PostForm("filmID"))
	genreID, _ := strconv.Atoi(c.PostForm("genreID"))
	fmt.Println(filmID)
	fmt.Println(genreID)

	storage.Exec(fmt.Sprintf(`update Film_Genre set Film_ID = %d Genre_ID = %d where id = %d`, filmID, genreID, id))
}

func updateFilmAuthor(c *gin.Context) {
	//Ловит форму и зависывает данные
	id, _ := strconv.Atoi(c.PostForm("id"))
	filmID, _ := strconv.Atoi(c.PostForm("filmID"))
	authorID, _ := strconv.Atoi(c.PostForm("authorID"))
	fmt.Println(filmID)
	fmt.Println(authorID)

	storage.Exec(fmt.Sprintf(`update Film_Author set Film_ID = %d Author_ID = %d where id = %d`, filmID, authorID, id))
}

func updateFilmActor(c *gin.Context) {
	//Ловит форму и зависывает данные
	id, _ := strconv.Atoi(c.PostForm("id"))
	filmID, _ := strconv.Atoi(c.PostForm("filmID"))
	actorID, _ := strconv.Atoi(c.PostForm("actorID"))
	fmt.Println(filmID)
	fmt.Println(actorID)

	storage.Exec(fmt.Sprintf(`update Film_Actor set Film_ID = %d Actor_ID = %d where id = %d`, filmID, actorID, id))
}

func updateFilmComment(c *gin.Context) {
	//Ловит форму и зависывает данные
	id, _ := strconv.Atoi(c.PostForm("id"))
	filmID, _ := strconv.Atoi(c.PostForm("filmID"))
	name := c.PostForm("name")
	text := c.PostForm("text")
	fmt.Println(filmID)
	fmt.Println(name)
	fmt.Println(text)

	storage.Exec(fmt.Sprintf(`
		update Film_Comments set 
		film_id = %d commentator_name = %s comment_text = %s
		where id = %d
	`, filmID, name, text, id))
}
