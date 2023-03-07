package database

import (
	"database/sql"

	"github.com/felipedias-dev/gointensivo2/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		Db: db,
	}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	_, error := r.Db.Exec("INSERT INTO orders (id, price, tax, final_price) VALUES(?, ?, ?, ?)",
		order.ID, order.Price, order.Tax, order.FinalPrice)
	if error != nil {
		return error
	}
	return nil
}

func (r *OrderRepository) GetTotal(order *entity.Order) (int, error) {
	var total int
	error := r.Db.QueryRow("SELECT COUNT(*) FROM orders").Scan(&total)
	if error != nil {
		return 0, error
	}
	return total, nil
}
