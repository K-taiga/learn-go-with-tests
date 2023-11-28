package main

import "reflect"

// fnで関数型の引数を受け取る
func walk(x interface{}, fn func(input string)) {
	// xの値のreflect値を取得 ポインターの場合はその実際の値を取得
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	// valの種類によってcaseで異なる処理をする
	switch val.Kind() {
	// 文字列ならfn内の関数を文字列に適用
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		// 初期化; 継続条件; 後処理
		// val.Recv()でチャネルから値を取得 v = 値,ok = 受信成功のbool
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walk(v.Interface(), fn)
		}
	case reflect.Func:
		// valの関数を実行しその中身を走査
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walk(res.Interface(), fn)
		}
	}

}

func getValue(x interface{}) reflect.Value {
	// 受け取ったxの構造体のReflection値を取得しxの型や値を取得
	val := reflect.ValueOf(x)

	// Ptr(ポインター)
	if val.Kind() == reflect.Ptr {
		// reflectで受け取った値がポインタの場合にそのポインタの実際の構造体にアクセスする(ポインタの逆参照のようなもの)
		val = val.Elem()
	}

	return val
}
