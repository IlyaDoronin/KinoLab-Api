package storage

import (
	"fmt"
)

//Actor структура описывающая таблицу actor в БД
type Actor struct {
	ID    int
	FName string
	LName string
}

//Select - метод для выборки данных из бд
func (a *Actor) Select(id int) Actor {

	actor := Actor{}

	sql := fmt.Sprintf("select * from actor where id = %d", id)

	err := conn.QueryRow(sql).Scan(&actor.ID, &actor.FName, &actor.LName)
	if err != nil {
		fmt.Println(err)
	}

	return actor
}

//SelectRange - метод для выборки данных из бд
func (p *Actor) SelectRange(pageNumber int) []Actor {

	actors := []Actor{}

	fromID, toID := GetIDBorders(pageNumber)

	sql := fmt.Sprintf(`
		select row_number() over() as num, ac.id, ac.lname, ac.fname
		from actor ac
		order by num asc limit %d offset %d
	`, toID, fromID)

	rows, err := conn.Query(sql)
	if err != nil {
		fmt.Println(err)
	}

	//Проход по всем элементом результата запроса и запись результата в объедок структуры
	for rows.Next() {

		actor := Actor{}

		err = rows.Scan(nil, &actor.ID, &actor.FName, &actor.LName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		actors = append(actors, actor)
	}

	//Выводим все данные для проверки
	for _, v := range actors {
		fmt.Println("-------------------Данные для проверки ёпт-----------------------")
		fmt.Println(v.ID)
		fmt.Println(v.FName)
		fmt.Println("-------------------Проверка окончена-----------------------")
	}

	// rows.Close вызывается rows.Next когда все строки прочитаны
	// или если произошла ошибка в методе Next или Scan.
	defer rows.Close()

	return actors
}

//SelectAll - метод для выборки данных из бд
func (p *Actor) SelectAll() []Actor {

	actors := []Actor{}

	rows, err := conn.Query("select * from actor")
	if err != nil {
		fmt.Println(err)
	}

	//Проход по всем элементом результата запроса и запись результата в объедок структуры
	for rows.Next() {

		actor := Actor{}

		err = rows.Scan(&actor.ID, &actor.FName, &actor.LName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		actors = append(actors, actor)
	}

	// rows.Close вызывается rows.Next когда все строки прочитаны
	// или если произошла ошибка в методе Next или Scan.
	defer rows.Close()

	return actors
}
