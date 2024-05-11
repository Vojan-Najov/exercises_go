/*
 * 4. Sync patterns
 *
 * Дан сервер доступный по адресу localhost:8082.
 * По запросу localhost:8082/mark?name=<имя студента> сервер возвращает:
 * - код 200 и значение оценки студента, если все прошло успешно
 * - код 404, если студент не найден
 * - код 500, если при с сервером случилась проблема
 *
 * Напишите функцию CompareList(names []string) (map[string]string, error),
 * выводит карту, где ключом является имя студента из списка name, а значением
 * является
 * > (оценка студента больше средней оценки студентов),
 * < (оценка студента меньше средней оценки студентов),
 * = (оценка студента равна средней оценки студентов).
 * Функция возвращает ошибку в случае невозможности получения оценки хотя бы
 * одного студента.
 *
 * Примечания
 * Используйте WaitGroup
 *
 */

package main

import (
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

func CompareList(names []string) (map[string]string, error) {
	marks := make([]int, len(names))
	errs := make([]error, len(names))
	wg := sync.WaitGroup{}
	wg.Add(len(names))

	for i := range names {
		go func(idx int) {
			defer wg.Done()
			resp, err := http.Get(
				"http://localhost:8082/mark?name=" + names[idx],
			)
			if err != nil {
				errs[idx] = err
				return
			}
			defer resp.Body.Close()
			if resp.StatusCode != 200 {
				errs[idx] = errors.New(names[idx] + resp.Status)
				return
			}
			buf := new(strings.Builder)
			_, err = io.Copy(buf, resp.Body)
			if err != nil {
				errs[idx] = err
				return
			}
			marks[idx], err = strconv.Atoi(buf.String())
			if err != nil {
				errs[idx] = err
				return
			}
		}(i)
	}

	wg.Wait()

	var avg int
	for i := range names {
		if errs[i] != nil {
			return nil, errs[i]
		}
		avg += marks[i]
	}
	avg /= len(names)

	m := make(map[string]string)
	for i := range names {
		if marks[i] > avg {
			m[names[i]] = ">"
		} else if marks[i] < avg {
			m[names[i]] = "<"
		} else {
			m[names[i]] = "="
		}
	}

	return m, nil
}
