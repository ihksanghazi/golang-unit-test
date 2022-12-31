package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("sandy")
	if result != "Hello sandy" {
		// unit test failed
		panic("Result is not OK")
	}
}

func TestHelloWorldIhksan(t *testing.T) {
	result := HelloWorld("ihksan")
	if result != "Hello ihksan" {
		// unit test failed
		panic("Result is not OK")
	}
}
