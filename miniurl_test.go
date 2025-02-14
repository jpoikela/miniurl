package miniurl_test

import (
	"fmt"
	"testing"

	"github.com/jpoikela/miniurl"
	"github.com/stretchr/testify/assert"
)

func TestHashLength(t *testing.T) {
	const (
		input = "https://github.com/jpoikela/miniurl"
		expectedLength = 32
	)
	output := miniurl.Hash(input)
	assert.Len(t, output, expectedLength)
}

func TestHashIsDeterministic(t *testing.T) {
	const input = "https://github.com/jpoikela/miniurl"
    output1 := miniurl.Hash(input)
	output2 := miniurl.Hash(input)
	assert.Equal(t, output1, output2)
}

func ExampleHash() {
	const input = "https://github.com/jpoikela/miniurl"
	output := miniurl.Hash(input)
	fmt.Println(output)
	// output:
	// c7cf04b5508b043b25ed6f46da3945cc
}

func BenchmarkHash(b *testing.B) {
	const input = "https://github.com/jpoikela/miniurl"
	for n:=0; n<b.N; n++ {
		miniurl.Hash(input)
	}
}

func FuzzHash(f *testing.F) {
	f.Add("arbitrary string")
	f.Fuzz(func(t *testing.T, a string) {
		output1 := miniurl.Hash(a)
		output2 := miniurl.Hash(a)
		assert.Equal(t, output1, output2)	
	})
}
