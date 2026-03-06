package logger

import (
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
)

func NewOTELSDKProvider(res *resource.Resource) *sdklog.LoggerProvider {
	return sdklog.NewLoggerProvider(
		sdklog.WithResource(res),
		// sdklog.WithProcessor(exp,
		//	batchOptions()...,
		// ),
		// sdklog.WithSampler(sampler),
	)
}
