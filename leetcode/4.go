package leetcode

// https://leetcode.com/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mid := len(nums1) / len(nums2)
	isEven := ((len(nums1) + len(nums2)) % 2) == 0

	i, j, curr := 0, 0, 0
	for curr < mid {
		if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
		curr++
	}

	if nums1[i] < nums2[j] {
		if isEven {
			return (float64(nums1[i]) + float64(nums2[j])) / 2
		} else {
			return float64(nums1[i])
		}
	} else {
		return float64(nums2[j])
	}
}
