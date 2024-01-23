package mytypes

import "strings"

// MyInt is a local type (that is to say, it’s defined in this package).
type MyInt int
type MyString string

// MyBuilder  won't be able to use strings.Builder methods via following
// type MyBuilder strings.Builder
// but here's a solution
type MyBuilder struct {
	Contents strings.Builder
}

type StringUpperCaser struct {
	Contents strings.Builder
}

func (i MyInt) Twice() MyInt {
	return i * 2
}

func (s MyString) Slen() int {

	return len(s)
}

func (mb MyBuilder) Hello() string {
	return "Hello, Gophers!"
}

// local type, defined by us, we can also add any methods to it that we want
func (su StringUpperCaser) ToUpper() string {
	return strings.ToUpper(su.Contents.String())
}

// type here is not just “pointer”, but specifically “pointer to int”
func Double(x *int) {
	// dereferencing the pointer *x then multiply with 2
	*x *= 2
}

func (input *MyInt) Double() {
	*input *= 2
}
