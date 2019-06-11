package main

import "testing"

func TestFib(t *testing.T) {
	if fib(1) != 1 {
		t.Fatal("fib(n)")
	}
}
