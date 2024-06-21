package nor 

// median to call the medal number
func Median(data []int) float64 {
	datalen := len(data)
	var medianindex int
	var median float64
	if datalen % 2 == 0 { // in the len is even there is 2 median
		medianindex = datalen/2
		median = float64(data[medianindex]+data[medianindex-1]) / 2.0
	} else { // if len odd there is only one median
		medianindex = datalen/2
		median = float64(data[medianindex])
	}
	return median
}