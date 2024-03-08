package main

import (
  "fmt"
)

type Person struct {
  Name string
  Age int
  Address string
}

func (p Person) Print() {
  fmt.Printf("Name: %s\nAge: %d\nAddress: %s\n", p.Name, p.Age, p.Address)
}

func main() {
  p := Person{Name: "Гоша", Age: 21, Address: "Ясногорск"}
  p.Print()
}
