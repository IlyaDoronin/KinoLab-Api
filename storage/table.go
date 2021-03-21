package storage

import (
	"fmt"
	"log"
)

//Exec выполняет запрос в БД (В основном DELETE или INSERT)
func Exec(sql string) {

	rows, err := conn.Exec(sql)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(rows.RowsAffected())
}

//Fetch возвращает одну строку результата запроса БД
func Fetch(sql string) string {
	var result string
	err := conn.QueryRow(sql).Scan(&result)
	if err != nil {
		fmt.Println(err)
	}
	return result
}

func FetchAll(sql string) []string {

	var rows []string

	rows_result, err := conn.Query(sql)
	if err != nil {
		fmt.Println(err)
	}

	for rows_result.Next() {

		var row string

		err = rows_result.Scan(&row)
		if err != nil {
			fmt.Println(err)
			continue
		}
		rows = append(rows, row)
	}

	defer rows_result.Close()

	return rows
}

//GetIDBorders возвращает границы ID'шников постов
func GetIDBorders(pageNumber int) (left, right int) {
	right = pageNumber * 10
	left = right - 10
	return
}

//GetPageCount Получает количество страниц
func GetPageCount(table string) int {
	var rowCount int
	err := conn.QueryRow(fmt.Sprintf("select count(id) from %s", table)).Scan(&rowCount)
	if err != nil {
		fmt.Println(err)
	}
	pageCount := rowCount / 10
	if rowCount%10 != 0 {
		pageCount++
	}
	return pageCount
}
