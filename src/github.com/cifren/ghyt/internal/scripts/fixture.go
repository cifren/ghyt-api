package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/cifren/ghyt/internal/model"
)

func main() {
	// DB
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Create
	db.Create(&model.Product{Code: "L1212", Price: 1000})
}
