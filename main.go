package main

import (
	"fmt"
	"os"

	"truecode.com/challenge1/bmi"
	"truecode.com/challenge1/evenodd"
	"truecode.com/challenge1/foobar"
	"truecode.com/challenge1/helper"
	"truecode.com/challenge1/mario"
	"truecode.com/challenge1/ohm"
)

func main() {
	args := os.Args
	functionName := args[1]

	switch functionName {
		case "evenodd":
			var num int
			fmt.Println("Type a number to evaluate if it is even or odd: ")
			fmt.Scan(&num)
			fmt.Printf("%v is a %v number", num, evenodd.EvenOrOdd(num))
		
		case "foobar":
			var num int
			fmt.Println("Type a number to convert it to foobar format")
			fmt.Scan(&num)
			fmt.Printf("%v in foobar format is: %q",num,foobar.Foobar(num))
		
		case "ohm":
			var v float32
			var r float32
			var i float32
			fmt.Println("Ohm calculation. Set 0 to the variable you want to find a result")
			fmt.Println("Example: if you set 'v' to 0, is to find voltage")
			fmt.Println("Type the values for v, r and i")
			fmt.Print("v: ")
			fmt.Scan(&v)
			fmt.Print("r: ")
			fmt.Scan(&r)
			fmt.Print("i: ")
			fmt.Scan(&i)
			fmt.Printf("Ohm calculation for given values, v:%f, r:%f, i:%f is: %f", v, r, i, ohm.Ohm(v,r,i))
		
		case "bmi":
			var weight float64
			var height float64
			fmt.Println("--BMI CALCULATOR--")
			fmt.Println("How much do you weigh? (don't lie)")
			fmt.Scan(&weight)
			fmt.Println("How tall are you? (barefoot)")
			fmt.Scan(&height)
			fmt.Println(bmi.BmiCalculator(weight, height))
		
		case "mario":
			var height int
			fmt.Println("Lets print a  right-aligned pyramid made of hashes")
			height = helper.InputInRange(1, 8, "Pyramid height: ")
			fmt.Print(mario.PrintMarioPyramid(height))
	}
}

