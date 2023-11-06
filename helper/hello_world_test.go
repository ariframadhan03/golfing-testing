package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("Before Unit Test")
	m.Run()
	fmt.Println("After Unit Test")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("rama")

	if result != "Hello rama" {
		t.Fatal("Result must be 'Hello rama'")
	}

	fmt.Println("Test Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("rama")

	require.Equal(t, "Hello rama", result, "Result must be 'Hello rama'")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("rama")

	assert.Equal(t, "Hello rama", result, "Result must be 'Hello rama'")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Cant do test on Mac OS")
	}

	result := HelloWorld("rama")
	assert.Equal(t, "Hello rama", result, "Result must be 'Hello rama'")
}

func TestSubTest(t *testing.T) {
	t.Run("SubTest 1", func(t *testing.T) {
		result := HelloWorld("arif")
		assert.Equal(t, "Hello arif", result, "Result must be 'Hello arif'")
	})

	t.Run("SubTest 2", func(t *testing.T) {
		result := HelloWorld("rama")
		assert.Equal(t, "Hello rama", result, "Result must be 'Hello rama'")
	})
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "arif",
			request:  "arif",
			expected: "Hello arif",
		},
		{
			name:     "rama",
			request:  "rama",
			expected: "Hello rama",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

// benchmark
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("rama")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("arif", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("arif")
		}
	})

	b.Run("rama", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("rama")
		}
	})
}

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name string
		req  string
	}{
		{
			name: "rama",
			req:  "rama",
		},
		{
			name: "arif",
			req:  "arif",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.req)
			}
		})
	}
}
