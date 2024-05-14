package pool

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestWorkerPool_NewPool(t *testing.T) {
	if _, err := NewWorkerPool(0, 0); err == nil {
		t.Fatalf("expected error when creating pool with 0 workers, got: %v", err)
	}
	if _, err := NewWorkerPool(-1, 0); err == nil {
		t.Fatalf("expected error when creating pool with -1 workers, got: %v", err)
	}
	if _, err := NewWorkerPool(1, -1); err == nil {
		t.Fatalf("expected error when creating pool with -1 channel size, got: %v", err)
	}

	if _, err := NewWorkerPool(5, 0); err != nil {
		t.Fatalf("expected no error creating pool, got: %v", err)
	}
}

func TestPool_MultipleStartStop(t *testing.T) {
	p, err := NewWorkerPool(5, 0)
	if err != nil {
		t.Fatal("error creating pool:", err)
	}

	p.Start()
	p.Start()

	p.Stop()
	p.Stop()
}

type testT struct {
	executeFunc func() error

	shouldErr bool
	wg        *sync.WaitGroup

	mFailure       *sync.Mutex
	failureHandled bool
}

func newTestTask(executeFunc func() error, shouldErr bool, wg *sync.WaitGroup) *testT {
	return &testT{
		executeFunc: executeFunc,
		shouldErr:   shouldErr,
		wg:          wg,
		mFailure:    &sync.Mutex{},
	}
}

func (t *testT) Execute() error {
	if t.wg != nil {
		defer t.wg.Done()
	}

	if t.executeFunc != nil {
		return t.executeFunc()
	}

	time.Sleep(50 * time.Millisecond)
	if t.shouldErr {
		return fmt.Errorf("planned Execute() error")
	}
	return nil
}

func (t *testT) OnFailure(e error) {
	t.mFailure.Lock()
	defer t.mFailure.Unlock()

	t.failureHandled = true
}

func (t *testT) hitFailureCase() bool {
	t.mFailure.Lock()
	defer t.mFailure.Unlock()

	return t.failureHandled
}

func TestPool_Work(t *testing.T) {
	var tasks []*testT
	wg := &sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		tasks = append(tasks, newTestTask(nil, false, wg))
	}

	p, err := NewWorkerPool(5, len(tasks))
	if err != nil {
		t.Fatal("error making worker pool:", err)
	}
	p.Start()

	for _, j := range tasks {
		p.AddWork(j)
	}

	wg.Wait()

	for taskNum, task := range tasks {
		if task.hitFailureCase() {
			t.Fatalf("error function called on task %d when it shouldn't be", taskNum)
		}
	}
}

func TestWorkerPool_WorkWithErrors(t *testing.T) {
	var tasks []*testT
	wg := &sync.WaitGroup{}

	// first 10 workers succeed
	for i := 0; i < 10; i++ {
		wg.Add(1)
		tasks = append(tasks, newTestTask(nil, false, wg))
	}

	// second 10 workers fail
	for i := 0; i < 10; i++ {
		wg.Add(1)
		tasks = append(tasks, newTestTask(nil, true, wg))
	}

	p, err := NewWorkerPool(5, len(tasks))
	if err != nil {
		t.Fatal("error making worker pool:", err)
	}
	p.Start()

	for _, j := range tasks {
		p.AddWork(j)
	}

	// we'll get a timeout failure if the tasks weren't processed
	wg.Wait()

	for taskNum, task := range tasks {
		if task.hitFailureCase() {
			// the first 10 tasks succeed, the second 10 fail
			if taskNum >= 10 {
				continue
			}

			t.Fatalf("error function called on task %d when it shouldn't be", taskNum)
		}
	}
}

func TestWorkerPool_BlockedAddWorkReleaseAfterStop(t *testing.T) {
	p, err := NewWorkerPool(1, 0)
	if err != nil {
		t.Fatal("error making worker pool:", err)
	}

	p.Start()

	wg := &sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		// the first should start processing right away, the second two should hang
		wg.Add(1)
		go func() {
			p.AddWork(newTestTask(func() error {
				time.Sleep(20 * time.Second)
				return nil
			}, false, nil))
			wg.Done()
		}()
	}

	done := make(chan struct{})
	p.Stop()
	go func() {
		// wait on our AddWork calls to complete, then signal on the done channel
		wg.Wait()
		done <- struct{}{}
	}()

	// wait until either we hit our timeout, or we're told the AddWork calls completed
	select {
	case <-time.After(1 * time.Second):
		t.Fatal("failed because still hanging on AddWork")
	case <-done:
		// this is the success case
	}
}
