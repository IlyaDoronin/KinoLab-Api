package storage

import (
	"context"
	"fmt"
)

//Film структура описывающая таблицу film в БД
type Film struct {
	ID          int
	FilmName    string
	Description string
	FilmYear    string
	Budget      int
	FileURL     string
	PosterURL   string
	BannerURL   string
	Genres      []string
}

type Year struct {
	ID       int
	FilmYear int
}

//Select - метод для выборки данных из бд
func (a *Film) Select(id int) Film {

	film := Film{}

	sql := fmt.Sprintf("select id, Film_name, Description, Film_year::date::varchar, Budget::int, File_URL, Poster_URL, Banner_URL from film where id = %d", id)

	row, err := conn.Query(context.Background(), sql)
	if err != nil {
		fmt.Println(err)
	}

	for row.Next() {
		row.Scan(&film.ID, &film.FilmName, &film.Description, &film.FilmYear, &film.Budget, &film.FileURL, &film.PosterURL, &film.BannerURL)
		if err != nil {
			fmt.Println(err)
		}
	}

	return film
}

//SelectRange - метод для выборки данных из бд
func (f *Film) SelectRange(pageNumber int) []Film {

	films := []Film{}

	fromID, toID := GetIDBorders(pageNumber)

	sql := fmt.Sprintf(`
		select row_number() over() as num, f.id, f.film_name, 
		f.description, f.film_year::date::varchar, f.budget::int, File_URL, f.poster_url, f.banner_url
		from film f
		order by num asc limit %d offset %d
	`, toID, fromID)

	rows, err := conn.Query(context.Background(), sql)
	if err != nil {
		fmt.Println(err)
	}

	//Проход по всем элементом результата запроса и запись результата в объедок структуры
	for rows.Next() {

		film := Film{}

		err = rows.Scan(nil, &film.ID, &film.FilmName, &film.Description, &film.FilmYear, &film.Budget, &film.FileURL, &film.PosterURL, &film.BannerURL)
		if err != nil {
			fmt.Println(err)
			continue
		}
		films = append(films, film)
	}

	//Выводим все данные для проверки
	for _, v := range films {
		fmt.Println("-------------------Данные для проверки фильмов-----------------------")
		fmt.Println(v.ID)
		fmt.Println(v.FilmName)
		fmt.Println(v.FilmYear)
		fmt.Println(v.BannerURL)
		fmt.Println("-------------------Проверка окончена-----------------------")
	}

	return films
}

//SelectRange - метод для выборки данных из бд
func (f *Film) SelectAllYears() []Year {

	years := []Year{}

	rows, err := conn.Query(context.Background(), "select row_number() over() as num, (extract(year from film_year)) as film_year from film")
	if err != nil {
		fmt.Println(err)
	}

	//Проход по всем элементом результата запроса и запись результата в объедок структуры
	for rows.Next() {

		year := Year{}

		err = rows.Scan(&year.ID, &year.FilmYear)
		if err != nil {
			fmt.Println(err)
			continue
		}
		years = append(years, year)
	}

	return years
}

//SelectAllWeb - метод для выборки данных из бд
func (f *Film) SelectAllWeb() []Film {

	films := []Film{}

	rows, err := conn.Query(context.Background(), `
		select row_number() over() as num, f.id, f.film_name, 
		f.description, f.film_year::date::varchar, f.budget::int, File_URL, f.poster_url, f.banner_url
		from film f
		order by num asc
	`)
	if err != nil {
		fmt.Println(err)
	}

	//Проход по всем элементом результата запроса и запись результата в объедок структуры
	for rows.Next() {

		film := Film{}

		err = rows.Scan(nil, &film.ID, &film.FilmName, &film.Description, &film.FilmYear, &film.Budget, &film.FileURL, &film.PosterURL, &film.BannerURL)
		if err != nil {
			fmt.Println(err)
			continue
		}
		films = append(films, film)
	}

	return films
}

//SelectAll - метод для выборки данных из бд
func (f *Film) SelectRangeWeb(pageNumber int) []Film {

	films := []Film{}

	fromID, toID := GetIDBorders(pageNumber)

	rows, err := conn.Query(context.Background(), fmt.Sprintf(`
		select
		row_number() over() as num,
		id,
		f.Film_name,
		f.Poster_URL,
		g.g_array
		from Film as f
		join (
			select f_g.film_id as id, array_agg(g.genre_name) as g_array
			from film_genre as f_g
			join genre g on g.id = f_g.genre_id 
			group by f_g.film_id
		) g using (id)
		order by num asc limit %d offset %d
	`, toID, fromID))
	if err != nil {
		fmt.Println(err)
	}

	//Проход по всем элементом результата запроса и запись результата в объедок структуры
	for rows.Next() {

		film := Film{}

		err = rows.Scan(nil, &film.ID, &film.FilmName, &film.PosterURL, &film.Genres)
		if err != nil {
			fmt.Println(err)
			continue
		}
		films = append(films, film)
	}

	return films
}
