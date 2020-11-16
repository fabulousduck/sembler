package byte

import (
	"testing"
	"strconv"
	"fmt"
)

type TestIO struct {
	input          string
	expectedResult []int
}

func sliceEq(s1 []int, s2 []int) bool {
	if(len(s1) != len(s2)) {
		return false
	}

	for idx, v := range s1 {
		if(v != s2[idx]) {
			return false
		}
	}
	
	return true
}

func formatSliceToString(s1 []int) string {
	var outputString string
	for _, v := range s1 {
		strVal := strconv.Itoa(v)
		outputString += strVal + ", "
	}
	return outputString
}

func TestGetByteForChar(T *testing.T) {
	for i := 0; i < 0xF; i++ {
		stringForm := strconv.Itoa(i)
		result := getByteForChar(stringForm)
		stringResult := strconv.Itoa(result)
		if(result != i) {
			T.Errorf("input: %s\nerr: output did not match expected output\nexpected: %s\ngot: %s\n",
				stringForm,
				stringForm,
				stringResult)
		} else {
			fmt.Printf("byte for char pass: %s\n", stringResult)
		}
	}
}

func TestStringToByteSequence(T *testing.T) {
	testStringCaps := TestIO{input: "FF", expectedResult: []int{255}}
	testStringLowerCase := TestIO{input: "ff", expectedResult: []int{255}}
	testStringCombined := TestIO{input: "f3", expectedResult: []int{243}}
	testStringDoubleByte := TestIO{input: "FFAA", expectedResult: []int{255, 170}}

	cases := []TestIO{testStringCaps, testStringLowerCase, testStringCombined, testStringDoubleByte}

	for _, test := range cases {
		result := StringToByteSequence(test.input)
		if(!sliceEq(test.expectedResult, result)) {
			T.Errorf("\ninput: %s\nerr: result did not match expected result\nexpected: %s\ngot: %s\n", 
				test.input, 
				formatSliceToString(test.expectedResult), 
				formatSliceToString(result))
			T.FailNow()
		} else {
			fmt.Printf("str to byte pass: %s\n", test.input)
		}
	}



}
