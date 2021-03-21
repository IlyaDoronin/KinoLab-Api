package storage

import (
	"fmt"
)

//Film структура описывающая таблицу film в БД
type Film struct {
	ID          int
	FilmName    string
	Description string
	FilmYear    string
	Budget      float32
	FileURL     string
	PosterURL   string
	BannerURL   string
}

type Year struct {
	ID       int
	FilmYear int
}

//Select - метод для выборки данных из бд
func (a *Film) Select(id int) Film {

	film := Film{}

	sql := fmt.Sprintf("select id, Film_name, Description, Film_year::date::varchar, Budget, File_URL, Poster_URL, Banner_URL from film where id = %d", id)

	err := conn.QueryRow(sql).Scan(&film.ID, &film.FilmName, &film.Description, &film.FilmYear, &film.Budget, &film.FileURL, &film.PosterURL, &film.BannerURL)
	if err != nil {
		fmt.Println(err)
	}

	return film
}

//SelectRange - метод для выборки данных из бд
func (f *Film) SelectRange(pageNumber int) []Film {

	films := []Film{}

	fromID, toID := GetIDBorders(pageNumber)

	sql := fmt.Sprintf(`
		select row_number() over() as num, f.id, f.film_name, 
		f.description, f.film_year::date::varchar, f.budget, File_URL, f.poster_url, f.banner_url
		from film f
		order by num asc limit %d offset %d
	`, toID, fromID)

	rows, err := conn.Query(sql)
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

	// rows.Close вызывается rows.Next когда все строки прочитаны
	// или если произошла ошибка в методе Next или Scan.
	defer rows.Close()

	return films
}

//SelectRange - метод для выборки данных из бд
func (f *Film) SelectAllYears() []Year {

	years := []Year{}

	rows, err := conn.Query("select row_number() over() as num, (extract(year from film_year)) as film_year from film")
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

	// rows.Close вызывается rows.Next когда все строки прочитаны
	// или если произошла ошибка в методе Next или Scan.
	defer rows.Close()

	return years
}

//SelectAll - метод для выборки данных из бд
func (f *Film) SelectAll() []Film {

	films := []Film{}

	rows, err := conn.Query(`
		select row_number() over() as num, f.id, f.film_name, 
		f.description, f.film_year::date::varchar, f.budget, File_URL, f.poster_url, f.banner_url
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

	// rows.Close вызывается rows.Next когда все строки прочитаны
	// или если произошла ошибка в методе Next или Scan.
	defer rows.Close()

	return films
}
