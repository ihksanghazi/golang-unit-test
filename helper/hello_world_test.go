package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("sandy")
	if result != "Hello sandy" {
		// unit test failed
		t.Error("Must be 'Hello sandy'")
	}
	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldIhksan(t *testing.T) {
	result := HelloWorld("ihksan")
	if result != "Hello ihksan" {
		// unit test failed
		t.Fatal("Must be 'Hello ihksan'")
	}
	fmt.Println("TestHelloWorldIhksan Done")
}
