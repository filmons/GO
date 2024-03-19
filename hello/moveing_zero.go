package main

// package main
 
// import "testing"
 
// func TestMovingZeroes(t *testing.T) {
// 	scenarii := []struct {
// 		input    []int
// 		expected []int
// 	}{
// 		{
// 			input: []int{
// 				0,
// 			},
// 			expected: []int{
// 				0,
// 			},
// 		},
// 		{
// 			input: []int{
// 				0, 0, 0, 1,
// 			},
// 			expected: []int{
// 				1, 0, 0, 0,
// 			},
// 		},
// 		{
// 			input: []int{
// 				0, 1, 0, 3, 12,
// 			},
// 			expected: []int{
// 				1, 3, 12, 0, 0,
// 			},
// 		},
// 		{
// 			input: []int{
// 				0, 1, 0, 33, 12,
// 			},
// 			expected: []int{
// 				1, 33, 12, 0, 0,
// 			},
// 		},
// 	}
 
// 	for scindex, scenario := range scenarii {
// 		res := MovingZeroes(scenario.input)
// 		if len(res) != len(scenario.expected) {
// 			t.Error("Moving Zeroes wrong lengths")
// 		}
 
// 		for i := range res {
// 			if res[i] != scenario.expected[i] {
// 				t.Errorf("MovingZeroes scen %d range %d expected %d got %d",
// 					scindex, i, scenario.expected[i], res[i],
// 				)
// 			}
// 		}
// 	}
// }
 