package main

import (
	"fmt"
	"testing"
)

func TestCreateUserData2(t *testing.T) {
	result := createUserData("paul", "07656500000655", 333)
	if result.userName != "aul" {
		t.Errorf("expected name 'aul' but got %s", result.userName)
	}
	if result.phoneNumber == "07656500000655" {
		fmt.Printf("Correct number '%s' registered", result.phoneNumber)

	} else {
		fmt.Print("Eiii we need program needs work")
	}

}
