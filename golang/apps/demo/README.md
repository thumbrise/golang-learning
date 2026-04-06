# Demo — High-Throughput Comment Service

Go backend service demonstrating **event-driven architecture**, **Redis Streams buffering** and **full-stack OpenTelemetry observability** — built for high write throughput.

## Architecture

```
                          ┌──────────────┐
  HTTP POST ──────────────▶  Gin Server  │
  /api/comments/publish   │  (OTel mid.) │
                          └──────┬───────┘
                                 │ TxPipeline
                          ┌──────▼───────┐
                          │ Redis Stream │  ◄── write buffer
                          │ + Hash cache │      (comments_unprocessed)
                          └──────┬───────┘
                                 │ XReadGroup (consumer group)
                          ┌──────▼───────┐
                          │   Batcher    │  ◄── separate process
                          │  (rate-lim.) │      3 000 ops/s budget
                          └──────┬───────┘
                                 │ batch persist (planned)
                          ┌──────▼───────┐
                          │  PostgreSQL  │
                          └──────────────┘
```

**Write path.** The API handler writes a comment into a Redis Stream (`XADD`) and a Hash cache inside a single `TxPipeline` — atomic, no round-trip overhead. The client gets a response immediately; the heavy work (validation, enrichment, DB persist) is deferred.

**Batch consumer.** A dedicated long-running process (`app_comments_batcher`) reads the stream via `XReadGroup` consumer groups. It first drains pending (unacknowledged) messages, then switches to new ones — crash-safe reprocessing out of the box. A `rate.Limiter` (3 000 ops/s) protects downstream storage from burst overload.

**Why Redis Streams, not Kafka.** For this throughput range Redis Streams give the same consumer-group semantics (at-least-once, fan-out, pending tracking) with zero operational overhead — no brokers, no ZooKeeper, no topic partitions to manage.

## Observability

The service is instrumented with **OpenTelemetry SDK** across all three pillars — traces, metrics, logs — plus continuous profiling.

| Signal | Implementation | Export target |
|---|---|---|
| **Traces** | `otel/sdk/trace` → OTLP/gRPC | Uptrace / Jaeger / any OTLP backend |
| **Metrics** | `otel/sdk/metric` → OTLP/gRPC | Uptrace / Prometheus / VictoriaMetrics |
| **Logs** | `slog` → `otelslog` bridge → OTLP/gRPC | Uptrace / Loki |
| **Profiling** | Pyroscope (CPU, alloc, goroutines, mutex, block) | Grafana Pyroscope |
| **Infra metrics** | postgres_exporter, redis_exporter → vmagent remote-write | any Prometheus-compatible TSDB |

### What is measured

- **HTTP middleware** — per-request spans with semantic conventions (`httpconv`), request/response size histograms, duration histogram (custom bucket boundaries), status-aware span status (OK / CLIENT_ERROR / SERVER_ERROR).
- **Redis** — auto-instrumented via `redisotel` (every command gets a child span + latency metric).
- **Batcher worker** — each batch cycle is a traced span; errors are recorded on the span.
- **Business metrics** — `business_users_count` observable gauge, pulled from PostgreSQL on scrape.

### Scraping & remote write

Infrastructure metrics (Postgres, Redis stream depth) are scraped by **vmagent** and remote-written to the OTLP backend — no local Prometheus instance needed.

## Tech Stack

| Layer | Tech |
|---|---|
| Language | Go 1.24 |
| HTTP | Gin, Swagger (swaggo) |
| Streaming | Redis 8 Streams (consumer groups, TxPipeline) |
| Database | PostgreSQL 17, GORM |
| Observability | OpenTelemetry SDK, Pyroscope, vmagent |
| DI | Google Wire |
| CLI | Cobra |
| Load testing | k6 (constant-arrival-rate, 3 000 rps) |
| Infra | Docker Compose (app, batcher, Redis, Postgres, exporters) |

## Project Structure

```
internal/
├── app/                        # config, core kernel
├── bootstrap/                  # runner (errgroup + graceful shutdown), DI container
├── modules/
│   ├── analytics/              # business metrics (OTel gauges)
│   ├── auth/                   # authentication, user repository
│   ├── comments/
│   │   ├── application/
│   │   │   ├── usecases/       # CommentsCommandPublish (Redis TxPipeline)
│   │   │   └── workers/        # CommentsBatcher (XReadGroup consumer)
│   │   └── endpoints/
│   │       ├── cmd/            # CLI: `comments batch`
│   │       └── http/           # REST: POST /api/comments/publish
│   ├── plugins/
│   │   ├── http/               # Gin engine, OTel middleware, HTTP metrics
│   │   ├── observability/      # OTel registrar, resource, tracer/meter/logger providers
│   │   ├── redis/              # client, OTel instrumentation (redisotel)
│   │   ├── database/           # PostgreSQL, GORM
│   │   └── ...
│   └── shared/                 # error mapping, contracts
└── tests/
    └── load/                   # k6 scenarios
```

## Quick Start

```bash
cp .env.example .env
# adjust ports and OTLP endpoint in .env

docker compose up -d

# publish a comment
curl -X POST http://localhost:${DOCKER_PORT_APP}/api/comments/publish \
  -H 'Content-Type: application/json' \
  -d '{"userUUID":"u1","postUUID":"p1","content":"hello"}'

# run load test (3 000 rps, 30s)
docker compose run --rm k6 run api/comments/publish.ts
```

## License

MIT
