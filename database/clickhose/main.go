package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/ClickHouse/clickhouse-go"
)

func main() {
	connect, err := sql.Open("clickhouse", "tcp://10.3.95.52:8123?username=default&password=fjklajfk2313qqqs&database=qryn")
	if err != nil {
		log.Fatal(err)
	}
	defer connect.Close()

	if err := connect.Ping(); err != nil {
		log.Fatal(err)
	}

	// Giả sử bảng `user` có các cột `id`, `name`, và `age`
	tx, err := connect.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("INSERT INTO user (id, name, age) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Thay thế `1`, `John Doe`, và `30` bằng dữ liệu thực tế của bạn
	if _, err := stmt.Exec(
		1,
		"John Doe",
		30,
	); err != nil {
		log.Fatal(err)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Dữ liệu đã được insert thành công vào bảng `user`.")
}
