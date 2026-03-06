package histogram

import "math"

// DefBuckets returns the default bucket boundaries suitable for typical
// HTTP/gRPC request latencies (from 5ms to 10s).
//
// Literally:
//
// 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10
//
// or
//
// 5ms, 10ms, 25ms, 50ms, 100ms, 250ms, 500ms, 1s, 2.5s, 5s, 10s
//
// These are the same bucket boundaries that Prometheus uses as default
// for histograms. They are empirically chosen to cover the range from very fast
// (cached) requests up to slow ones, with logarithmic scale giving roughly
// equal relative resolution across the range.
//
// When to use:
//   - You are measuring request duration, operation latency, or any other
//     metric that typically falls within 1ms to 10s range.
//   - You want a well-understood, battle-tested bucket layout that works
//     for most services.
//   - You need compatibility with existing dashboards that expect these
//     boundaries.
//
// When NOT to use:
//   - Your values are extremely small (microseconds) or extremely large
//     (minutes/hours). In such cases, custom boundaries are required.
//   - You need higher precision in a specific region (e.g., between 100-200ms).
//   - You are measuring non-time metrics like bytes, counts, etc.
func DefBuckets() []float64 {
	return []float64{
		0.005, // 5ms
		0.01,  // 10ms
		0.025, // 25ms
		0.05,  // 50ms
		0.1,   // 100ms
		0.25,  // 250ms
		0.5,   // 500ms
		1.0,   // 1s
		2.5,   // 2.5s
		5.0,   // 5s
		10.0,  // 10s
	}
}

// ExponentialBuckets generates bucket boundaries that grow exponentially
// by a fixed factor.
//
// Parameters:
//   - start: The first bucket boundary (must be > 0).
//   - factor: The multiplier applied to each subsequent boundary (must be > 1).
//   - count: Number of buckets to generate (must be >= 1).
//
// Example: ExponentialBuckets(100, 10, 4) -> [100, 1000, 10000, 100000]
//
// When to use:
//   - You need to cover a wide range of magnitudes (e.g., file sizes from
//     bytes to gigabytes).
//   - The data spans several orders of magnitude and you want roughly
//     equal relative precision across the range.
//   - You don't know the exact distribution but know the scale (e.g.,
//     response size can be 1KB, 10KB, 100KB, 1MB, 10MB...).
//
// When NOT to use:
//   - You need very fine control over specific regions.
//   - Your data is concentrated in a narrow range – linear buckets may be
//     more appropriate.
//   - Factor is too small (close to 1) – this will generate many buckets
//     and increase storage cost.
func ExponentialBuckets(start, factor float64, count int) []float64 {
	if start <= 0 {
		panic("start must be positive")
	}

	if factor <= 1 {
		panic("factor must be greater than 1")
	}

	if count < 1 {
		panic("count must be at least 1")
	}

	buckets := make([]float64, count)

	buckets[0] = start
	for i := 1; i < count; i++ {
		buckets[i] = buckets[i-1] * factor
	}

	return buckets
}

// LogarithmicBuckets generates bucket boundaries using a logarithmic scale.
// This is similar to exponential buckets but expressed as powers of a base.
//
// Parameters:
//   - base: The logarithmic base (must be > 1). Common choices: 2, 10, e.
//   - minPower: The smallest exponent (start = base^minPower).
//   - maxPower: The largest exponent (inclusive).
//
// Example: LogarithmicBuckets(10, 2, 6) -> [100, 1000, 10000, 100000, 1e6]
//
//	because 10^2 = 100, 10^3 = 1000, ..., 10^6 = 1,000,000.
//
// When to use:
//   - You want mathematically clean boundaries based on powers (10^3, 10^4...).
//   - You are measuring values that naturally follow a power law
//     (e.g., disk I/O sizes, network packet sizes).
//   - You need to align with human-readable units (e.g., 1KB, 10KB, 100KB, 1MB).
//
// When NOT to use:
//   - base is too large – you'll lose resolution in the low range.
//   - Your data doesn't span multiple orders of magnitude – linear or
//     fine-grained exponential may be better.
//   - You need non-power boundaries (e.g., 0.5, 1, 2, 4 – use exponential
//     with factor 2 instead).
func LogarithmicBuckets(base float64, minPower, maxPower int) []float64 {
	if base <= 1 {
		panic("base must be greater than 1")
	}

	if minPower > maxPower {
		panic("minPower must be less than or equal to maxPower")
	}

	count := maxPower - minPower + 1

	buckets := make([]float64, count)
	for i := range count {
		power := minPower + i
		buckets[i] = math.Pow(base, float64(power))
	}

	return buckets
}
