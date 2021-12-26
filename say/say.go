package say

import (
	"fmt"
	"strings"
)

type NumStruct struct {
	billions int64
	millions int64
	thousands int64
	hundreds int64
}

func CreateNumsMap() map[int64]string {
	return map[int64]string{
		0: "zero",
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "fifteen",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
		20: "twenty",
		30: "thirty",
		40: "forty",
		50: "fifty",
		60: "sixty",
		70: "seventy",
		80: "eighty",
		90: "ninety",
	}
}

func Ones2Tens(n int64) string {
	var result string
	nums := CreateNumsMap()
	// First dividing factor is if we're single or two digit
	if n>=10 { // 10-99
		ones := n % 10
		tens := n / 10
		// Now, less than 20 is treated different thanks to english
		if tens == 1 { // 10-19
			return nums[n]
		} else { // 20-99
			result += nums[tens*10] 
			// Now, if ones = 0, we shouldn't add ones
			if ones != 0 {
				result +=  "-" + nums[ones]
			}
		}
	} else if n<10 { // 0-9
		return nums[n]
	}
	return result
}

func Ones2Hundreds(n int64) string {
	var result string
	nums := CreateNumsMap()
	// First dividing factor is if we're three digit or less
	if n>=100 { // 100-999
		hundreds := n / 100
		tensandones := n-(hundreds*100)
		result += nums[hundreds] + " hundred"
		// Now, if tensandones = 0, we shouldn't add anything
		if tensandones != 0 {
			result += " " + Ones2Tens(tensandones)
		}
	} else if n<100 { // 0-99
		result = Ones2Tens(n) // Delegate to ones2tens
	}
	return result
}

func SplitNum(n int64) NumStruct {
	var result NumStruct

	billions := n / 1000000000
	millions := n / 1000000
	thousands := n / 1000

	result.billions = billions
	result.millions = millions - 1000 * billions
	result.thousands = thousands - 1000 * millions
	result.hundreds = n - 1000 * thousands  
	
	return result
}

func GenerateName(n NumStruct) string {
	var result string
	
	if n.billions != 0 {
		result += Ones2Hundreds(n.billions) + " billion "
	}
	if n.millions != 0 {
		result += Ones2Hundreds(n.millions) + " million "
	}
	if n.thousands != 0 {
		result += Ones2Hundreds(n.thousands) + " thousand "
	}
	if n.thousands != 0 {
		result += Ones2Hundreds(n.hundreds)
	}

	return strings.TrimSpace(result)
}

func Say(n int64) (string, bool) {
	// Error handling
	if n < 0 || n > 999999999999 {
		return "", false
	}
	// 0 - 99
	fmt.Println(Ones2Hundreds(0))
	fmt.Println(Ones2Hundreds(9))
	fmt.Println(Ones2Hundreds(10))
	fmt.Println(Ones2Hundreds(14))
	fmt.Println(Ones2Hundreds(19))
	fmt.Println(Ones2Hundreds(20))
	fmt.Println(Ones2Hundreds(50))
	fmt.Println(Ones2Hundreds(51))
	fmt.Println(Ones2Hundreds(69))
	fmt.Println(Ones2Hundreds(90))
	fmt.Println(Ones2Hundreds(98))
	// 100 - 999
	fmt.Println(Ones2Hundreds(100))
	fmt.Println(Ones2Hundreds(123))
	fmt.Println(Ones2Hundreds(999))

	// Split logic
	fmt.Println(SplitNum(987654321123))
	fmt.Println(SplitNum(1000000000))
	fmt.Println(SplitNum(1002345))
	fmt.Println(SplitNum(1000000))
	fmt.Println(SplitNum(1234))
	fmt.Println(SplitNum(1000))
	fmt.Println(SplitNum(123))
	fmt.Println(SplitNum(100))
	fmt.Println(SplitNum(22))
	fmt.Println(SplitNum(1))

	// Test Name
	fmt.Println(GenerateName(SplitNum(987654321123)))
	fmt.Println(GenerateName(SplitNum(1000000000)))

	return "", false
}
