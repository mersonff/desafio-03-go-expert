package repository

import (
	"database/sql"
	"github.com/mersonff/20-CleanArch/internal/domain"
)

type MySQLOrderRepository struct {
	db *sql.DB
}

func NewMySQLOrderRepository(db *sql.DB) domain.OrderRepository {
	return &MySQLOrderRepository{
		db: db,
	}
}

func (r *MySQLOrderRepository) Save(order *domain.Order) error {
	stmt, err := r.db.Prepare(`
		INSERT INTO orders (id, price, tax, final_price, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		order.ID,
		order.Price,
		order.Tax,
		order.FinalPrice,
		order.CreatedAt,
		order.UpdatedAt,
	)
	return err
}

func (r *MySQLOrderRepository) List() ([]*domain.Order, error) {
	rows, err := r.db.Query(`
		SELECT id, price, tax, final_price, created_at, updated_at
		FROM orders
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*domain.Order
	for rows.Next() {
		order := &domain.Order{}
		err := rows.Scan(
			&order.ID,
			&order.Price,
			&order.Tax,
			&order.FinalPrice,
			&order.CreatedAt,
			&order.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
} 