package internal

import (
	"fmt"

	"github.com/savita-kumari13/grpc-example/protogen/golang/orders"
)

type DB struct {
	collection []*orders.Order
}

func NewDB() *DB {
	return &DB{
		collection: make([]*orders.Order, 0),
	}
}

func (d *DB) AddOrder(order *orders.Order) (*orders.Order, error) {
	for _, o := range d.collection {
		if o.OrderId == order.OrderId {
			return nil, fmt.Errorf("duplicate order id: %d", order.GetOrderId())

		}
	}
	d.collection = append(d.collection, order)
	return order, nil
}

func (d *DB) GetOrders() ([]*orders.Order, error) {
	fmt.Println("d.collection = ", d.collection)
	return d.collection, nil
}

func (d *DB) GetOrderById(orderId int64) (*orders.Order, error) {
	for _, o := range d.collection {
		if o.OrderId == orderId {
			return o, nil
		}
	}
	return nil, fmt.Errorf("duplicate order id: %d", orderId)
}
func (d *DB) UpdateOrder(order *orders.Order) (*orders.Order, error) {
	for idx, o := range d.collection {
		if o.OrderId == order.OrderId {
			orders := append(d.collection[:idx], d.collection[idx+1:]...)
			d.collection = append(orders, order)
			return order, nil
		}
	}
	return nil, fmt.Errorf("order not found with id: %d", order.OrderId)
}
