package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	// Thay đổi thông tin kết nối cho phù hợp với cấu hình MySQL của bạn
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/mydatabase")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Kiểm tra kết nối
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Kết nối thành công đến MySQL!")
	// Thực hiện truy vấn để lấy dữ liệu từ bảng "user"
	//inster
	//start1 := time.Now()
	////_, _ = db.ExecContext(context.Background(), "insert users set id = ?,name=?", 2, "data ")
	//for i := 0; i < 10000; i++ {
	//	_, _ = db.ExecContext(context.Background(), "insert users set id = ?,name=?", i, "data "+strconv.Itoa(i))
	//	fmt.Println(i)
	//	time.Sleep(time.Second / 10000)
	//}
	//fmt.Println(time.Since(start1))
	start1 := time.Now()
	for i := 0; i < 10000; i++ {
		var id int64
		var name string
		_ = db.QueryRowContext(context.Background(), `select id, name
		from users
		where id = ?`, i).Scan(&id, &name)
		fmt.Println(id, name)
		time.Sleep(time.Second / 10000)
	}
	fmt.Println(time.Since(start1))
}
