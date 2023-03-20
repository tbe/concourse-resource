package types

// InOutput is the result of an `in` operation
type InOutput struct {
	Version  any            `json:"version" validate:"required"`
	Metadata map[string]any `json:"metadata,omitempty"`
}

// OutOutput is the result of an `out` operation
type OutOutput struct {
	Version  any            `json:"version" validate:"required"`
	Metadata map[string]any `json:"metadata,omitempty"`
}

// CheckOutput is the result of an `out` operation
type CheckOutput []any

// A NewResource function creates a resource object based on the given config
type NewResource[T any] func(config T) (any, error)

// A CheckResource is able to handle the `check` step
type CheckResource interface {
	// VersionPtr must return a pointer to the target version struct.
	VersionPtr() any
	// Check implements the `check` logic
	Check() (CheckOutput, error)
}

// A InResource is able to handle the `in` step
type InResource interface {
	// VersionPtr must return a pointer to the target version struct.
	VersionPtr() any

	// In implements the `in` logic
	In(targetDir string) (*InOutput, error)
}

// A OutResource os able to handle the `out` step
type OutResource interface {
	// Out implements the `out` logic
	Out(sourceDir string) (*OutOutput, error)
}

// A ParametrizedResource takes optional parameters to in and out
type ParametrizedResource interface {
	// ParamsPtr must return a pointer to the target parameters struct
	ParamsPtr() any
}
