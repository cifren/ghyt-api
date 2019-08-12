package main

import (
	"github.com/kataras/iris"
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
	//db.Create(&Product{Code: "L1212", Price: 1000})

	// Web Server
	app := iris.New()

	// Method:   GET
	// Resource: http://localhost:8080
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome on Ghyt API</h1>")
	})

	app.Get("/webhook", func(ctx iris.Context) {
		// Read
		var product model.Product
		db.First(&product, 1) // find product with id 1

		ctx.HTML("<h1>Webhook ok, code product " + product.Code + "</h1>" )
	})

	// http://localhost:8080
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
