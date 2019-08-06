package byte

import (
	"strconv"
)

/*
StringToByteSequence takes a raw string and turns it into a byte sequence of the same representation
*/
func StringToByteSequence(byteString string) []int {
	bytes := []int{}

	//left off here, need to swapthe first and second byte for them to be correct
	for i := 0; i < len(byteString); i += 2 {
		var currentByte int
		currentByte = currentByte<<4 | getByteForChar(string(byteString[i]))
		currentByte = currentByte<<4 | getByteForChar(string(byteString[i+1]))
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
