package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/vmlellis/go-hexagonal/application"
	"github.com/vmlellis/go-hexagonal/application/adapters/db"
)

func main() {
	dbInstance, _ := sql.Open("sqlite3", "sqlite.db")
	productDbAdapter := db.NewProductDb(dbInstance)
	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Product Exemplo", 30)

	productService.Enable(product)
}
