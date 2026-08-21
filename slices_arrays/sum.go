package slicesarrays

func Sum(intSlice []int) int {
	var res int
	for _, num := range intSlice {
		res += num
	}
	return res
}

func SumSlices(intSlices [][]int) []int {
	numberOfSlices := len(intSlices)
	res := make([]int, numberOfSlices)

	for i := range numberOfSlices {
		res[i] = Sum(intSlices[i])
	}
	return res
}

func SumAllTaills(intSlices [][]int) []int {
	numberOfSlices := len(intSlices)
	res := make([]int, numberOfSlices)

	for i := range numberOfSlices {
		if len(intSlices[i]) > 1 {
			res[i] = Sum(intSlices[i][1:])
		} else {
			res[i] = 0
		}
	}
	return res
}
