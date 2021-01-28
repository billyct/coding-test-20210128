package coding_test_20210128

func expandArray(inputs [][]int) []int {
	res, inputs := sortWithLast(inputs)

	for _, input := range inputs {
		res = merge(input, res)
	}

	return res
}

// 合并两个有序切片
func merge(a, b []int) []int {
	var pa, pb int
	var res []int

	for pa < len(a) && pb < len(b) {
		if a[pa] < b[pb] {
			res = append(res, a[pa])
			pa++
		} else {
			res = append(res, b[pb])
			pb++
		}
	}

	if pa < len(a) {
		res = append(res, a[pa:]...)
	} else {
		res = append(res, b[pb:]...)
	}

	return res
}

// 根据头尾排序
func sortWithLast(inputs [][]int) ([]int, [][]int) {
	var res []int

	// 存储头尾排序之后，剩下的数组
	var resInputs [][]int

	// 遍历数组，头尾进行排序
	for _, input := range inputs {
		if last(res) < input[0] {
			res = append(res, input...)
		} else {
			resInputs = append(resInputs, input)
		}
	}

	return res, resInputs
}

func last(arr []int) int {
	if len(arr) <= 0 {
		return 0
	}

	return arr[len(arr)-1]
}
