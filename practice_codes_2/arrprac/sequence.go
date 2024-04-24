package arrprac

// BracketSequence ...
func BracketSequence(arr []byte) (count int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == '(' {
			if i+1 < len(arr) {
				if arr[i+1] == ')' {
					count++
				}
			}

		}

	}
	return count
}
