package insertsort

func Sort(input []int) ([]int) {

    for i, _ := range input {
        this := input[i]
        j := i - 1

        for j >= 0 && input[j] > this {
            input[j+1] = input[j]
            j--
        }

        input[j + 1] = this
    }

    return input
}
