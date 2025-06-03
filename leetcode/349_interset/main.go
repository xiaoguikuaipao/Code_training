package main

func main() {

}

func intersection(nums1 []int, nums2 []int) []int {
	record := make(map[int]struct{})
	result := make([]int, 0)
	for _, e := range nums1 {
		record[e] = struct{}{}
	}
	for _, e := range nums2 {
		if _, has := record[e]; has {
			result = append(result, e)
			delete(record, e)
		}
	}
	return result
}
