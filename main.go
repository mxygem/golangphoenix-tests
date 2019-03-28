package calculator

// NumberCruncher runs calculations and verifications
type NumberCruncher interface {
	Add(x, y float64) float64
	Verify(got, want float64) bool
}

// Add sums two numbers
func Add(x, y float64) float64 {
	return x + y
}

// Verify is an "example" of a wrapper for an html call. In this example, the API could be thought
// of as not being made yet, but that doesn't prevent us from testing using mocks.
func Verify(got, want float64) bool {
	return true
}
