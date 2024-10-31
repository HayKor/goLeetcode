package main

func matrixLen(mat [][]int) int {
	var length int
	for _, row := range mat {
		length += len(row)
	}
	return length
}

func flatten(mat [][]int) []int {
	res := make([]int, 0, matrixLen(mat))
	for _, row := range mat {
		res = append(res, row...)
	}
	return res
}

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if matrixLen(mat) != r*c {
		return mat
	}
	res := make([][]int, 0, r)
	flattened := flatten(mat)
	for i := range r {
		tmp := flattened[i*c : i*c+c]
		res = append(res, tmp)
	}
	return res
}

// 1 2 3                        1 2
// 4 5 6  --> 1 2 3 4 5 6  -->  3 4
//	                            5 6

// func main() {
// 	mat2 := [][]int{{1, 2, 3}, {4, 5, 6}}
// 	newMat := matrixReshape(mat2, 3, 2)
// 	fmt.Println(newMat)
// 	fmt.Printf("cap(newMat): %v, len(newMat): %v\n", cap(newMat), len(newMat))
// }
