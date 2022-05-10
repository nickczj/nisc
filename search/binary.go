package search

func BinaryIterative(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right + left) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func BinaryRecursive(nums []int, target int) int {
	return binaryRecursive(nums, target, 0, len(nums)-1)
}

func binaryRecursive(nums []int, target, left, right int) int {
	if left <= right {
		mid := (right + left) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			return binaryRecursive(nums, target, mid+1, right)
		} else {
			return binaryRecursive(nums, target, left, mid-1)
		}
	} else {
		return -1
	}
}
