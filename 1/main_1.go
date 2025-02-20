package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
)

func main() {
	data := NewData(22, 0777, 0xFF, 3.14, "Hello", true, complex(float32(3), float32(4)))
	data.DisplayIntegers()
	str := data.ConvertDataToOneString()
	fmt.Println(str)

	runes := ConvertStringToRunes(str)
	fmt.Println("runes:", runes)

	saltAdded := AddSalt(runes)
	hash := HashRunes(saltAdded)
	fmt.Println("hash:", hash)
}

type Data struct {
	NumDecimal, NumOctal, NumHexadecimal int
	Pi                                   float64
	Name                                 string
	IsActive                             bool
	ComplexNum                           complex64
}

func NewData(
	numDecimal, numOctal, numHexadecimal int,
	pi float64,
	name string,
	isActive bool,
	complexNum complex64) *Data {
	return &Data{
		NumDecimal:     numDecimal,
		NumOctal:       numOctal,
		NumHexadecimal: numHexadecimal,
		Pi:             pi,
		Name:           name,
		IsActive:       isActive,
		ComplexNum:     complexNum}
}

func (d Data) DisplayIntegers() {
	fmt.Printf("%d type of %v\n", d.NumDecimal, reflect.TypeOf(d.NumDecimal))
	fmt.Printf("%d type of %v\n", d.NumOctal, reflect.TypeOf(d.NumOctal))
	fmt.Printf("%d type of %v\n", d.NumHexadecimal, reflect.TypeOf(d.NumHexadecimal))
	fmt.Printf("%f type of %v\n", d.Pi, reflect.TypeOf(d.Pi))
	fmt.Printf("%s type of %v\n", d.Name, reflect.TypeOf(d.Name))
	fmt.Printf("%t type of %v\n", d.IsActive, reflect.TypeOf(d.IsActive))
	fmt.Printf("%v type of %v\n", d.ComplexNum, reflect.TypeOf(d.ComplexNum))
}

func (d Data) ConvertDataToOneString() string {
	str1 := fmt.Sprintf("%d", d.NumDecimal)
	str2 := fmt.Sprintf("%o", d.NumOctal)
	str3 := fmt.Sprintf("%x", d.NumHexadecimal)
	str4 := fmt.Sprintf("%f", d.Pi)
	str5 := fmt.Sprintf("%s", d.Name)
	str6 := fmt.Sprintf("%t", d.IsActive)
	str7 := fmt.Sprintf("%v", d.ComplexNum)

	return str1 + str2 + str3 + str4 + str5 + str6 + str7
}

func ConvertStringToRunes(str string) []rune {
	return []rune(str)
}

func AddSalt(runes []rune) []rune {
	salt := []rune("go-2024")

	var newRunes []rune
	newRunes = append(newRunes, runes[:len(runes)/2]...)
	newRunes = append(newRunes, salt...)
	newRunes = append(newRunes, runes[len(runes)/2:]...)

	return newRunes
}

func HashRunes(runes []rune) string {
	data := []byte(string(runes))

	hash := sha256.Sum256(data)

	return hex.EncodeToString(hash[:])
}
