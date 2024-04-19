package array

// 数组有一个有趣的属性，它的大小也属于类型的一部分, 如果使用数组，需要写上大小
func Sum(numbers []int) int {
	sum := 0
	// 如果是数组， 所以不能使用for循环， 使用range 否则参数里面要写上数组的的长度
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	//for _, number := range numbers {
	//	sum += number
	//}
	return sum
}

// 可变参数
func SumAll(numbers ...[]int) (sums []int) {
	length := len(numbers)
	// make 可以在创建切片的时候指定我们需要的长度和容量。
	sums = make([]int, length)

	for i, n := range numbers {
		sums[i] = Sum(n)
	}

	return
}
