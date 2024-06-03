package pascal2

// Implementations directly came from this article.
// The goal of this exercise was to learn a bit more about GO.
// https://medium.com/@nicola.moro2312/pascals-triangles-daily-coding-problem-19-566af9a817d3

func PascalTriangle(row int) []int {
	// temp := []int{1}
	if row == 1 {
		return []int{1}
	}
	if row == 2 {
		return []int{1, 1}
	}

	out := []int{1, 1}
	for n := 3; n <= row; n++ {
		temp := []int{1}
		for i, j := range out[:len(out)-1] {
			temp = append(temp, j+out[i+1])
		}
		temp = append(temp, 1)
		out = temp
	}
	return out
}
