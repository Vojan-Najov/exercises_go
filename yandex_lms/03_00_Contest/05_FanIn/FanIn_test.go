/*
 * Тест FanIn
 *
 * Напишите тесты для следующего кода:
 *
 * func FanIn[T any](ctx context.Context, channels ...<-chan T) <-chan T {
 *    outputCh := make(chan T)
 *    wg := sync.WaitGroup{}
 *    for _, ch := range channels {
 *        wg.Add(1)
 *        go func(input <-chan T) {
 *            defer wg.Done()
 *            for {
 *                 select {
 *                 case data, ok := <-input:
 *                     if !ok {
 *                         return
 *                     }
 *                     outputCh <- data
 *                 case <-ctx.Done():
 *                     return
 *                 }
 *             }
 *         }(ch)
 *     }
 *     go func() {
 *         wg.Wait()       // дождёмся завершения обработки всех каналов
 *         close(outputCh) // закроем выходной канал
 *     }()
 *
 *     return outputCh // вернём канал
 * }
 *
 * Примечания
 *    покрытие кода функции должно быть 100%.
 */

package main

import (
	"context"
	"testing"
)

func TestFanInOk(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch1 := make(chan int, 10)
	ch2 := make(chan int, 20)
	ch3 := make(chan int, 30)
	for i := 0; i < 10; i++ {
		ch1 <- i
	}
	for i := 10; i < 30; i++ {
		ch2 <- i
	}
	for i := 30; i < 60; i++ {
		ch3 <- i
	}
	close(ch1)
	close(ch2)
	close(ch3)

	out := FanIn(ctx, ch1, ch2, ch3)
	m := make(map[int]bool)
	for v := range out {
		m[v] = true
	}

	for i := 0; i < 60; i++ {
		if !m[i] {
			t.Errorf("Incorrect output channel")
		}
	}
}

func TestFanInCancelCtx(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	ch1 := make(chan int, 10)
	ch2 := make(chan int, 20)
	ch3 := make(chan int, 30)
	for i := 0; i < 5; i++ {
		ch1 <- i
	}
	for i := 5; i < 15; i++ {
		ch2 <- i
	}
	for i := 15; i < 30; i++ {
		ch3 <- i
	}

	out := FanIn(ctx, ch1, ch2, ch3)
	m := make(map[int]bool)
	count := 0
	for v := range out {
		count++
		m[v] = true
		if count == 30 {
			cancel()
		}
	}

	for i := 30; i < 35; i++ {
		ch1 <- i
	}
	for i := 35; i < 45; i++ {
		ch2 <- i
	}
	for i := 45; i < 60; i++ {
		ch3 <- i
	}
	close(ch1)
	close(ch2)
	close(ch3)

	for i := 0; i < 30; i++ {
		if !m[i] {
			t.Errorf("Incorrect output channel %v", i)
		}
	}
	for i := 30; i < 60; i++ {
		if m[i] {
			t.Errorf("Incorrect output channel with cancel")

		}
	}
}
