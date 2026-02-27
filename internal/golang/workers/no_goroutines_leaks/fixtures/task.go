package fixtures

import (
	"fmt"

	"github.com/thumbrise/golang-learning/internal/golang/workers/no_goroutines_leaks/task"
)

func TasksChannel(tasks chan *task.Task) chan *task.Task {
	for i := range cap(tasks) {
		msg := fmt.Sprintf("Hello there %d", i)
		tasks <- &task.Task{Message: msg}
	}

	return tasks
}
