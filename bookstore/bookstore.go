package bookstore

import (
	"errors"
	"fmt"
)

type Catalog map[int]Book

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

const (
	CategoryAutobiography = iota
	CategoryLargePrintRomance
	CategoryParticlePhysics
)

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

func (c Catalog) BrowseCatalog() []Book {
	result := []Book{}
	for _, b := range c {
		result = append(result, b)
	}
	return result
}

func (c Catalog) GetBook(ID int) (Book, error) {
	//looking up a non‐existent key doesn’t cause an error: instead, it returns the zero value of the element type
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
	if !validCategory[category] {
		return fmt.Errorf("unknown category %v", category)
	}
	b.category = category
	return nil
}
