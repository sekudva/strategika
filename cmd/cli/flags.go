// cmd/cli/flags.go
package cli

import (
	"flag"
)

var (
	Silent bool
	Quick  bool
)

func init() {
	flag.BoolVar(&Silent, "s", false, "silent mode (no detailed logs)")
	flag.BoolVar(&Silent, "silent", false, "silent mode (no detailed logs)")

	flag.BoolVar(&Quick, "q", false, "quick mode (use all defaults, skip questions)")
	flag.BoolVar(&Quick, "quick", false, "quick mode (use all defaults, skip questions)")
}

func ParseFlags() {
	flag.Parse()
}
