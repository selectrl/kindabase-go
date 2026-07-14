package kindabase

import (
	"math"
	"slices"
	"strings"
)

type Alphabet []byte

var (
	Base26 Alphabet = []byte{
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
		'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
		'u', 'v', 'w', 'x', 'y', 'z'}
	Base36 Alphabet = []byte{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
		'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
		'u', 'v', 'w', 'x', 'y', 'z'}
	Base62 Alphabet = []byte{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
		'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
		'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd',
		'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n',
		'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x',
		'y', 'z'}
)

type Integer interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64
}

func Decode(s string, alphabet Alphabet) int {
	chars := []byte(s)

	var negative bool
	if chars[0] == '-' {
		negative = true
		chars = chars[1:]
	}

	slices.Reverse(chars)
	l := len(alphabet)
	var result, place int
	for _, char := range chars {
		result += (slices.Index(alphabet, char) + 1) * int(math.Pow(float64(l), float64(place)))
		place++
	}

	result--

	if negative {
		return -result
	}
	return result
}

func Encode[T Integer](integer T, alphabet Alphabet) string {
	var sign byte
	if integer < 0 {
		sign = '-'
		integer = -integer
	}

	l := T(len(alphabet))
	var chars []byte
	var newInteger, mod T
	for integer >= 0 {
		newInteger = integer / l
		mod = integer % l
		chars = append(chars, alphabet[mod])
		// Add 0 check so unsigned integers can be used without underflow
		if newInteger == 0 {
			break
		}
		integer = newInteger - 1
	}

	if sign == '-' {
		chars = append(chars, sign)
	}
	slices.Reverse(chars)
	return string(chars)
}

func Decode26(s string) int {
	// As this is case-insensitive lower string to match alphabet
	return Decode(strings.ToLower(s), Base26)
}

func Decode36(s string) int {
	// As this is case-insensitive lower string to match alphabet
	return Decode(strings.ToLower(s), Base36)
}

func Decode62(s string) int {
	return Decode(s, Base62)
}

func Encode26[T Integer](integer T) string {
	return Encode(integer, Base26)
}

func Encode36[T Integer](integer T) string {
	return Encode(integer, Base36)
}

func Encode62[T Integer](integer T) string {
	return Encode(integer, Base62)
}
