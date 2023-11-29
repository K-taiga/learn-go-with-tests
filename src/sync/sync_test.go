package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		// sync.WaitGroupのwgを定義 WaitGroupは複数のgoroutine待機用
		var wg sync.WaitGroup
		// 待機するカウントを追加
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			// goroutineでwg = wを使用(*sync.WaitGroupでwgのアドレスを取得することを指定)
			go func(w *sync.WaitGroup) {
				counter.Inc()
				// DoneでwantedCountを１つ減らす
				w.Done()
				// 参照渡ししているのでwgのメモリアドレスを無名関数に渡す
				// これで同じオブジェクトを関数内で利用できる
				// 値渡しだとgoroutine毎にオブジェクトのコピーを渡すため終了の感知ができなくなる
			}(&wg)
		}
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

// Mutexは使用後にそのコピーを渡すのはgo vetでよくないとでるためMutexを含んだCounterのポインター(参照しているメモリの中の値)で渡す
func assertCounter(t *testing.T, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
