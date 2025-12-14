package main

import (
	"math"
	"strings"
)

// 1
func bubbleSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		swap := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
}

// 2
func palindrome(str string) bool {
	word := []rune(str)

	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		if word[i] != word[j] {
			return false
		}
	}
	return true
}

// 3
func factorial(num int) int {
	if num == 0 {
		return 1
	}
	temp := 1
	for i := 1; i <= num; i++ {
		temp *= i
	}
	return temp
}

// 4
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 5
func findSubString(str, substr string) bool {
	return strings.Contains(str, substr)
}

// 6
func numberPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i++ {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// 7
func reverseMath32(x int) int {
	var result int32 = 0
	num := int32(x)

	for num != 0 {
		digit := num % 10
		num /= 10
		if result > math.MaxInt32/10 || result < math.MinInt32/10 {
			return 0
		}
		result = result*10 + digit
	}

	return int(result)
}

// 8
func intToRim(num int) string {
	if num <= 0 || num > 3999 {
		return ""
	}

	var result strings.Builder

	for num > 0 {
		switch {
		case num >= 1000:
			result.WriteString("M")
			num -= 1000
		case num >= 900:
			result.WriteString("CM")
			num -= 900
		case num >= 500:
			result.WriteString("D")
			num -= 500
		case num >= 400:
			result.WriteString("CD")
			num -= 400
		case num >= 100:
			result.WriteString("C")
			num -= 100
		case num >= 90:
			result.WriteString("XC")
			num -= 90
		case num >= 50:
			result.WriteString("L")
			num -= 50
		case num >= 40:
			result.WriteString("XL")
			num -= 40
		case num >= 10:
			result.WriteString("X")
			num -= 10
		case num >= 9:
			result.WriteString("IX")
			num -= 9
		case num >= 5:
			result.WriteString("V")
			num -= 5
		case num >= 4:
			result.WriteString("IV")
			num -= 4
		default:
			result.WriteString("I")
			num -= 1
		}
	}

	return result.String()
}

func main() {

}
