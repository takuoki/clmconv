package clmconv_test

import (
	"fmt"

	"github.com/takuoki/clmconv"
)

func Example() {

	// Default converter
	i, _ := clmconv.Atoi("A")
	a := clmconv.Itoa(0)
	fmt.Printf("i=%d, a='%s'\n", i, a)

	// Custom converter
	converter := clmconv.New(clmconv.WithStartFromOne(), clmconv.WithLowercase())
	i, _ = converter.Atoi("A")
	a = converter.Itoa(1)
	fmt.Printf("i=%d, a='%s'\n", i, a)

	// Output:
	// i=0, a='A'
	// i=1, a='a'
}

func ExampleAtoi() {
	ai, _ := clmconv.Atoi("A")
	bi, _ := clmconv.Atoi("B")
	_, err := clmconv.Atoi("error!")
	fmt.Printf("ai=%d, bi=%d, err='%v'", ai, bi, err)

	// Output:
	// ai=0, bi=1, err='must not contain non-alphabetic characters'
}

func ExampleMustAtoi() {
	ai := clmconv.MustAtoi("A")
	v := func() (v interface{}) {
		defer func() {
			if p := recover(); p != nil {
				v = p
			}
		}()
		clmconv.MustAtoi("panic!")
		return nil
	}()
	fmt.Printf("ai=%d, v='%v'", ai, v)

	// Output:
	// ai=0, v='must not contain non-alphabetic characters'
}

func ExampleItoa() {
	a0 := clmconv.Itoa(0)
	a1 := clmconv.Itoa(1)
	an := clmconv.Itoa(-1)
	fmt.Printf("a0='%s', a1='%s', an='%s'", a0, a1, an)

	// Output:
	// a0='A', a1='B', an=''
}
