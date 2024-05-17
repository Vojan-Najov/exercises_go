package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// создание контекста WithTimeout, который ограничивает
        // продолжительность в течение 2 секунд
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	//создание канала chan_c
	chan_c := make(chan string)

	//создание потока выполнения горутины
	go func() {
		// создание фиктивной операции, которая занимает 3 секунды с
                // использованием Sleep()
		time.Sleep(3 * time.Second)
		chan_c <- "Инструкции успешно завершены."
	}()

	select {
	case result := <-chan_c:
		fmt.Println(result)
	case <-ctx.Done():
		fmt.Println("Время операции истекло или было отменено.")
	}
}
