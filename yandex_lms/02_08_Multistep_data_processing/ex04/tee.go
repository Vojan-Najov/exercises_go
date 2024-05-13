/*
 * 4. Копия потока
 *
 * В некоторых случаях на одном из этапов обработки данных нужно сделать копию
 * входящего потока. Например, для записи лога. Напишите функцию
 * Tee[T any](done <-chan struct{}, in <-chan T) (<-chan T, <-chan T),
 * которая записывает данные из канала in в два выходящих канала. Функция
 * должна завершать работу при закрытии канала done, либо при закрытии канала
 * in.
 */

package main

func Tee[T any](done <-chan struct{}, in <-chan T) (<-chan T, <-chan T) {
	out1 := make(chan T)
	out2 := make(chan T)

	go func() {
		defer close(out1)
		defer close(out2)
		for v := range in {
			select {
			case <-done:
				return
			default:
				out1 <- v
				out2 <- v
			}
		}
	}()

	return out1, out2
}
