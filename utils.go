package main

import (
	"strconv"
	"strings"
)

func FormatCurrency(value float64) string {
	s := strconv.FormatFloat(value, 'f', 0, 64)
	parts := []string{}
	for i := len(s); i > 0; i -= 3 {
		start := i - 3
		if start < 0 {
			start = 0
		}
		parts = append([]string{s[start:i]}, parts...)
	}
	return strings.Join(parts, ".")
}
