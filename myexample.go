package golangexamples

import (
	"fmt"
	"strings"
)

func ConcatSlice(sliceToConcat []byte) string {
	var s []string
	for _, i := range sliceToConcat {
		s = append(s, string(i))
	}
	return fmt.Sprintf(strings.Join(s, "-"))
}

func Encrypt(sliceToEncrypt []byte, ceaserCount int) []byte {
	for i, _ := range sliceToEncrypt {
		sliceToEncrypt[i] = byte(int(sliceToEncrypt[i]) + ceaserCount)
	}
	return sliceToEncrypt
}

func PrintGreetings(name string) string {
	s := "Hello World-"
	s += name
	return s
}