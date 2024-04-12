package internal

import (
	"context"
	"fmt"
	"log"

	"github.com/savita-kumari13/grpc-example/protogen/golang/orders"
)

type OrderService struct {
	db *DB
	orders.UnimplementedOrdersServer
}

func NewOrderService(db *DB) OrderService {
	return OrderService{db: db}
}

func (o *OrderService) AddOrder(_ context.Context, req *orders.PayloadWithSingleOrder) (*orders.PayloadWithSingleOrder, error) {
	log.Printf("Received an add-order request")
	order, err := o.db.AddOrder(req.GetOrder())
	if err != nil {
		fmt.Println("error adding order : ", err)
		return nil, err
	}
	fmt.Println("order added - ", order)
	return &orders.PayloadWithSingleOrder{Order: order}, nil
}

func (o *OrderService) GetOrders(_ context.Context, req *orders.Empty) (*orders.GetAllOrderResponse, error) {
	log.Printf("Get all orders request")

	orderList, err := o.db.GetOrders()
	if err != nil {
		fmt.Println("error adding order : ", err)
		return nil, err
	}
	return &orders.GetAllOrderResponse{Order: orderList}, nil

}

func (o *OrderService) GetOrder(_ context.Context, req *orders.GetOrderByIdRequest) (*orders.PayloadWithSingleOrder, error) {
	log.Printf("Get order by ID request - %v", req.GetOrderId())

	order, err := o.db.GetOrderById(req.GetOrderId())
	if err != nil {
		fmt.Println("error adding order : ", err)
		return nil, err
	}
	fmt.Println("order FOUND - ", order)
	return &orders.PayloadWithSingleOrder{Order: order}, nil
}

func (o *OrderService) UpdateOrder(_ context.Context, req *orders.PayloadWithSingleOrder) (*orders.PayloadWithSingleOrder, error) {
	log.Printf("Update order request")
	order, err := o.db.UpdateOrder(req.Order)
	if err != nil {
		fmt.Println("error adding order : ", err)
		return nil, err
	}
	fmt.Println("order FOUND - ", order)
	return &orders.PayloadWithSingleOrder{Order: order}, nil
}
