package server

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yeahyeahcore/KinoLab-Api/storage"
)

func removeFiles(id int) {
	filmName := storage.Fetch(fmt.Sprintf("select film_name from film where id = %d", id))
	err := os.RemoveAll(fmt.Sprintf("film-store/%s", filmName))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func deleteAuthor(c *gin.Context) {
	//Ловит форму и зависывает данные
	id, _ := strconv.Atoi(c.PostForm("id"))
	storage.Exec(fmt.Sprintf(`delete from author where id = %d`, id))
}

func deleteActor(c *gin.Context) {
	//Ловит форму и зависывает данные
	id, _ := strconv.Atoi(c.PostForm("id"))
	storage.Exec(fmt.Sprintf(`delete from actor where id = %d`, id))
}

func deleteFilm(c *gin.Context) {
	//Ловит форму и зависывает данные
	id, _ := strconv.Atoi(c.PostForm("id"))
	sql := fmt.Sprintf(`delete from Film where id = %d`, id)
	removeFiles(id)
	storage.Exec(sql)
}

func deleteGenre(c *gin.Context) {
	//Ловит форму и зависывает данные
	id, _ := strconv.Atoi(c.PostForm("id"))
	storage.Exec(fmt.Sprintf(`delete from genre where id = %d`, id))
}

func deleteFilmGenre(c *gin.Context) {
	//Ловит форму и зависывает данные
	id, _ := strconv.Atoi(c.PostForm("id"))
	storage.Exec(fmt.Sprintf(`delete from Film_Genre where id = %d`, id))
}

func deleteFilmAuthor(c *gin.Context) {
	//Ловит форму и зависывает данные
	id, _ := strconv.Atoi(c.PostForm("id"))
	storage.Exec(fmt.Sprintf(`delete from Film_Author where id = %d`, id))
}

func deleteFilmActor(c *gin.Context) {
	//Ловит форму и зависывает данные
	id, _ := strconv.Atoi(c.PostForm("id"))
	storage.Exec(fmt.Sprintf(`delete from Film_Actor where id = %d`, id))
}

func deleteFilmComment(c *gin.Context) {
	//Ловит форму и зависывает данные
	id, _ := strconv.Atoi(c.PostForm("id"))
	storage.Exec(fmt.Sprintf(`delete from Film_Comments where id = %d`, id))
}
