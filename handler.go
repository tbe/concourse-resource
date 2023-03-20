package concourse

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/go-playground/validator/v10"
	"pkg.loki.codes/concourse-resource/internal"
	"pkg.loki.codes/concourse-resource/types"
)

var getCommand = func() string {
	return path.Base(os.Args[0])
}

var getArg = func() string {
	if len(os.Args) < 2 {
		panic("missing argument")
	}
	return os.Args[1]
}

// The Handler is responsible for all the communication with concourse
type Handler[cfgType any] struct {
	constructor types.NewResource[cfgType]
	data        inputData[cfgType]
	resource    any

	validate *validator.Validate
}

type inputData[T any] struct {
	Source  T               `json:"source"`
	Version json.RawMessage `json:"version"`
	Params  json.RawMessage `json:"params"`
}

// New creates a new Handler for the given configuration type cfgType
func New[T any](constructor types.NewResource[T]) *Handler[T] {
	return &Handler[T]{
		constructor: constructor,
		validate:    validator.New(),
	}
}

// Run executes the current command.
//
// The command is detected by the current binary name
func (h *Handler[cfgType]) Run() error {
	// read our input
	if err := json.NewDecoder(internal.StdIn).Decode(&h.data); err != nil {
		return err
	}

	// validate the input
	if err := h.validate.Struct(&h.data.Source); err != nil {
		return err
	}

	// input seems fine, create our resource
	r, err := h.constructor(h.data.Source)
	if err != nil {
		return err
	}

	h.resource = r

	switch getCommand() {
	case "check":
		return h.check()
	case "in":
		return h.in()
	case "out":
		return h.out()
	default:
		return fmt.Errorf("unknown mode %v", os.Args[0])
	}
}

func (h *Handler[cfgType]) check() error {
	versions, err := internal.RunCheck(h.resource, func(verPtr any) error {
		return json.Unmarshal(h.data.Version, verPtr)
	}, h.validate)

	if err != nil {
		return err
	}

	// encode the output
	return json.NewEncoder(internal.StdOut).Encode(versions)
}

func (h *Handler[cfgType]) in() error {
	output, err := internal.RunIn(h.resource, getArg(), func(verPtr any) error {
		return json.Unmarshal(h.data.Version, verPtr)
	}, func(paramsPtr any) error {
		return json.Unmarshal(h.data.Params, paramsPtr)
	}, h.validate)
	if err != nil {
		return err
	}

	// encode the output
	return json.NewEncoder(internal.StdOut).Encode(output)
}

func (h *Handler[cfgType]) out() error {
	output, err := internal.RunOut(h.resource, getArg(), func(paramsPtr any) error {
		return json.Unmarshal(h.data.Params, paramsPtr)
	}, h.validate)
	if err != nil {
		return err
	}
	return json.NewEncoder(internal.StdOut).Encode(output)
}
