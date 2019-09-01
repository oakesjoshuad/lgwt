package main

import "strings"

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

type RomanNumerals []RomanNumeral

var (
	romanNumerals = RomanNumerals{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
)

func (r RomanNumerals) ValueOf(symbol string) uint16 {
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}
	return 0
}

func ConvertToRoman(arabic uint16) string {
	var result strings.Builder

	for _, numeral := range romanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) uint16 {
	if len(roman) > 1 {
		value := romanNumerals.ValueOf(string(roman[:2]))
		if 0 != value {
			return value + ConvertToArabic(string(roman[2:]))
		} else {
			value = romanNumerals.ValueOf(roman[:1])
			return value + ConvertToArabic(string(roman[1:]))
		}
	}
	return romanNumerals.ValueOf(roman)
}
