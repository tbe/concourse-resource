package test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"pkg.loki.codes/concourse-resource/test/dummy"
)

func TestSuite(t *testing.T) {
	s := NewSuite[dummy.Config](dummy.New)

	s.SetInCases(map[string]InCase[dummy.Config]{
		"default": {
			Case: Case{
				ShouldFail:  false,
				ErrorString: "",
				Validation:  nil,
			},
			Input: InInput[dummy.Config]{
				Source:  dummy.Config{},
				Version: dummy.VersionExpected,
				Params:  dummy.ParamsExpected,
			},
			Output: dummy.InOutput,
		},
	})
	suite.Run(t, s)
}
