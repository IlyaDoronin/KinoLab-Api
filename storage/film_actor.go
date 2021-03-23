package storage

import (
	"fmt"
)

//FilmActor структура описывающая таблицу film_actor в БД
type FilmActor struct {
	ID        int
	FilmID    int
	ActorID   int
	FilmName  string
	ActorName string
}

//Select - метод для выборки данных из бд
func (fa *FilmActor) Select(id int) FilmActor {

	film_actor := FilmActor{}

	sql := fmt.Sprintf(`
		select f_ac.id, f_ac.film_id, f_ac.actor_id, f.film_name, (ac.lname || ' ' || ac.fname) as actor_name
		from film_actor f_ac join film f on f.id = f_ac.film_id join actor ac on ac.id = f_ac.actor_id
		where f_ac.id = %d
	`, id)

	row, err := conn.Query(sql)
	if err != nil {
		fmt.Println(err)
	}

	for row.Next() {
		row.Scan(&film_actor.ID, &film_actor.FilmID, &film_actor.ActorID, &film_actor.FilmName, &film_actor.ActorName)
		if err != nil {
			fmt.Println(err)
		}
	}

	return film_actor
}

//SelectRange - метод для выборки данных из бд
func (fa *FilmActor) SelectRange(pageNumber int) []FilmActor {

	films_actors := []FilmActor{}

	fromID, toID := GetIDBorders(pageNumber)

	sql := fmt.Sprintf(`
		select row_number() over() as num, f_ac.id, f_ac.film_id, f_ac.actor_id, f.film_name, (ac.lname || ' ' || ac.fname) as actor_name
		from film_actor f_ac join film f on f.id = f_ac.film_id join actor ac on ac.id = f_ac.actor_id
		order by num asc limit %d offset %d
	`, toID, fromID)

	rows, err := conn.Query(sql)
	if err != nil {
		fmt.Println(err)
	}

	//Проход по всем элементом результата запроса и запись результата в объедок структуры
	for rows.Next() {

		film_actor := FilmActor{}

		err = rows.Scan(nil, &film_actor.ID, &film_actor.FilmID, &film_actor.ActorID, &film_actor.FilmName, &film_actor.ActorName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		films_actors = append(films_actors, film_actor)
	}

	//Выводим все данные для проверки
	for _, v := range films_actors {
		fmt.Println("-------------------Данные для проверки-----------------------")
		fmt.Println(v.FilmID)
		fmt.Println(v.ActorID)
		fmt.Println("-------------------Проверка окончена-----------------------")
	}

	return films_actors
}

//SelectAll - метод для выборки данных из бд
func (fa *FilmActor) SelectAll() []FilmActor {

	films_actors := []FilmActor{}

	rows, err := conn.Query(`
		select row_number() over() as num, f_ac.id, f_ac.film_id, f_ac.actor_id, f.film_name, (ac.lname || ' ' || ac.fname) as actor_name
		from film_actor f_ac join film f on f.id = f_ac.film_id join actor ac on ac.id = f_ac.actor_id
		order by num asc
	`)
	if err != nil {
		fmt.Println(err)
	}

	//Проход по всем элементом результата запроса и запись результата в объедок структуры
	for rows.Next() {

		film_actor := FilmActor{}

		err = rows.Scan(nil, &film_actor.ID, &film_actor.FilmID, &film_actor.ActorID, &film_actor.FilmName, &film_actor.ActorName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		films_actors = append(films_actors, film_actor)
	}

	return films_actors
}
