package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// dictionary Rome Numbers -> Arabic Numbers
var romeNumbers map[string]int

// dictionary Arabic Numbers -> Rome Numbers
var arabicNumbers map[int]string

// makes arabic number from string
func makeNumber(rawNumber string) (int, bool) {
	var isRome bool = false
	number, err := strconv.Atoi(rawNumber)
	if err != nil {
		isRome = true
		for i := 0; i < len(rawNumber)-1; i++ {
			val, ok := romeNumbers[string(rawNumber[i])]
			if ok {
				if romeNumbers[string(rawNumber[i])] < romeNumbers[string(rawNumber[i+1])] {
					number -= int(val)
				} else {
					number += int(val)
				}
			} else {
				fmt.Fprintf(os.Stderr, "Something in your equation is not a number!\n")
				os.Exit(1)
			}
		}
		_, ok := romeNumbers[string(rawNumber[len(rawNumber)-1])]
		if !ok {
			fmt.Fprintf(os.Stderr, "Something in your equation is not a number!\n")
			os.Exit(1)
		}
		number += romeNumbers[string(rawNumber[len(rawNumber)-1])]
	}

	if number < 1 || number > 10 {
		fmt.Fprintf(os.Stderr, "Number should be from 1 to 10 inclusively!\n")
		os.Exit(1)
	}

	return number, isRome
}

// makes Roman number from Arabic number
func makeRoman(rawRes string) string {
	var res string = ""
	for i := 0; i < len(rawRes); i++ {
		number, _ := strconv.Atoi(string(rawRes[i]))
		position := len(rawRes) - i - 1
		switch {
		case number < 4 && number > 0:
			for j := 0; j < number; j++ {
				res += arabicNumbers[int(math.Pow(10, float64(position)))]
			}
		case number == 4:
			res += arabicNumbers[int(math.Pow(10, float64(position)))]
			res += arabicNumbers[5*int(math.Pow(10, float64(position)))]
		case number == 5:
			res += arabicNumbers[5*int(math.Pow(10, float64(position)))]
		case number < 9 && number > 5:
			res += arabicNumbers[5*int(math.Pow(10, float64(position)))]
			for j := 0; j < number-5; j++ {
				res += arabicNumbers[int(math.Pow(10, float64(position)))]
			}
		case number == 9:
			res += arabicNumbers[int(math.Pow(10, float64(position)))]
			res += arabicNumbers[int(math.Pow(10, float64(position+1)))]
		default:
			res += ""
		}
	}
	return res
}

// sum of Arabic numbers
func add(x int, y int) int {
	return x + y
}

// sum of Roman numbers
func addRoman(x int, y int) string {
	rawRes := strconv.Itoa(x + y)
	return makeRoman(rawRes)
}

// subtraction of Arabic numbers
func sub(x int, y int) int {
	return x - y
}

// subtraction of Roman numbers
func subRoman(x int, y int) string {
	if x - y < 1 {
		fmt.Fprintf(os.Stderr, "Subtraction of Roman numbers cannot be 0 or lower!\n")
		os.Exit(1)
	}
	rawRes := strconv.Itoa(x - y)
	return makeRoman(rawRes)
}

// multiplication of Arabic numbers
func mul(x int, y int) int {
	return x * y
}

// multiplication of Roman numbers
func mulRoman(x int, y int) string {
	rawRes := strconv.Itoa(x * y)
	return makeRoman(rawRes)
}

// division of Arabic numbers
func div(x int, y int) int {
	if y == 0 {
		fmt.Fprintf(os.Stderr, "You cannot divide by 0!\n")
		os.Exit(1)
	}
	return x / y
}

// division of Roman numbers
func divRoman(x int, y int) string {
	rawRes := strconv.Itoa(x / y)
	return makeRoman(rawRes)
}

func main() {

	romeNumbers = make(map[string]int)
	romeNumbers["I"] = 1
	romeNumbers["V"] = 5
	romeNumbers["X"] = 10
	romeNumbers["L"] = 50
	romeNumbers["C"] = 100
	romeNumbers["D"] = 500
	romeNumbers["M"] = 1000

	arabicNumbers = make(map[int]string)
	arabicNumbers[1] = "I"
	arabicNumbers[5] = "V"
	arabicNumbers[10] = "X"
	arabicNumbers[50] = "L"
	arabicNumbers[100] = "C"
	arabicNumbers[500] = "D"
	arabicNumbers[1000] = "M"

	//To make possible reading from console
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter smth to calculate")

	//Processing of input data
	inputData, _ := reader.ReadString('\n')
	inputData = strings.TrimSuffix(inputData, "\n")
	var rawResult []string = strings.Split(inputData, " ")

	if len(rawResult) != 3 {
		fmt.Fprintf(os.Stderr, "It's not mathematical operation or number of operations greater, than 1!\n")
		os.Exit(1)
	}

	//converts string to Arabic numbers and check, were they Roman or not
	number1, isRomanNumber1 := makeNumber(rawResult[0])
	number2, isRomanNumber2 := makeNumber(rawResult[2])

	if isRomanNumber1 != isRomanNumber2 {
		fmt.Fprintf(os.Stderr, "Numbers should be both Roman or Arabic!\n")
		os.Exit(1)
	}

	//Just Operations with Numbers
	switch rawResult[1] {
	case "+":
		if isRomanNumber1 {
			fmt.Println(addRoman(number1, number2))
		} else {
			fmt.Println(add(number1, number2))
		}
	case "-":
		if isRomanNumber1 {
			fmt.Println(subRoman(number1, number2))
		} else {
			fmt.Println(sub(number1, number2))
		}
	case "*":
		if isRomanNumber1 {
			fmt.Println(mulRoman(number1, number2))
		} else {
			fmt.Println(mul(number1, number2))
		}
	case "/":
		if isRomanNumber1 {
			fmt.Println(divRoman(number1, number2))
		} else {
			fmt.Println(div(number1, number2))
		}
	default:
		fmt.Fprintf(os.Stderr, "There is no logical operation!\n")
		os.Exit(1)
	}
	fmt.Println("Here is your result above!!!")
}
