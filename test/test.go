package test

import (
	"bytes"
	"os"
	"reflect"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"pkg.loki.codes/concourse-resource/internal"
	resource "pkg.loki.codes/concourse-resource/types"
)

var resStdin = new(bytes.Buffer)
var resStdout = new(bytes.Buffer)

func init() {
	// we set up our environment variables required for testing
	_ = os.Setenv("ATC_EXTERNAL_URL", "http://example.com")
	_ = os.Setenv("BUILD_TEAM_NAME", "main")
	_ = os.Setenv("BUILD_PIPELINE_NAME", "testPipeline")
	_ = os.Setenv("BUILD_JOB_NAME", "testJob")
	_ = os.Setenv("BUILD_NAME", "build42")
	_ = os.Setenv("BUILD_ID", "42")

	internal.StdIn = resStdin
	internal.StdOut = resStdout
}

// A Handler for your tests.
type Handler[T any] struct {
	constructor resource.NewResource[T]
	assert      *assert.Assertions
}

// NewHandler creates a new testing Handler
func NewHandler[T any](t *testing.T, constructor resource.NewResource[T]) *Handler[T] {
	return &Handler[T]{
		constructor: constructor,
		assert:      assert.New(t),
	}
}

// RunConfig executes a single ConfigCase
func (h *Handler[T]) RunConfig(c ConfigCase[T]) bool {
	res, err := h.constructor(c.Input)
	if c.ShouldFail {
		return h.assert.Error(err) && h.assert.Equal(c.ErrorString, err.Error())
	}
	if c.Validation != nil {
		return c.Validation(h.assert, res)
	}
	return true
}

// RunCheck executes a single CheckCase
func (h *Handler[T]) RunCheck(c CheckCase[T]) bool {
	res, err := h.constructor(c.Input.Source)
	if !h.assert.NoError(err) {
		return false
	}
	checkRes, err := internal.RunCheck(res, func(verPtr any) error {
		v := reflect.ValueOf(verPtr)
		d := reflect.ValueOf(c.Input.Version)

		v.Elem().Set(d)
		return nil
	}, validator.New())
	return h.check(c.Case, err, c.Output, checkRes)
}

// RunIn executes a single InCase
func (h *Handler[T]) RunIn(c InCase[T]) bool {
	res, err := h.constructor(c.Input.Source)
	if !h.assert.NoError(err) {
		return false
	}
	checkRes, err := internal.RunIn(res, "/dummy", func(verPtr any) error {
		v := reflect.ValueOf(verPtr)
		d := reflect.ValueOf(c.Input.Version)

		v.Elem().Set(d)
		return nil
	}, func(paramsPtr any) error {
		v := reflect.ValueOf(paramsPtr)
		d := reflect.ValueOf(c.Input.Params)

		v.Elem().Set(d)
		return nil
	}, validator.New())
	return h.check(c.Case, err, &c.Output, checkRes)
}

// RunOut executes a single OutCase
func (h *Handler[T]) RunOut(c OutCase[T]) bool {
	res, err := h.constructor(c.Input.Source)
	if !h.assert.NoError(err) {
		return false
	}
	checkRes, err := internal.RunOut(res, "/dummy", func(paramsPtr any) error {
		v := reflect.ValueOf(paramsPtr)
		d := reflect.ValueOf(c.Input.Params)

		v.Elem().Set(d)
		return nil
	}, validator.New())
	return h.check(c.Case, err, &c.Output, checkRes)
}

func (h *Handler[T]) check(c Case, err error, expected any, res any) bool {
	if c.ShouldFail {
		return h.assert.Error(err) && h.assert.Equal(c.ErrorString, err.Error())
	}
	if !h.assert.Equal(expected, res) {
		return false
	}

	if c.Validation != nil {
		return c.Validation(h.assert, res)
	}
	return true
}
