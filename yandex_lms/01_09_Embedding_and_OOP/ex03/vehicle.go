/*
Управление транспортом

Вам нужно создать систему управления транспортными средствами, такими как
автомобили и мотоциклы. Каждое транспортное средство должно иметь метод для
расчета времени в пути до определенного пункта назначения.

Создайте интерфейс Vehicle, который будет представлять транспортное средство
и иметь метод CalculateTravelTime(distance float64) float64 для расчета
времени в пути.

Реализуйте две структуры, Car (автомобиль) и Motorcycle (мотоцикл), обе
должны реализовывать интерфейс Vehicle и иметь соответствующие поля для
хранения данных о транспортных средствах (например, скорость и тип
транспортного средства).

Создайте функцию
EstimateTravelTime(vehicles []Vehicle, distance float64) map[string]float64,
которая принимает список транспортных средств и расстояние до пункта
назначения, а затем возвращает карту (map), где ключи - это типы
транспортных средств, а значения - время в пути для каждого транспортного
средства. Используйте полиморфизм для вызова метода CalculateTravelTime()
на каждом транспортном средстве независимо от его типа.
*/

package main

import (
  "fmt"
)

type Vehicle interface {
  GetType() string
  CalculateTravelTime(distance float64) float64
}

type Car struct {
  Type string
  Speed float64
  FuelType string
}

func NewCar(speed float64) Car {
  return Car{Type: "Car", Speed: speed}
}

func (c Car) GetType() string {
  return c.Type
}

func (c Car) CalculateTravelTime(distance float64) float64 {
  return distance / c.Speed
}

type Motorcycle struct {
  Type string
  Speed float64
}

func NewMotorcycle(speed float64) Motorcycle {
  return Motorcycle{Type: "Motorcycle", Speed: speed}
}

func (m Motorcycle) GetType() string {
  return m.Type
}


func (m Motorcycle) CalculateTravelTime(distance float64) float64 {
  return distance / m.Speed
}

func EstimateTravelTime(vehicles []Vehicle, distance float64) map[string]float64 {
  m := make(map[string]float64)
  for _, v := range vehicles {
    m[v.GetType()] = v.CalculateTravelTime(distance)
  }
  return m
}

func main() {
  c := NewCar(100)
  m := NewMotorcycle(150)

  s := make([]Vehicle, 0, 2)
  s = append(s, c)  
  s = append(s, m)

  mm := EstimateTravelTime(s, 300.0)
  fmt.Println(mm)  
}
