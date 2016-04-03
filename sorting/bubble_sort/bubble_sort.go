package bubble_sort 

type IntSlice []int

func (a *IntSlice) swap(leftIndex, rightIndex int) {
    slice := *a
    temp := slice[leftIndex]

    slice[leftIndex] = slice[rightIndex]
    slice[rightIndex] = temp

    *a = slice
}

func Sort(arr []int) ([]int) {
    array := IntSlice(arr)
    swapped := true

    for swapped {
        swapped = false
        for i := 0; i < len(array) - 1; i++ {
            leftElement := array[i]
            rightElement := array[i+1]

            if (leftElement > rightElement) {
                array.swap(i, i+1)
                swapped = true
            }
        }
    }

    return array
}
