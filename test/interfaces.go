package test

import (
	"github.com/stretchr/testify/assert"
	"pkg.loki.codes/concourse-resource/types"
)

type MockResource[cfgType any] interface {
	Validate(assert *assert.Assertions, expected cfgType) bool
}

type MockCheckResource interface {
	types.InResource
	ValidateCheck(assert *assert.Assertions, expectedVersion any) bool
}

type MockInResource interface {
	types.InResource
	ValidateIn(assert *assert.Assertions, expectedVersion, expectedParams any) bool
}

type MockOutResource interface {
	types.InResource
	ValidateOut(assert *assert.Assertions, expectedParams any) bool
}
