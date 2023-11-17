package main

import (
	"fmt"
	"net/http"
	"time"
)

// func Racer(a, b string) (winner string) {
// 	startA := time.Now()
// 	http.Get(a)
// 	aDuration := time.Since(startA)

// 	startB := time.Now()
// 	http.Get(b)
// 	bDuration := time.Since(startB)

// 	if aDuration < bDuration {
// 		return a
// 	}

// 	return b
// }

// // 小文字で始まるメソッドはprivate(他のパッケージからは呼べない)
// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	// selectで2つの非同期プロセスの結果を待つ 共有のチャネルに通知を先に送ったのはどっちかを判断
	// 先に送った方のcaseを実行
	select {
	// pingのチャネルからの完了通知を受信
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	// 10秒経ったことをチャネルに通知
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	// chan structはnilの構造体でメモリを消費せず単にチャネルの完了を通知するのに適している
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		// chが閉じたらゼロ値を返す(chan struct{}の構造体)
		close(ch)
	}()

	// このreturnのタイミングではchは開いていてgo funcの非同期処理でcloseされ次第閉じる
	return ch
}
