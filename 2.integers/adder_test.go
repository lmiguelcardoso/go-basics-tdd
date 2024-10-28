package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2,2)
	expected := 4

assertAdderResult(t,sum,expected)
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
func assertAdderResult(t testing.TB,got, expected int){
	if got != expected{
		t.Errorf("expected %v but got %v",expected,got)
	}
}