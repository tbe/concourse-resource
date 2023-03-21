package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"

	_ "pkg.loki.codes/concourse-resource/test"
)

func TestBuildURL(t *testing.T) {
	assert.Equal(t, "http://example.com/teams/main/pipelines/testPipeline/jobs/testJob/builds/build42", BuildURL())
}
