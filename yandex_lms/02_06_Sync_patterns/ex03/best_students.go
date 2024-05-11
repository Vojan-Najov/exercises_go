/*
 * 3. Sync patterns
 *
 * Дан сервер доступный по адресу localhost:8082.
 * По запросу localhost:8082/mark?name=<имя студента> сервер возвращает:
 * - код 200 и значение оценки студента, если все прошло успешно
 * - код 404, если студент не найден
 * - код 500, если при с сервером случилась проблема
 *
 * Напишите функцию BestStudents(names []string) (string, error), выводяющую
 * список студентов с оценками выше средней успеваемости студентов из списка
 * names в алфавитном порядке через запятую. Функция возвращает ошибку в случае
 * невозможности получения оценки хотя бы одного студента.
 *
 * Примечания
 * Используйте WaitGroup
 */

package main

import (
	"errors"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func BestStudents(names []string) (string, error) {
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
			return "", errs[i]
		}
		avg += marks[i]
	}
	avg /= len(names)

	var bestStudents []string
	for i := range names {
		if marks[i] > avg {
			bestStudents = append(bestStudents, names[i])
		}
	}
	sort.Strings(bestStudents)

	return strings.Join(bestStudents, ","), nil
}
