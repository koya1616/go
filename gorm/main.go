package main

import (
	"context"
	"fmt"
	"net/http"

	"database/sql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"koya/dao/model"
	"koya/dao/query"
)

func main() {
	sqlDB, err := sql.Open("pgx", "postgresql://postgres:password@db:5432/db")
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	// gen
	q := query.Use(gormDB)
	tableNames, err := q.TableName.WithContext(context.Background()).Where(q.TableName.ID.Eq(2)).First()

	// 通常Gorm
	var tableName model.TableName
	gormDB.First(&tableName)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("User: %v\n", tableNames)
		fmt.Printf("User: %v\n", tableName)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	fmt.Println("Server starting on localhost:8008")
	if err := http.ListenAndServe(":8008", nil); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
