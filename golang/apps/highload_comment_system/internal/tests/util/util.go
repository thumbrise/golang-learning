//go:build test

package util

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/steinfletcher/apitest"
	"github.com/thumbrise/demo/golang-demo/internal/bootstrap/modules"
	http2 "github.com/thumbrise/demo/golang-demo/internal/modules/shared/http"
	"github.com/thumbrise/demo/golang-demo/pkg/env"
)

var cfg *http2.Config

func Uri(uri string) string {
	if cfg == nil {
		c := http2.NewConfig(env.NewLoader())
		cfg = &c
	}

	port := cfg.Port
	//nolint:godox
	// TODO: Вынести в http конфиг
	host := "http://localhost"

	uri = strings.TrimLeft(uri, "/")
	host = strings.TrimRight(host, "/")

	result := host + ":" + port + "/" + uri

	return result
}

var handler http.Handler

func Handler(ctx context.Context) http.Handler {
	if handler == nil {
		c, err := modules.InitializeContainer(ctx)
		if err != nil {
			panic(err)
		}

		err = c.Boot(ctx)
		if err != nil {
			panic(err)
		}

		handler = c.HttpKernel.Gin().Handler()
	}

	return handler
}

var ErrCantMarshal = errors.New("can't marshal")

func MustToJson(v interface{}) string {
	marshal, err := json.Marshal(v)
	if err != nil {
		log.Fatalf("%s: %s", ErrCantMarshal, err)
	}

	return string(marshal)
}

func ObserveResponse(t *testing.T) func(res *http.Response, req *http.Request, test *apitest.APITest) {
	t.Helper()

	return func(res *http.Response, req *http.Request, test *apitest.APITest) {
		// Захватываем тело ответа
		if res != nil && res.Body != nil {
			bodyBytes, _ := io.ReadAll(res.Body)
			t.Logf("Body:\n %s\n", string(bodyBytes))
		}
	}
}
