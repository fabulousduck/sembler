package byte

import (
	"strconv"
)

/*
PrependBytes prepends length number of bytes to a byte sequence
*/
func PrependBytes(byteSequence []int, length int) []int {
	for i := 0; i < length; i++ {
		byteSequence = append([]int{0x00}, byteSequence...)
	}
	return byteSequence
}

/*
StringToByteSequence takes a raw string and turns it into a byte sequence of the same representation
*/
func StringToByteSequence(byteString string) []int {
	bytes := []int{}

	//left off here, need to swapthe first and second byte for them to be correct
	for i := 0; i < len(byteString); i += 2 {
		var currentByte int
		currentByte = currentByte<<4 | getByteForChar(string(byteString[i]))
		if len(byteString) > 1 {
			currentByte = currentByte<<4 | getByteForChar(string(byteString[i+1]))
		}
		bytes = append(bytes, currentByte)
	}

	return bytes
}

func getByteForChar(char string) int {
	intForm, _ := strconv.Atoi(char)
	if intForm < 10 {
		return intForm
	}

	switch char {
	case "A":
		return 10
	case "B":
		return 11
	case "C":
		return 12
	case "D":
		return 13
	case "E":
		return 14
	case "F":
		return 15
	}

	return 0
}
