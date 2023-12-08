package main

import (
	"context"
	"fmt"
	"net/http"
)

// Fetchメソッドを持つStoreというinterface
type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// Store interfaceを満たす引数を受け取る
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return // todo: log error however you like
		}

		fmt.Fprint(w, data)
	}
}
