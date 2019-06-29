package snippets_test

import (
	"testing"

	"github.com/practigo/snippets"
)

type otherErrType string

func (e otherErrType) Error() string {
	return string(e)
}

const otherConstErr = otherErrType("demo constant error")

func TestConstErr(t *testing.T) {
	cases := map[string]struct {
		exp bool
	}{
		"function return":        {exp: snippets.DemoConstErr == snippets.GetDemoErr()},
		"fungible":               {exp: snippets.DemoConstErr == snippets.FungibleErr},
		"fungible (wrapped)":     {exp: snippets.DemoConstErr == error(snippets.FungibleErr)},
		"no unexpected equality": {exp: snippets.DemoConstErr != error(otherConstErr)},
		// "compile error": {exp: snippets.DemoConstErr == otherConstErr}, // mismatched types
		"string does not implement error tho": {exp: snippets.DemoConstErr == "demo constant error"},
	}

	for name, c := range cases {
		if !c.exp {
			t.Error(name, "failed")
		}
	}
}
