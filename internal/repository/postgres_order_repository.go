package repository

import (
	"database/sql"

	"github.com/fabricioviapiana/orders-app/internal/domain"
)

type postgresOrderRepository struct {
	db *sql.DB
}

func NewPostgresOrderRepository(db *sql.DB) *postgresOrderRepository {
	return &postgresOrderRepository{
		db: db,
	}
}

func (r *postgresOrderRepository) Create(userID string, items []domain.OrderItem, totalAmount float64) (domain.Order, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return domain.Order{}, err
	}
	defer tx.Rollback()

	var orderID string
	queryOrder := `INSERT INTO orders (user_id, total_amount) VALUES ($1, $2) RETURNING id`
	err = tx.QueryRow(queryOrder, userID, totalAmount).Scan(&orderID)
	if err != nil {
		return domain.Order{}, err
	}

	queryItem := `INSERT INTO order_items (order_id, product_id, quantity, unit_price) VALUES ($1, $2, $3, $4)`
	for _, item := range items {
		_, err = tx.Exec(queryItem, orderID, item.ProductID, item.Quantity, item.UnitPrice)
		if err != nil {
			return domain.Order{}, err
		}
	}

	if err := tx.Commit(); err != nil {
		return domain.Order{}, err
	}

	return domain.Order{
		ID:          orderID,
		UserID:      userID,
		Items:       items,
		TotalAmount: totalAmount,
	}, nil
}

func (r *postgresOrderRepository) List() ([]domain.Order, error) {
	query := `
		SELECT o.id, o.user_id, o.total_amount, oi.product_id, oi.quantity, oi.unit_price
		FROM orders o
		LEFT JOIN order_items oi ON o.id = oi.order_id
		ORDER BY o.created_at DESC
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ordersMap := make(map[string]*domain.Order)
	var orderIDs []string // Para manter a ordem original do banco

	for rows.Next() {
		var (
			orderID     string
			userID      string
			totalAmount float64
			productID   sql.NullString
			quantity    sql.NullInt64
			unitPrice   sql.NullFloat64
		)

		err := rows.Scan(&orderID, &userID, &totalAmount, &productID, &quantity, &unitPrice)
		if err != nil {
			return nil, err
		}

		if _, ok := ordersMap[orderID]; !ok {
			order := &domain.Order{
				ID:          orderID,
				UserID:      userID,
				TotalAmount: totalAmount,
				Items:       []domain.OrderItem{},
			}
			ordersMap[orderID] = order
			orderIDs = append(orderIDs, orderID)
		}

		if productID.Valid {
			ordersMap[orderID].Items = append(ordersMap[orderID].Items, domain.OrderItem{
				ProductID: productID.String,
				Quantity:  int(quantity.Int64),
				UnitPrice: unitPrice.Float64,
			})
		}
	}

	result := make([]domain.Order, 0, len(orderIDs))
	for _, id := range orderIDs {
		result = append(result, *ordersMap[id])
	}

	return result, nil
}

func (r *postgresOrderRepository) FindByID(id string) (domain.Order, error) {
	queryOrder := `SELECT id, user_id, total_amount FROM orders WHERE id = $1`
	var order domain.Order
	err := r.db.QueryRow(queryOrder, id).Scan(&order.ID, &order.UserID, &order.TotalAmount)
	if err != nil {
		return domain.Order{}, err
	}

	queryItems := `SELECT product_id, quantity, unit_price FROM order_items WHERE order_id = $1`
	rows, err := r.db.Query(queryItems, id)
	if err != nil {
		return domain.Order{}, err
	}
	defer rows.Close()

	var items []domain.OrderItem
	for rows.Next() {
		var item domain.OrderItem
		if err := rows.Scan(&item.ProductID, &item.Quantity, &item.UnitPrice); err != nil {
			return domain.Order{}, err
		}
		items = append(items, item)
	}
	order.Items = items

	return order, nil
}
