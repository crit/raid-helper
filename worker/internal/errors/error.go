package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"
)

type (
	// Error is a standard error with extra information.
	Error struct {
		// Code is a http status code indicating the severity of the error. Allows handling code to do different things
		// with this returned error.
		Code int `json:"code"`

		// Message is the human-readable error message.
		Message string `json:"message"`

		// File is the origin folder and file of the error when New was called.
		File string `json:"file"`

		// Line is the origin file line of the error when New was called.
		Line int `json:"line"`
	}
)

// Error satisfies the error interface.
func (e Error) Error() string {
	return fmt.Sprintf("%s:%d %s", e.File, e.Line, e.Message)
}

// New creates a new Error backed error while recording the file and line of the caller.
func New(code int, format string, args ...any) error {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}

	return Error{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
		File:    srcFileParse(file),
		Line:    line,
	}
}

func Unwrap(err error) Error {
	if err == nil {
		return Error{}
	}

	e, ok := err.(Error)
	if !ok {
		return Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			File:    "",
			Line:    0,
		}
	}

	return e
}

func Wrap(err error) error {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}

	return Error{
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
		File:    srcFileParse(file),
		Line:    line,
	}
}

// Json attempts to cast a standard error to a json string version of Error. Output will always have a "message" property
// which is a blank string for nil.
func Json(err error) string {
	if err == nil {
		return `{"message": ""}`
	}

	e, ok := err.(Error)
	if !ok {
		return fmt.Sprintf(`{"message": "%s"}`, err.Error())
	}

	raw, jErr := json.Marshal(e)
	if jErr != nil {
		return fmt.Sprintf(`{"message": "%s"}`, e.Message)
	}

	return string(raw)
}

// Code returns the recorded Error.Code, OK for nil, and 500 for a regular error.
func Code(err error) int {
	if err == nil {
		return http.StatusOK
	}

	e, ok := err.(Error)
	if !ok {
		return http.StatusInternalServerError
	}

	return e.Code
}

// Message returns the recorded Error.Message, blank for nil, and string output for a regular error.
func Message(err error) string {
	if err == nil {
		return ""
	}

	e, ok := err.(Error)
	if !ok {
		return err.Error()
	}

	return e.Message
}

// File returns the recorded Error.File, blank for nil and a regular error.
func File(err error) string {
	if err == nil {
		return ""
	}

	e, ok := err.(Error)
	if !ok {
		return ""
	}

	return e.File
}

// Line returns the recorded Error.Line, 0 for nil and a regular error.
func Line(err error) int {
	if err == nil {
		return 0
	}

	e, ok := err.(Error)
	if !ok {
		return 0
	}

	return e.Line
}

// srcFileParse returns either the filename and extension, or the last directory
// (which is also usually the package name in Go) with the filename and extension.
// "project/src/model/user.go" => "model/user.go"
// "main.go" => "main.go
func srcFileParse(filename string) string {
	// "project/src/model/user.go" => "project/src/model", "user.go"
	dir, file := filepath.Split(filename)

	// "project/src/model" => ["project", "src", "model"]
	parts := strings.FieldsFunc(dir, func(c rune) bool {
		return c == filepath.Separator
	})

	if len(parts) > 0 {
		// => "model/user.go"
		return filepath.Join(parts[len(parts)-1], file)
	}

	// "user.go"
	return file
}
