// Package metadata provides access to the metadata provided by concourse in the environment of
// the resources.
//
// See https://concourse-ci.org/implementing-resource-types.html#resource-metadata
package metadata

import (
	"encoding/json"
	"os"
)

// BuildID returns the ID of the current build
func BuildID() string {
	return os.Getenv("BUILD_ID")
}

// BuildName returns the name of the current build
func BuildName() string {
	return os.Getenv("BUILD_NAME")
}

// BuildJobName returns the job name of the current build
func BuildJobName() string {
	return os.Getenv("BUILD_JOB_NAME")
}

// BuildPipelineName returns the pipeline name of the current build
func BuildPipelineName() string {
	return os.Getenv("BUILD_PIPELINE_NAME")
}

// BuildPipelineInstanceVarsRaw returns the instance variables for the current pipeline as string
func BuildPipelineInstanceVarsRaw() string {
	return os.Getenv("BUILD_PIPELINE_INSTANCE_VARS")
}

// BuildPipelineInstanceVars returns the instance variables for the current pipeline as a map
func BuildPipelineInstanceVars() map[string]any {
	res := make(map[string]any)
	// we ignore any error here. As we trust concourse to deliver correct json, and if it does not
	//  there is nothing we can do.
	_ = json.Unmarshal([]byte(BuildPipelineInstanceVarsRaw()), &res)
	return res
}

// BuildTeamName returns the name of the current team
func BuildTeamName() string {
	return os.Getenv("BUILD_TEAM_NAME")
}

// BuildCreatedBy returns the username that created the build
func BuildCreatedBy() string {
	return os.Getenv("BUILD_CREATED_BY")
}

// ConcourseURL returns the external URL of the concourse instance. The original variable name is `ATC_EXTERNAL_URL`
func ConcourseURL() string {
	return os.Getenv("ATC_EXTERNAL_URL")
}

// BuildURL returns the full URL to the build
func BuildURL() string {
	return os.ExpandEnv("${ATC_EXTERNAL_URL}/teams/${BUILD_TEAM_NAME}/pipelines/${BUILD_PIPELINE_NAME}" +
		"/jobs/${BUILD_JOB_NAME}/builds/${BUILD_NAME}")
}

// ExpandEnv is mainly a wrapper around os.ExpandEnv, but limited to the
// variables that are provided by concourse.
//
// The following variables are expanded
// - BUILD_ID
// - BUILD_NAME
// - BUILD_JOB_NAME
// - BUILD_PIPELINE_NAME
// - BUILD_PIPELINE_INSTANCE_VARS
// - BUILD_TEAM_NAME
// - BUILD_CREATED_BY
// - ATC_EXTERNAL_URL (CONCOURSE_URL is also supported)
// - BUILD_URL (see BuildURL)
//
// See https://concourse-ci.org/implementing-resource-types.html#resource-metadata
func ExpandEnv(s string) string {
	return os.Expand(s, func(s string) string {
		switch s {
		case "BUILD_ID":
			return BuildID()
		case "BUILD_NAME":
			return BuildName()
		case "BUILD_JOB_NAME":
			return BuildJobName()
		case "BUILD_PIPELINE_NAME":
			return BuildPipelineName()
		case "BUILD_PIPELINE_INSTANCE_VARS":
			return BuildPipelineInstanceVarsRaw()
		case "BUILD_TEAM_NAME":
			return BuildTeamName()
		case "BUILD_CREATED_BY":
			return BuildCreatedBy()
		case "ATC_EXTERNAL_URL", "CONCOURSE_URL":
			return ConcourseURL()
		case "BUILD_URL":
			return BuildURL()
		default:
			return ""
		}
	})
}
