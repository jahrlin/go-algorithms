package selectsort

func Sort(input *[]int) ([]int) {
    sorted := []int{}
    unsorted := *input

    for len(unsorted) > 0 {
        smallestIndex := 0
        
        for i := 0; i < len(unsorted); i++ {
            if unsorted[smallestIndex] > unsorted[i] {
                smallestIndex = i
            }
        }

        sorted = append(sorted, unsorted[smallestIndex])
        unsorted = append(unsorted[:smallestIndex], unsorted[smallestIndex+1:]...)
    }
	
	*input = sorted
	return sorted;
}
