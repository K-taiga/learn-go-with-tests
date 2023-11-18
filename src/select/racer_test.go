package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		// deferで関数の最後に実行するようにする サーバーを作成したのと近くに置くことで可読性を上げる
		// 実行順はスタックで逆順のため後ろのfastServer.Closeから実行される
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server does'nt respond within 10s", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)

		defer server.Close()

		// タイムアウト時間を引数で渡す
		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but did'nt get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	// テスト用のHTTPサーバーを立ててエミュレートする
	// http.HandlerFuncはhttp.ResponseWriterとhttp.Requestを受け取る
	// HandlerFuncのパッケージにはServeHTTP(ResponseWriter, *Request)メソッドが実装されている
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		// clientに200を返す
		w.WriteHeader(http.StatusOK)
	}))
}
