package byte

import (
	"testing"
)

type TestIO struct {
	input          string
	expectedResult []int
}

func TestStringToByteSequence(T *testing.T) {
	testStringCaps := TestIO{input: "FF", expectedResult: []int{255}}
	testStringLowerCase := TestIO{input: "ff", expectedResult: []int{255}}
	testStringCombined := TestIO{input: "f3", expectedResult: []int{243}}
	testStringDoubleByte := TestIO{input: "FFAA", expectedResult: []int{255, 250}}

	cases := []TestIO{testStringCaps, testStringLowerCase, testStringCombined, testStringDoubleByte}

	for _, test := range cases {
		result := StringToByteSequence(test.input)

	}

}
