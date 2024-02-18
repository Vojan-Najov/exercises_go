/*
Отчетики

Вам нужно создать программу для создания отчетов о пользователях. У вас
есть структура User с данными о пользователях и структура Report, которая
встраивает в себя структуру User. Ваша задача - создать отчеты о
пользователях на основе их данных.

Создайте структуру User со следующими полями:
ID (уникальный идентификатор пользователя)
Name (имя пользователя)
Email (электронная почта пользователя)
Age (возраст пользователя)

Создайте структуру Report, которая встраивает в себя структуру User и
добавляет следующие поля:
ReportID (уникальный идентификатор отчета)
Date (дата создания отчета)

Создайте функцию CreateReport(user User, reportDate string) Report,
которая принимает пользователя и дату и возвращает отчет с заполненными
данными. Уникальный ReportID можно сгенерировать, например, с
использованием времени.

Создайте функцию PrintReport(report Report), которая выводит информацию из
отчета, включая данные о пользователе и дату создания отчета.

Создайте функцию
GenerateUserReports(users []User, reportDate string) []Report,
которая принимает список пользователей и дату и возвращает список отчетов
о пользователях. Для каждого пользователя в списке создайте отчет,
используя функцию CreateReport, и добавьте его в результирующий список.
*/

package main

import (
  "fmt"
)

type User struct {
  ID int
  Name string
  Email string
  Age int
}

type Report struct {
  User
  ReportID int
  Date string
}

func CreateReport(user User, reportDate string) Report {
  return Report{User: user, ReportID: user.ID, Date: reportDate}
}

func PrintReport(report Report) {
  fmt.Printf("%d %s %s %d %s\n", report.ID, report.Name, report.Email,
                                 report.Age, report.Date)
}

func GenerateUserReports(users []User, reportDate string) []Report {
  s := make([]Report, 0, len(users))
  for _, u := range users {
    s = append(s, CreateReport(u, reportDate))
  }
  return s
}

func main() {
  u1 := User{1, "Bob", "bob@t.com", 30}
  u2 := User{2, "John", "john@t.com", 40}

  s1 := []User{u1, u2}
  s2 := GenerateUserReports(s1, "10-11-12")
  for _, r := range s2 {
    PrintReport(r)
  }
}
