package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func validateCC(ccNum []int) (valid bool, cardType string) {
	valid = luhnCheck(ccNum)
	cardType = getCardType(ccNum)

	return valid, cardType
}

func getCardType(ccNum []int) string {
	sSlice := make([]string, len(ccNum))
	for i, n := range ccNum {
		sSlice[i] = strconv.Itoa(n)
	}
	s := strings.Join(sSlice, "")

	match, _ := regexp.MatchString("^4[0-9]{12}(?:[0-9]{3})?$", s)
	if match {
		return "Visa"
	}

	match, _ = regexp.MatchString("^5[1-5][0-9]{14}$", s)
	if match {
		return "MasterCard"
		
	}

	return "Unknown"
}

func luhnCheck(ccNum []int) bool {
	checksum := 0
	for i, n := range ccNum {
		if (i + len(ccNum) % 2) % 2 == 0 {
			checksum += sumDigits(2 * n)
		} else {
			checksum += n
		}
	}

	return checksum % 10 == 0
}

func sumDigits(n int) int {
	hundreds := n / 100
	tens := (n - 100 * hundreds) / 10
	ones := n - 100 * hundreds - 10 * tens

	return ones + tens + hundreds
}

func main() {
	fmt.Println("Please, enter the credit card number: ")
	in := bufio.NewReader(os.Stdin)
	input, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	ccNum := make([]int, 19)
	i := 0
	for _, rune := range input {
		n, err := strconv.Atoi(string(rune))
		if err == nil {
			ccNum[i] = n
			i++
		}

	}
	ccNum = ccNum[:i]
	
	valid, cardType := validateCC(ccNum)

	if valid && cardType != "Unknown" {
		fmt.Println("The credit card number is a valid", cardType, "card.")
	} else {
		fmt.Println("Credit card number is NOT valid!")
	}
}