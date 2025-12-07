package main

// merge merges nums2 into nums1 in-place.
// nums1 has a size of m+n, with the first m elements valid
// and nums2 having n elements. Both arrays are sorted.
//
// The merge is performed from the back to avoid overwriting
// initial elements in nums1.
func merge(nums1 []int, m int, nums2 []int, n int) {
    k := len(nums1) - 1
    i := m - 1
    j := n - 1

    for k >= 0 {
        if i >= 0 && j >= 0 && nums1[i] >= nums2[j] {
            nums1[k] = nums1[i]
            i--

        } else if i >= 0 && j >= 0 && nums1[i] < nums2[j] {
            nums1[k] = nums2[j]
            j--

        } else if j >= 0 {
            nums1[k] = nums2[j]
            j--
        }

        k--
    }
}
