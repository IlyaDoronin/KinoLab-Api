package storage

import (
	"context"
	"fmt"
	"log"
)

//FilmAuthor структура описывающая таблицу film_author в БД
type FilmGenre struct {
	ID        int
	FilmID    int
	GenreID   int
	FilmName  string
	GenreName string
}

//Select - метод для выборки данных из бд
func (fg *FilmGenre) Select(id int) FilmGenre {

	film_genre := FilmGenre{}

	sql := fmt.Sprintf(`
		select f_g.id, f_g.film_id, f_g.genre_id, f.film_name, g.genre_name
		from film_genre f_g join film f on f.id = f_g.film_id join genre g on g.id = f_g.genre_id
		where f_g.id = %d
	`, id)

	row, err := conn.Query(context.Background(), sql)
	if err != nil {
		fmt.Println(err)
	}

	for row.Next() {
		row.Scan(&film_genre.ID, &film_genre.FilmID, &film_genre.GenreID, &film_genre.FilmName, &film_genre.GenreName)
		if err != nil {
			log.Fatal(err)
		}
	}

	return film_genre
}

//SelectRange - метод для выборки данных из бд
func (fg *FilmGenre) SelectRange(pageNumber int) []FilmGenre {

	films_genres := []FilmGenre{}

	fromID, toID := GetIDBorders(pageNumber)

	sql := fmt.Sprintf(`
		select row_number() over() as num, f_g.id, f_g.film_id, f_g.genre_id, f.film_name, g.genre_name
		from film_genre f_g join film f on f.id = f_g.film_id join genre g on g.id = f_g.genre_id
		order by num asc limit %d offset %d
	`, toID, fromID)

	rows, err := conn.Query(context.Background(), sql)
	if err != nil {
		fmt.Println(err)
	}

	//Проход по всем элементом результата запроса и запись результата в объедок структуры
	for rows.Next() {

		film_genre := FilmGenre{}

		err = rows.Scan(nil, &film_genre.ID, &film_genre.FilmID, &film_genre.GenreID, &film_genre.FilmName, &film_genre.GenreName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		films_genres = append(films_genres, film_genre)
	}

	//Выводим все данные для проверки
	for _, v := range films_genres {
		fmt.Println("-------------------Данные для проверки-----------------------")
		fmt.Println(v.FilmID)
		fmt.Println(v.GenreID)
		fmt.Println(v.FilmName)
		fmt.Println(v.GenreName)
		fmt.Println("-------------------Проверка окончена-----------------------")
	}

	return films_genres
}

//SelectAll - метод для выборки данных из бд
func (fg *FilmGenre) SelectAll() []FilmGenre {

	films_genres := []FilmGenre{}

	rows, err := conn.Query(context.Background(), `
		select row_number() over() as num, f_g.id, f_g.film_id, f_g.genre_id, f.film_name, g.genre_name
		from film_genre f_g join film f on f.id = f_g.film_id join genre g on g.id = f_g.genre_id
		order by num asc 
	`)
	if err != nil {
		fmt.Println(err)
	}

	//Проход по всем элементом результата запроса и запись результата в объедок структуры
	for rows.Next() {

		film_genre := FilmGenre{}

		err = rows.Scan(nil, &film_genre.ID, &film_genre.FilmID, &film_genre.GenreID, &film_genre.FilmName, &film_genre.GenreName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		films_genres = append(films_genres, film_genre)
	}

	return films_genres
}
