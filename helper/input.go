package helper

import "fmt"

func InputInRange(limInf, limSup int, helpMessage string) int {
	isValid := false
	var num int
	for !isValid {
		fmt.Print(helpMessage)
		fmt.Scan(&num)

		if num >= limInf && num <= limSup {
			isValid = true
		}
	}

	return num
}