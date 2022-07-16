package errors_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/crit/raid-helper/worker/internal/errors"
	"github.com/stretchr/testify/assert"
)

func TestError_Error(t *testing.T) {
	var err errors.Error

	err.Message = "oh no!"
	err.File = "error/error_test.go"
	err.Line = 6
	err.Code = 500

	assert.Equal(t, "error/error_test.go:6 oh no!", err.Error())
}

func TestNew(t *testing.T) {
	err := errors.New(http.StatusNotImplemented, "oh no: %s", "test")

	assert.Equal(t, http.StatusNotImplemented, errors.Code(err))
	assert.Equal(t, "oh no: test", errors.Message(err))
	assert.Equal(t, "errors/error_test.go", errors.File(err))
	assert.Equal(t, 24, errors.Line(err))
}

func TestJson(t *testing.T) {
	err := errors.New(http.StatusNotImplemented, "oh no: %s", "test")

	assert.Equal(t, `{"code":501,"message":"oh no: test","file":"errors/error_test.go","line":33}`, errors.Json(err))
}

func TestUnwrap(t *testing.T) {
	a := errors.Error{
		Code:    http.StatusNotImplemented,
		Message: "oh no: test",
		File:    "errors/error_test.go",
		Line:    46,
	}

	err := errors.New(http.StatusNotImplemented, "oh no: %s", "test")

	assert.Equal(t, a, errors.Unwrap(err))

	err = fmt.Errorf("oh no")
	b := errors.Error{
		Code:    http.StatusInternalServerError,
		Message: "oh no",
		File:    "",
		Line:    0,
	}

	assert.Equal(t, b, errors.Unwrap(err))
}
