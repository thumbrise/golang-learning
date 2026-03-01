package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/thumbrise/demo/golang/internal/golang/workers/no_goroutines_leaks/fixtures"
	"github.com/thumbrise/demo/golang/internal/golang/workers/no_goroutines_leaks/task"
	"github.com/thumbrise/demo/golang/internal/golang/workers/no_goroutines_leaks/workers"
)

const (
	workersCount = 100
	tasksCount   = 10000
	timeout      = time.Second * 10
)

// Слой верхнего уровня
func main() {
	tasks := fixtures.TasksChannel(
		make(chan *task.Task, tasksCount),
	)

	// Graceful Shutdown - Обработка сигналов
	ctx, cancel := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
	defer cancel()

	consumer := workers.NewTaskConsumer(
		task.NewProcessor(),
		timeout,
		workersCount,
	)

	if err := consumer.Run(ctx, tasks); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			err = fmt.Errorf("timeout exceeded (%d): %w", timeout, err)
		}

		fmt.Printf("workers fail: %s", err)
	}
}
