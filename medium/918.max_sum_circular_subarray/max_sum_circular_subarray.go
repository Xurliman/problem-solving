package max_sum_circular_subarray

func maxSubarraySumCircular(nums []int) int {
	var totalSum int
	for _, n := range nums {
		totalSum += n
	}

	maxS := maxSubArray(nums)
	minS := minSubArray(nums)

	if totalSum == minS {
		return maxS
	}

	return max(maxS, totalSum-minSubArray(nums))
}

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

func minSubArray(nums []int) int {
	var minSum = nums[0]
	var curSum int

	for _, n := range nums {
		curSum = min(curSum, 0)
		curSum += n
		minSum = min(minSum, curSum)
	}

	return minSum
}
