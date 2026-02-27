// Package task
// Слой бизнес логики, операция не переживающая о сложной настройке горутин и всего остального
package task

import (
	"context"
	"fmt"
)

type Processor struct{}

func NewProcessor() *Processor {
	return &Processor{}
}

func (p *Processor) Process(_ context.Context, task *Task) (string, error) {
	// Some business logic
	return fmt.Sprintf("Message is = <%s>", task.Message), nil
}
