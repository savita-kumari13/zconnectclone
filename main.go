package main

import (
	"fmt"
	"log"

	"github.com/savita-kumari13/grpc-example/protogen/golang/orders"
	"github.com/savita-kumari13/grpc-example/protogen/golang/product"
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	fmt.Println("HELLO from main")

	orderItem := orders.Order{
		OrderId:    10,
		CustomerId: 11,
		IsActive:   true,
		OrderDate:  &date.Date{Year: 2024, Month: 1, Day: 1},
		Products: []*product.Product{
			{ProductId: 1, ProductName: "CocaCola", ProductType: product.ProductType_DRINK},
		},
	}

	bytes, err := protojson.Marshal(&orderItem)
	if err != nil {
		log.Fatal("deserialization error:", err)
	}
	fmt.Println(string(bytes))
}
