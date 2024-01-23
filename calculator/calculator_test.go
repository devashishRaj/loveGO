package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

type TestCase struct {
	a, b     float64
	function func(float64, float64) float64
	want     float64
}

func TestUsingStruct(t *testing.T) {
	testCases := []TestCase{
		{a: 2, b: 2, function: calculator.Add, want: 4},
		{a: 3, b: 3, function: calculator.Multiply, want: 9},
		{a: 10, b: 2, function: calculator.Subtract, want: 8},
	}
	for _, test := range testCases {
		got := test.function(test.a, test.b)
		if got != test.want {
			t.Errorf("Expected %f got %f ", got, test.want)
		}
	}
}

func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 3, want: 4},
		{a: 4, b: 9, want: 13},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		basicMathAssert(t, got, tc.want)
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 0},
		{a: 1, b: 3, want: -2},
		{a: 4, b: 9, want: -5},
	}
	// _ is a blank identifier , use this when you don't want the assigned value
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		basicMathAssert(t, got, tc.want)
	}
}

func TestMultipy(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 3, want: 3},
		{a: 4, b: 9, want: 36},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		basicMathAssert(t, got, tc.want)
	}
}

func TestDivide(t *testing.T) {
	// one behaviour one test
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 1},
		{a: -1, b: -1, want: 1},
		{a: 10, b: 2, want: 5},
		{a: 1, b: 3, want: 0.333333},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			// unlike Errorf, Fatalf exits the test immediately
			t.Fatalf("want no error for valid input, got %v", err)
		}
		//basicMathAssert(t, got, tc.want)
		// floating point numbers have a limited precison thus we can for approximate too , how aproximate ?
		// is it flight navigation or brain surgery by a robotic arm
		closeEnough(t, got, tc.want, 0.001)
	}
}

func TestDivideInvalid(t *testing.T) {
	// one behaviour one test
	t.Parallel()
	type testCase struct {
		a, b float64
		want string
	}
	testCases := []testCase{
		{a: 1, b: 0, want: "Inf"},
		{a: -1, b: 0, want: "-Inf"},
		{a: -1, b: -0, want: "Inf"},
	}
	for _, tc := range testCases {
		_, err := calculator.Divide(tc.a, tc.b)
		if err == nil {
			// unlike Errorf, Fatalf exits the test immediately
			t.Fatalf("expected error for Invalid inputs but got none for input %f , %f", tc.a, tc.b)
		}

	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a    float64
		want float64
	}
	testCases := []testCase{
		{a: 4, want: 2},
		{a: 6, want: 2.44948},
		{a: 11, want: 3.316624},
	}
	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)
		if err != nil {
			t.Errorf("Did not expect error for valid inputs but got one : %v", err)
		}
		closeEnough(t, got, tc.want, 0.1)
	}
}

func TestSqrtInvalid(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a    float64
		want float64
	}
	testCases := []testCase{
		{a: -1, want: 0},
		{a: -0.1, want: 0},
		{a: -9, want: 0},
	}
	for _, tc := range testCases {
		_, err := calculator.Sqrt(tc.a)
		if err == nil {
			t.Fatalf("expected error for Invalid inputs but got none for input %f", tc.a)
		}
	}
}

func basicMathAssert(t testing.TB, got, want float64) {
	t.Helper()
	if want != got {
		// errorf marks the tests as fail but continues the rest of Test
		t.Errorf("want %f got %f", want, got)
	}
}

func closeEnough(t testing.TB, got, want, tolerance float64) {
	t.Helper()
	if !(math.Abs(got-want) <= tolerance) {
		t.Errorf("want %f got %f , tolerance given : %f", want, got, tolerance)
	}
}
