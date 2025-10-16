package main

func main() {
	input := []int{1, 0}
	output := productExceptSelf(input)
	for _, value := range output {
		print(value, ", ")
	}
}

func productExceptSelf(nums []int) []int {
	//O(n)
	totalProduct := 1
	zeroCount := 0
	for _, num := range nums {
		if num != 0 {
			totalProduct *= num
		} else {
			zeroCount++
		}
	}

	output := make([]int, len(nums))
	if zeroCount > 1 {
		return output
	}

	for idx, num := range nums {
		if zeroCount > 0 {
			if num == 0 {
				output[idx] = totalProduct
			} else {
				output[idx] = 0
			}
		} else {
			output[idx] = totalProduct / num
		}
	}

	return output
}
