//go:generate mockery -name=NumberCruncher

package calculator

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jaysonesmith/golangphoenix-tests/mocks"
)

// Units can be any grouping of your code that makes sense given the context that you're working
// with. In this case, we'll be testing individual functions. While this can be useful, I'd
// recommend testing only at your exported functions/functions on interfaces so as to help help
// prevent you from writing too many brittle tests. (though this isn't a failsafe, just one good
// method.) Additionally, try to make sure that your tests follow a clear pattern of arrange, act,
// assert. This greatly helps with readability and keeps style consistent.
// Arrange - set up your state, open/create test files, create variables, etc.
// Act - trigger the action or part of your code that's being tested. This can be calling
// 	functions, api's, etc.
// Assert - check that the outcome you're expecting actually happened.

// Basic tests are simple and useful, but repetitive, having the same boilerplate over and over for
// small iterations for different cases
func TestAddZeros(t *testing.T) {
	var expected float64

	actual := Add(0, 0)

	if expected != actual {
		t.Fatalf("expected %f found %f", expected, actual)
	}
}

// More of the same, just some different parameters
func TestAddNegativeNumbers(t *testing.T) {
	expected := -10.0

	actual := Add(-5, -5)

	if expected != actual {
		t.Fatalf("expected %f found %f", expected, actual)
	}
}

// Lets combine them together by using table tests AND make sure that all run even if some fail by
// introducing subtests. These are triggered with the t.Run() you'll see below.
// More: https://golang.org/pkg/testing/#hdr-Subtests_and_Sub_benchmarks
// Not only do you get info about all your tests, but you also get it in a much more consice way.
// Additionally, we can introduce the assert() function from github.com/stretchr/testify to clean
// up our assertions and give us some more informative output when things fail.
func TestAdd(t *testing.T) {
	testCases := []struct {
		name     string
		x        float64
		y        float64
		expected float64
	}{
		{name: "Zeros", x: 0, y: 0, expected: 0.0},
		{name: "Negative numbers", x: -5, y: -5, expected: -10.0},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := Add(tc.x, tc.y)

			// This is a MUCH cleaner way of
			// running our assertions
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

// If you set up test helpers to open files, set ENV variables or so on, here's some neat ways to
// set them up that'll help remove clutter from your tests, allowing you and your team to focus on
// what's actually important: the tests.

// First, let's take a look at the 'standard' way of calling a test helper that has the chance to
// fail. This first helper example will return an error if any is returned. This results in
// cluttering up the test example
func TestExampleHelper(t *testing.T) {
	// Here our example will parse some 'data' that's required for the test to run. In practice,
	// this could be a test file for instance. This particular example takes up three lines when
	// one could suffice.
	expected, err := testFloatParserReturnsError("10")
	if err != nil {
		t.FailNow()
	}

	// Do something for our test
	actual := Add(5, 5)

	// Make our assertion with the data from the helper
	assert.Equal(t, expected, actual)
}

// testFloatParserReturnsError attempts to parse the input string into a float64 value and returns
// an error if one is found
func testFloatParserReturnsError(input string) (float64, error) {
	out, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return out, err
	}

	return out, nil
}

// Let's instead do basically the same thing but we'll modify the test helper and the way we're
// using it so that it'll be cleaner
func TestExampleHelperTwo(t *testing.T) {
	expected := testFloatParserFails(t, "10")

	out := Add(5, 5)

	assert.Equal(t, expected, out)
}

// testFloatParserFails attempts to parse the input string into a float64 value and fails on its
// own if an error is returned. This is done as if this file is needed for the test to be
// successful, then it makes sense that the test should fail if something goes wrong. Failing in
// the test provides cleaner tests!
func testFloatParserFails(t *testing.T, input string) float64 {
	out, err := strconv.ParseFloat(input, 64)
	if err != nil {
		t.Fatal(err)
	}

	return out
}

// In addition to having test helpers fail tests for required data, we can also write test helpers
// that have cleanup to do in a slick way to futher clean up our test setup!
// For this example, our testSetENV helper will set the desired environment variable and also
// automatically reset it to whatever value might have been in place when this test started. In
// order to do this, we'll wrap the reset behavior in a closure, return it as the functions return
// value, and call the helper with defer.
func TestExampleHelperThree(t *testing.T) {
	defer testSetENV("foo", "bar")

	actual := Add(0, 0)

	assert.Equal(t, 0.0, actual)
}

func testSetENV(key, value string) func() {
	// Store the original env var
	ogENV := os.Getenv(key)

	// Set the new value
	os.Setenv(key, value)

	// Return our reset function. This could be a call to another function if you'd like as well!
	return func() { os.Setenv(key, ogENV) }
}

// Mocking is an extremely powerful tool for being able to test various things from http responses
// to function data without having to actually have those things working or even exist fully! For
// our example here, we'll be using mockery to generate a mock of our NumberCruncher interface.
// https://github.com/vektra/mockery
// There are many ways to generate a mock of an interface, but here we'll use `make generate` which
// will prompt Go to see our `//go:generate` note at the top of the file and call mockery for us.
// Mocked interfaces utilize the same signature of our functions but allow us to specify the
// response behavior as we see fit. With the mock created we can use it!
func TestVerify(t *testing.T) {
	// Create a new instance of our mocked interface
	mockNumberCruncher := &mocks.NumberCruncher{}
	// Specify that whenever the Verify function is called with the arguments 0, 1, then return false
	mockNumberCruncher.On("Verify", 0, 1).Return(false)
	expected := false

	// Call Verify on our mock
	actual := mockNumberCruncher.Verify(0, 1)

	// Verify like normal!
	assert.Equal(t, expected, actual)
}
