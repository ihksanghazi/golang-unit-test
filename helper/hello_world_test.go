package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("sandy")
	if result != "Hello sandy" {
		// unit test failed
		t.Fail()
	}
	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldIhksan(t *testing.T) {
	result := HelloWorld("ihksan")
	if result != "Hello ihksan" {
		// unit test failed
		t.FailNow()
	}
	fmt.Println("TestHelloWorldIhksan Done")
}
