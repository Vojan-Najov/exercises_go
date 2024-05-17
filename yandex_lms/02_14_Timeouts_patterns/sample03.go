
package main

import (
	"fmt"
	"net/http"
	"time"
)

func fetchURL(url string, c chan string) {
	//создание http-клиента с таймаутом
	client := http.Client{
		/*
		 * Установка таймаута для HTTP-запроса, вы можете уменьшить
		 * или увеличить это время.
                 * в случае уменьшения вы можете часто сталкиваться с таймаутом
		 * в случае увеличения получение данных займет больше времени
		*/
		Timeout: 5 * time.Second,
	}

	//получение ответа по URL
	resp, err := client.Get(url)
	if err != nil {
		c <- fmt.Sprintf("Ошибка при получении %s: %s", url, err)
		return
	}
	defer resp.Body.Close()
	c <- fmt.Sprintf("Ответ от %s: Статус - %s", url, resp.Status)
}

func main() {
	//создание массива URL-адресов
	urls := []string{
		"https://yandex.ru",
		"https://lyceum.yandex.ru",
		"https://translate.yandex.com",
		// Симуляция несуществующего URL
		//"https://ihumaunkabir.com",
	}

	//создание канала c
	c := make(chan string, len(urls))

	//итерация по массиву URL-адресов
	for _, url := range urls {
		go fetchURL(url, c)
	}

	//установка общего таймаута для всех запросов
	timeout := time.After(15 * time.Second)

	//итерация до конца массива URL-адресов
	for i := 0; i < len(urls); i++ {
		select {
		case result := <-c:
			fmt.Println(result)
		case <-timeout:
			fmt.Println("Произошел таймаут. Прерывание остальных запросов.")
			return
		}
	}
}
