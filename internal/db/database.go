// internal/db/database.go
package db

import (
    "database/sql"
    "fmt"
    "os"

    _ "github.com/go-sql-driver/mysql" // Dấu gạch dưới để import driver mà không cần gọi trực tiếp
)

// ConnectDatabase khởi tạo và trả về một kết nối tới MySQL.
func ConnectDatabase() (*sql.DB, error) {
    // Tạo chuỗi kết nối (DSN - Data Source Name) từ biến môi trường
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"),
    )

    // Mở kết nối
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err // Trả về lỗi nếu không thể mở kết nối
    }

    // Kiểm tra xem kết nối có thực sự thành công không
    err = db.Ping()
    if err != nil {
        return nil, err // Trả về lỗi nếu ping thất bại
    }

    fmt.Println("Successfully connected to the database!")
    return db, nil
}