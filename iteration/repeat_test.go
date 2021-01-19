package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	makeAssertion := func(t testing.TB, repeated, expected string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("should repeat the character five times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		makeAssertion(t, repeated, expected)
	})

	t.Run("should be a empty string", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""

		makeAssertion(t, repeated, expected)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", i)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 3)

	fmt.Println(repeated)
	//Output: aaa
}
