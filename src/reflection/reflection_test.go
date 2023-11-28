package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	// expected := "Chris"
	// var got []string

	// // 無名の構造体を定義しNameをChrisで生成
	// x := struct {
	// 	Name string
	// }{expected}

	// // 第二引数で無名関数を渡しそれを実行している
	// walk(x, func(input string) {
	// 	got = append(got, input)
	// })

	// if len(got) != 1 {
	// 	t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
	// }

	// if got[0] != expected {
	// 	t.Errorf("got %q, want %q", got[0], expected)
	// }

	// 構造体を含んだスライス 各構造体は１つのテストケースを表す
	cases := []struct {
		// Name
		Name string
		// interfaceのinput
		Input interface{}
		// 実行結果
		ExpectedCall []string
	}{
		{
			// Name
			"Struct with one string field",
			// 構造体 Name
			struct{ Name string }{"Chris"},
			// 返り値
			[]string{"Chris"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"Nested fields",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
		{
			"Maps",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
		},
	}

	// rangeの返り値はindex,value indexは不要なため_
	for _, test := range cases {
		// 各テストケース名はName
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			// walk関数にInputと無名関数を渡しその結果をinputに入れてこの無名関数を実行
			walk(test.Input, func(input string) {
				// もとのwalk関数のinputの返り値をgotにappend
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCall) {
				t.Errorf("got %v, want %v", got, test.ExpectedCall)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with channels", func(t *testing.T) {
		// Profile型のチャネル作成
		aChannel := make(chan Profile)

		go func() {
			// aChannelへProfileを渡す
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "Katowice"}
			// closeしたらaChannelへはもう送信<-はできず->受信のみ
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v,want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		// Profileを返す関数を変数に入れる
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Katowice"}
		}

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

// haystack = 干し草（=検索対象) needle = 一本の針(検索対象で調べるもの)
func assertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()
	contains := false

	// sliceの中身とneedleがあっていればtrue
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
