package tests

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/rand"
	"testing"
)

/**
See here for details about testing in Go:
https://pkg.go.dev/testing@go1.20
https://blog.jetbrains.com/go/2022/11/22/comprehensive-guide-to-testing-in-go/

1. Test file name must end in "_test.go"
2. Test functions must have signature:
     func TestXXX(t *testing.T) {...}
You write a test by creating a file with a name ending in _test.go that contains functions
named TestXXX with signature func (t *testing.T).
The test framework runs each such function; if the function calls a failure function such as
t.Error or t.Fail, the test is considered to have failed.

You can run tests in two ways:
1. In IntelliJ, click on the green arrow next to the left of the test func and run,
   or click on the double green arrow on top next to "package xxx" and run all the tests.
2. In terminal, navigate to a folder that contains tests, and run `go test`
*/

func TestSquare1(t *testing.T) {
	res := Square1(2)
	if res != 4 {
		// Errorf is equivalent to Logf followed by Fail.
		t.Errorf("Incorrect. Got %d, but want %d", res, 4)
	}

}

func TestSquare2(t *testing.T) {
	res := Square2(2)
	if res != 4 {
		t.Errorf("Incorrect. Got %d, but want %d", res, 4)
	}

}

// Example code: Go can run and compare the output with a given value specified in comment
func ExampleHello() {
	fmt.Println("hello")
	// Output: hello
}

// Benchmark testing. This is so cool.
// See https://pkg.go.dev/testing@go1.20
func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Int()
	}
}

// fuzzing, a testing technique where a function is called with randomly generated inputs to find bugs
// not anticipated by unit tests.
func FuzzHex(f *testing.F) {
	for _, seed := range [][]byte{{}, {0}, {9}, {0xa}, {0xf}, {1, 2, 3, 4}} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, in []byte) {
		enc := hex.EncodeToString(in)
		out, err := hex.DecodeString(enc)
		if err != nil {
			t.Fatalf("%v: decode: %v", in, err)
		}
		if !bytes.Equal(in, out) {
			t.Fatalf("%v: not equal after round trip: %v", in, out)
		}
	})
}

// Skip a test depending on testing mode
// Mode needs to be set through flag: see https://blog.jbowen.dev/2019/08/using-go-flags-in-tests/
func TestTimeConsuming(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	res := Square2(2)
	if res != 4 {
		t.Errorf("Incorrect. Got %d, but want %d", res, 4)
	}
}
