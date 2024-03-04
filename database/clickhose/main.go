package main

import (
	"database/sql"
	"fmt"

	_ "github.com/ClickHouse/clickhouse-go"
)

func main() {
	// Thay đổi thông tin kết nối dưới đây thành thông tin của ClickHouse của bạn
	dsn := "tcp://10.3.54.220:8123?username=default&password=zxc123456zxc&database=default"

	// Kết nối đến ClickHouse
	db, err := sql.Open("clickhouse", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Kiểm tra kết nối
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Kết nối thành công đến ClickHouse!")

	// Thực hiện truy vấn
	rows, err := db.Query("SELECT * FROM <table>")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Đọc dữ liệu từ truy vấn
	for rows.Next() {
		var column1 string
		var column2 int

		err := rows.Scan(&column1, &column2)
		if err != nil {
			panic(err)
		}

		fmt.Println(column1, column2)
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}
}
