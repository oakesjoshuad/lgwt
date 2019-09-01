package main

import (
	"fmt"
	"testing"
	"testing/quick"
)

func TestRomanNumerals(t *testing.T) {

	for _, test := range cases {
		t.Run(descitor(test.Arabic, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			assertRoman(t, got, test.Roman)
		})
	}
}

func TestRomanToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(descrtoi(test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			assertNumeral(t, got, test.Arabic)
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log("testing", arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{MaxCount: 10000}); err != nil {
		t.Error("failed check", err)
	}
}

func descitor(arabic uint16, roman string) string {
	return fmt.Sprintf("%d gets converted to %q", arabic, roman)
}

func descrtoi(roman string, arabic uint16) string {
	return fmt.Sprintf("%q gets converted to %d", roman, arabic)
}

func assertNumeral(t *testing.T, got, want uint16) {
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func assertRoman(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

var (
	cases = []struct {
		Arabic uint16
		Roman  string
	}{
		{Arabic: 1, Roman: "I"},
		{Arabic: 2, Roman: "II"},
		{Arabic: 3, Roman: "III"},
		{Arabic: 4, Roman: "IV"},
		{Arabic: 5, Roman: "V"},
		{Arabic: 6, Roman: "VI"},
		{Arabic: 7, Roman: "VII"},
		{Arabic: 8, Roman: "VIII"},
		{Arabic: 9, Roman: "IX"},
		{Arabic: 10, Roman: "X"},
		{Arabic: 14, Roman: "XIV"},
		{Arabic: 18, Roman: "XVIII"},
		{Arabic: 20, Roman: "XX"},
		{Arabic: 39, Roman: "XXXIX"},
		{Arabic: 40, Roman: "XL"},
		{Arabic: 47, Roman: "XLVII"},
		{Arabic: 49, Roman: "XLIX"},
		{Arabic: 50, Roman: "L"},
		{Arabic: 100, Roman: "C"},
		{Arabic: 90, Roman: "XC"},
		{Arabic: 400, Roman: "CD"},
		{Arabic: 500, Roman: "D"},
		{Arabic: 900, Roman: "CM"},
		{Arabic: 1000, Roman: "M"},
		{Arabic: 1984, Roman: "MCMLXXXIV"},
		{Arabic: 3999, Roman: "MMMCMXCIX"},
		{Arabic: 2014, Roman: "MMXIV"},
		{Arabic: 1006, Roman: "MVI"},
		{Arabic: 798, Roman: "DCCXCVIII"},
	}
)
