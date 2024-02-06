/*
Файлы. Копирование части файла

Напишите функцию
CopyFilePart(inputFilename, outFileName string, startpos int) error,
которая открывает файл с именем inputFilename на чтение, создает файл с
именем outFileName и записывает содержимое файла inputFilename с позиции
startPos и до конца в файл outFileName.

Не забудьте закрыть файлы после обработки.

Заметьте, что функция возвращает ошибку.
Если все операции прошли без ошибок, то верните nil.
*/

package main

import (
  "os"
  "io"
  "fmt"
)

func CopyFilePart(inFilename, outFilename string, startpos int) error {
  input, readErr := os.Open(inFilename)
  if readErr != nil {
    return readErr
  }
  defer input.Close()

  if _, seekErr := input.Seek(int64(startpos), 0); seekErr != nil {
    return seekErr
  }

  output, createErr := os.Create(outFilename)
  if createErr != nil {
    return createErr
  }
  defer output.Close()

  buf := make([]byte, 1024 * 1024)
  for {
    n, err := input.Read(buf)
    if err == io.EOF {
      break
    }
    if err != nil {
      return err
    }

    _, err = output.Write(buf[:n])
    if err != nil {
      return nil
    }
  }

  return nil
}

func main() {
  err := CopyFilePart("CopyFilePart.go", "tmp.txt", 698)
  if err != nil {
    fmt.Println(err)
  }
}
