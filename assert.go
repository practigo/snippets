package snippets

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// ObjectsAreEqual checks two interfaces with reflect.DeepEqual.
func ObjectsAreEqual(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	return reflect.DeepEqual(expected, actual)
}

// IsNil checks an interface{} with the reflect package.
func IsNil(object interface{}) bool {
	if object == nil {
		return true
	}

	value := reflect.ValueOf(object)
	kind := value.Kind()
	if kind >= reflect.Chan && kind <= reflect.Slice && value.IsNil() {
		return true
	}

	return false
}

// errorSingle fails and prints the single object
// along with the message.
func errorSingle(t testing.TB, msg string, obj interface{}) {
	_, file, line, _ := runtime.Caller(2)
	fmt.Printf("\033[31m\t%s:%d: %s\n\n\t\t%#v\033[39m\n\n", filepath.Base(file), line, msg, obj)
	t.Fail()
}

// errorCompare fails and prints both the compared objects
// along with the message.
func errorCompare(t testing.TB, msg string, expected, actual interface{}) {
	_, file, line, _ := runtime.Caller(2)
	fmt.Printf("\033[31m\t%s:%d: %s\n\n\t\tgot: %#v\n\033[32m\t\texp: %#v\033[39m\n\n", filepath.Base(file), line, msg, actual, expected)
	t.Fail()
}

// Assert wraps a testing.TB for convenient asserting calls.
type Assert struct {
	t testing.TB
}

// True tests if the cond is true and prints the msg for failure.
func (a *Assert) True(cond bool, msg string) {
	if !cond {
		errorSingle(a.t, msg, cond)
	}
}

// Equal tests if the two interfaces provided is equal
// and prints the msg for failure.
func (a *Assert) Equal(expected, actual interface{}, msg string) {
	if !ObjectsAreEqual(expected, actual) {
		errorCompare(a.t, msg, expected, actual)
	}
}

func (a *Assert) noError(err error, msg string, fatal bool) {
	if err != nil {
		_, file, line, _ := runtime.Caller(2)
		fmt.Printf("\033[31m\t%s:%d: %s\n\n\t\t%s\033[39m\n\n", filepath.Base(file), line, msg, err)
		if fatal {
			a.t.FailNow()
		} else {
			a.t.Fail()
		}
	}
}

// NoError fails the test and prints the msg if err != nil.
func (a *Assert) NoError(err error, msg string) {
	a.noError(err, msg, false)
}

// MustNoError fails & stops the test if err != nil.
func (a *Assert) MustNoError(err error, msg string) {
	a.noError(err, msg, true)
}

// NotNil fails the test and prints the msg if the obj is nil.
func (a *Assert) NotNil(obj interface{}, msg string) {
	if IsNil(obj) {
		errorSingle(a.t, msg, obj)
	}
}

// FloatEqual compares two floats with an epsilon 1e-9.
func (a *Assert) FloatEqual(expected, actual float64, msg string) {
	if math.Abs(expected-actual) > 1e-9 {
		errorCompare(a.t, msg, expected, actual)
	}
}

// RequireEnv checks if a specific env variable is set, and
// skip the test if it is not set.
func (a *Assert) RequireEnv(key string) string {
	s := os.Getenv(key)
	if s == "" {
		a.t.Skip("set " + key + " to run this test")
	}
	return s
}

// JSON logs v in json-encoded form.
func (a *Assert) JSON(v interface{}) {
	bs, _ := json.Marshal(v)
	a.t.Log(string(bs))
}

// NewAssert provides an Assert instance.
func NewAssert(t testing.TB) *Assert {
	return &Assert{t}
}
