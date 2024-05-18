/*
 * Синхронизация Запись из очереди
 *
 * Для функции processTempCh дополните функционал сканирования элементов из
 * буфера queuedData. Если элемент queuedData[i] - очередной для записи в
 * выходной канал, то должны быть реализованы следующие действия:
 * - запись в выходной канал queuedData[i]
 * - разблокировка исходного канала с помощью pause.
 *
 * Код решения должен содержать объявление интерфейса sequenced и типа
 * fanInRecord.
 *
 * type sequenced interface {
 *   getSequence() int
 * }
 *
 * type fanInRecord[T sequenced] struct {
 *   index int
 *   data T
 *   pause chan struct{}
 * }
 *
 */

package main

import (
	"context"
)

type sequenced interface {
	getSequence() int
}

type fanInRecord[T sequenced] struct {
	index int
	data  T
	pause chan struct{}
}

func processTempCh[T sequenced](
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
			if in.data.getSequence() == expected {
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
							if ptr != nil && ptr.data.getSequence() == expected {
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
