package main

import (
	"fmt"
	"strings"
)

func FormatCurrency(amount float64) string {
	formatted := fmt.Sprintf("%.2f", amount)

	parts := strings.Split(formatted, ".")
	integerPart := parts[0]
	decimalPart := ""
	if len(parts) > 1 {
		decimalPart = "," + parts[1]
	} else {
		decimalPart = ",00"
	}

	var formattedInteger strings.Builder
	n := len(integerPart)
	for i := 0; i < n; i++ {
		formattedInteger.WriteByte(integerPart[i])
		if (n-i-1)%3 == 0 && i < n-1 {
			formattedInteger.WriteByte('.')
		}
	}

	return formattedInteger.String() + decimalPart
}
