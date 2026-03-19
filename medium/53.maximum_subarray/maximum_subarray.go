package maximum_subarray

func maxSubArray(nums []int) int {
	var maxSum = nums[0]
	var curSum int

	for _, n := range nums {
		curSum = max(curSum, 0)
		curSum += n
		maxSum = max(maxSum, curSum)
	}

	return maxSum
}
