/*
 * 2. timeouts_patterns
 *
 * Напишите программу для чтения JSON-файла и отображения содержимого в
 * терминале, используйте концепцию таймаута вместе с контекстом и отменой при
 * чтении JSON-файла. Реализуйте функцию
 * func readJSON(ctx context.Context, path string, result chan<- []byte)
 * ctx - контекст
 * path - путь к json
 * result - канал, в который нужно вывести прочитанное значение
 */

package main

import (
	"context"
	"os"
)

func readJSON(ctx context.Context, path string, result chan<- []byte) {
	data, err := os.ReadFile(path)
	select {
	case <-ctx.Done():
		return
	default:
		if err != nil {
			result <- data
		}
	}
}
