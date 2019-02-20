package code

// TwoSum : 两数之和
func TwoSum(nums []int, target int) []int {

	if len(nums) == 0 {
		return nil
	}

	data := make(map[int]int)
	for index, value := range nums {

		if v, ok := data[target-value]; ok {
			return []int{v, index}
		}

		data[value] = index
	}

	return nil
}
