/*
 * 1. Sync patterns
 *
 * Дан сервер доступный по адресу localhost:8082. По запросу
 * localhost:8082/mark?name=<имя студента> сервер возвращает:
 * - код 200 и значение оценки студента, если все прошло успешно
 * - код 404, если студент не найден
 * - код 500, если с сервером случилась проблема
 *
 * Напишите функцию Compare(name1, name2 string) (string, error), которая
 * сравнивает оценки двух студентов с именами name1 и name2 и выводит
 * > (оценка студента 1 больше оценки студента 2)
 * < (оценка студента 1 меньше оценки студента 2),
 * = (оценка студента 1 равна оценке студента 2).
 * Функция возвращает ошибку в случае невозможности получения оценки хотя бы
 * одного студента.
 *
 * Примечания
 * Используйте WaitGroup
 */

package main

import (
	"io"

	"errors"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

func Compare(name1, name2 string) (string, error) {
	var resp1, resp2 *http.Response
	var err1, err2 error

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		resp1, err1 = http.Get("http://localhost:8082/mark?name=" + name1)
	}()
	go func() {
		defer wg.Done()
		resp2, err2 = http.Get("http://localhost:8082/mark?name=" + name2)
	}()

	wg.Wait()

	if err1 != nil {
		return "", err1
	} else if err2 != nil {
		return "", err2
	}

	defer resp1.Body.Close()
	defer resp2.Body.Close()

	if resp1.StatusCode != 200 {
		return "", errors.New(name1 + resp1.Status)
	}
	if resp2.StatusCode != 200 {
		return "", errors.New(name2 + resp2.Status)
	}

	buf1 := new(strings.Builder)
	_, err1 = io.Copy(buf1, resp1.Body)
	if err1 != nil {
		return "", err1
	}
	buf2 := new(strings.Builder)
	_, err1 = io.Copy(buf2, resp2.Body)
	if err1 != nil {
		return "", err2
	}

	res1, err1 := strconv.Atoi(buf1.String())
	if err1 != nil {
		return "", err1
	}
	res2, err2 := strconv.Atoi(buf2.String())
	if err2 != nil {
		return "", err2
	}

	if res1 < res2 {
		return "<", nil
	} else if res1 > res2 {
		return ">", nil
	}
	return "=", nil
}
