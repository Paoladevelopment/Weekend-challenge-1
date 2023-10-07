package foobar

import "strconv"

func parseNumberToFoobar(num int) string {
	if num%3 == 0 && num%5 == 0 {
		return "FooBar"
	} else if num%3 == 0 {
		return "Foo"
	} else if num%5 == 0 {
		return "Bar"
	} else {
		return strconv.Itoa(num)
	}
}

func Foobar(num int) string {
	var fooBarRes string
	for i := 1; i <= num; i++ {
		fooBarRes += parseNumberToFoobar(i)
		if i < num {
			fooBarRes += "->"
		}
	}
	return fooBarRes

}