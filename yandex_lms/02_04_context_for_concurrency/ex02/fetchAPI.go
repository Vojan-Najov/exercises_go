/*
Ограничение времени запроса

Напишите функцию
fetchAPI(ctx context.Context, url string, timeout time.Duration) (*APIResponse, error)
которая запрашивает данные по адресу url (метод GET) и возвращает код ответа и само
тело ответа. Используйте контекст для ограничения времени запроса и отмены ожидания
свыше timeout.
В случае возникновения ошибок - возвращайте nil, error. При превышении таймаута
ожидания - верните nil, context.DeadlineExceeded.

Примечания
Код должет содержать структуру:

type APIResponse struct {
  Data string // тело ответа
  StatusCode int // код ответа
}
*/

package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

type APIResponse struct {
	Data       string
	StatusCode int
}

var tmp APIResponse

func fetchAPI(
	ctx context.Context,
	url string,
	timeout time.Duration,
) (*APIResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	client := http.Client{}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	answ := APIResponse{Data: string(buf), StatusCode: resp.StatusCode}
	return &answ, nil
}
