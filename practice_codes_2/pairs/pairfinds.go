package pairs

// FindPossiblePairs ...
func FindPossiblePairs(arr []int) (totalPairs int) {
	arrLen := len(arr)

	intMap := make(map[int]int)
	_ = intMap //will do this operation later
	for i := 0; i < arrLen; i++ {
		for j := i; j < arrLen; j++ {
			if arr[i] == arr[j+1] || arr[j] == arr[i+1] {
				totalPairs++
			}
		}
	}
	return totalPairs
}
