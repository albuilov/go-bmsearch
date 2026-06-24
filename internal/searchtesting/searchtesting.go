// Package searchtesting предоставляет общий набор тест-кейсов и вспомогательную
// функцию Run для единообразного тестирования всех реализаций internal.Searcher.
package searchtesting

import (
	"slices"
	"testing"

	"github.com/albuilov/go-bmsearch/internal"
)

// Case описывает один тест-кейс для алгоритма поиска подстроки.
type Case struct {
	Name    string
	Text    string
	Pattern string
	Want    []int
}

// Cases — полный набор кейсов, покрывающих граничные условия и специфику каждого алгоритма.
var Cases = []Case{
	// --- базовые ---
	{"простое совпадение", "hello world", "world", []int{6}},
	{"два непересекающихся вхождения", "abcabc", "abc", []int{0, 3}},
	{"совпадение в начале", "abcdef", "abc", []int{0}},
	{"совпадение в конце", "abcdef", "def", []int{3}},
	{"одиночный символ", "abcabc", "b", []int{1, 4}},
	{"точное совпадение всего текста", "abc", "abc", []int{0}},

	// --- граничные ---
	{"пустой паттерн", "abc", "", nil},
	{"пустой текст", "", "abc", nil},
	{"паттерн длиннее текста", "a", "abc", nil},
	{"текст и паттерн равны по длине без совпадения", "abc", "xyz", nil},
	{"совпадений нет", "hello", "xyz", nil},
	{"одиночный символ — совпадение", "a", "a", []int{0}},
	{"одиночный символ — нет совпадения", "a", "b", nil},

	// --- перекрывающиеся вхождения ---
	{"перекрывающиеся вхождения aa", "aaaa", "aa", []int{0, 1, 2}},
	{"перекрывающиеся вхождения aaa", "aaaaa", "aaa", []int{0, 1, 2}},

	// --- повторяющиеся символы ---
	{"повторяющиеся символы", "aaabaaabaaab", "aaab", []int{0, 4, 8}},
	{"worst case — почти совпадение", "aaaaab", "aaaab", []int{1}},

	// --- специфика Good Suffix (Boyer-Moore) ---
	{"good suffix case 1", "abcxxxabcabc", "abcabc", []int{6}},
	{"good suffix case 2", "xyzabcabczabc", "abcabc", []int{3}},

	// --- длинные тексты ---
	{"длинный текст без совпадений", "abcdefghijklmnop", "xyz", nil},
	{"длинный паттерн без совпадений в коротком тексте", "abcde", "abcdef", nil},

	// --- однобайтовые символы ---
	{"паттерн из одного байта — несколько вхождений", "xaxbxcx", "x", []int{0, 2, 4, 6}},
	{"текст из одного символа — совпадение", "z", "z", []int{0}},
	{"текст из одного символа — нет совпадения", "z", "a", nil},

	// --- бинарные данные ---
	{"нулевые байты в тексте", "\x00\x01\x00", "\x00", []int{0, 2}},
	{"паттерн — нулевой байт", "\x00\x01\x02", "\x01", []int{1}},
}

// Run прогоняет все Cases через переданный Searcher.
// Вызывается из TestSearch каждого алгоритмового пакета.
func Run(t *testing.T, s internal.Searcher) {
	t.Helper()

	for _, tc := range Cases {
		t.Run(tc.Name, func(t *testing.T) {
			got := s.Search([]byte(tc.Text), []byte(tc.Pattern))
			if !slices.Equal(got, tc.Want) {
				t.Errorf("Search(%q, %q) = %v, want %v",
					tc.Text, tc.Pattern, got, tc.Want)
			}
		})
	}
}
