package arrayupdate

// CalculateTime ...
func CalculateTime(arr []int, updateVal int) (totaltime int) {
	//totaltime = 0
	// arrlen := len(arr)
	// var incrementTime bool
	minValInArr := arr[0]
	for k := 0; k < len(arr); k++ {
		if arr[k] < minValInArr {
			minValInArr = arr[k]
		}
	}

	if minValInArr < updateVal {
		return updateVal - minValInArr
	}
	if minValInArr >= updateVal {
		return 0
	}

	// for arrlen > 0 {
	// 	for i := 0; i < len(arr); i++ {
	// 		if arr[i] < updateVal {
	// 			arr[i]++
	// 			incrementTime = true
	// 		}
	// 	}
	// 	if incrementTime {
	// 		totaltime++
	// 	}
	// 	arrlen--
	// }

	// return totaltime
	return 0
}
