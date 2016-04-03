package bubble_sort

import (
    "testing"
    "reflect"
)

func TestSort(t *testing.T) {
    input1 := []int{5,4,3,2,1}
    expected1 := []int{1,2,3,4,5}
    sorted1 := Sort(input1)

    if (!reflect.DeepEqual(sorted1, expected1)) {
        t.Errorf("\nResult   %v\nExpected %v", sorted1, expected1)
    }

    input2 := []int{17,2,196,20175,1,646,235,18,86,127,1005}
    expected2 := []int{1,2,17,18,86,127,196,235,646,1005,20175}
    sorted2 := Sort(input2)

    if (!reflect.DeepEqual(sorted2, expected2)) {
        t.Errorf("\nResult   %v\nExpected %v", sorted2, expected2)
    }
}
