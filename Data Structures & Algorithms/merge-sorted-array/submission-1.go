func merge(nums1 []int, m int, nums2 []int, n int) {
    i, j := 0, 0
    res := make([]int, 0 , m + n)

    for i < m && j < n {
        if nums1[i] < nums2[j] {
            res = append(res, nums1[i])
            i++
        } else {
            res = append(res, nums2[j])
            j++
        }
    }

    if i < m {
        res = append(res, nums1[i:]...)
    }

    if j < n {
        res = append(res, nums2[j:]...)
    }

    copy(nums1, res)
}
