// Package internal определяет общий интерфейс для всех реализаций поиска подстроки.
package internal

// Searcher определяет контракт для алгоритмов поиска подстроки.
type Searcher interface {
	// Search возвращает все позиции вхождения pattern в text.
	// Возвращает nil если вхождений не найдено.
	Search(text, pattern []byte) []int
}
