package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	result, roman := expressionParser()
	if roman {
		fmt.Println(IToRoman(result))
	} else {
		fmt.Println(result)
	}
}

func expressionParser() (int, bool) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите выражение:")
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	text = strings.Replace(text, " ", "", -1)
	expr := ExpressionSplitter(text)
	return Etoi(expr)

}

func ExpressionSplitter(text string) [3]string {
	exprLen := len(text)
	for ind := 0; ind < exprLen; ind++ {
		if text[ind] == '+' || text[ind] == '-' || text[ind] == '*' || text[ind] == '/' {
			return [3]string{text[:ind], text[ind : ind+1], text[ind+1 : exprLen-1]}
		}

	}
	panic("Unknown Operator!")
}

func Etoi(expr [3]string) (result int, roman bool) {
	roman, arabic := false, false
	nums := [2]int{}

	for i := 0; i < 3; i += 2 {
		num, err := strconv.Atoi(expr[i])
		if err != nil {
			num = RomanToI(expr[i])
			roman = true
		} else {
			arabic = true
		}
		nums[i/2] = num
	}

	if roman == arabic {
		panic("Can't process, arabic and roman at once!")
	}
	// исскуственное ограничение на числа с величиной от 1 до 10
	if nums[0] > 10 || nums[0] < 1 || nums[1] > 10 || nums[1] < 1 {
		panic("Can't process, num out of domain [1, 10]")
	}

	switch expr[1] {
	case "+":
		result = nums[0] + nums[1]
	case "-":
		result = nums[0] - nums[1]
	case "*":
		result = nums[0] * nums[1]
	case "/":
		result = nums[0] / nums[1]
	}

	return

}

func RomanToI(numRaw string) (arabic int) {

	numRaw = strings.ToLower(numRaw)

	for ind, digitLast := len(numRaw)-1, 0; ind >= 0; ind-- {
		digit := 0
		switch numRaw[ind] {
		case 'i':
			digit = 1
		case 'v':
			digit = 5
		case 'x':
			digit = 10
		case 'l':
			digit = 50
		case 'c':
			digit = 100
		case 'd':
			digit = 500
		case 'm':
			digit = 1000
		default:
			panic("Unknown numeral literal!")
		}
		if digit < digitLast {
			arabic -= digit
		} else {
			arabic += digit
		}
		digitLast = digit

	}
	return
}

func IToRoman(num int) (roman string) {
	romans := [13]string{"m", "cm", "d", "cd", "c", "xc", "l", "xl", "x", "ix", "v", "iv", "i"}
	values := [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	for i := 0; i < 13; i++ {
		times := num / values[i]
		num -= times * values[i]
		roman += strings.Repeat(romans[i], times)
	}
	return
}
