package main


/*=====================

ARRAYS AND SLICES

======================*/

func Sum(nums []int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }

    return sum
}

func SumAll(numsToSum ...[]int) []int {
    /*
    lengthOfNumbers := len(numsToSum)
    sums := make([]int, lengthOfNumbers)

    for i, numbers := range numsToSum {
        sums[i] = Sum(numbers)
    }

    return sums
    */

    var sums []int
    for _, numbers := range numsToSum {
        sums = append(sums, Sum(numbers))
    }

    return sums
}

func SumAllTails(nums ...[]int) []int {
    var sums []int
    for _, numbers := range nums {
        if len(nums) == 0 {
            sums = append(sums, 0)
        } else {
            tail := numbers[1:]
            sums = append(sums, Sum(tail))
        }
    }

    return sums
}
