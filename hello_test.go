package hello

import "testing"

func TestHello() {
    expected = "Hello W"
    if ret := hello(); ret != expected {
        t.Errorf("hello() is %q want %q", ret, expected)
    }
}
