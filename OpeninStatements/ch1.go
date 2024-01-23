package main

import (
	"fmt"
	"log"
	"os"
	"sort"
)
// Alternatives to init : function runs before any other in the package
var thing = initialiseThing()

func initialiseThing() int { // intialise some stuff
	return 999
}

func getPositon() (int, int, int) {
	return 1, 2, 4
}

func bigger(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func doStuff() error {
	return nil
}

func somethingWrong() bool { return true }

// Employee struct representing an employee
type Employee struct {
	Name      string
	IsCurrent bool
}

// PrintCheck method to print the employee's check
func (e *Employee) PrintCheck() {
	fmt.Printf("Printing check for %s\n", e.Name)
}

func cleanup1() { return }
func cleanup2() { return }

func main() {
	// declare multiple variables in one go using tuple assignment
	// prefer not to use just for declarion.
	a, b := 1, 2
	fmt.Println(a, b)
	// more usefull when a function returns mutliple values
	x, y, z := getPositon()
	fmt.Println(x, y, z)
	//blank identifier
	_ = "Hello"
	//fmt.Println(_)
	// compile error: cannot use _ as value
	//You can use the blank identifier more than once in the same statement. x
	_, _, z = getPositon()

	//x := y++
	// syntax error: unexpected ++ at end of statement

	// special if syntax
	if err := doStuff(); err != nil {
		// it restricts the scope of err to just this block
		fmt.Println(err)
	}

	// switch
	switch {
	case x < 0:
		fmt.Println("negative")
	case x > 0:
		fmt.Println("positive")
	default:
		fmt.Println("zero")
	}

	// match aganist one of multiple values
	switch x {
	case 1, 2, 3:
		fmt.Println("one, two, or three")
	case 4:
		//want to leave early, without executing the rest of the case. Similar NOT same as C++
		if somethingWrong() {
			break
		}
		// otherwise carry on

	}

	name := "Dev"
	switch name {
	case "Dev":
		fmt.Println("Hey Dev")
	default:
		fmt.Println("Supppp")
	}

	//Using range to loop over collections
	fruits := []string{"Apple", "Mango", "Orange", "Guava", "Banana"}
	// range returns two values index and the element itself , you can you _ for index if it's of no
	// no use or just give one var for index if the accesing the element is not necessary
	for index, fruit := range fruits {
		fmt.Printf("Index %d fruit %s\n", index, fruit)
	}

	// conditonal for loops
	for x < 10 {
		fmt.Println("x is less than 10")
		//network servers such as HTTP webservers run an endless loop that waits for a
		//request to arrive, services the request, and then goes back to waiting.
		x++
	}
	// compact version of above but rarley used in GO, mostly GO uses range for collections or forever
	for e := 0; e <= 100; e += 2 {
		fmt.Printf(" %d ", e)
		if e%10 == 0 && e != 0 {
			fmt.Println("")
		}
	}

	// Sample list of employees
	employees := []Employee{
		{"Alice", true},
		{"Bob", false},
		{"Charlie", true},
		{"David", true},
		{"Eve", false},
	}

	//Jumping to the next element with continue
	// Iterate over employees using range
	for _, e := range employees {
		// happy path instead nested if using continue to maintain readibility of code
		///happy path runs in a straight line down the left margin, and the only indented code is
		// for handling exceptional or invalid situations
		if !e.IsCurrent {
			continue
		}
		fmt.Printf("%s is current employee\n", e.Name)
	}
	signedNums := []int{2, -1, 0, -4}
	for _, i := range signedNums {
		if i < 0 {
			continue
		}
		fmt.Printf(" %d ", i)
	}
	//NOTE:Exit loops with break

	// Controlling nested loops with labels
	// labels are rarely used as there missue can easily lead to control flow liek bowl of noodles.
outer:
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			fmt.Println(x, y)
			if y == 5 {
				// use can use continue outer to continue the outer loop straight away
				break outer
			}
		}
	}

	// function literals
	nums := []int{2, 3, 1, 6, 9, 4}
	sort.Slice(nums, func(i, j int) bool {
		// calling function, can “see” all variables that are in scope in that sort.Slice(),
		// including nums. thus we can access nums[i] here. This is called closure.
		// Any func‐ tion has access to the variables available in the scope where it’s defined,
		// and when this is a function literal, we call it a closure over those variables.
		return nums[i] < nums[j]
	})
	// you will 1 , 2 , 3
	for _, v := range []int{1, 2, 3} {
		func() {
			fmt.Println(v)
		}() //We don’t defer a function, but a function call, so we add the parentheses to call it.
	}
	// but here
	funcs := []func(){}
	for _, v := range []int{1, 2, 3} {
		funcs = append(funcs, func() { fmt.Println(v) })
	}
	//  what do you get ?
	for _, f := range funcs {
		f()
		//we set the variable v to the values 1, 2, and 3, in succession.
		//So after the first loop has finished, v has the value 3
	}

	// defer
	defer cleanup1() // executed last on exit
	defer cleanup2() // executed first on exit

	// defere and closure power combo
	err := Hellotxt()
	if err != nil {
		log.Fatalln(err)
	}

	// vardiac functions , function with variable number of parameters like pritnln()
	fmt.Println(AddMany(1, 2, 3, 4, 5, 7, 3, 7, 9), AddMany(1, 1, 1, 1, 1, 1, 1, 1, 1, 1))
	fmt.Println(AddMany(1, 1, 1, 1, 1, 1, 1, 1, 1, 1))
}

// If we named this function’s error result parameter
func Hellotxt() (err error) {
	// Open a file for writing
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	//
	defer func() {
		closeErr := file.Close()
		if closeErr != nil {

			err = closeErr
		}
	}()
	// Write some data to the file
	_, err = file.WriteString("Hello, Golang!")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}
	return err
}

// vardiac functions
// ... indicates that there can be zero, one, two, or any number of parameters to the function.
func AddMany(inputs ...float64) float64 {
	var sum float64
	for _, input := range inputs {
		sum += input
	}
	return sum
}
