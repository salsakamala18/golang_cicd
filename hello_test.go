package hello

import "testing"

func TestHello(t *testing.T) {
    expected = "Hello W"
    if ret := hello(); ret != expected {
        t.Errorf("hello() is %q want %q", ret, expected)
    }
}
