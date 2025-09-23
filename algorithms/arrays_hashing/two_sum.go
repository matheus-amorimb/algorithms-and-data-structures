package main

func main() {
	nums := []int{5, 5}
	target := 10
	println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	// nums[i] + nums[j] == target, i != j
	// nums = [4, 5, 6], target = 10
	complementsMap := map[int]int{}
	for idx, num := range nums {
		complement := target - num
		idxComplement, complementExists := complementsMap[complement]
		if complementExists {
			return []int{idxComplement, idx}
		}
		complementsMap[num] = idx
	}
	return nil
}
