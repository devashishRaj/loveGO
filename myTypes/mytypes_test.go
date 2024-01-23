package mytypes_test

import (
	"mytypes"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTwice(t *testing.T) {
	t.Parallel()
	want := mytypes.MyInt(8)
	got := mytypes.MyInt(4).Twice()
	if got != want {
		t.Errorf("Wanted %v got %v", want, got)
	}
}

func TestMyStringLen(t *testing.T) {
	want := 5
	got := mytypes.MyString("Hello").Slen()
	if got != want {
		t.Errorf("Wanted %v got %v", want, got)
	}
}

func TestStringBuilder(t *testing.T) {
	t.Parallel()
	var sb strings.Builder
	sb.WriteString("Hello, ")
	sb.WriteString("Gophers!")
	want := "Hello, Gophers!"
	got := sb.String()
	AssertValues(t, got, want, "sb.String() test")
	wantLen := 15
	gotLen := sb.Len()
	AssertValues(t, gotLen, wantLen, "string length test")
}

func TestMyBuilder(t *testing.T) {
	t.Parallel()
	var mb mytypes.MyBuilder
	mb.Contents.WriteString("Hello, ")
	mb.Contents.WriteString("Gophers!")
	got := mb.Contents.String()
	want := "Hello, Gophers!"
	AssertValues(t, got, want, "Mybuilder unable to use string.Builders method:  String()")
	wantLen := 15
	gotLen := mb.Contents.Len()
	AssertValues(t, gotLen, wantLen, "Mybuidler unable to use string.Builders method: Len()")
}

func TestStringUpcaser(t *testing.T) {
	t.Parallel()
	var su mytypes.StringUpperCaser
	su.Contents.WriteString("Hello, Gophers!")
	want := "HELLO, GOPHERS!"
	got := su.ToUpper()
	AssertValues(t, got, want, "Test : String Upper Caser")
}

func TestDouble(t *testing.T) {
	t.Parallel()
	var x int = 2
	want := 4
	mytypes.Double(&x)
	AssertValues(t, x, want, "Test double function")

	xi := mytypes.MyInt(12)
	wanti := mytypes.MyInt(24)
	p := &xi
	p.Double()
	AssertValues(t, wanti, xi, "Test method with recevier of pointer type ")

}

func AssertValues(t testing.TB, got interface{}, want interface{}, s string) {
	t.Helper()
	if !cmp.Equal(got, want) {
		t.Errorf("Test : %s \n %v ", s, cmp.Diff(got, want))
	}
}
