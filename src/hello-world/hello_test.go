package main

import "testing"

func TestHello(t *testing.T) {
	// 関数内で関数を宣言し変数に割り当てられる *testing.T と *testing.B の両方が満たすインターフェース
	assertCorrectMessage := func(t testing.TB, got, want string) {
		// このメソッドがヘルパーであるとテストスイートに伝えるもの
		t.Helper()
		if got != want {
			// fは、プレースホルダー値％qに値が挿入された文字列を作成できる形式を表す
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Christiane", "French")
		want := "Bondule, Christiane"
		assertCorrectMessage(t, got, want)
	})
}
