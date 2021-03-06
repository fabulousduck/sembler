package byte

import (
	"strconv"
	"strings"
)

/*
PrependBytes prepends length number of bytes to a byte sequence
*/
func PrependBytes(byteSequence []int, length int) []int {
	for i := 0; i < length-len(byteSequence); i++ {
		byteSequence = append([]int{0x00}, byteSequence...)
	}
	return byteSequence
}

/*
AppendBytes prepends length number of bytes to a byte sequence
*/
func AppendBytes(byteSequence []int, length int) []int {
	for i := 0; i < length-len(byteSequence); i++ {
		byteSequence = append(byteSequence, 0x00)
	}
	return byteSequence
}

/*
StringToByteSequence takes a raw string and turns it into a byte sequence of the same representation
*/
func StringToByteSequence(byteString string) []int {
	bytes := []int{}

	//left off here, need to swap the first and second byte for them to be correct
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
	intForm, err := strconv.Atoi(char)
	if err != nil {
		char = strings.ToLower(char)
		switch char {
		case "a":
			return 0xA
		case "b":
			return 0xB
		case "c":
			return 0xC
		case "d":
			return 0xD
		case "e":
			return 0xE
		case "f":
			return 0xF
		}
	}

	return intForm
	
}
