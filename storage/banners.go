package storage

import "fmt"

type Banner struct {
	FilmName  string
	BannerURL string
}

//Select - метод для выборки данных из бд
func (b *Banner) SelectRange() []Banner {

	banners := []Banner{}

	sql := fmt.Sprintf(`
		select row_number() over() as num, Banner_URL, Film_name from film order by num asc limit 5 offset 0
	`)

	rows, err := conn.Query(sql)
	if err != nil {
		fmt.Println(err)
	}

	//Проход по всем элементом результата запроса и запись результата в объедок структуры
	for rows.Next() {

		banner := Banner{}

		err = rows.Scan(nil, &banner.BannerURL, &banner.FilmName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		banners = append(banners, banner)
	}

	// rows.Close вызывается rows.Next когда все строки прочитаны
	// или если произошла ошибка в методе Next или Scan.
	defer rows.Close()

	return banners
}
