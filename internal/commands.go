package internal

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"go.loki.codes/concourse-resource/types"
)

func RunCheck(resource any, setVersion func(any) error, validate *validator.Validate) (types.CheckOutput, error) {
	checkRes, ok := resource.(types.CheckResource)
	if !ok {
		return nil, fmt.Errorf("`check` mode not supported")
	}

	verPtr := checkRes.VersionPtr()
	if err := setVersion(verPtr); err != nil {
		return nil, err
	}

	if err := validate.Struct(verPtr); err != nil {
		return nil, err
	}

	return checkRes.Check()
}

func RunIn(resource any, path string, setVersion, setParams func(any) error, validate *validator.Validate) (*types.InOutput, error) {
	inRes, ok := resource.(types.InResource)
	if !ok {
		return nil, fmt.Errorf("`in` mode not supported")
	}

	verPtr := inRes.VersionPtr()
	if err := setVersion(verPtr); err != nil {
		return nil, err
	}

	if err := validate.Struct(verPtr); err != nil {
		return nil, err
	}

	if pRes, ok := resource.(types.ParametrizedResource); ok {
		paramsPtr := pRes.ParamsPtr()
		if err := setParams(paramsPtr); err != nil {
			return nil, err
		}

		if err := validate.Struct(paramsPtr); err != nil {
			return nil, err
		}
	}
	return inRes.In(path)
}

func RunOut(resource any, path string, setParams func(any) error, validate *validator.Validate) (*types.OutOutput, error) {
	outRes, ok := resource.(types.OutResource)
	if !ok {
		return nil, fmt.Errorf("`out` mode not supported")
	}

	if pRes, ok := resource.(types.ParametrizedResource); ok {
		paramsPtr := pRes.ParamsPtr()
		if err := setParams(paramsPtr); err != nil {
			return nil, err
		}

		if err := validate.Struct(paramsPtr); err != nil {
			return nil, err
		}
	}
	return outRes.Out(path)
}
