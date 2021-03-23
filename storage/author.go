package storage

import (
	"fmt"
)

//Author структура описывающая таблицу author в БД
type Author struct {
	ID    int
	FName string
	LName string
}

//Select - метод для выборки данных из бд
func (a *Author) Select(id int) Author {

	author := Author{}

	sql := fmt.Sprintf("select * from author where id = %d", id)

	row, err := conn.Query(sql)
	if err != nil {
		fmt.Println(err)
	}

	for row.Next() {

		row.Scan(&author.ID, &author.FName, &author.LName)
		if err != nil {
			fmt.Println(err)
		}
	}

	return author
}

//SelectRange - метод для выборки данных из бд
func (p *Author) SelectRange(pageNumber int) []Author {

	authors := []Author{}

	fromID, toID := GetIDBorders(pageNumber)

	sql := fmt.Sprintf(`
		select row_number() over() as num, au.id, au.lname, au.fname
		from author au
		order by num asc limit %d offset %d
	`, toID, fromID)

	rows, err := conn.Query(sql)
	if err != nil {
		fmt.Println(err)
	}

	//Проход по всем элементом результата запроса и запись результата в объедок структуры
	for rows.Next() {

		author := Author{}

		err = rows.Scan(nil, &author.ID, &author.FName, &author.LName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		authors = append(authors, author)
	}

	//Выводим все данные для проверки
	for _, v := range authors {
		fmt.Println("-------------------Данные для проверки ёпт-----------------------")
		fmt.Println(v.ID)
		fmt.Println(v.FName)
		fmt.Println("-------------------Проверка окончена-----------------------")
	}

	return authors
}

//SelectAll - метод для выборки данных из бд
func (p *Author) SelectAll() []Author {

	authors := []Author{}

	rows, err := conn.Query("select * from author")
	if err != nil {
		fmt.Println(err)
	}

	//Проход по всем элементом результата запроса и запись результата в объедок структуры
	for rows.Next() {

		author := Author{}

		err = rows.Scan(&author.ID, &author.FName, &author.LName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		authors = append(authors, author)
	}

	return authors
}
