package sort

func insertion(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i] < arr[j] {
				temp := arr[i]
				arr[i] = *&arr[j]
				arr[j] = temp
			}
		}
	}
}

func Binary(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		j := i - 1
		selected := arr[i]
		loc := binarySearch(arr[:j], selected)

		for loc <= j {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = selected
	}
	return arr
}

func binarySearch(nums []int, target int) int {
	println("here1 ", target)
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right + left) / 2
		if nums[mid] == target {
			println("here ", mid)
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	println("here ", left)
	return left
}
