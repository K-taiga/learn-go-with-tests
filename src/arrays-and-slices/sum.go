package arrays_and_slices

func Sum(numbers []int) int {
	sum := 0
	// - でインデックスを無視 numbersのrangeの数の分だけ繰り返す
	for _, numbers := range numbers {
		sum += numbers
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// // numbersToSumの長さを取り出し
	// lengthOfNumbers := len(numbersToSum)
	// // lenの長さの開始容量のsliceを作成
	// sums := make([]int, lengthOfNumbers)

	// // sumにかける
	// for i, numbers := range numbersToSum {
	// 	sums[i] = Sum(numbers)
	// }

	var sums []int
	for _, numbers := range numbersToSum {
		// appendでsliceの中身を追加して返せる
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTrails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		// 空のsliceが渡されたら0
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			// slice[low:high]でlowからhighのどこまでsliceするか 省略すると全て
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
