package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("Before Testing")

	m.Run()

	fmt.Println("After Testing")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("can not run on windows")
	}
	result := HelloWorld("Sandy")
	require.Equal(t, "Hello Sandy", result, "Result Must be Hello Sandy")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Sandy")
	require.Equal(t, "Hello Sandy", result, "Result Must be Hello Sandy")
	fmt.Println("TestHelloWorld with Assert Done")
}
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Sandy")
	assert.Equal(t, "Hello Sandy", result, "Result Must be Hello Sandy")
	fmt.Println("TestHelloWorld with Assert Done")
}

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
