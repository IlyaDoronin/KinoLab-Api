package storage

import (
	"context"
	"fmt"
)

//WebSiteFilm структура описывающая таблицу film в БД
type WebSiteFilm struct {
	ID          int
	FilmName    string
	Description string
	FilmYear    string
	Budget      int
	FileURL     string
	PosterURL   string
	BannerURL   string
	Authors     []string
	Actors      []string
	Genres      []string
}

type WebFiteFilms struct {
	ID        int
	FilmName  string
	PosterURL string
	Genres    []string
}

//Select - метод для выборки данных из бд
func (w *WebSiteFilm) Select(id int) WebSiteFilm {

	film := WebSiteFilm{}

	row, err := conn.Query(context.Background(),
		fmt.Sprintf(`
		select id, f.Film_name, f.Description, f.Film_year::date::varchar, f.Budget::bigint, f.File_URL, f.Poster_URL, f.Banner_URL, au.au_array, ac.ac_array, g.g_array
		from Film as f
		join (
			select f_au.film_id as id, array_agg(au.lname || ' '  || au.fname) as au_array
			from film_author as f_au
			join author au on au.id = f_au.author_id 
			group by f_au.film_id
		) au using (id)
		join (
			select f_ac.film_id as id, array_agg(ac.lname || ' ' || ac.fname) as ac_array
			from film_actor as f_ac
			join actor ac on ac.id = f_ac.actor_id 
			group by f_ac.film_id
		) ac using (id)
		join (
			select f_g.film_id as id, array_agg(g.genre_name) as g_array
			from film_genre as f_g
			join genre g on g.id = f_g.genre_id 
			group by f_g.film_id
		) g using (id)
		where id = %d
		`, id))
	if err != nil {
		fmt.Println(err)
	}

	for row.Next() {
		row.Scan(
			&film.ID, &film.FilmName, &film.Description, &film.FilmYear, &film.Budget,
			&film.FileURL, &film.PosterURL, &film.BannerURL, &film.Authors, &film.Actors, &film.Genres,
		)
		if err != nil {
			fmt.Println(err)
		}
	}

	return film

}

//SelectRange - метод для выборки данных из бд
func (f *WebSiteFilm) SelectRange(pageNumber int) []WebFiteFilms {

	films := []WebFiteFilms{}

	fromID, toID := GetIDBorders(pageNumber)

	sql := fmt.Sprintf(`
		select row_number() over() as num, id, f.Film_name, f.Poster_URL, g.g_array
		from Film as f
		join (
			select f_g.film_id as id, array_agg(g.genre_name) as g_array
			from film_genre as f_g
			join genre g on g.id = f_g.genre_id 
			group by f_g.film_id
		) g using (id)
		order by num desc limit %d offset %d
	`, toID, fromID)

	rows, err := conn.Query(context.Background(), sql)
	if err != nil {
		fmt.Println(err)
	}

	//Проход по всем элементом результата запроса и запись результата в объедок структуры
	for rows.Next() {

		film := WebFiteFilms{}

		err = rows.Scan(nil, &film.ID, &film.FilmName, &film.PosterURL, &film.Genres)
		if err != nil {
			fmt.Println(err)
			continue
		}
		films = append(films, film)
	}

	return films
}
