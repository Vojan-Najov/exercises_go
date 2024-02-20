/*
Логгирование HTTP сервера

Напишите веб сервер, который логгирует информацию о запросе вида
2023/11/01 09:41:19 GET / [::1]:55250 {} 7.708µs
2023/11/01 09:41:19 GET / [::1]:55250 {} 1.166µs
В функции main должен быть запуск веб-сервера на порту 8080

*/

package main

import (
	"log"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		//log.Printf("Запрос: %s %s", r.Method, r.URL.Path)

		// Передаем управление следующему обработчику
		next.ServeHTTP(w, r)

		// Вычисляем время выполнения запроса
		duration := time.Since(start)
		// Логируем информацию о запросе
		log.Printf("%s %s %s {} %s", r.Method, r.URL.Path, r.Host, duration)
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Привет, мир!"))
}

func main() {
	mux := http.NewServeMux()

	// Создаем обработчик для маршрута "/"
	hello := http.HandlerFunc(helloHandler)

	// Применяем logging middleware к обработчику "/"
	mux.Handle("/", loggingMiddleware(hello))

	// Запускаем сервер на порту 8080
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

