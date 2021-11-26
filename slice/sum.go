package slice

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(numsToSum ...[]int) []int {
	lengthOfNums := len(numsToSum)

	ans := make([]int, lengthOfNums)

	for i, nums := range numsToSum {
		ans[i] = Sum(nums)
	}
	return ans
}

func SumAllTails(numsToSum ...[]int) []int {

	var ans []int

	for _, nums := range numsToSum {
		if len(nums) == 0 {
			ans = append(ans, 0)
		} else {
			tail := nums[1:]
			ans = append(ans, Sum(tail))
		}
	}
	return ans
}
