package main

import (
	"crypto/sha256"
	"fmt"
)

var numDecimal int = 42
var numOctal int = 042
var numHexadecimal int = 0x5F
var pi float64 = 3.14
var title string = "my program"
var isPrime bool = true
var complexNum complex64 = 1 + 2i

func main() {
	fmt.Printf("numDecimal = %d Type: %s\n", numDecimal, typeOf(numDecimal))
	fmt.Printf("numOctal = %o Type: %s\n", numOctal, typeOf(numOctal))
	fmt.Printf("numHexadecimal = %#X Type: %s\n", numHexadecimal, typeOf(numHexadecimal))
	fmt.Printf("pi = %g Type: %s\n", pi, typeOf(pi))
	fmt.Printf("title = %s Type: %s\n", title, typeOf(title))
	fmt.Printf("isPrime = %t Type: %s\n", isPrime, typeOf(isPrime))
	fmt.Printf("complexNum = %g Type: %s\n", complexNum, typeOf(complexNum))

	stringNumDecimal := fmt.Sprintf("%d", numDecimal)
	stringNumOctal := fmt.Sprintf("%o", numOctal)
	stringNumHexadecimal := fmt.Sprintf("%#X", numHexadecimal)
	stringPi := fmt.Sprintf("%g", pi)
	stringTitle := fmt.Sprintf("%s", title)
	stringIsPrime := fmt.Sprintf("%t", isPrime)
	stringComplex := fmt.Sprintf("%g", complexNum)

	s := stringNumDecimal + stringNumOctal + stringNumHexadecimal + stringPi + stringTitle + stringIsPrime + stringComplex
	fmt.Println("Соединенная строка: ", s)

	salt := "go-2024"
	fmt.Println("Получившийся хэш с солью: ", hashSha256WithSalt(s, salt))
}

func typeOf(t any) string {
	return fmt.Sprintf("%T", t)
}

func hashSha256WithSalt(s string, salt string) string {
	runes := []rune(s)
	runes = append(append(runes[:len(runes)/2], []rune(salt)...), runes[len(runes)/2:]...)
	return fmt.Sprintf("%X", sha256.Sum256([]byte(string(runes))))
}
