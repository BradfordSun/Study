package main

import (
	"encoding/json"
	"fmt"
)

type OrderItem struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Order struct {
	ID         string       `json:"id"`
	Item       *OrderItem   `json:"item,omitempty"`
	Quantity   int          `json:"quantity"`
	TotalPrice float64      `json:"total_price"`
	Items      []*OrderItem `json:"items"`
}

func main() {
	marshal()
	unmarshal()
}

func marshal() {
	o := Order{
		ID: "1234",
		Item: &OrderItem{
			ID:    "item_1",
			Name:  "learn go",
			Price: 15,
		},
		Quantity:   3,
		TotalPrice: 30,
		Items: []*OrderItem{
			{
				ID:    "item_3",
				Name:  "learn python",
				Price: 11,
			},
			{
				ID:    "item_4",
				Name:  "learn java",
				Price: 12,
			},
		},
	}
	b, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
}

func unmarshal() {
	s := `{"id":"1234","item":{"id":"item_1","name":"learn go","price":15},"quantity":3,"total_price":30,"items":[{"id":"item_3","name":"learn python","price":11},{"id":"item_4","name":"learn java","price":12}]}`
	var o Order
	err := json.Unmarshal([]byte(s), &o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", o)
	for _, item := range o.Items {
		fmt.Println(*item)
	}
}
