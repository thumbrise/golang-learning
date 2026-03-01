//go:build test

package e2e_test

import (
	"testing"

	"github.com/steinfletcher/apitest"
	"github.com/steinfletcher/apitest-jsonpath"
	"github.com/thumbrise/demo/golang-demo/internal/tests/util"
)

func TestAuthSignIn(t *testing.T) {
	expectedJson := util.MustToJson(map[string]interface{}{
		"message": "Check you email",
	})

	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(util.Handler(t.Context())).
		Observe(util.ObserveResponse(t)).
		Post(util.Uri("api/auth/sign-in")).
		JSON(map[string]string{
			"username": "user",
			"email":    "test@test.test",
		}).
		Expect(t).
		Status(200).
		Body(expectedJson).
		Assert(jsonpath.Present(`$.message`)).
		End()
}
