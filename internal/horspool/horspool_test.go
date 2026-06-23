package horspool

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	cases := []struct {
		name          string
		text, pattern string
		want          []int
	}{
		{"простое совпадение", "hello world", "world", []int{6}},
		{"перекрывающиеся вхождения", "aaaa", "aa", []int{0, 1, 2}},
		{"два не пересекающихся вхождения", "abcabc", "abc", []int{0, 3}},
		{"совпадений нет", "hello", "xyz", nil},
		{"паттерн длиннее текста", "a", "abc", nil},
		{"пустой паттерн", "abc", "", nil},
		{"точное совпадение всего текста", "abc", "abc", []int{0}},
		{"worst case naive", "aaaaab", "aaaab", []int{1}},
		{"большой сдвиг по bad character", "abcdefghij", "xyz", nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := New()
			got := s.Search([]byte(tc.text), []byte(tc.pattern))
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Search(%q, %q) = %v, want %v",
					tc.text, tc.pattern, got, tc.want)
			}
		})
	}
}
