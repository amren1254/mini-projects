package main

import (
	"fmt"
	"strings"
)

var ones = map[int]string{
	0: "zero",
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

var teens = map[int]string{
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
}

var tens = map[int]string{
	0: "",
	2: "twenty",
	3: "thirty",
	4: "forty",
	5: "fifty",
	6: "sixty",
	7: "seventy",
	8: "eighty",
	9: "ninety",
}

var powersOfTen = map[int]string{
	0:  "",
	1:  "thousand",
	2:  "lakh",
	3:  "crore",
	4:  "arab",
	5:  "kharab",
	6:  "neel",
	7:  "padma",
	8:  "shankh",
	9:  "mahashankh",
	10: "vrind",
	11: "jaladhi",
	12: "antahstha",
	13: "shishira",
	14: "shankh",
	15: "vriddhi",
	16: "sanku",
	17: "megh",
	18: "pran",
	19: "mahapran",
	20: "antarbindu",
}

func convertLessThanThousand(num int) string {
	var result string

	hundreds := num / 100
	tensAndOnes := num % 100

	if hundreds > 0 {
		result += ones[hundreds] + " hundred"
		if tensAndOnes > 0 {
			result += " and "
		}
	}

	if tensAndOnes < 10 {
		result += ones[tensAndOnes]
	} else if tensAndOnes < 20 {
		result += teens[tensAndOnes]
	} else {
		result += tens[tensAndOnes/10]
		if tensAndOnes%10 > 0 {
			result += "-" + ones[tensAndOnes%10]
		}
	}

	return result
}

func numberToWords(num int) string {
	if num == 0 {
		return "zero"
	}

	var result string
	var i int

	for num > 0 {
		if num%1000 != 0 {
			group := convertLessThanThousand(num % 1000)
			if group != "" {
				result = group + " " + powersOfTen[i] + " " + result
			}
		}
		num /= 1000
		i++
	}

	result = strings.TrimSpace(result)
	return result
}

func main() {
	num := 4784732
	words := numberToWords(num)
	fmt.Printf("%d in words: %s\n", num, words)
}
