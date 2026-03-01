package tracer

import "fmt"

type ErrorHandler struct{}

func (l ErrorHandler) Handle(err error) {
	fmt.Printf("OpenTelemetry error: %v\n", err)
}
