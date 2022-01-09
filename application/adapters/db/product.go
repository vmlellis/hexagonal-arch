package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/vmlellis/go-hexagonal/application/contract"
	"github.com/vmlellis/go-hexagonal/application/entity"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (contract.ProductInterface, error) {
	var product entity.Product
	stmt, err := p.db.Prepare("select id, name, price, status from products where id=?")
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductDb) Save(product contract.ProductInterface) (contract.ProductInterface, error) {
	var rows int
	p.db.QueryRow("SELECT COUNT(1) FROM products where id = ?", product.GetId()).Scan(&rows)
	var err error
	if rows == 0 {
		_, err = p.create(product)
	} else {
		_, err = p.update(product)
	}
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductDb) create(product contract.ProductInterface) (contract.ProductInterface, error) {
	stmt, err := p.db.Prepare(`INSERT INTO products(id, name, price, status) values(?,?,?,?)`)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(
		product.GetId(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
	)
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductDb) update(product contract.ProductInterface) (contract.ProductInterface, error) {
	_, err := p.db.Exec(
		`UPDATE products SET name=?, price=?, status=? WHERE id = ?`,
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
		product.GetId(),
	)
	if err != nil {
		return nil, err
	}
	return product, nil
}
