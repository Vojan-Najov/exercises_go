/*
 * Синхронизация OrderedFanIn
 *
 * На вход программы поступают сообщения в виде объектов, тип которых
 * удовлетворяет интерфейсу:
 *
 * type ordered interface {
 *   getIndex() int
 *   getData() string
 * }
 *
 * Напишите функцию
 * OrderedFanIn[T ordered](ctx context.Context, channels ...<-chan T) <-chan T,
 * которая запишет данные в выходной канал в соответствии со значением
 * getIndex()(в порядке увеличения).
 * Код программы должен содержать определение интерфейса.
 *
 */

package main

import (
	"context"
	"sync"
)

type ordered interface {
	getIndex() int
	getData() string
}

type fanInRecord[T ordered] struct {
	index int           // порядковый номер горутины, из которой получено сообщение
	data  T             // непосредственно данные
	pause chan struct{} // канал для синхронизации
}

func inTemp[T ordered](
	ctx context.Context,
	channels ...<-chan T,
) <-chan fanInRecord[T] {
	// канал для ожидания
	fanInCh := make(chan fanInRecord[T])
	// для синхронизации
	wg := sync.WaitGroup{}
	// перебор всех входных каналов
	for i := range channels {
		wg.Add(1)
		// запустим горутину для получения данных из канала
		go func(index int) {
			defer wg.Done()
			// канал для синхронизации
			pauseCh := make(chan struct{})
			// цикл для получения данных из канала
			for {
				select {
				// получим данные из канала
				case data, ok := <-channels[index]:
					if !ok {
						return // канал закрыт - выходим
					}
					// положим во временный канал вместе с индексом
					fanInCh <- fanInRecord[T]{
						// индекс канала, откуда пришли данные
						index: index,
						// данные из канала
						data: data,
						// канал для синхронизации
						pause: pauseCh,
					}
				case <-ctx.Done():
					return
				}
				// ждём, пока в канал pause не будет передан сигнал
				// о получении очередного элемента из канала
				select {
				case <-pauseCh:
				// сняли с паузы
				// продолжим обработку данных из входного канала
				case <-ctx.Done():
					return
				}
			}
		}(i)
	}
	go func() {
		// ожидаем завершения
		wg.Wait()
		close(fanInCh)
	}()
	// вернём канал с неотсортированными элементами
	return fanInCh
}

func processTempCh[T ordered](
	ctx context.Context,
	inputChannelsNum int, // количество входных каналов
	fanInCh <-chan fanInRecord[T], // временный канал с данными
) <-chan T {
	// выходной канал с упорядоченными данными
	outputCh := make(chan T)
	go func() {
		defer close(outputCh)
		// порядковый номер очередного элемента
		expected := 0
		// буфер для ожидания элементов по количеству входных каналов
		queuedData := make([]*fanInRecord[T], inputChannelsNum)
		for in := range fanInCh {
			// если получили элемент с номером, который ожидаем
			if in.data.getIndex() == expected {
				select {
				// запишем элемент в выходной канал
				case outputCh <- in.data:
					// снимем с паузы исходный канал
					// для продолжения обработки из входного канала
					in.pause <- struct{}{}
					// инкремент номера очередного элемента
					expected++
					// здесь нужно реализовать запись в выходной канал
					// из буфера queuedData (задача для домашней работы)
					for {
						endFlag := true
						for i, ptr := range queuedData {
							if ptr != nil && ptr.data.getIndex() == expected {
								endFlag = false
								select {
								case outputCh <- ptr.data:
									ptr.pause <- struct{}{}
									expected++
									queuedData[i] = nil
								case <-ctx.Done():
									return
								}
							}
							if !endFlag {
								break
							}
						}
						if endFlag {
							break
						}
					}
				case <-ctx.Done():
					return
				}
			} else {
				// если НЕ получили элемент с номером, который ожидаем
				// запишем элемент в буфер
				in := in
				queuedData[in.index] = &in
			}
		}
	}()
	return outputCh
}

func OrderedFanIn[T ordered](ctx context.Context, channels ...<-chan T) <-chan T {
	inCh := inTemp(ctx, channels...)
	outCh := processTempCh(ctx, len(channels), inCh)
	return outCh
}
