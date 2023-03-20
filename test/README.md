# Testing

This subpackage provides a test suite for resources.

## Single Tests

The `Handler` provides a way to execute single test cases. This may be useful if you need the full control over the test
flow.


```go
package main

import (
	"testing"
	
	"pkg.loki.codes/concourse-resource/test"

	
	"your/resource/implementation"
)

func TestResource_Check(t *testing.T) {
	h := test.NewHandler[implementation.Config](t, implementation.New)
	testCase := test.CheckCase[implementation.Config]{
		Case: test.Case{
			ShouldFail:  false,
			ErrorString: "" /* of ShouldFail is true, compare against this error string */,
			Validation:  nil /* optional validation function */,
		},
		Input: test.CheckInput[implementation.Config]{
			Source:  implementation.Config{/* config */},
			Version: /* your version definition */,
		},
		Output: /* the expected output */,
	}
	h.RunCheck(testCase)
}
```

## Multiple Tests

For convenience, a `suite.Suit` is provided that tests all resource functions. 

```go
package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"pkg.loki.codes/concourse-resource/test"

	"your/resource/implementation"
)

func TestSuite(t *testing.T) {
	s := NewSuite[implementation.Config](implementation.New)

	s.SetInCases(map[string]InCase[implementation.Config]{
		"default": test.CheckCase[implementation.Config]{
			Case: test.Case{
				ShouldFail:  false,
				ErrorString: "" /* of ShouldFail is true, compare against this error string */,
				Validation:  nil /* optional validation function */,
			},
			Input: test.CheckInput[implementation.Config]{
				Source:  implementation.Config{/* config */},
				Version: /* your version definition */,
			},
			Output: /* the expected output */,
		},
	})
	suite.Run(t, s)
}

```

## Examples

Have a look at the `*_test.go` files for more details 