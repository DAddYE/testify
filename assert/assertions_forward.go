// auto genrated file, do not edit
package assert

import "time"

type Assertions struct {
	t TestingT
}

func New(t TestingT) *Assertions {
	return &Assertions{t: t}
}

// Fail reports a failure through
func (a *Assertions) Fail(failureMessage string, msgAndArgs ...interface{}) bool {
	return Fail(a.t, failureMessage, msgAndArgs)
}

// Implements asserts that an object is implemented by the specified interface.
//
//    assert.Implements((*MyInterface)(nil), new(MyObject), "MyObject")
func (a *Assertions) Implements(interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) bool {
	return Implements(a.t, interfaceObject, object, msgAndArgs)
}

// IsType asserts that the specified objects are of the same type.
func (a *Assertions) IsType(expectedType interface{}, object interface{}, msgAndArgs ...interface{}) bool {
	return IsType(a.t, expectedType, object, msgAndArgs)
}

// Equal asserts that two objects are equal.
//
//    assert.Equal(123, 123, "123 and 123 should be equal")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Equal(expected, actual interface{}, msgAndArgs ...interface{}) bool {
	return Equal(a.t, expected, actual, msgAndArgs)
}

// Exactly asserts that two objects are equal is value and type.
//
//    assert.Exactly(int32(123), int64(123), "123 and 123 should NOT be equal")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Exactly(expected, actual interface{}, msgAndArgs ...interface{}) bool {
	return Exactly(a.t, expected, actual, msgAndArgs)
}

// NotNil asserts that the specified object is not nil.
//
//    assert.NotNil(err, "err should be something")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NotNil(object interface{}, msgAndArgs ...interface{}) bool {
	return NotNil(a.t, object, msgAndArgs)
}

// Nil asserts that the specified object is nil.
//
//    assert.Nil(err, "err should be nothing")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Nil(object interface{}, msgAndArgs ...interface{}) bool {
	return Nil(a.t, object, msgAndArgs)
}

// Empty asserts that the specified object is empty.  I.e. nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
// assert.Empty(obj)
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Empty(object interface{}, msgAndArgs ...interface{}) bool {
	return Empty(a.t, object, msgAndArgs)
}

// NotEmpty asserts that the specified object is NOT empty.  I.e. not nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
// if assert.NotEmpty(obj) {
//   assert.Equal("two", obj[1])
// }
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NotEmpty(object interface{}, msgAndArgs ...interface{}) bool {
	return NotEmpty(a.t, object, msgAndArgs)
}

// Len asserts that the specified object has specific length.
// Len also fails if the object has a type that len() not accept.
//
//    assert.Len(mySlice, 3, "The size of slice is not 3")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Len(object interface{}, length int, msgAndArgs ...interface{}) bool {
	return Len(a.t, object, length, msgAndArgs)
}

// True asserts that the specified value is true.
//
//    assert.True(myBool, "myBool should be true")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) True(value bool, msgAndArgs ...interface{}) bool {
	return True(a.t, value, msgAndArgs)
}

// False asserts that the specified value is true.
//
//    assert.False(myBool, "myBool should be false")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) False(value bool, msgAndArgs ...interface{}) bool {
	return False(a.t, value, msgAndArgs)
}

// NotEqual asserts that the specified values are NOT equal.
//
//    assert.NotEqual(obj1, obj2, "two objects shouldn't be equal")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NotEqual(expected, actual interface{}, msgAndArgs ...interface{}) bool {
	return NotEqual(a.t, expected, actual, msgAndArgs)
}

// Contains asserts that the specified string or list(array, slice...) contains the
// specified substring or element.
//
//    assert.Contains("Hello World", "World", "But 'Hello World' does contain 'World'")
//    assert.Contains(["Hello", "World"], "World", "But ["Hello", "World"] does contain 'World'")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Contains(s, contains interface{}, msgAndArgs ...interface{}) bool {
	return Contains(a.t, s, contains, msgAndArgs)
}

// NotContains asserts that the specified string or list(array, slice...) does NOT contain the
// specified substring or element.
//
//    assert.NotContains("Hello World", "Earth", "But 'Hello World' does NOT contain 'Earth'")
//    assert.NotContains(["Hello", "World"], "Earth", "But ['Hello', 'World'] does NOT contain 'Earth'")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NotContains(s, contains interface{}, msgAndArgs ...interface{}) bool {
	return NotContains(a.t, s, contains, msgAndArgs)
}

// Condition uses a Comparison to assert a complex condition.
func (a *Assertions) Condition(comp Comparison, msgAndArgs ...interface{}) bool {
	return Condition(a.t, comp, msgAndArgs)
}

// Panics asserts that the code inside the specified PanicTestFunc panics.
//
//   assert.Panics(func(){
//     GoCrazy()
//   }, "Calling GoCrazy() should panic")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Panics(f PanicTestFunc, msgAndArgs ...interface{}) bool {
	return Panics(a.t, f, msgAndArgs)
}

// NotPanics asserts that the code inside the specified PanicTestFunc does NOT panic.
//
//   assert.NotPanics(func(){
//     RemainCalm()
//   }, "Calling RemainCalm() should NOT panic")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NotPanics(f PanicTestFunc, msgAndArgs ...interface{}) bool {
	return NotPanics(a.t, f, msgAndArgs)
}

// WithinDuration asserts that the two times are within duration delta of each other.
//
//   assert.WithinDuration(time.Now(), time.Now(), 10*time.Second, "The difference should not be more than 10s")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) WithinDuration(expected, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) bool {
	return WithinDuration(a.t, expected, actual, delta, msgAndArgs)
}

// InDelta asserts that the two numerals are within delta of each other.
//
// 	 assert.InDelta(math.Pi, (22 / 7.0), 0.01)
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) InDelta(expected, actual interface{}, delta float64, msgAndArgs ...interface{}) bool {
	return InDelta(a.t, expected, actual, delta, msgAndArgs)
}

// InEpsilon asserts that expected and actual have a relative error less than epsilon
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) InEpsilon(expected, actual interface{}, epsilon float64, msgAndArgs ...interface{}) bool {
	return InEpsilon(a.t, expected, actual, epsilon, msgAndArgs)
}

// NoError asserts that a function returned no error (i.e. `nil`).
//
//   actualObj, err := SomeFunction()
//   if assert.NoError(err) {
//	   assert.Equal(actualObj, expectedObj)
//   }
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NoError(err error, msgAndArgs ...interface{}) bool {
	return NoError(a.t, err, msgAndArgs)
}

// Error asserts that a function returned an error (i.e. not `nil`).
//
//   actualObj, err := SomeFunction()
//   if assert.Error(err, "An error was expected") {
//	   assert.Equal(err, expectedError)
//   }
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Error(err error, msgAndArgs ...interface{}) bool {
	return Error(a.t, err, msgAndArgs)
}

// EqualError asserts that a function returned an error (i.e. not `nil`)
// and that it is equal to the provided error.
//
//   actualObj, err := SomeFunction()
//   if assert.Error(err, "An error was expected") {
//	   assert.Equal(err, expectedError)
//   }
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) EqualError(theError error, errString string, msgAndArgs ...interface{}) bool {
	return EqualError(a.t, theError, errString, msgAndArgs)
}

// Regexp asserts that a specified regexp matches a string.
//
//  assert.Regexp(regexp.MustCompile("start"), "it's starting")
//  assert.Regexp("start...$", "it's not starting")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) Regexp(rx interface{}, str interface{}) bool {
	return Regexp(a.t, rx, str)
}

// NotRegexp asserts that a specified regexp does not match a string.
//
//  assert.NotRegexp(regexp.MustCompile("starts"), "it's starting")
//  assert.NotRegexp("^start", "it's not starting")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) NotRegexp(rx interface{}, str interface{}) bool {
	return NotRegexp(a.t, rx, str)
}
