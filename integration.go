package testutils

import (
	"flag"
)

var (
	integration = flag.Bool("test.integration", false, "Run integration tests.")
)

func Integration() bool {
	return *integration
}
