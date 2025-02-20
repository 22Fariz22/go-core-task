package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"testing"
)

// Тест для NewIntegers
func Test_NewData(t *testing.T) {
	data := NewData(22, 0777, 0xFF, 3.14, "Hello", true, complex(float32(3), float32(4)))

	if data.NumDecimal != 22 {
		t.Errorf("Expected NumDecimal to be 22, got %d", data.NumDecimal)
	}

	if data.NumOctal != 0777 {
		t.Errorf("Expected NumOctal to be 0777, got %d", data.NumOctal)
	}

	if data.NumHexadecimal != 0xFF {
		t.Errorf("Expected NumHexadecimal to be 0xFF, got %d", data.NumHexadecimal)
	}

	if data.Pi != 3.14 {
		t.Errorf("Expected Pi to be 3.13, got %f", data.Pi)
	}

	if data.Name != "Hello" {
		t.Errorf("Expected Name to be 'Hello', got %s", data.Name)
	}

	if data.IsActive != true {
		t.Errorf("Expected IsActive to be true, got %t", data.IsActive)
	}

	if data.ComplexNum != complex(float32(3), float32(4)) {
		t.Errorf("Expected ComplexNum to be (3+4i), got %v", data.ComplexNum)
	}
}

func Test_Display(t *testing.T) {
	data := NewData(22, 0777, 0xFF, 3.14, "Hello", true, complex(float32(3), float32(4)))
	data.DisplayIntegers()

	if reflect.TypeOf(data.NumDecimal) != reflect.TypeOf(1) {
		t.Errorf("NumDecimal should be int")
	}
	if reflect.TypeOf(data.NumOctal) != reflect.TypeOf(1) {
		t.Errorf("NumOctal should be int")
	}
	if reflect.TypeOf(data.NumHexadecimal) != reflect.TypeOf(1) {
		t.Errorf("NumHexadecimal should be int")
	}
	if reflect.TypeOf(data.Pi) != reflect.TypeOf(1.0) {
		t.Errorf("Pi should be float64")
	}
	if reflect.TypeOf(data.Name) != reflect.TypeOf("") {
		t.Errorf("Name should be string")
	}
	if reflect.TypeOf(data.IsActive) != reflect.TypeOf(false) {
		t.Errorf("IsActive should be boolean")
	}
	if reflect.TypeOf(data.ComplexNum) != reflect.TypeOf(complex64(0)) {
		t.Errorf("ComplexNum should be complex64")
	}
}

func Test_ConvertDataToOneString(t *testing.T) {
	data := NewData(22, 0777, 0xFF, 3.14, "Hello", true, complex(float32(3), float32(4)))

	res := data.ConvertDataToOneString()

	str1 := fmt.Sprintf("%d", data.NumDecimal)
	if str1 != "22" {
		t.Errorf("Expected str1 to be '22' , got %s", str1)
	}
	str2 := fmt.Sprintf("%o", data.NumOctal)
	if str2 != "777" {
		t.Errorf("Expected str2 to be '777', got %s", str2)
	}

	str3 := fmt.Sprintf("%x", data.NumHexadecimal)
	if str3 != "ff" {
		t.Errorf("Expected str3 to be 'ff', got %s", str3)
	}

	str4 := fmt.Sprintf("%f", data.Pi)
	if str4 != "3.140000" {
		t.Errorf("Expected str4 to be '3.140000', got %s", str4)
	}

	str5 := fmt.Sprintf("%s", data.Name)
	if str5 != "Hello" {
		t.Errorf("Expected str5 to be 'Hello', got %s", str5)
	}

	str6 := fmt.Sprintf("%t", data.IsActive)
	if str6 != "true" {
		t.Errorf("Expected str6 to be 'true', got %s", str6)
	}

	str7 := fmt.Sprintf("%v", data.ComplexNum)
	if str7 != "(3+4i)" {
		t.Errorf("Expected str7 to be '(3+4i), got %s'", str7)
	}

	if res != str1+str2+str3+str4+str5+str6+str7 {
		t.Errorf("Expected result string '22777ff3.140000Hellotrue(3+4i)', got %s", res)
	}
}

func Test_ConvertStringToRunes(t *testing.T) {
	data := NewData(22, 0777, 0xFF, 3.14, "Hello", true, complex(float32(3), float32(4)))
	str := data.ConvertDataToOneString()

	runes := ConvertStringToRunes(str)

	type1 := reflect.TypeOf(runes)
	type2 := reflect.TypeOf([]rune(""))

	if type1 != type2 {
		t.Errorf("Expected type runes, got %v", type1)
	}
}

func Test_AddSalt(t *testing.T) {
	runes := []rune("22777ff3.140000Hellotrue(3+4i)")
	saltAdded := AddSalt(runes)

	var newRunes []rune

	left := runes[:len(runes)/2]
	expectedSaltInTheMiddle := []rune("go-2024")
	right := runes[len(runes)/2:]

	newRunes = append(newRunes, left...)
	newRunes = append(newRunes, expectedSaltInTheMiddle...)
	newRunes = append(newRunes, right...)

	if !reflect.DeepEqual(saltAdded, newRunes) {
		t.Errorf("The slice is not what is needed")
	}
}

func Test_HashRunes(t *testing.T) {
	runes := []rune("test")

	result := HashRunes(runes)

	hash := sha256.Sum256([]byte(string(runes)))
	expected := hex.EncodeToString(hash[:])

	if expected != result {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	lenght := len(result)
	if lenght != 64 {
		t.Errorf("Expected hash length 64, got %d", lenght)
	}
}
