package bench

import (
	"bytes"
	"math/rand"
	"testing"

	"github.com/albuilov/go-bmsearch/internal"
	"github.com/albuilov/go-bmsearch/internal/bm"
	"github.com/albuilov/go-bmsearch/internal/horspool"
	"github.com/albuilov/go-bmsearch/internal/naive"
)

// scenario описывает входные данные для бенчмарка.
type scenario struct {
	name    string
	text    []byte
	pattern []byte
}

// algorithms перечисляет все реализации поисковика для сравнения.
var algorithms = []struct {
	name     string
	searcher internal.Searcher
}{
	{"Naive", naive.New()},
	{"Horspool", horspool.New()},
	{"BM", bm.New()},
}

// buildScenarios формирует набор сценариев, покрывающих разные классы входов.
func buildScenarios() []scenario {
	return []scenario{
		{
			name:    "ShortText_ShortPattern",
			text:    []byte("the quick brown fox jumps over the lazy dog"),
			pattern: []byte("fox"),
		},
		{
			name:    "LongText_ShortPattern",
			text:    bytes.Repeat([]byte("abcdefghij"), 1000),
			pattern: []byte("hij"),
		},
		{
			name:    "LongText_LongPattern",
			text:    bytes.Repeat([]byte("abcdefghij"), 1000),
			pattern: []byte("abcdefghij"),
		},
		{
			name:    "WorstCaseNaive",
			text:    append(bytes.Repeat([]byte("a"), 10000), 'b'),
			pattern: append(bytes.Repeat([]byte("a"), 50), 'b'),
		},
		{
			name:    "NoMatch",
			text:    bytes.Repeat([]byte("abcdefghij"), 1000),
			pattern: []byte("xyz"),
		},
		{
			name:    "RandomText",
			text:    randomBytes(10000, 42),
			pattern: []byte("xyzqwerty"),
		},
	}
}

// BenchmarkSearch прогоняет все алгоритмы по всем сценариям.
// Результаты можно сравнивать через benchstat между ветками алгоритмов.
func BenchmarkSearch(b *testing.B) {
	scenarios := buildScenarios()

	for _, sc := range scenarios {
		for _, algo := range algorithms {
			name := sc.name + "/" + algo.name
			b.Run(name, func(b *testing.B) {
				b.ReportAllocs()
				b.SetBytes(int64(len(sc.text)))
				b.ResetTimer()

				for range b.N {
					_ = algo.searcher.Search(sc.text, sc.pattern)
				}
			})
		}
	}
}

// randomBytes генерирует детерминированный набор байтов для воспроизводимых бенчмарков.
func randomBytes(n int, seed int64) []byte {
	r := rand.New(rand.NewSource(seed))
	buf := make([]byte, n)
	for i := range buf {
		buf[i] = byte('a' + r.Intn(26))
	}
	return buf
}
