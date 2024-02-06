/*
Файлы. Чтение лог файла

Представьте, что другая программа пишет лог-файлы, где каждая строка
начинается с даты формата dd.MM.YYYY.

Ваша задача - написать функцию
ExtractLog(inputFileName string, start, end time.Time) ([]string, error),
которая вернет строки "лога", созданные в указанный диапазон времени
[start..end]

Формат вывода
Если в процессе работы возникает любая ошибка, то необходимо вернуть
nil, err.

Если ошибок нет, то возвращайте nil в качестве значения ошибки.

Если ни одна строка не попала в указанный диапазон, то должна также
вернуться произвольная ошибка.

Примечания

Например, для исходного файла:
12.12.2022 info
13.12.2022 info
14.12.2022 info
15.12.2022 info
16.12.2022 info
Если start = 13.12.2022, end = 15.12.2022, то функция должна вернуть:
13.12.2022 info
14.12.2022 info
15.12.2022 info

Если ни одна строка не попала в указанный диапазон, то должна вернуться
ошибка.
*/

package main

import (
  "os"
  "bufio"
  "time"
  "fmt"
  "errors"
)

func ExtractLog(inputFile string, start, end time.Time) ([]string, error) {
  lines := make([]string, 0)

  file, err := os.Open(inputFile)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  fileScanner := bufio.NewScanner(file)
  for fileScanner.Scan() {
    str := fileScanner.Text()

    if (len(str) < 10) {
      continue
    }
    date, err := time.Parse("02.01.2006", str[:10])
    if err != nil {
      return nil, err
    }
    
    if (date.Equal(start) || date.Equal(end) || 
        date.After(start) && date.Before(end)) {
      lines = append(lines, str)
    }
  }

  if len(lines) == 0 {
    return nil, errors.New("no lines")
  }

  return lines, nil
}

func main() {
  start, _ := time.Parse("02.01.2006", "13.12.2022")
  end, _ := time.Parse("02.01.2006", "15.12.2022")
  lines, err := ExtractLog("tmp.log", start, end)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(lines)
}

