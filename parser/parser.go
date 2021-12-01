package parser

import (
	"strings"

	"github.com/kyuki3rain/simple_cpu/order"
)

func Parse(program string) []order.Order {
	var orders []order.Order
	order_strings := strings.Split(program, "\n")

	for _, s := range order_strings {
		if s == "" {
			continue
		}

		order_string := strings.Fields(s)
		orders = append(orders, order.Init(order_string))
	}

	return orders
}
