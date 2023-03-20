# concourse-resource

***a Go module to simplify the development of concourse resources***

## Features

### minimal dependencies

Only `github.com/stretchr/testify` and  `github.com/go-playground/validator/v10` are directly required by this module.

### config validation

To validate your input, simply add the correct [`validator` tags](https://pkg.go.dev/github.com/go-playground/validator/v10) 

### flexible

Use whatever library you want, do what ever you want. This modules takes only care about the communication with concourse.
No assumptions beyond that are made.


### testable

Test your resources with `go test`. A full set of testing helpers and a `testify` suite are provided.

## Usage

To develop a new resource, you must provide a type, that fulfills at least one of the following interfaces:

- `types.CheckResource`
- `types.InResource`
- `types.OutResource`

If your resource takes parameters for `in` or `out`, the type must implement the `types.ParametrizedResource` interface.

In addition, a factory function must be provided.

While all arguments to `check`, `in` and `out` are provided by the resource at runtime, the configuration (`source`)
must be provided statically.

All communication with *concourse* is done by the `Handler`. For an example of a resource, have a look at the [`test/dummy`](test/dummy)
implementation.

Inside your main, put all the stuff together and call the `Run` method.

```go
package main

import (
	"log"

	resource "pkg.loki.codes/concourse-resource"

	"your/resource/implementation"
)

func main() {
	if err := resource.New[implementation.Config](implementation.New).Run(); err != nil {
		log.Fatal(err)
	}
}

```

### Container build

To build your container image, you may just symlink the binary to the correct locations. For example:

```Dockerfile
FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

ADD . /app/
WORKDIR /app
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o resource .
RUN mkdir -p /target/opt/resource/
RUN cp resource /target/opt/resource/
RUN ln -s resource /target/opt/resource/in
RUN ln -s resource /target/opt/resource/out
RUN ln -s resource /target/opt/resource/check


FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /target/opt /opt
```

## Testing

There is an extensive test framework. See [test](test)