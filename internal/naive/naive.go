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
	return nil
}
