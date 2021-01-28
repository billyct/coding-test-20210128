package coding_test_20210128

func expandArrayWithBucket(inputs [][]int) []int {
	var res []int
	// 构建一个大小为 101 小桶子，可以存储的最大数字为 100
	bucket := make([]int, 101)

	for _, input := range inputs {
		for _, num := range input {
			// 每一次命中，命中的数字都 + 1
			bucket[num]++
		}
	}

	// 遍历这个小桶子中的数据
	for num, v := range bucket {
		if v == 0 {
			continue
		}

		// 如果数字不为零，则命中过多少次，就把数字加入到结果中多少次
		for i := 0; i < v; i++ {
			res = append(res, num)
		}
	}

	return res
}
