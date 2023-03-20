package dummy

import (
	"github.com/stretchr/testify/assert"
	"pkg.loki.codes/concourse-resource/types"
)

type Config struct {
	URI        string `json:"uri"`
	Branch     string `json:"branch"`
	PrivateKey string `json:"private_key"`
}

type version struct {
	Ref string `json:"ref"`
}

type params struct {
	Param string `json:"param"`
}

type resource struct {
	cfg  Config
	v    version
	p    params
	path string
}

func New(config Config) (any, error) {
	return &resource{cfg: config}, nil
}

func (r *resource) Validate(assert *assert.Assertions, expected Config) bool {
	return assert.Equal(expected, r.cfg)
}

func (r *resource) ParamsPtr() any {
	return &r.p
}

func (r *resource) VersionPtr() any {
	return &r.v
}

func (r *resource) Check() (types.CheckOutput, error) {
	return types.CheckOutput{
		r.v,
		version{Ref: "d74e01"},
		version{Ref: "7154fe"},
	}, nil
}

func (r *resource) ValidateCheck(assert *assert.Assertions, expectedVersion any) bool {
	return assert.Equal(expectedVersion, r.v)
}

func (r *resource) In(path string) (*types.InOutput, error) {
	r.path = path
	return &types.InOutput{
		Version:  r.v,
		Metadata: map[string]any{"test": "value"},
	}, nil
}

func (r *resource) ValidateIn(assert *assert.Assertions, expectedVersion, expectedParams any) bool {
	return assert.Equal(expectedVersion, r.v) && assert.Equal(expectedParams, r.p)
}

func (r *resource) Out(path string) (*types.OutOutput, error) {
	r.path = path
	return &types.OutOutput{
		Version:  version{Ref: "61cbef"},
		Metadata: map[string]any{"test": "value"},
	}, nil
}

func (r *resource) ValidateOut(assert *assert.Assertions, expectedParams any) bool {
	return assert.Equal(expectedParams, r.p)
}
