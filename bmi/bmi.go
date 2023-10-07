package bmi

import (
	"math"
	"strconv"
)

func calculateBmi(weight float64, height float64) float64 {
	return weight/math.Pow(height, 2)
}

func BmiCalculator(weight float64, height float64) string {

	bmi := calculateBmi(weight, height)
	finalMessage := "Right now your BMI is " + strconv.FormatFloat(calculateBmi(weight, height), 'f', 2, 64) + "\n"
	if bmi < 18.5 {
		finalMessage += "You are underweight, add more potato to the broth"
	} else if bmi >= 18.5 && bmi < 25 {
		finalMessage += "You have a normal weight, I have healthy envy of you"
	} else if bmi >= 25 {
		finalMessage += "You are overweight, I know, the pandemic has affected us all"
	}

	return finalMessage
}