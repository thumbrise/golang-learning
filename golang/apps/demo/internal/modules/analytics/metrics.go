package analytics

import (
	"context"

	"github.com/thumbrise/demo/golang-demo/internal/modules/auth/infrastructure/dal"
	"github.com/thumbrise/demo/golang-demo/internal/modules/plugins/observability/components/meter"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

type Metrics struct {
	meterProvider  *meter.Provider
	userRepository *dal.UserRepository
	usersTotal     metric.Int64ObservableGauge
}

func NewMetrics(p *meter.Provider, repository *dal.UserRepository) *Metrics {
	return &Metrics{meterProvider: p, userRepository: repository}
}

func (m *Metrics) GaugeUsersTotal() {
	mtr := m.meterProvider.Meter()

	var err error

	m.usersTotal, err = mtr.Int64ObservableGauge(
		"business_users_count",
		metric.WithDescription("Total registered users"),
		metric.WithUnit("users"),
		metric.WithInt64Callback(func(ctx context.Context, observer metric.Int64Observer) error {
			count, err := m.userRepository.Count(ctx)
			if err != nil {
				return err
			}

			observer.Observe(int64(count))

			return nil
		}),
	)
	if err != nil {
		otel.Handle(err)
	}
}
