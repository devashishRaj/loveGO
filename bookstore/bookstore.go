package bookstore

import (
	"errors"
	"fmt"
)

// Starting a literal with Captial letter will make accessible in other packages too : exported
// local == unexported , can be used by files withing same package only.

// Catalog is a local type(user defined type)
// map of int to book
type Catalog map[int]Book

// In Go maps looking up a non‐existent key returns the zero value of the element type NOT error
// A data type made up of several values that you can treat as a single unit is called a composite type
// in Go is called a struct, short for “structured record”.
// A struct groups together related pieces of data, called fields like Title , ID ...
// they can be of any type even another struct
type Book struct {
	Title           string
	Author          string
	Copies          int
	ID              int
	PriceCents      int
	DiscountPercent int
	category        Category
}

type Category int

// definig a category as a const helps in helps in validation by compiler
// and also they can by easily refered anywhere we want instead of some arbitary strings
// if they should misspell the name of the category, the compiler will pick this
const (
	// when you assign iota to a const, it will get value zero and successive const will get value
	// 1,2,3.... , it's used here as the actual value of the constant does not matter the name of
	// constant is enough
	CategoryAutobiography Category = iota
	CategoryLargePrintRomance
	CategoryParticlePhysics
)

// a way to ensure that some value is one of a predefined set of possibile values
var validCategory = map[Category]bool{
	CategoryAutobiography:     true,
	CategoryLargePrintRomance: true,
	CategoryParticlePhysics:   true,
}

func Buy(b *Book) (Book, error) {
	// Book{} is an empty struct zero type value for type struct
	//return Book{}
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies -= 1
	return *b, nil
}

// here (c Catalog) is parameter for BrowseCatalog and acts as a receiver thus meaning
// that method is about this receiver, and it’s part of the definition of the receiver’s type.
func (c Catalog) BrowseCatalog() []Book {
	result := []Book{}
	for _, b := range c {
		result = append(result, b)
	}
	return result
}

func (c Catalog) GetBook(ID int) (Book, error) {
	//looking up a non‐existent key doesn’t cause an error: instead, it returns the zero value of the element type
	//  range over a map returns the key and the element
	b, ok := c[ID]
	if !ok {
		//Since we have the ID value, we can include it in the message, using fmt.Errorf instead of errors.New.
		return Book{}, fmt.Errorf("ID %d doesn't exist", ID)
	}
	return b, nil
}

// object’s methods (such as NetPriceCents) are closely asso‐ ciated with its type (Book) and
// must be defined in the same package as the type.
func (b Book) NetPriceCents() int {
	return (b.PriceCents - (b.PriceCents*b.DiscountPercent)/100)
}

func (b *Book) SetPrice(NewPrice int) error {
	if NewPrice <= 0 {
		// Try not to  use sentinel errors as they are fixed and not give usefull info
		// eg: you made a sentinel error var FileError = errors.New("FileError")
		// now depending upon case you might have a read acces porblem or file does not exits
		// so insted of getting the right error from io.reader library you will get a vague one .
		return errors.New("Negative input for setPrice not allowed")
	}
	b.PriceCents = NewPrice
	return nil
}

func (b Book) Category() Category {
	return b.category
}

func (b *Book) SetCategory(category Category) error {
	// if category is not present , the map will return 0 (false)
	if !validCategory[category] {
		return fmt.Errorf("unknown category %v", category)
	}
	b.category = category
	return nil
}
