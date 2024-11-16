package assert

import (
        "fmt"
        "testing"
        "strings"
)

func TestEqual(t *testing.T) {
	type myType string

	mockT := new(testing.T)
	var m map[string]interface{}

	cases := []struct {
		expected interface{}
		actual   interface{}
		result   bool
		remark   string
	}{
		{"Hello World", "Hello World", true, ""},
		{123, 123, true, ""},
		{123.5, 123.5, true, ""},
		{[]byte("Hello World"), []byte("Hello World"), true, ""},
		{nil, nil, true, ""},
		{int32(123), int32(123), true, ""},
		{uint64(123), uint64(123), true, ""},
		{myType("1"), myType("1"), true, ""},
		{&struct{}{}, &struct{}{}, true, "pointer equality is based on equality of underlying value"},

		// Not expected to be equal
		{m["bar"], "something", false, ""},
		{myType("1"), myType("2"), false, ""},

		// A case that might be confusing, especially with numeric literals
		{10, uint(10), false, ""},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Equal(%#v, %#v)", c.expected, c.actual), func(t *testing.T) {
			res := Equal(mockT, c.expected, c.actual)

			if res != c.result {
				t.Errorf("Equal(%#v, %#v) should return %#v: %s", c.expected, c.actual, c.result, c.remark)
			}
		})
	}
}


func TestObjectsAreEqual(t *testing.T) {
	cases := []struct {
		expected interface{}
		actual   interface{}
		result   bool
	}{
		// cases that are expected to be equal
		{"Hello World", "Hello World", true},
		{123, 123, true},
		{123.5, 123.5, true},
		{[]byte("Hello World"), []byte("Hello World"), true},
		{nil, nil, true},

		// cases that are expected not to be equal
		{map[int]int{5: 10}, map[int]int{10: 20}, false},
		{'x', "x", false},
		{"x", 'x', false},
		{0, 0.1, false},
		{0.1, 0, false},
		{time.Now, time.Now, false},
		{func() {}, func() {}, false},
		{uint32(10), int32(10), false},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("ObjectsAreEqual(%#v, %#v)", c.expected, c.actual), func(t *testing.T) {
			res := ObjectsAreEqual(c.expected, c.actual)

			if res != c.result {
				t.Errorf("ObjectsAreEqual(%#v, %#v) should return %#v", c.expected, c.actual, c.result)
			}

		})
	}
}

func TestFormatUnequalValues(t *testing.T) {
	expected, actual := formatUnequalValues("foo", "bar")
	Equal(t, `"foo"`, expected, "value should not include type")
	Equal(t, `"bar"`, actual, "value should not include type")

	expected, actual = formatUnequalValues(123, 123)
	Equal(t, `123`, expected, "value should not include type")
	Equal(t, `123`, actual, "value should not include type")

	expected, actual = formatUnequalValues(int64(123), int32(123))
	Equal(t, `int64(123)`, expected, "value should include type")
	Equal(t, `int32(123)`, actual, "value should include type")

	expected, actual = formatUnequalValues(int64(123), nil)
	Equal(t, `int64(123)`, expected, "value should include type")
	Equal(t, `<nil>(<nil>)`, actual, "value should include type")

	type testStructType struct {
		Val string
	}

	expected, actual = formatUnequalValues(&testStructType{Val: "test"}, &testStructType{Val: "test"})
	Equal(t, `&assert.testStructType{Val:"test"}`, expected, "value should not include type annotation")
	Equal(t, `&assert.testStructType{Val:"test"}`, actual, "value should not include type annotation")
}

func Test_truncatingFormat(t *testing.T) {

	original := strings.Repeat("a", bufio.MaxScanTokenSize-102)
	result := truncatingFormat(original)
	Equal(t, fmt.Sprintf("%#v", original), result, "string should not be truncated")

	original = original + "x"
	result = truncatingFormat(original)
	NotEqual(t, fmt.Sprintf("%#v", original), result, "string should have been truncated.")

	if !strings.HasSuffix(result, "<... truncated>") {
		t.Error("truncated string should have <... truncated> suffix")
	}
}

func Test_validateEqualArgs(t *testing.T) {
	if validateEqualArgs(func() {}, func() {}) == nil {
		t.Error("non-nil functions should error")
	}

	if validateEqualArgs(func() {}, func() {}) == nil {
		t.Error("non-nil functions should error")
	}

	if validateEqualArgs(nil, nil) != nil {
		t.Error("nil functions are equal")
	}
}
