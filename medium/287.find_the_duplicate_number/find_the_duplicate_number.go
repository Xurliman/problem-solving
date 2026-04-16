package main

// Input: nums = [3, 1, 3, 4, 2] & 2
func findDuplicate(nums []int) int {
	// Phase 1: Find intersection point inside the cycle
	// Both start at index 0 — slow moves 1 step, fast moves 2 steps
	slow := nums[0]       // slow takes one step from 0
	fast := nums[nums[0]] // fast takes two steps from 0

	for slow != fast { // loop until they meet (same position)
		slow = nums[slow]       // move 1 step: follow one link
		fast = nums[nums[fast]] // move 2 steps: follow two links
	}

	// Phase 2: Find the entrance to the cycle (= the duplicate value)
	// Reset slow to the start, keep fast at the meeting point
	// Move both 1 step at a time — they will meet at the cycle entrance
	slow = 0
	for slow != fast {
		slow = nums[slow] // 1 step
		fast = nums[fast] // 1 step (not 2 anymore!)
	}

	return slow // slow == fast == the duplicate number
}
