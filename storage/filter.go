package storage

import (
	"fmt"
	"reflect"
)

//Filter структура для фильтрации информации
type Filter struct {
	Genres  []string `json:"genres,omitempty"`
	Actors  []string `json:"actors,omitempty"`
	Authors []string `json:"authors,omitempty"`
	Years   []string `json:"years,omitempty"`
	Search  string   `json:"search,omitempty"`
}

type FilterResult struct {
	ID        int
	FilmName  string
	PosterURL string
	Genres    []string
}

//QueryGeneration - фунцкия для динамической генерации запроса исходя из получаемых данных
func (f *Filter) FilmFilterQueryGeneration(filter Filter) string {

	sql := `
		select 
		id,
		f.Film_name,
		f.Poster_URL,
		g.g_array
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
	`

	var where string

	//Пробежка по всем полям структуры объедка filter
	val := reflect.ValueOf(filter)
	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).Len() != 0 {
			sql += " where "
			break
		}
	}

	if len(filter.Genres) != 0 {
		for _, v := range filter.Genres {
			if where != "" {
				where += " and "
			}
			where += fmt.Sprintf("g.g_array && '{%s}'", v)
		}
	}

	if len(filter.Actors) != 0 {
		for _, v := range filter.Actors {
			if where != "" {
				where += " and "
			}
			where += fmt.Sprintf("ac.ac_array && '{%s}'", v)
		}
	}

	if len(filter.Authors) != 0 {
		for _, v := range filter.Authors {
			if where != "" {
				where += " and "
			}
			where += fmt.Sprintf("au.au_array && '{%s}'", v)
		}
	}

	if len(filter.Years) != 0 {
		for _, v := range filter.Years {
			if where != "" {
				where += " and "
			}
			where += fmt.Sprintf("(extract(year from f.film_year)) = '%s'", v)
		}
	}

	if len(filter.Search) != 0 {
		if where != "" {
			where += " and "
		}
		percent := string([]byte("%"))
		where += fmt.Sprintf("f.film_name like '%s%s%s'", percent, filter.Search, percent)
	}

	sql += where

	return sql
}

//SelectFilter - метод для выборки данных фильтрованных данных из БД
func (w *Filter) SelectFilter(sql string) []FilterResult {

	films := []FilterResult{}

	rows, err := conn.Query(sql)
	if err != nil {
		fmt.Println(err)
	}

	//Проход по всем элементом результата запроса и запись результата в объедок структуры
	for rows.Next() {

		film := FilterResult{}

		err = rows.Scan(
			&film.ID, &film.FilmName, &film.PosterURL, &film.Genres,
		)
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
