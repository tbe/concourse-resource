package test

import (
	"github.com/stretchr/testify/assert"
	"pkg.loki.codes/concourse-resource/types"
)

// CheckInput is a dummy input for a `check` test
type CheckInput[T any] struct {
	Source  T
	Version any
}

// InInput is a dummy input for a `in` test
type InInput[cfgType any] struct {
	Source  cfgType
	Version any
	Params  any
}

// OutInput is a dummy input for a `out` test
type OutInput[cfgType any] struct {
	Source cfgType
	Params any
}

// Case defines the input and expected result of a single test
type Case struct {
	ShouldFail  bool                                               // defines if the testcase should fail
	ErrorString string                                             // an optional error message to with
	Validation  func(assert *assert.Assertions, resource any) bool // an optional validation function
}

// ConfigCase defines a testcase for the construction of a new resource
type ConfigCase[cfgType any] struct {
	Case
	Input cfgType
}

// CheckCase defines a testcase for the `check` action
type CheckCase[cfgType any] struct {
	Case
	Input  CheckInput[cfgType]
	Output types.CheckOutput
}

// InCase defines a testcase for the `in` action
type InCase[cfgType any] struct {
	Case
	Input  InInput[cfgType]
	Output types.InOutput
}

// OutCase defines a testcase for the `out` action
type OutCase[cfgType any] struct {
	Case
	Input  OutInput[cfgType]
	Output types.OutOutput
}
