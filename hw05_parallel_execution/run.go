package hw05parallelexecution

import (
	"errors"
	"fmt"
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
	var handlerTasks = make(chan Task)
	var doneCh = make(chan struct{})
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
					fmt.Println("this run task")
					if err := t(); err != nil {
						fmt.Println("this error", err)
						atomic.AddInt64(&errCount, 1)
					}
				}

			}
		}()
	}

	for _, task := range tasks {
		if atomic.LoadInt64(&errCount) >= int64(m) {
			close(doneCh)
			return ErrErrorsLimitExceeded
		}
		handlerTasks <- task
	}
	close(handlerTasks)
	wg.Wait()

	return nil
}
