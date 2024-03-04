package postges_connect

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "postgres"
)

var Postgres *pgx.Conn

func GetDBConnection() error {
	// Tạo chuỗi kết nối đến PostgreSQL
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Tạo cấu hình kết nối
	config, err := pgx.ParseConfig(connStr)
	if err != nil {
		return err
	}
	// Mở kết nối đến PostgreSQL
	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		return err
	}
	// Kiểm tra kết nối
	err = conn.Ping(context.Background())
	if err != nil {
		return err
	}
	Postgres = conn
	return nil
}
