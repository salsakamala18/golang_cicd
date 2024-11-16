package main

import "testing"

func TestHello(t *testing.T) {
    expected := "Hello W"
    if ret := Hello(); ret != expected {
        t.Errorf("hello() is %q want %q", ret, expected)
    }
}
