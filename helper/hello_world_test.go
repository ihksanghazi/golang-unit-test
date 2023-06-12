package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkSub(b *testing.B) {
	b.Run("Sandy",func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Sandy")
		}	
	})
	b.Run("Ihksan",func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Ihksan")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Sandy")
	}
}

func TestTableTest(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Nur",
			request:  "Nur",
			expected: "Hello Nur",
		}, {
			name:     "Sandy",
			request:  "Sandy",
			expected: "Hello Sandy",
		}, {
			name:     "Ihksan",
			request:  "Ihksan",
			expected: "Hello Ihksan",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result, "Result Must be Hello "+test.request)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Sandy", func(t *testing.T) {
		result := HelloWorld("Sandy")
		require.Equal(t, "Hello Sandy", result, "Result Must be Hello Sandy")
	})

	t.Run("Ihksan", func(t *testing.T) {
		result := HelloWorld("Ihksan")
		require.Equal(t, "Hello Ihksan", result, "Result Must be Hello Ihksan")
	})
}

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
