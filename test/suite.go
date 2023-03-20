package test

import (
	"github.com/stretchr/testify/suite"
	resource "pkg.loki.codes/concourse-resource/types"
)

// NewSuite creates a new suite.Suite
func NewSuite[cfgType any](constructor resource.NewResource[cfgType]) *ResourceTestSuite[cfgType] {
	return &ResourceTestSuite[cfgType]{
		handler: &Handler[cfgType]{
			constructor: constructor,
		},
	}
}

// ResourceTestSuite provides the ability to run multiple tests, simply by providing maps of test cases.
//
// If no case is given for a test type, the subtest will return a SUCCESS status
type ResourceTestSuite[cfgType any] struct {
	suite.Suite
	configCases map[string]ConfigCase[cfgType]
	checkCases  map[string]CheckCase[cfgType]
	inCases     map[string]InCase[cfgType]
	outCases    map[string]OutCase[cfgType]

	handler *Handler[cfgType]
}

// SetConfigCases sets the list of ConfigCase tests
func (s *ResourceTestSuite[cfgType]) SetConfigCases(configCases map[string]ConfigCase[cfgType]) {
	s.configCases = configCases
}

// SetCheckCases sets the list of CheckCase tests
func (s *ResourceTestSuite[cfgType]) SetCheckCases(checkCases map[string]CheckCase[cfgType]) {
	s.checkCases = checkCases
}

// SetInCases sets the list of InCase tests
func (s *ResourceTestSuite[cfgType]) SetInCases(inCases map[string]InCase[cfgType]) {
	s.inCases = inCases
}

// SetOutCases sets the list of OutCase tests
func (s *ResourceTestSuite[cfgType]) SetOutCases(outCases map[string]OutCase[cfgType]) {
	s.outCases = outCases
}

func (s *ResourceTestSuite[cfgType]) SetupSuite() {
	s.handler.assert = s.Assert()
}

func (s *ResourceTestSuite[cfgType]) TestConfig() {
	for name, c := range s.configCases {
		s.Run(name, func() {
			s.handler.RunConfig(c)
		})
	}
}

func (s *ResourceTestSuite[cfgType]) TestCheck() {
	for name, c := range s.checkCases {
		s.Run(name, func() {
			s.handler.RunCheck(c)
		})
	}
}

func (s *ResourceTestSuite[cfgType]) TestIn() {
	for name, c := range s.inCases {
		s.Run(name, func() {
			s.handler.RunIn(c)
		})
	}
}

func (s *ResourceTestSuite[cfgType]) TestOut() {
	for name, c := range s.outCases {
		s.Run(name, func() {
			s.handler.RunOut(c)
		})
	}
}
