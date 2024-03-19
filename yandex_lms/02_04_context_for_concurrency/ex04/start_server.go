/*
timeoutHandler

Напишите функцию StartServer(maxTimeout time.Duration) которая запусает веб-сервер
по адресу http://localhost:8080. При обращении по URL
http://localhost:8080/readSource сервер должен сделать запрос по другому адресу:
http://localhost:8081/provideData (код запуска сервера localhost:8081 писать не нужно)
и вернуть полученные данные. Используйте http.TimeoutHandler для ограничения времени
ожидания данных с сервера localhost:8081. При привышении лимита maxTimeout
пользователю должна веруться ошибка с кодом StatusServiceUnavailable, иначе -
полученные данные.
*/

package main

import (
	"io"
	"net/http"
	"time"
)

type TmpHandler struct{}

func (h TmpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest(
		r.Method,
		"http://localhost:8081/provideData",
		r.Body,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func StartServer(maxTimeout time.Duration) {
	http.Handle(
		"/readSource",
		http.TimeoutHandler(TmpHandler{}, maxTimeout, ""),
	)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		return
	}
}
