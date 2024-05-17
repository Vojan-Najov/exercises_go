/*
 * 2. timeouts
 *
 * Напишите программу, имитирующую портал викторин.
 * Имеется несколько вопросов и правильных ответов на них. Вопросы будут
 * приходить один за другим. Время, которое отводится на ответ - 1 секунда. Если
 * время истекло, то переходите к следующему вопросу. В конце покажите
 * количество правильных ответов завершенного.
 * Сигнатура проверяемой функции
 * QuizRunner(questions, answers []string, answerCh chan string) int
 * questions - список вопросов
 * answers - список ответов
 * answerCh - канал с пользовательским вводом
 * результат - число правильных ответов
 */

package main

import (
	"time"
)

func QuizRunner(questions, answers []string, answerCh chan string) int {
	var goodAnswers int
	for i := 0; i < len(questions); i++ {
		select {
		case answer := <-answerCh:
			if answer == answers[i] {
				goodAnswers++
			}
		case <-time.After(time.Second):
		}
	}

	return goodAnswers
}
