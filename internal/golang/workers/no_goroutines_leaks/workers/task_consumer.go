// Package workers
// Слой посредник инкапсулирующий чтение канала
// Не знает про мейн, не знает точной бизнес логики
package workers

import (
	"context"
	"fmt"
	"time"

	"github.com/thumbrise/golang-learning/internal/golang/workers/no_goroutines_leaks/task"
	"golang.org/x/sync/errgroup"
)

type TaskConsumer struct {
	processor    *task.Processor
	workersCount int
	timeout      time.Duration
}

func NewTaskConsumer(processor *task.Processor, timeout time.Duration, workersCount int) *TaskConsumer {
	return &TaskConsumer{processor: processor, timeout: timeout, workersCount: workersCount}
}

func (t *TaskConsumer) Run(ctx context.Context, tasks chan *task.Task) error {
	grp, ctx := errgroup.WithContext(ctx)
	for i := range t.workersCount {
		grp.Go(func() error {
			name := fmt.Sprintf("Worker %d", i)

			ctx, cancel := context.WithTimeout(ctx, t.timeout)
			defer cancel()

			return t.consume(ctx, name, tasks)
		})
	}

	return grp.Wait()
}

func (t *TaskConsumer) consume(ctx context.Context, name string, tasks chan *task.Task) error {
	fmt.Printf("worker %s is started!\n", name)
	defer fmt.Printf("worker %s is done BYE BYE\n", name)

	for {
		select {
		// Graceful Shutdown - Обработка отмены
		case <-ctx.Done():
			return fmt.Errorf("%s context canceled: %w", name, ctx.Err())
		case tsk, ok := <-tasks:
			if !ok {
				fmt.Printf("%s task channel closed\n", name)

				return nil
			}

			result, err := t.processor.Process(ctx, tsk)
			if err != nil {
				return err
			}

			fmt.Printf("worker %#v process result %#v\n", name, result)
		}
	}
}
