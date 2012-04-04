// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// strings 패키지에 대한 테스트
// $GOROOT/src/pkg/strings/*_test.go 참조
package strings_test

import (
	"fmt"
	"testing"
	"strings"
	"unicode"
)

// Test
// func Index(s, sep string) int
var indexTests = []struct {
	s, sep string
	out int
}{
	{"", "", 0},
	{"", "foo", -1}, // 없으면 -1
	{"fo", "foo", -1},
	{"foo", "foo", 0},
	{"oofofoofooo", "f", 2},
	{"barfoobarfoo", "foo", 3},
	{"foo", "", 0},
	{"foo", "o", 1},
	{"abcABCabc", "A", 3},
	{"대한민국", "한", 3}, // 한글은 utf-8 인코딩 기준
}

func TestIndex(t *testing.T) {
	for _, tc := range indexTests {
		actual := strings.Index(tc.s, tc.sep)
		if actual != tc.out {
			t.Errorf("strings.Index(%q, %q) = %v; want %v", tc.s, tc.sep, actual, tc.out)
		}
	}

}

func ExampleIndex() {
	fmt.Printf("%d\n", len("한"))
	// Output:
	// 3
}



// Test
// func LastIndex(s, sep string) int
// s 안에서 sep가 여러 개 발견되더라도 sep가 마지막에 발견된 인덱스를 리턴
var lastIndexTests = []struct {
	s, sep string
	out int
}{
	{"", "", 0},
	{"", "a", -1}, // 없으면 -1
	{"", "foo", -1},
	{"fo", "foo", -1},
	{"foo", "foo", 0},
	{"foo", "f", 0},
	{"oofofoofooo", "f", 7},
	{"oofofoofooo", "foo", 7},
	{"barfoobarfoo", "foo", 9},
	{"foo", "", 3}, // 마지막 인덱스 + 1 부분 리턴
	{"foo", "o", 2},
	{"abcABCabc", "A", 3},
	{"abcABCabc", "a", 6},
	{"구글고고랭", "고", 9}, // 마지막 '고'가 발견된 인덱스 리턴
}

func TestLastIndex(t *testing.T) {
	for _, tc := range lastIndexTests {
		actual := strings.LastIndex(tc.s, tc.sep)
		if actual != tc.out {
			t.Errorf("strings.LastIndex(%q, %q) = %v; want %v", tc.s, tc.sep, actual, tc.out)
		}
	}
}


// Test
// func IndexRune(s string, r rune) int
var indexRuneTests = []struct {
	s string
	r rune
	out int
}{
	{"a A x", 'A', 2},
	{"some_text=some_value", '=', 9},
	{"☺a", 'a', 3},
	{"a☻☺b", '☺', 4},
	{"대한민국", '한', 3},
}

func TestIndexRune(t *testing.T) {
	for _, tc := range indexRuneTests {
		actual := strings.IndexRune(tc.s, tc.r)
		if actual != tc.out {
			t.Errorf("strings.IndexRune(%q, %q) = %v; want %v", tc.s, tc.r, actual, tc.out)
		}
	}
}

func ExampleIndexRune() {
	fmt.Println(len("☺"))
	fmt.Println(len("☻"))
	// Output:
	// 3
	// 3
}


// Test
// func IndexAny(s, chars string) int
//
// IndexAny returns the index of the first instance of any Unicode code point
// from chars in s, or -1 if no Unicode code point from chars is present in s.
var indexAnyTests = []struct {
	s, chars string
	out int
}{
	{"", "", -1}, // Note
	{"", "a", -1},
	{"", "abc", -1},
	{"a", "", -1}, // Note
	{"a", "a", 0},
	{"aaa", "a", 0},
	{"abc", "xyz", -1},
	{"abc", "xcz", 2}, // x, z는 없고 c가 발견된 인덱스 2 리턴
	{"abcxyz", "xyab", 0}, // 모두 다 발견되나 a가 발견된 인덱스 0 리턴
	{"a☺b☻c☹d", "uvw☻xyz", 2 + len("☺")},
	{"aRegExp*", ".(|)*+?^$[]", 7},
}

func TestIndexAny(t *testing.T) {
	for _, tc := range indexAnyTests {
		actual := strings.IndexAny(tc.s, tc.chars)
		if actual != tc.out {
			t.Errorf("strings.IndexAny(%q, %q) = %v; want %v", tc.s, tc.chars, actual, tc.out)
		}
	}
}


// Test
// func LastIndexAny(s, chars string) int
var lastIndexAnyTests = []struct {
	s, chars string
	out int
}{
	{"", "", -1},
	{"", "a", -1},
	{"", "abc", -1},
	{"a", "", -1},
	{"a", "a", 0},
	{"aaa", "a", 2}, // a가 발견된 마지막 인덱스 2 리턴
	{"abc", "xyz", -1},
	{"abc", "ab", 1}, // b가 발견된 마지막 인덱스 1 리턴
	{"a☺b☻c☹d", "uvw☻xyz", 2 + len("☺")}, // chars의 나머지 글자들은 s에 없고, ☻ <- 이게 발견된 5 리턴
	{"a.RegExp*", ".(|)*+?^$[]", 8},
}

func TestLastIndexAny(t *testing.T) {
	for _, tc := range lastIndexAnyTests {
		actual := strings.LastIndexAny(tc.s, tc.chars)
		if actual != tc.out {
			t.Errorf("strings.LastIndexAny(%q, %q) = %v; want %v", tc.s, tc.chars, actual, tc.out)
		}
	}
}


// Test
// func IndexFunc(s string, f func(rune) bool) int
// func LastIndexFunc(s string, f func(rune) bool) int
//
// IndexFunc 함수는 f(c)를 만족하는 s 속의 첫번째 인덱스 리턴
// rune, unicode에 대한 더 자세한 예제는
// $GOROOT/src/pkg/strings/strings_test.go 참조
const space = "\t\v\r\f\n\u0085\u00a0\u2000\u3000"

var indexFuncTests = []struct {
	s string
	f func(rune) bool
	fname string
	first, last int
}{
	{"abc", unicode.IsDigit, "IsDigit", -1, -1},
	{"0123", unicode.IsDigit, "IsDigit", 0, 3},
	{"a1b", unicode.IsDigit, "IsDigit", 1, 1},
	{"Republic of Korea", unicode.IsUpper, "IsUpper", 0, 12},
	{space, unicode.IsSpace, "IsSpace", 0, len(space) - 3}, // last rune in space is 3 bytes. len(space) == 15
}

func TestIndexFunc(t *testing.T) {
	for _, tc := range indexFuncTests {
		first := strings.IndexFunc(tc.s, tc.f)
		if first != tc.first {
			t.Errorf("strings.IndexFunc(%q, %s) = %d; want %d", tc.s, tc.fname, first, tc.first)
		}
		last := strings.LastIndexFunc(tc.s, tc.f)
		if last != tc.last {
			t.Errorf("strings.LastIndexFunc(%q, %s) = %d; want %d", tc.s, tc.fname, last, tc.last)
		}
	}
}


// Test
// func Contains(s, substr string) bool
// s 안에 substr이 있으면 true 리턴. 내부적으로 strings.Index() 호출
var containsTests = []struct {
	s, substr string
	out bool
}{
	{"abc", "bc", true},
	{"abc", "bcd", false},
	{"abc", "", true},
	{"", "a", false},
	{"seafood", "", true},
	{"", "", true},
	{"대한 민국 동해물과", "동해", true},
}

func TestContains(t *testing.T) {
	for _, tc := range containsTests {
		if strings.Contains(tc.s, tc.substr) != tc.out {
			t.Errorf("strings.Contains(%s, %s) = %v, want %v",
				tc.s, tc.substr, !tc.out, tc.out)
		}
	}
}


// Test
// func ContainsAny(s, chars string) bool
// s 안에 chars의 어떤 유니코드 코드 포인드라도 포함되어 있으면 true 리턴
// 내부적으로 strings.IndexAny() 호출
var containsAnyTests = []struct {
	s, chars string
	out bool
}{
	{"", "", false}, // Note
	{"a", "", false}, // Note
	{"a", "a", true},
	{"aaa", "a", true},
	{"abc", "xyz", false},
	{"abc", "xcz", true},
}

func TestContainsAny(t *testing.T) {
	for _, tc := range containsAnyTests {
		if strings.ContainsAny(tc.s, tc.chars) != tc.out {
			t.Errorf("strings.ContainsAny(%s, %s) = %v, want %v",
				tc.s, tc.chars, !tc.out, tc.out)
		}
	}
}


// Test
// func ContainsRune(s string, r rune) bool
var containsRuneTests = []struct {
	str      string
	r        rune
	out bool
}{
	{"", 'a', false},
	{"a", 'a', true},
	{"aaa", 'a', true},
	{"abc", 'y', false},
	{"abc", 'c', true},
	{"a☺b☻c☹d", 'x', false},
	{"a☺b☻c☹d", '☻', true},
	{"aRegExp*", '*', true},
}

func TestContainsRune(t *testing.T) {
	for _, tc := range containsRuneTests {
		if strings.ContainsRune(tc.str, tc.r) != tc.out {
			t.Errorf("strings.ContainsRune(%s, %s) = %v, want %v",
				tc.str, tc.r, !tc.out, tc.out)
		}
	}
}


// Test
// func Count(s, sep string) int
var countTests = []struct {
	s, sep string
	out int
}{
	{"cheese", "e", 3},
	{"five", "", 5},
	{"뽀로로", "로", 2},
}

func TestCount(t *testing.T) {
	for _, tc := range countTests {
		actual := strings.Count(tc.s, tc.sep)
		if actual != tc.out {
			t.Errorf("strings.Count(%s, %s) = %v, want %v", tc.s, tc.sep, actual, tc.out)
		}
	}
}


// Test
// func EqualFold(s, t string) bool
// UTF-8로 해석된 s, t가 유니코드 case-folding에 매핑된 같은 문자인지 비교
var equalFoldTests = []struct {
	s, t string
	out  bool
}{
	{"abc", "abc", true},
	{"123abc", "123ABC", true},
	{"αβδ", "ΑΒΔ", true},
	{"abc", "xyz", false},
	{"abcdefghijk", "abcdefghijX", false},
	{"abcdefghijk", "abcdefghij\u212A", true},
	{"abcdefghijK", "abcdefghij\u212A", true}, // \u212A는 대문자 K가 아닌 KELVIN SIGN
	{"abcdefghijK", "abcdefghij\u004B", true}, // \u004B는 대문자 K
	{"abcdefghijkz", "abcdefghij\u212Ay", false},
}

func TestEqualFold(t *testing.T) {
	for _, tc := range equalFoldTests {
		if out := strings.EqualFold(tc.s, tc.t); out != tc.out {
			t.Errorf("strings.EqualFold(%#q, %#q) = %v, want %v", tc.s, tc.t, out, tc.out)
		}
		if out := strings.EqualFold(tc.t, tc.s); out != tc.out {
			t.Errorf("strings.EqualFold(%#q, %#q) = %v, want %v", tc.t, tc.s, out, tc.out)
		}
	}
}


// Test
// func Fields(s string) []string
// 내부적으로 FieldsFunc(s, unicode.IsSpace)를 호출. 
// unicode.IsSpace는 인자로 주어진 rune이 Unicode White Space property에 속한 것이면
// true 리턴. 아래 외에도 더 있음
//   '\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP)
// see http://golang.org/src/pkg/unicode/graphic.go?s=3547:3572#L103
// see $GOROOT/src/pkg/unicode/tables.go:_White_Space 참조
var fieldTests = []struct {
	s string
	out []string
}{
	{"", []string{}},
	{" ", []string{}},
	{" \t ", []string{}},
	{"  abc  ", []string{"abc"}},
	{"1 2 3 4", []string{"1", "2", "3", "4"}},
	{"1  2  3  4", []string{"1", "2", "3", "4"}},
	{"1\t\t2\t\t3\t4", []string{"1", "2", "3", "4"}},
	{"1\u20002\u20013\u20024", []string{"1", "2", "3", "4"}},
	{"\u2000\u2001\u2002", []string{}},
	{"\n™\t™\n", []string{"™", "™"}},
}

func eq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestFields(t *testing.T) {
	for _, tc := range fieldTests {
		actual := strings.Fields(tc.s)
		if !eq(actual, tc.out) {
			t.Errorf("strings.Fields(%q) = %v; want %v", tc.s, actual, tc.out)
		}
	}
}


// Test
// func FieldsFunc(s string, f func(rune) bool) []string
//    FieldsFunc splits the string s at each run of Unicode code points c
//    satisfying f(c) and returns an array of slices of s. If all code points
//    in s satisfy f(c) or the string is empty, an empty slice is returned.
var fieldsFuncTests = []struct {
	s string
	out []string
}{
	{"", []string{}},
	{"XX", []string{}},
	{"XXhiXXX", []string{"hi"}},
	{"aXXbXXXcX", []string{"a", "b", "c"}},
}

func TestFieldsFunc(t *testing.T) {
	pred := func(c rune) bool { return c == 'X' }
	for _, tc := range fieldsFuncTests {
		actual := strings.FieldsFunc(tc.s, pred)
		if !eq(actual, tc.out) {
			t.Errorf("strings.FieldsFunc(%q) = %v, want %v", tc.s, actual, tc.out)
		}
	}
}
