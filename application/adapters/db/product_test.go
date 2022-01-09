package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vmlellis/go-hexagonal/application/adapters/db"
	"github.com/vmlellis/go-hexagonal/application/entity"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
		"id" string,
		"name" string,
		"price" float,
		"status" string
	);`

	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products VALUES("abc", "Product Test", 0, "disabled")`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestProductDB_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("abc")
	require.Nil(t, err)
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, float64(0), product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDB_Save(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)
	product := entity.NewProduct("Product Test", float64(25))

	productResult, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, "Product Test", productResult.GetName())
	require.Equal(t, float64(25), productResult.GetPrice())
	require.Equal(t, "disabled", productResult.GetStatus())

	product.Enable()

	productResult, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, "Product Test", productResult.GetName())
	require.Equal(t, float64(25), productResult.GetPrice())
	require.Equal(t, "enabled", productResult.GetStatus())
}
