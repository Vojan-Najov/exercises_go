/*

Напишите программу для управления банковским счетом. Создайте структуру
Account с приватными полями balance (баланс) и owner (владелец).

Реализуйте методы для установки баланса и получения баланса, а также
методы для внесения и снятия денег с счета. Убедитесь, что баланс не может
быть отрицательным.

Примечания

Код программы должен содержать описание струкрутры Account:
type Account struct { owner string balance float64 }

Методы, которые нужно реализовать

NewAccount(owner string) *Account
SetBalance(amount float64) error
GetBalance() float64
Deposit(amount float64) error
Withdraw(amount float64) error

*/

package main

import (
  "fmt"
  "errors"
)

type Account struct {
  owner string
  balance float64
}

func NewAccount(owner string) *Account {
  return &Account{owner: owner, balance: 0.0}
}

func (ac *Account) SetBalance(amount float64) error {
  if amount < 0.0 {
    return errors.New("negative amount")
  }
  ac.balance = amount
  return nil
}

func (ac *Account) GetBalance() float64 {
  return ac.balance
}

func (ac *Account) Deposit(amount float64) error {
  if amount < 0 {
    return errors.New("negative amount")
  }
  ac.balance += amount
  return nil
}

func (ac *Account) Withdraw(amount float64) error {
  if amount < 0 {
    return errors.New("negative amount")
  }
  if amount > ac.balance {
    return errors.New("underflow balance")
  }
  ac.balance -= amount
  return nil
}

func main() {
  ac := NewAccount("Bob")
  if err := ac.Withdraw(10); err != nil {
    fmt.Println(err)
  }
  ac.SetBalance(100)
  if err := ac.Deposit(-10); err != nil {
    fmt.Println(err)
  }
  ac.Deposit(10)
  ac.Withdraw(20)
  fmt.Println(ac.GetBalance())
}
