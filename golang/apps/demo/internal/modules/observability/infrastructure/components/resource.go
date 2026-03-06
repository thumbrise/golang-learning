package components

import (
	"context"
	"errors"
	"fmt"

	"github.com/thumbrise/demo/golang-demo/internal/app"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

var ErrResourceNew = errors.New("failed to create resource")

func NewResource(ctx context.Context, cfgApp app.Config) (*resource.Resource, error) {
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(cfgApp.Name),
			semconv.ServiceVersion(cfgApp.Version),
			semconv.DeploymentEnvironment(cfgApp.Environment),
		),
		resource.WithFromEnv(),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
		resource.WithContainer(),
	)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrResourceNew, err)
	}

	return res, err
}
