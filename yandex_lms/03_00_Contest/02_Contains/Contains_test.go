/*
 * Тест Contains
 *
 * Функция
 * Contains(ctx context.Context, r io.Reader, seq []byte) (bool, error)
 * находит первое вхождение байт seq в данных, доступных через Reader r.
 * Если последовательность найдена - она возвращает true, nil,
 * иначе false, nil.
 * В случае возникновения ошибки - false и возникшую ошибку.
 * В случае отмены контекста - функция возвращает false и ошибку -причину отмены контекста.
 * Напишите тесты для проверки корректности работы.
 * Примечания
 *
 * покрытие кода функции должно быть 100%.
 */

package main

import (
	"bytes"
	"context"
	"errors"
	"testing"
)

type WrongReader struct{}

func (r *WrongReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("tmp")
}

func TestContainsOk1(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	buf := bytes.NewBuffer([]byte("123456789"))
	seq := []byte("1")

	got, err := Contains(ctx, buf, seq)
	if got != true || err != nil {
		t.Errorf("Contains: got %v, %v; want %v, %v", got, err, true, nil)
	}
}

func TestContainsOk2(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	buf := bytes.NewBuffer([]byte("123456789"))
	seq := []byte("4")

	got, err := Contains(ctx, buf, seq)
	if got != true || err != nil {
		t.Errorf("Contains: got %v, %v; want %v, %v", got, err, true, nil)
	}
}

func TestContainsOk3(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	buf := bytes.NewBuffer([]byte("123456789"))
	seq := []byte("9")

	got, err := Contains(ctx, buf, seq)
	if got != true || err != nil {
		t.Errorf("Contains: got %v, %v; want %v, %v", got, err, true, nil)
	}
}

func TestContainsOk4(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	buf := bytes.NewBuffer([]byte("123456789"))
	seq := []byte("45")

	got, err := Contains(ctx, buf, seq)
	if got != true || err != nil {
		t.Errorf("Contains: got %v, %v; want %v, %v", got, err, true, nil)
	}
}

func TestContainsFail(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	buf := bytes.NewBuffer([]byte("123456789"))
	seq := []byte("a")

	got, err := Contains(ctx, buf, seq)
	if got != false || err != nil {
		t.Errorf("Contains: got %v, %v; want %v, %v", got, err, false, nil)
	}
}

func TestContainsCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	buf := bytes.NewBuffer([]byte("123456789"))
	seq := []byte("4")

	got, err := Contains(ctx, buf, seq)
	if got != false || err == nil {
		t.Errorf("Contains: got %v, %v; want %v, %v", got, err, true, errors.New(""))
	}
}

func TestContainsFalse(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	buf := bytes.NewBuffer([]byte{})
	seq := []byte("4")

	got, err := Contains(ctx, buf, seq)
	if got != false || err != nil {
		t.Errorf("Contains: got %v, %v; want %v, %v", got, err, true, errors.New(""))
	}
}

func TestContainsReadError(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	buf := new(WrongReader)
	seq := []byte("4")

	got, err := Contains(ctx, buf, seq)
	if got != false || err == nil {
		t.Errorf(
			"Contains: got %v, %v; want %v, %v",
			got,
			err,
			true,
			errors.New("tmp"),
		)
	}
}
