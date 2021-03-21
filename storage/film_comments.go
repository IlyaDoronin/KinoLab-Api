package storage

import (
	"fmt"
)

//FilmComments структура описывающая таблицу film_comments в БД
type FilmComments struct {
	ID        int
	FilmID    int
	Name      string
	Text      string
	CreatedAt string
}

//Select - метод для выборки данных из бд
func (f_c *FilmComments) Select(id int) FilmComments {

	film_comment := FilmComments{}

	sql := fmt.Sprintf(`
	select id commentator_name, comment_text, created_at::timestamp(0)::varchar from film_comments where film_comments.film_id = %d
	`, id)

	err := conn.QueryRow(sql).Scan(&film_comment.ID, &film_comment.FilmID, &film_comment.Name, &film_comment.Text, &film_comment.CreatedAt)
	if err != nil {
		fmt.Println(err)
	}

	return film_comment
}

//SelectRange - метод для выборки данных из бд
func (p *FilmComments) SelectRange(pageNumber int) []FilmComments {

	film_comments := []FilmComments{}

	fromID, toID := GetIDBorders(pageNumber)

	sql := fmt.Sprintf(`
		select row_number() over() as num, f_c.id , f_c.film_id, 
		f_c.commentator_name, f_c.comment_text, f_c.created_at::timestamp(0)::varchar
		from film_comments f_c
		order by num asc limit %d offset %d
	`, toID, fromID)

	rows, err := conn.Query(sql)
	if err != nil {
		fmt.Println(err)
	}

	//Проход по всем элементом результата запроса и запись результата в объедок структуры
	for rows.Next() {

		film_comment := FilmComments{}

		err = rows.Scan(nil, &film_comment.ID, &film_comment.FilmID, &film_comment.Name, &film_comment.Text, &film_comment.CreatedAt)
		if err != nil {
			fmt.Println(err)
			continue
		}
		film_comments = append(film_comments, film_comment)
	}

	//Выводим все данные для проверки
	for _, v := range film_comments {
		fmt.Println("-------------------Данные для проверки ёпт-----------------------")
		fmt.Println(v.Name)
		fmt.Println(v.Text)
		fmt.Println("-------------------Проверка окончена-----------------------")
	}

	// rows.Close вызывается rows.Next когда все строки прочитаны
	// или если произошла ошибка в методе Next или Scan.
	defer rows.Close()

	return film_comments
}

//SelectAll - метод для выборки данных из бд
func (fa *FilmComments) SelectAll() []FilmComments {

	film_comments := []FilmComments{}

	rows, err := conn.Query(`
		select row_number() over() as num, f_c.id , f_c.film_id, 
		f_c.commentator_name, f_c.comment_text, f_c.created_at::timestamp(0)::varchar
		from film_comments f_c
		order by num asc
	`)
	if err != nil {
		fmt.Println(err)
	}

	//Проход по всем элементом результата запроса и запись результата в объедок структуры
	for rows.Next() {

		film_comment := FilmComments{}

		err = rows.Scan(nil, &film_comment.ID, &film_comment.FilmID, &film_comment.Name, &film_comment.Text, &film_comment.CreatedAt)
		if err != nil {
			fmt.Println(err)
			continue
		}
		film_comments = append(film_comments, film_comment)
	}

	// rows.Close вызывается rows.Next когда все строки прочитаны
	// или если произошла ошибка в методе Next или Scan.
	defer rows.Close()

	return film_comments
}
