/*
Логи в терминал

Напишите функцию (logger *OrderLogger) AddOrder(order Order), которая
пишет в терминал информацию о добавленом заказе вида
Добавлен заказ #%d, Имя клиента: %s, Сумма заказа: $%.2f

Примечания

// Order представляет информацию о заказе.
type Order struct {
    OrderNumber  int
    CustomerName string
    OrderAmount  float64
}

// OrderLogger представляет журнал заказов и хранит записи о заказах.
type OrderLogger struct{}

// NewOrderLogger создает новый экземпляр OrderLogger.
func NewOrderLogger() *OrderLogger {
    return &OrderLogger{}
}
*/

package main

import (
  "log"
)

type Order struct {
  OrderNumber  int
  CustomerName string
  OrderAmount  float64
}

type OrderLogger struct {}

func NewOrderLogger() *OrderLogger {
  return &OrderLogger{}
}

func (logger *OrderLogger) AddOrder(order Order) {
  flags := log.Flags()
  log.SetFlags(0)
  log.Printf("Добавлен заказ #%d, Имя клиента: %s, Сумма заказа: $%.2f\n",
             order.OrderNumber, order.CustomerName, order.OrderAmount)
  log.SetFlags(flags)
}

func main() {
  logger := NewOrderLogger()
  order := Order{OrderNumber: 10, CustomerName: "Bob", OrderAmount: 2.5}
  logger.AddOrder(order)
}
