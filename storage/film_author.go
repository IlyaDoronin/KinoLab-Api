package storage

import (
	"context"
	"fmt"
)

//FilmAuthor структура описывающая таблицу film_author в БД
type FilmAuthor struct {
	ID         int
	FilmID     int
	AuthorID   int
	FilmName   string
	AuthorName string
}

//Select - метод для выборки данных из бд
func (fa *FilmAuthor) Select(id int) FilmAuthor {

	film_author := FilmAuthor{}

	sql := fmt.Sprintf(`
		select f_au.id, f_au.film_id, f_au.author_id, f.film_name, (au.lname || ' ' || au.fname) as author_name
		from film_author as f_au join film f on f.id = f_au.film_id join author au on au.id = f_au.author_id
		where f_au.id = %d
	`, id)

	row, err := conn.Query(context.Background(), sql)
	if err != nil {
		fmt.Println(err)
	}

	for row.Next() {
		row.Scan(&film_author.ID, &film_author.FilmID, &film_author.AuthorID, &film_author.FilmName, &film_author.AuthorName)
		if err != nil {
			fmt.Println(err)
		}
	}

	return film_author
}

//SelectRange - метод для выборки данных из бд
func (fa *FilmAuthor) SelectRange(pageNumber int) []FilmAuthor {

	films_authors := []FilmAuthor{}

	fromID, toID := GetIDBorders(pageNumber)

	sql := fmt.Sprintf(`
		select row_number() over() as num, f_au.id, f_au.film_id, f_au.author_id, f.film_name, (au.lname || ' ' || au.fname) as author_name
		from film_author as f_au join film f on f.id = f_au.film_id join author au on au.id = f_au.author_id
		order by num asc limit %d offset %d
	`, toID, fromID)

	rows, err := conn.Query(context.Background(), sql)
	if err != nil {
		fmt.Println(err)
	}

	//Проход по всем элементом результата запроса и запись результата в объедок структуры
	for rows.Next() {

		film_author := FilmAuthor{}

		err = rows.Scan(nil, &film_author.ID, &film_author.FilmID, &film_author.AuthorID, &film_author.FilmName, &film_author.AuthorName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		films_authors = append(films_authors, film_author)
	}

	//Выводим все данные для проверки
	for _, v := range films_authors {
		fmt.Println("-------------------Данные для проверки ёпт-----------------------")
		fmt.Println(v.FilmID)
		fmt.Println(v.AuthorID)
		fmt.Println("-------------------Проверка окончена-----------------------")
	}

	return films_authors
}

//SelectAll - метод для выборки данных из бд
func (fa *FilmAuthor) SelectAll() []FilmAuthor {

	films_authors := []FilmAuthor{}

	rows, err := conn.Query(context.Background(), `
		select row_number() over() as num, f_au.id, f_au.film_id, f_au.author_id, f.film_name, (au.lname || ' ' || au.fname) as author_name
		from film_author as f_au join film f on f.id = f_au.film_id join author au on au.id = f_au.author_id
		order by num asc
	`)
	if err != nil {
		fmt.Println(err)
	}

	//Проход по всем элементом результата запроса и запись результата в объедок структуры
	for rows.Next() {

		film_author := FilmAuthor{}

		err = rows.Scan(nil, &film_author.ID, &film_author.FilmID, &film_author.AuthorID, &film_author.FilmName, &film_author.AuthorName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		films_authors = append(films_authors, film_author)
	}

	return films_authors
}
