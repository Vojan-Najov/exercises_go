/*
Ограничение времени запроса (concurrently)

Напишите функцию
FetchAPI(ctx context.Context, urls []string, timeout time.Duration) []*APIResponse,\
которая одновременно (concurrently) получает данные из переданных urls (метод GET).
Используйте контекст для ограничения времени запроса и отмены ожидания свыше timeout.
В случае возникновения ошибки - верните её в соответсвующем объекте APIResponse. При
превышении таймаута ожидания должна быть ошибка context.DeadlineExceeded.

Примечания
Код должет содержать структуру:

	type APIResponse struct {
	  URL string // запрошенный URL
	  Data string // тело ответа
	  StatusCode int // код ответа
	  Err error // ошибка, если возникла
	}
*/
package main

import (
	"context"
	"io"
	"net/http"
	"sync"
	"time"
)

type APIResponse struct {
	URL        string
	Data       string
	StatusCode int
	Err        error
}

func FetchAPI(
	ctx context.Context,
	urls []string,
	timeout time.Duration,
) []*APIResponse {
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(len(urls))

	client := http.Client{}
	answs := make([]*APIResponse, len(urls))

	for i, url := range urls {
		go func(idx int, url string) {
			defer wg.Done()
			answ := APIResponse{URL: url}
			ctx, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()

			req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
			if err != nil {
				answ.Err = err
				mu.Lock()
				answs[idx] = &answ
				mu.Unlock()
				return
			}

			resp, err := client.Do(req)
			if err != nil {
				answ.Err = err
				mu.Lock()
				answs[idx] = &answ
				mu.Unlock()
				return
			}
			defer resp.Body.Close()

			answ.StatusCode = resp.StatusCode

			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				answ.Err = err
				mu.Lock()
				answs[idx] = &answ
				mu.Unlock()
				return
			}

			answ.Data = string(buf)
			mu.Lock()
			answs[idx] = &answ
			mu.Unlock()
		}(i, url)
	}

	wg.Wait()

	return answs
}
