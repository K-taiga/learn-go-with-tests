package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	// resultの構造体を書き込めるチャネルを作成
	resultChannel := make(chan result)

	// rangeの返り値はindexと要素の値 _ で index不要なら無視
	for _, url := range urls {
		go func(u string) {
			// resultの結果を<-でチャネルに書き込み
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	// string => boolのmapを作成
	results := make(map[string]bool)

	for i := 0; i < len(urls); i++ {
		// <-でチャネルから値を取り出して結果のmapを作成
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}
