package main

import (
	"errors"
	"fmt"
)

func example(code string) (int, error) {
	if code == "hoge" {
		a := uint(1 << 63)
		fmt.Printf("1<<63=%d\n", a)
		return 1, nil
	}
	return 0, errors.New("code must be hoge")
}
