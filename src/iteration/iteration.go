package iteration

const repeatCount = 5

func Repeat(character string) string {
	// 変数の定義のみ
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}

	return repeated
}
