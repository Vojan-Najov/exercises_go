/*
 * 1. Генератор строк
 *
 * Напишите функцию StringsGen(lines ...string) <-chan string,
 * которая принимает на вход набор строк и возвращает их через канал.
 */

package main

func StringsGen(lines ...string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for _, line := range lines {
			out <- line
		}
	}()
	return out
}
