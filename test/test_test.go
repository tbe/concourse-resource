package test

import (
	"testing"

	"pkg.loki.codes/concourse-resource/test/dummy"
	"pkg.loki.codes/concourse-resource/types"
)

func TestHandler_RunCheck(t *testing.T) {
	h := NewHandler[dummy.Config](t, dummy.New)
	testCase := CheckCase[dummy.Config]{
		Case: Case{
			ShouldFail:  false,
			ErrorString: "",
			Validation:  nil,
		},
		Input: CheckInput[dummy.Config]{
			Source:  dummy.Config{},
			Version: dummy.VersionExpected,
		},
		Output: dummy.CheckOutput,
	}
	h.RunCheck(testCase)
}

func TestHandler_RunIn(t *testing.T) {
	h := NewHandler[dummy.Config](t, dummy.New)
	testCase := InCase[dummy.Config]{
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
	}
	h.RunIn(testCase)
}

func TestHandler_RunOut(t *testing.T) {
	h := NewHandler[dummy.Config](t, dummy.New)
	testCase := OutCase[dummy.Config]{
		Case: Case{
			ShouldFail:  false,
			ErrorString: "",
			Validation:  nil,
		},
		Input: OutInput[dummy.Config]{
			Source: dummy.Config{},
			Params: dummy.ParamsExpected,
		},
		Output: types.OutOutput(dummy.InOutput),
	}
	h.RunOut(testCase)
}
