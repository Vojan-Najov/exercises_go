/*
Каналы

Напишите программу, где в горутину func sendDataToRoutine(ch chan string) передается
канал и эта горутина передает сообщение "Hello from the channel!" через этот канал.
*/

package main

func main() {
	ch := make(chan string)
	go sendDataToRoutine(ch)
	_ = <-ch
}

func sendDataToRoutine(ch chan string) {
	ch <- "Hello from the channel!"
}
