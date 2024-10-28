package arraysandslices

import "testing"

func TestSum(t *testing.T) {

	t.Run("slice of any size", func(t *testing.T) {
		numbers:= []int{3,4,5,6}
		
		got:=Sum(numbers)
		want:= 18

		if got != want{
			t.Errorf("got %v want %v",got, want)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1,2},[]int{0,9,10})
	want:= []int{3,19}

	if got != want{
		t.Errorf("got %v want %v",got, want)
	}
}