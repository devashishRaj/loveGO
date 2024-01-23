package bookstore_test

import (
	"bookstore"
	"log"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

// one behavouir one test is imporatn as one function one test can lead to writing useles test just
// for 100% test coverage
//If there’s no real behaviour in the untested code, it’s probably not worth testing.
//80‐90% coverage is probably a reasonable figure to aim for.

// TestBook is compile only test , used to test for existence and design of the struct.
func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "Happy Joy",
		Author: "Raj",
		Copies: 33,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "x",
		Author: "y",
		Copies: 12,
	}
	want := 11
	result, err := bookstore.Buy(&b)
	if err != nil {
		t.Fatal(err)
	}
	got := result.Copies
	AssertValues(t, got, want, "Buy() test")
}

func TestBuyErrorsIfNoCopiesLeft(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 0,
	}
	_, err := bookstore.Buy(&b)
	if err == nil {
		t.Error("want error buying from zero copies, got nil")
	}
}

func TestBrowseCatalog(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}
	want := []bookstore.Book{
		{ID: 1, Title: "For the Love of Go"},
		{ID: 2, Title: "The Power of Go: Tools"},
	}
	got := catalog.BrowseCatalog()
	sort.Slice(got, func(i, j int) bool {
		return got[i].ID < got[j].ID
	})
	AssertValues(t, got, want, "Browse catalog test")
}

func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}
	// to add new element to map
	//if you do this with a key that already exists in the map,
	//the new element will overwrite the old one
	catalog[3] = bookstore.Book{ID: 3, Title: "Spark Joy"}
	//catalog[1].Title = "For the Love of Go"
	// this doesn't work but you can do this
	//b := catalog[1]
	//b.Title = "For the Love of Go"
	//catalog[1] = b

	want := bookstore.Book{
		ID:    2,
		Title: "The Power of Go: Tools",
	}
	got, err := catalog.GetBook(2)
	if err != nil {
		t.Fatal(err)
	}
	AssertValues(t, got, want, "GetBook test")
}

func TestGetBookBadIDReturnsError(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{}
	_, err := catalog.GetBook(999)
	if err == nil {
		t.Fatal("want error for non-existent ID, got nil")
	}
}

func TestNetPriceCents(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:           "For the Love of Go",
		PriceCents:      4000,
		DiscountPercent: 25,
	}
	want := 3000
	got := b.NetPriceCents()
	AssertValues(t, got, want, "Net percent test")

}

func TestSetPrice(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:      "For the love of Go",
		PriceCents: 4000,
	}
	want := 3000
	err := b.SetPrice(want)
	AssertValues(t, err, nil, "Testing setPrice valid input")
	err = b.SetPrice(-1)
	if err == nil {
		t.Fatalf("expected error got nil ")
	}
	// unable to use errors.New() in AssertValues() func
}

func TestSetCategory(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title: "For the love of Go",
	}
	cats := []bookstore.Category{
		bookstore.CategoryAutobiography,
		bookstore.CategoryLargePrintRomance,
		bookstore.CategoryParticlePhysics,
	}
	for _, cat := range cats {
		err := b.SetCategory(cat)
		if err != nil {
			log.Fatalln(err)
		}
		got := b.Category()
		if cat != got {
			t.Errorf("want category %q, got %q", cat, got)
		}
	}
}

func TestSetCategoryInvalid(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title: "For the Love of Go",
	}
	err := b.SetCategory(999)
	if err == nil {
		t.Fatal("want error for invalid category, got nil")
	}
}

func AssertValues(t testing.TB, got interface{}, want interface{}, s string) {
	t.Helper()
	// ignore the unexported fields to test for match as access to them is not allowed
	if !cmp.Equal(want, got,
		cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Errorf("Test : %s \n %v ", s, cmp.Diff(got, want))
	}
}
