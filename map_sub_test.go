package snippets

import (
	"reflect"
	"strings"
	"testing"
)

var split = strings.Split // alias for test subject

func TestSplit(t *testing.T) {
	// map based table-testing allows keys as test names,
	// and cases in a map means they are run in an *undefined* order
	tests := map[string]struct {
		input string
		sep   string
		want  []string
	}{
		"simple":       {input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
		"wrong sep":    {input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
		"no sep":       {input: "abc", sep: "/", want: []string{"abc"}},
		"trailing sep": {input: "a/b/c/", sep: "/", want: []string{"a", "b", "c"}},
	}

	// subtests so that t.Fatalf can be used
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := split(tc.input, tc.sep)
			if !reflect.DeepEqual(tc.want, got) {
				// #v for clearer output
				t.Fatalf("expected: %#v, got: %#v", tc.want, got)
			}
		})
	}
}
