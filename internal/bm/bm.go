// Package bm реализует полный алгоритм Boyer-Moore с правилами
// Bad Character и Good Suffix, обеспечивающий сублинейный поиск в среднем случае.
package bm

// alphabetSize задаёт размер таблицы для всех возможных значений байта.
const alphabetSize = 256

// Searcher реализует полный алгоритм Boyer-Moore
// с правилами Bad Character и Good Suffix.
type Searcher struct{}

// New возвращает новый экземпляр поисковика Boyer-Moore.
func New() *Searcher {
	return &Searcher{}
}

// Search возвращает все позиции вхождения pattern в text.
// На каждом шаге выбирается максимальный сдвиг из двух правил:
// Bad Character и Good Suffix.
// Возвращает nil если вхождений не найдено.
func (s *Searcher) Search(text, pattern []byte) []int {
	t := len(text)
	p := len(pattern)

	if p == 0 || p > t {
		return nil
	}

	badChar := buildBadCharTable(pattern)
	goodSuffix := buildGoodSuffixTable(pattern)

	var indices []int
	i := 0

	for i <= t-p {
		j := p - 1

		// Сравнение выполняется справа налево.
		for j >= 0 && pattern[j] == text[i+j] {
			j--
		}

		if j < 0 {
			// Полное совпадение зафиксировано.
			indices = append(indices, i)
			// Сдвиг после матча берётся из таблицы good suffix для позиции 0.
			i += goodSuffix[0]
		} else {
			// При несовпадении выбирается максимальный сдвиг.
			bcShift := j - badChar[text[i+j]]
			gsShift := goodSuffix[j+1]
			i += max(bcShift, gsShift)
		}
	}

	return indices
}

// buildBadCharTable строит таблицу последних позиций каждого байта в паттерне.
// Для символов, не встречающихся в паттерне, значение остаётся -1,
// что даёт максимальный сдвиг при использовании в основном цикле.
func buildBadCharTable(pattern []byte) [alphabetSize]int {
	var table [alphabetSize]int
	for i := range table {
		table[i] = -1
	}
	for i, b := range pattern {
		table[b] = i
	}
	return table
}

// buildGoodSuffixTable строит таблицу сдвигов по правилу Good Suffix.
// Алгоритм построения состоит из двух проходов:
// первый обрабатывает Case 1 (повторное вхождение суффикса в паттерне),
// второй — Case 2 (префикс паттерна совпадает с суффиксом совпавшей части).
func buildGoodSuffixTable(pattern []byte) []int {
	p := len(pattern)
	shift := make([]int, p+1)
	border := make([]int, p+1)

	// Первый проход: вычисляются позиции "широких границ" суффиксов.
	i := p
	j := p + 1
	border[i] = j

	for i > 0 {
		for j <= p && pattern[i-1] != pattern[j-1] {
			if shift[j] == 0 {
				shift[j] = j - i
			}
			j = border[j]
		}
		i--
		j--
		border[i] = j
	}

	// Второй проход: для позиций без сдвига применяется правило префикса.
	j = border[0]
	for i := 0; i <= p; i++ {
		if shift[i] == 0 {
			shift[i] = j
		}
		if i == j {
			j = border[j]
		}
	}

	return shift
}
