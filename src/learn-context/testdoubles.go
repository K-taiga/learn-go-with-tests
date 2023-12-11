package main

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	// 文字列でバッファ1のchanを作成
	data := make(chan string, 1)

	go func() {
		var result string
		// レスポンスの文字数だけfor
		for _, c := range s.response {
			select {
			// ctxがDoneしていたらreturnで抜ける
			case <-ctx.Done():
				s.t.Log("spy store got cancelled")
				return
			// それ以外は10秒待って結果に1文字入れる
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		// 終わり次第結果をdataに入れる
		data <- result
	}()

	// データが取得できたかどうか
	select {
	case <-ctx.Done():
		// キャンセルされたらerr
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}
