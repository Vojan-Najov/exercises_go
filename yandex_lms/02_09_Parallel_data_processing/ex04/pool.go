/*
 * Параллельная обработка worker pool *
 *
 * Напишите собственный worker pool (пакет pool). Он будет обрабатывать задачи,
 * реализующие интерфейс PoolTask:
 * type PoolTask interface {
 *   // Execute запускает выполнение задачи и возвращает nil,
 *   // либо возникшую ошибку.
 *   Execute() error
 *   // OnFailure будет обрабатывать ошибки, возникшие в Execute(), то есть
 *   // пул должен вызывать OnFailure в случае, если Execute возвращает ошибку.
 *   OnFailure(error)
 * }
 *
 * Для этого создайте структуру MyPool, которая удовлетворяет следующему
 * интерфейсу:
 * type WorkerPool interface {
 *   // Start подготавливает пул для обработки задач. Должен вызываться один раз
 *   // перед использованием пула. Очередные вызовы должны игнорироваться.
 *   Start()
 *   // Stop останавливает обработку в пуле. Должен вызываться один раз.
 *   // Очередные вызовы должны игнорироваться.
 *   Stop()
 *   // AddWork добавляет задачу для обработки пулом. Добавлять задачи
 *   // можно после вызова Start() и до вызова Stop().
 *   // Если на момент добавления в пуле нет
 *   // свободных ресурсов (очередь заполнена) -
 *   // эту функция ожидает их освобождения (либо вызова Stop).
 *   AddWork(PoolTask)
 * }
 *
 * Код должен содержать конструктор для MyPool:
 * // NewWorkerPool возвращает новый пул
 * // numWorkers - количество воркеров
 * // channelSize - размер очереди ожидания
 * // В случае ошибок верните nil и описание ошибки
 * func NewWorkerPool(numWorkers int, channelSize int) (*MyPool, error){
 *   // ваша реализация
 * }
 */

package pool

import (
	"errors"
	"sync"
)

type PoolTask interface {
	Execute() error
	OnFailure(error)
}

type MyPool struct {
	numberWorkers int
	tasks         chan PoolTask
	wg            sync.WaitGroup
	isActive      bool
	mu            sync.RWMutex
}

func (p *MyPool) Start() {
	p.mu.RLock()
	if p.isActive {
		defer p.mu.RUnlock()
		return
	}
	p.mu.RUnlock()

	p.wg.Add(p.numberWorkers)
	for i := 0; i < p.numberWorkers; i++ {
		go func() {
			defer p.wg.Done()
			for w := range p.tasks {
				err := w.Execute()
				if err != nil {
					w.OnFailure(err)
				}
			}
		}()
	}

	p.mu.Lock()
	p.isActive = true
	p.mu.Unlock()
}

func (p *MyPool) Stop() {
	p.mu.Lock()
	if !p.isActive {
		defer p.mu.Unlock()
		return
	}
	p.isActive = false
	p.mu.Unlock()

	close(p.tasks)
	p.wg.Wait()
}

func (p *MyPool) AddWork(task PoolTask) {
	p.mu.RLock()
	if !p.isActive {
		defer p.mu.RUnlock()
		return
	}
	p.mu.RUnlock()

	for {

		p.mu.RLock()
		if !p.isActive {
			defer p.mu.RUnlock()
			return
		}
		select {
		case p.tasks <- task:
			return
		default:
			continue
		}
		p.mu.RUnlock()
	}
}

func NewWorkerPool(numWorkers int, channelSize int) (*MyPool, error) {
	if numWorkers <= 0 {
		return nil, errors.New("non-positive number of workers")
	}
	if channelSize < 0 {
		return nil, errors.New("negative size of the channel")
	}

	var p *MyPool
	if channelSize == 0 {
		p = &MyPool{
			tasks:         make(chan PoolTask),
			numberWorkers: numWorkers,
		}
	} else {
		p = &MyPool{
			tasks:         make(chan PoolTask, channelSize),
			numberWorkers: numWorkers,
		}
	}

	return p, nil
}
