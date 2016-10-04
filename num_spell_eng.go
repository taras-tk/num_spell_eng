package num_spell_eng

import (
	"math"
	"regexp"
	"strings"
)

const TEN = 10
const ONE_HUNDRED = 100
const ONE_THOUSAND = 1000
const ONE_MILLION = 1000000
const ONE_BILLION = 1000000000           //         1.000.000.000
const ONE_TRILLION = 1000000000000       //     1.000.000.000.000
const ONE_QUADRILLION = 1000000000000000 // 1.000.000.000.000.000
const MAX = 9007199254740992             // 9.007.199.254.740.992

var LESS_THAN_TWENTY = []string{
	"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten",
	"eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen",
}

var TENTHS_LESS_THAN_HUNDRED = []string{
	"zero", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety",
}

func NumberSpell(number int, arguments ...[]string) string {
	var remainder int
	var word string
	var words []string

	if len(arguments) == 0 {
		words = []string{}
	} else {
		words = arguments[0]
	}

	// Recursion's done
	if (number == 0) {

		if len(words) == 0 {
			return string("zero")
		} else {
			result := strings.Join(words, " ")

			re := regexp.MustCompile(`,$`)
			result = re.ReplaceAllString(result, "")

			re = regexp.MustCompile(` and$`)
			result = re.ReplaceAllString(result, "")

			return result
		}
	}

	// If number is negative, prepend "minus"
	if (number < 0) {
		words = append(words, "minus")
		number = int(math.Abs(float64(number)))
	}

	if (number < 20) {
		remainder = 0
		word = LESS_THAN_TWENTY[number]

	} else if (number < ONE_HUNDRED) {
		remainder = number % TEN
		word = TENTHS_LESS_THAN_HUNDRED[int(math.Floor(float64(number / TEN)))]

		// In case of remainder, we need to handle it here to be able to add the “-”
		if (remainder > 0) {
			word += "-" + LESS_THAN_TWENTY[remainder]
			remainder = 0
		}

	} else if (number < ONE_THOUSAND) {
		remainder = number % ONE_HUNDRED
		word = NumberSpell(int(math.Floor(float64(number / ONE_HUNDRED)))) + " hundred and"

	} else if (number < ONE_MILLION) {
		remainder = number % ONE_THOUSAND
		word = NumberSpell(int(math.Floor(float64(number / ONE_THOUSAND)))) + " thousand,"

	} else if (number < ONE_BILLION) {
		remainder = number % ONE_MILLION
		word = NumberSpell(int(math.Floor(float64(number / ONE_MILLION)))) + " million,"

	} else if (number < ONE_TRILLION) {
		remainder = number % ONE_BILLION
		word = NumberSpell(int(math.Floor(float64(number / ONE_BILLION)))) + " billion,"

	} else if (number < ONE_QUADRILLION) {
		remainder = number % ONE_TRILLION
		word = NumberSpell(int(math.Floor(float64(number / ONE_TRILLION)))) + " trillion,"

	} else if (number <= MAX) {
		remainder = number % ONE_QUADRILLION
		word = NumberSpell(int(math.Floor(float64(number / ONE_QUADRILLION)))) + " quadrillion,"
	}

	words = append(words, word)
	return NumberSpell(remainder, words)
}

