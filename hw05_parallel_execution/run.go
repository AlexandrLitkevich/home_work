package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	if len(tasks) == 0 || tasks == nil {
		return nil
	}

	handlerTasks := make(chan Task)

	doneCh := make(chan struct{})
	var errCount int64
	var wg sync.WaitGroup

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()

			for {
				select {
				case <-doneCh:
					return
				case t, ok := <-handlerTasks:
					if !ok {
						return
					}
					if err := t(); err != nil {
						atomic.AddInt64(&errCount, 1)
					}
				}
			}
		}()
	}

	for _, task := range tasks {
		if atomic.LoadInt64(&errCount) >= int64(m) {
			break
		}
		handlerTasks <- task
	}
	close(doneCh)
	wg.Wait()

	if atomic.LoadInt64(&errCount) >= int64(m) {
		return ErrErrorsLimitExceeded
	}

	return nil
}
