package main
 
import "testing"
 
func TestContainsDuplicate(t *testing.T) {
	type scenario struct {
		input   []int
		outcome bool
	}
	scenarii := []scenario{
		{
			input: []int{
				1, 2, 3,
			},
			outcome: false,
		},
		{
			input: []int{
				1, 2, 3, 1,
			},
			outcome: true,
		},
		{
			input:   []int{},
			outcome: false,
		},
		{
			input: []int{
				1, 1, 1, 3, 3, 4, 3, 2, 4, 2,
			},
			outcome: true,
		},
		{
			input: []int{
				1,
			},
			outcome: false,
		},
	}
 
	for _, scenario := range scenarii {
		res := ContainsDuplicate(scenario.input)
		if res != scenario.outcome {
			t.Errorf("%v expected %t got %t",
				scenario.input,
				scenario.outcome,
				res,
			)
		}
	}
}
 
func ContainsDuplicate(input []int) bool {
	return false
}