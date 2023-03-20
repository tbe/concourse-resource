package internal

import (
	"io"
	"os"
)

var StdIn io.Reader = os.Stdin
var StdOut io.Writer = os.Stdout
