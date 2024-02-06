package Luhn

import (
	_"fmt"
	"strconv"
)

func LuhnAlgorithm(cardNumber string) bool {
	sum := 0
	isSecondDigit := false

	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(cardNumber[i]))

		if isSecondDigit {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		isSecondDigit = !isSecondDigit
	}

	return sum%10 == 0
}


