/*
Сортировка сотрудников предприятия

На предприятии работают несколько cсотрудников. Каждый из них имеет свою
должность, фиксированную месячную заработную плату, и стаж работы. Напишите 
программу в котором тип Company реализует следующий интерфейс:
type CompanyInterface interface{
AddWorkerInfo(name, position string, salary, experience uint) error
SortWorkers() ([]string, error)
}

AddWorkerInfo - метод добавления информации о новых сотрудниках, где name - 
имя сотрудника, position - должность, salary - месячная зароботная плата,
experience - стаж работы (месяцев).

SortWorkers - метод сортировки, который сортирует список сотрудников по
следующим свойствам: доход за время работы на предприятии(по убыванию),
должность (от высокой) и возвращает слайсл формата: имя - доход - должность.
Допустимые должности в порядке убывания: "директор", "зам. директора",
"начальник цеха", "мастер", "рабочий".

Примечания

Пример отсортированного вывода:
Михаил - 12000 - директор
Андрей - 11800 - мастер
Игорь - 11000 - зам. директора
*/

package main

import (
  "fmt"
  "sort"
  "errors"
)

type CompanyInterface interface {
  AddWorkerInfo(name, position string, salary, experience uint) error
  SortWorkers() ([]string, error)
}

type Company struct {
  names []string
  positions []string
  salaries []uint
  experiences []uint
}

func (c Company) Len() int {
  return len(c.names)
}

func (c Company) Swap(i, j int) {
  c.names[i], c.names[j] = c.names[j], c.names[i]
  c.positions[i], c.positions[j] = c.positions[j], c.positions[i]
  c.salaries[i], c.salaries[j] = c.salaries[j], c.salaries[i]
  c.experiences[i], c.experiences[j] = c.experiences[j], c.experiences[i]
}

func (c Company) Less(i, j int) bool {
  sort_positions := map[string]int{
    "директор": 1, "зам. директора": 2,
    "начальник цеха": 3, "мастер": 4, "рабочий": 5,
  }
  iprofit := c.salaries[i] * c.experiences[i]
  jprofit := c.salaries[j] * c.experiences[j]
  return iprofit > jprofit ||
         iprofit > jprofit &&
         sort_positions[c.positions[i]] < sort_positions[c.positions[j]]
}

func (c *Company) AddWorkerInfo(name, position string,
                                salary, experience uint) error {
  true_positions := map[string]bool{
    "директор": true, "зам. директора": true,
    "начальник цеха": true, "мастер": true, "рабочий": true,
  }
  if !true_positions[position] {
    return errors.New("Incorrect position: " + position)
  }
  c.names = append(c.names, name)
  c.positions = append(c.positions, position)
  c.salaries = append(c.salaries, salary)
  c.experiences = append(c.experiences, experience)
  return nil
}

func (c *Company) SortWorkers() ([]string, error) {
  sort.Sort(c)
  result := make([]string, len(c.names))
  for i := 0; i < len(c.names); i++ {
    result[i] = fmt.Sprintf("%s - %d - %s",
                            c.names[i],
                            c.salaries[i] * c.experiences[i],
                            c.positions[i])
  }
  return result, nil
}

func main() {
  var c Company
  c.AddWorkerInfo("Михаил", "директор", 200, 5 * 12)
  c.AddWorkerInfo("Игорь", "зам. директора", 180, 3 * 12)
  c.AddWorkerInfo("Николай", "начальник цеха", 120, 2 * 12)
  c.AddWorkerInfo("Андрей", "мастер", 90, 10 * 12)
  c.AddWorkerInfo("Виктор", "рабочий", 80, 3 * 12)

  strs, _ := c.SortWorkers()
  for _, str := range strs {
    fmt.Println(str)
  }
}
