// Package naive реализует наивный алгоритм поиска подстроки со сложностью O(n*m).
// Используется как baseline для сравнения с более быстрыми алгоритмами.
package naive

// Searcher реализует наивный алгоритм поиска подстроки O(n * m).
type Searcher struct{}

// New возвращает новый экземпляр наивного поисковика.
func New() *Searcher {
	return &Searcher{}
}

// Search возвращает все позиции вхождения pattern в text.
// Сравнение выполняется побайтово слева направо.
// Возвращает nil если вхождений не найдено.
func (s *Searcher) Search(text, pattern []byte) []int {
	t := len(text)
	p := len(pattern)

	if p == 0 || p > t {
		return nil
	}

	var indices []int

	for i := 0; i <= t-p; i++ {
		j := 0
		for j < p && text[i+j] == pattern[j] {
			j++
		}
		if j == p {
			indices = append(indices, i)
		}
	}

	return indices
}
