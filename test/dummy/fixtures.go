package dummy

import (
	"pkg.loki.codes/concourse-resource/types"
)

const SourceFixture = `
"source": {
"uri": "git://some-uri",
"branch": "develop",
"private_key": "..."
}`

var SourceExpected = Config{
	URI:        "git://some-uri",
	Branch:     "develop",
	PrivateKey: "...",
}

const VersionFixture = `"version": { "ref": "61cbef" }`

var VersionExpected = version{Ref: "61cbef"}

const ParamsFixture = `"params": {"param": "test1234"}`

var ParamsExpected = params{Param: "test1234"}

var CheckOutput = types.CheckOutput{
	version{Ref: "61cbef"},
	version{Ref: "d74e01"},
	version{Ref: "7154fe"},
}

var InOutput = types.InOutput{
	Version:  version{Ref: "61cbef"},
	Metadata: map[string]any{"test": "value"},
}
