package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"strconv"
	"sync"
	"time"
)

type User struct {
	Name string
	Age  int
}

func InsertManyUsers(ctx context.Context, conn *pgx.Conn, tableName string, users []User) error {
	// Tạo một transaction
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// Chuẩn bị câu lệnh INSERT
	stmtName := fmt.Sprintf("insert_statement_%s", tableName)
	_, err = tx.Prepare(ctx, stmtName, fmt.Sprintf("INSERT INTO %s (name, age) VALUES ($1, $2)", tableName))
	if err != nil {
		return err
	}
	// Chèn từng dòng vào PostgreSQL
	for _, user := range users {
		_, err := conn.Exec(ctx, stmtName, user.Name, user.Age)
		if err != nil {
			return err
		}
	}

	// Commit transaction
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
func main() {

	for i := 0; i < 10; i++ {
		wg := sync.WaitGroup{}
		for j := 0; j < 80; j++ {
			go func() {
				// Kết nối đến PostgreSQL
				conn, err := pgx.Connect(context.Background(), "postgresql://postgres:mysecretpassword@localhost:5432/postgres")
				if err != nil {
					fmt.Println("Error connecting to PostgreSQL:", err)
					return
				}
				defer conn.Close(context.Background())

				// Tạo bảng user (nếu chưa tồn tại)
				_, err = conn.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS users (
			name TEXT,
			age INT
		)
	`)
				if err != nil {
					fmt.Println("Error creating table:", err)
					return
				}
				defer wg.Done()
				wg.Add(1)
				users := []User{}
				for i := 0; i < 1000000; i++ {
					users = append(users, User{
						Name: "User " + strconv.Itoa(i),
						Age:  i,
					})
				}
				err = InsertManyUsers(context.Background(), conn, "users", users)
				if err != nil {
					fmt.Println("Error inserting rows:", err)
					return
				}
				fmt.Println("Inserted", len(users), "rows into the users table.")
			}()
			time.Sleep(time.Second * 8)
		}
		wg.Wait()
	}
}
