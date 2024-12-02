package kata

import (
	"fmt"
	"strconv"
	"strings"
)

func EncryptThis(text string) string {
	if text == "" {
		return ""
	}
	result := ""
	words := strings.Split(text, " ")
	fmt.Println(words)

	for _, word := range words {
		switch len(word) {
		case 1:
			asciiCode := int(word[0])
			asciiString := strconv.Itoa(asciiCode)
			result += asciiString + " "
		case 2:
			asciiCode := int(word[0])
			asciiString := strconv.Itoa(asciiCode)
			result += asciiString + string(word[1]) + " "
		default:
			asciiCode := int(word[0])
			asciiString := strconv.Itoa(asciiCode)
			result += asciiString + string(word[len(word)-1]) + string(word[2:len(word)-1]) + string(word[1]) + " "

		}
	}
	return strings.Trim(result, " ")
}

/*
You want to create secret messages which can be deciphered by the Decipher this! kata. Here are the conditions:

Your message is a string containing space separated words.
You need to encrypt each word in the message using the following rules:
The first letter must be converted to its ASCII code.
The second letter must be switched with the last letter
Keepin' it simple: There are no special characters in the input.
Examples:

EncryptThis("Hello") == "72olle"
EncryptThis("good") == "103doo"
EncryptThis("hello world") == "104olle 119drlo"
*/
