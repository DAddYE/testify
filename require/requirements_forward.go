// auto generated file, do not edit
package require

import (
	"time"

	"github.com/stretchr/testify/assert"
)

type Requirements struct {
	t TestingT
}

func New(t TestingT) *Requirements {
	return &Requirements{t: t}
}

// Fail reports a failure through
func (r *Requirements) Fail(failureMessage string, msgAndArgs ...interface{}) {
	Fail(r.t, failureMessage, msgAndArgs)
}

// Implements asserts that an object is implemented by the specified interface.
//
//    require.Implements((*MyInterface)(nil), new(MyObject), "MyObject")
func (r *Requirements) Implements(interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) {
	Implements(r.t, interfaceObject, object, msgAndArgs)
}

// IsType asserts that the specified objects are of the same type.
func (r *Requirements) IsType(expectedType interface{}, object interface{}, msgAndArgs ...interface{}) {
	IsType(r.t, expectedType, object, msgAndArgs)
}

// Equal asserts that two objects are equal.
//
//    require.Equal(123, 123, "123 and 123 should be equal")
//
// Returns whether the assertion was successful (true) or not (false).
func (r *Requirements) Equal(expected, actual interface{}, msgAndArgs ...interface{}) {
	Equal(r.t, expected, actual, msgAndArgs)
}

// Exactly asserts that two objects are equal is value and type.
//
//    require.Exactly(int32(123), int64(123), "123 and 123 should NOT be equal")
//
// Returns whether the assertion was successful (true) or not (false).
func (r *Requirements) Exactly(expected, actual interface{}, msgAndArgs ...interface{}) {
	Exactly(r.t, expected, actual, msgAndArgs)
}

// NotNil asserts that the specified object is not nil.
//
//    require.NotNil(err, "err should be something")
//
// Returns whether the assertion was successful (true) or not (false).
func (r *Requirements) NotNil(object interface{}, msgAndArgs ...interface{}) {
	NotNil(r.t, object, msgAndArgs)
}

// Nil asserts that the specified object is nil.
//
//    require.Nil(err, "err should be nothing")
//
// Returns whether the assertion was successful (true) or not (false).
func (r *Requirements) Nil(object interface{}, msgAndArgs ...interface{}) {
	Nil(r.t, object, msgAndArgs)
}

// Empty asserts that the specified object is empty.  I.e. nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
// require.Empty(obj)
//
// Returns whether the assertion was successful (true) or not (false).
func (r *Requirements) Empty(object interface{}, msgAndArgs ...interface{}) {
	Empty(r.t, object, msgAndArgs)
}

// NotEmpty asserts that the specified object is NOT empty.  I.e. not nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
// if require.NotEmpty(obj) {
//   require.Equal("two", obj[1])
// }
//
// Returns whether the assertion was successful (true) or not (false).
func (r *Requirements) NotEmpty(object interface{}, msgAndArgs ...interface{}) {
	NotEmpty(r.t, object, msgAndArgs)
}

// Len asserts that the specified object has specific length.
// Len also fails if the object has a type that len() not accept.
//
//    require.Len(mySlice, 3, "The size of slice is not 3")
//
// Returns whether the assertion was successful (true) or not (false).
func (r *Requirements) Len(object interface{}, length int, msgAndArgs ...interface{}) {
	Len(r.t, object, length, msgAndArgs)
}

// True asserts that the specified value is true.
//
//    require.True(myBool, "myBool should be true")
//
// Returns whether the assertion was successful (true) or not (false).
func (r *Requirements) True(value bool, msgAndArgs ...interface{}) {
	True(r.t, value, msgAndArgs)
}

// False asserts that the specified value is true.
//
//    require.False(myBool, "myBool should be false")
//
// Returns whether the assertion was successful (true) or not (false).
func (r *Requirements) False(value bool, msgAndArgs ...interface{}) {
	False(r.t, value, msgAndArgs)
}

// NotEqual asserts that the specified values are NOT equal.
//
//    require.NotEqual(obj1, obj2, "two objects shouldn't be equal")
//
// Returns whether the assertion was successful (true) or not (false).
func (r *Requirements) NotEqual(expected, actual interface{}, msgAndArgs ...interface{}) {
	NotEqual(r.t, expected, actual, msgAndArgs)
}

// Contains asserts that the specified string or list(array, slice...) contains the
// specified substring or element.
//
//    require.Contains("Hello World", "World", "But 'Hello World' does contain 'World'")
//    require.Contains(["Hello", "World"], "World", "But ["Hello", "World"] does contain 'World'")
//
// Returns whether the assertion was successful (true) or not (false).
func (r *Requirements) Contains(s, contains interface{}, msgAndArgs ...interface{}) {
	Contains(r.t, s, contains, msgAndArgs)
}

// NotContains asserts that the specified string or list(array, slice...) does NOT contain the
// specified substring or element.
//
//    require.NotContains("Hello World", "Earth", "But 'Hello World' does NOT contain 'Earth'")
//    require.NotContains(["Hello", "World"], "Earth", "But ['Hello', 'World'] does NOT contain 'Earth'")
//
// Returns whether the assertion was successful (true) or not (false).
func (r *Requirements) NotContains(s, contains interface{}, msgAndArgs ...interface{}) {
	NotContains(r.t, s, contains, msgAndArgs)
}

// Condition uses a Comparison to assert a complex condition.
func (r *Requirements) Condition(comp assert.Comparison, msgAndArgs ...interface{}) {
	Condition(r.t, comp, msgAndArgs)
}

// Panics asserts that the code inside the specified PanicTestFunc panics.
//
//   require.Panics(func(){
//     GoCrazy()
//   }, "Calling GoCrazy() should panic")
//
// Returns whether the assertion was successful (true) or not (false).
func (r *Requirements) Panics(f assert.PanicTestFunc, msgAndArgs ...interface{}) {
	Panics(r.t, f, msgAndArgs)
}

// NotPanics asserts that the code inside the specified PanicTestFunc does NOT panic.
//
//   require.NotPanics(func(){
//     RemainCalm()
//   }, "Calling RemainCalm() should NOT panic")
//
// Returns whether the assertion was successful (true) or not (false).
func (r *Requirements) NotPanics(f assert.PanicTestFunc, msgAndArgs ...interface{}) {
	NotPanics(r.t, f, msgAndArgs)
}

// WithinDuration asserts that the two times are within duration delta of each other.
//
//   require.WithinDuration(time.Now(), time.Now(), 10*time.Second, "The difference should not be more than 10s")
//
// Returns whether the assertion was successful (true) or not (false).
func (r *Requirements) WithinDuration(expected, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) {
	WithinDuration(r.t, expected, actual, delta, msgAndArgs)
}

// InDelta asserts that the two numerals are within delta of each other.
//
// 	 require.InDelta(math.Pi, (22 / 7.0), 0.01)
//
// Returns whether the assertion was successful (true) or not (false).
func (r *Requirements) InDelta(expected, actual interface{}, delta float64, msgAndArgs ...interface{}) {
	InDelta(r.t, expected, actual, delta, msgAndArgs)
}

// InEpsilon asserts that expected and actual have a relative error less than epsilon
//
// Returns whether the assertion was successful (true) or not (false).
func (r *Requirements) InEpsilon(expected, actual interface{}, epsilon float64, msgAndArgs ...interface{}) {
	InEpsilon(r.t, expected, actual, epsilon, msgAndArgs)
}

// NoError asserts that a function returned no error (i.e. `nil`).
//
//   actualObj, err := SomeFunction()
//   if require.NoError(err) {
//	   require.Equal(actualObj, expectedObj)
//   }
//
// Returns whether the assertion was successful (true) or not (false).
func (r *Requirements) NoError(err error, msgAndArgs ...interface{}) {
	NoError(r.t, err, msgAndArgs)
}

// Error asserts that a function returned an error (i.e. not `nil`).
//
//   actualObj, err := SomeFunction()
//   if require.Error(err, "An error was expected") {
//	   require.Equal(err, expectedError)
//   }
//
// Returns whether the assertion was successful (true) or not (false).
func (r *Requirements) Error(err error, msgAndArgs ...interface{}) {
	Error(r.t, err, msgAndArgs)
}

// EqualError asserts that a function returned an error (i.e. not `nil`)
// and that it is equal to the provided error.
//
//   actualObj, err := SomeFunction()
//   if require.Error(err, "An error was expected") {
//	   require.Equal(err, expectedError)
//   }
//
// Returns whether the assertion was successful (true) or not (false).
func (r *Requirements) EqualError(theError error, errString string, msgAndArgs ...interface{}) {
	EqualError(r.t, theError, errString, msgAndArgs)
}

// Regexp asserts that a specified regexp matches a string.
//
//  require.Regexp(regexp.MustCompile("start"), "it's starting")
//  require.Regexp("start...$", "it's not starting")
//
// Returns whether the assertion was successful (true) or not (false).
func (r *Requirements) Regexp(rx interface{}, str interface{}) {
	Regexp(r.t, rx, str)
}

// NotRegexp asserts that a specified regexp does not match a string.
//
//  require.NotRegexp(regexp.MustCompile("starts"), "it's starting")
//  require.NotRegexp("^start", "it's not starting")
//
// Returns whether the assertion was successful (true) or not (false).
func (r *Requirements) NotRegexp(rx interface{}, str interface{}) {
	NotRegexp(r.t, rx, str)
}
