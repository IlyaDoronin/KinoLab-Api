package storage

import (
	"context"
	"fmt"
)

//Genre структура описывающая таблицу genre в БД
type Genre struct {
	ID        int
	GenreName string
}

//Select - метод для выборки данных из бд
func (a *Genre) Select(id int) Genre {

	genre := Genre{}

	sql := fmt.Sprintf("select * from genre where id = %d", id)

	row, err := conn.Query(context.Background(), sql)
	if err != nil {
		fmt.Println(err)
	}

	for row.Next() {
		row.Scan(&genre.ID, &genre.GenreName)
		if err != nil {
			fmt.Println(err)
		}
	}

	return genre
}

//SelectRange - метод для выборки данных из бд
func (p *Genre) SelectRange(pageNumber int) []Genre {

	genres := []Genre{}

	fromID, toID := GetIDBorders(pageNumber)

	sql := fmt.Sprintf(`
		select row_number() over() as num, g.id, g.genre_name 
		from genre g
		order by num desc limit %d offset %d
	`, toID, fromID)

	rows, err := conn.Query(context.Background(), sql)
	if err != nil {
		fmt.Println(err)
	}

	//Проход по всем элементом результата запроса и запись результата в объедок структуры
	for rows.Next() {

		genre := Genre{}

		err = rows.Scan(nil, &genre.ID, &genre.GenreName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		genres = append(genres, genre)
	}

	//Выводим все данные для проверки
	for _, v := range genres {
		fmt.Println("-------------------Данные для проверки ёпт-----------------------")
		fmt.Println(v.ID)
		fmt.Println(v.GenreName)
		fmt.Println("-------------------Проверка окончена-----------------------")
	}

	return genres
}

//SelectAll - метод для выборки данных из бд
func (p *Genre) SelectAll() []Genre {

	genres := []Genre{}

	rows, err := conn.Query(context.Background(), "select * from genre")
	if err != nil {
		fmt.Println(err)
	}

	//Проход по всем элементом результата запроса и запись результата в объедок структуры
	for rows.Next() {

		genre := Genre{}

		err = rows.Scan(&genre.ID, &genre.GenreName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		genres = append(genres, genre)
	}

	return genres
}
