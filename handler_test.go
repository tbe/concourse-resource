package concourse

// TODO: we need to move stdin/stdout to a sub package. Somewhere in internal, to break this cycle
import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"pkg.loki.codes/concourse-resource/internal"
	"pkg.loki.codes/concourse-resource/test"
	"pkg.loki.codes/concourse-resource/test/dummy"
)

var resStdin = new(bytes.Buffer)
var resStdout = new(bytes.Buffer)

func init() {
	internal.StdIn = resStdin
	internal.StdOut = resStdout
}

func TestConfig(t *testing.T) {
	resStdin.Reset()
	resStdout.Reset()

	assert := assert.New(t)

	resStdin.Write([]byte(fmt.Sprintf(`{%s}`, dummy.SourceFixture)))
	getCommand = func() string {
		return "check"
	}

	var tRes any

	h := New[dummy.Config](func(config dummy.Config) (res any, err error) {
		res, err = dummy.New(config)
		tRes = res
		return
	})

	_ = h.Run()
	tRes.(test.MockResource[dummy.Config]).Validate(assert, dummy.SourceExpected)
}

func TestCheck(t *testing.T) {
	resStdin.Reset()
	resStdout.Reset()

	assert := assert.New(t)

	resStdin.Write([]byte(fmt.Sprintf(`{%s,  %s}`, dummy.SourceFixture, dummy.VersionFixture)))
	getCommand = func() string {
		return "check"
	}

	var tRes any

	h := New[dummy.Config](func(config dummy.Config) (res any, err error) {
		res, err = dummy.New(config)
		tRes = res
		return
	})

	if !assert.NoError(h.Run()) {
		return
	}
	tRes.(test.MockCheckResource).ValidateCheck(assert, dummy.VersionExpected)
	assert.JSONEq(`[{"ref":"61cbef"},{"ref":"d74e01"},{"ref":"7154fe"}]`, resStdout.String())
}

func TestIn(t *testing.T) {
	resStdin.Reset()
	resStdout.Reset()

	assert := assert.New(t)

	resStdin.Write([]byte(fmt.Sprintf(`{%s, %s, %s}`, dummy.SourceFixture, dummy.VersionFixture, dummy.ParamsFixture)))
	getCommand = func() string {
		return "in"
	}

	getArg = func() string {
		return "/dummy"
	}

	var tRes any

	h := New[dummy.Config](func(config dummy.Config) (res any, err error) {
		res, err = dummy.New(config)
		tRes = res
		return
	})

	if !assert.NoError(h.Run()) {
		return
	}
	tRes.(test.MockInResource).ValidateIn(assert, dummy.VersionExpected, dummy.ParamsExpected)
	assert.JSONEq(`{"metadata":{"test":"value"},"version":{"ref":"61cbef"}}`, resStdout.String())
}

func TestOut(t *testing.T) {
	resStdin.Reset()
	resStdout.Reset()

	assert := assert.New(t)

	resStdin.Write([]byte(fmt.Sprintf(`{%s, %s}`, dummy.SourceFixture, dummy.ParamsFixture)))
	getCommand = func() string {
		return "out"
	}

	getArg = func() string {
		return "/dummy"
	}

	var tRes any

	h := New[dummy.Config](func(config dummy.Config) (res any, err error) {
		res, err = dummy.New(config)
		tRes = res
		return
	})

	if !assert.NoError(h.Run()) {
		return
	}
	tRes.(test.MockOutResource).ValidateOut(assert, dummy.ParamsExpected)
	assert.JSONEq(`{"metadata":{"test":"value"},"version":{"ref":"61cbef"}}`, resStdout.String())
}
