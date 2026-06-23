package horspool

// alphabetSize задаёт размер таблицы сдвигов для всех возможных байтов.
const alphabetSize = 256

// Searcher реализует алгоритм Boyer-Moore-Horspool —
// упрощённую версию Boyer-Moore, использующую только правило Bad Character.
type Searcher struct{}

// New возвращает новый экземпляр поисковика Horspool.
func New() *Searcher {
	return &Searcher{}
}

// Search возвращает все позиции вхождения pattern в text.
// Сравнение выполняется справа налево, сдвиг вычисляется по таблице
// последних позиций символов в паттерне.
// Возвращает nil если вхождений не найдено.
func (s *Searcher) Search(text, pattern []byte) []int {
	t := len(text)
	p := len(pattern)

	if p == 0 || p > t {
		return nil
	}

	// Таблица сдвигов строится один раз перед основным циклом.
	shift := buildShiftTable(pattern)

	var indices []int
	i := 0

	// Внешний цикл двигает окно паттерна по тексту.
	for i <= t-p {
		// Сравнение выполняется справа налево начиная с последнего символа.
		j := p - 1
		for j >= 0 && pattern[j] == text[i+j] {
			j--
		}

		if j < 0 {
			// Достигнуто полное совпадение паттерна с текущим окном.
			indices = append(indices, i)
		}

		// Сдвиг определяется по символу текста под последней позицией окна.
		i += shift[text[i+p-1]]
	}

	return indices
}

// buildShiftTable строит таблицу сдвигов для правила Bad Character.
// Для каждого байта таблица хранит расстояние от его последнего вхождения
// в паттерне (без учёта последнего символа) до конца паттерна.
// Символы, не встречающиеся в паттерне, получают сдвиг равный длине паттерна.
func buildShiftTable(pattern []byte) [alphabetSize]int {
	p := len(pattern)

	var shift [alphabetSize]int

	// По умолчанию сдвиг равен длине паттерна — символ не встречался.
	for i := range shift {
		shift[i] = p
	}

	// Последний символ паттерна намеренно пропускается:
	// иначе при совпадении последнего байта сдвиг был бы равен 0
	// и алгоритм зациклился бы.
	for i := 0; i < p-1; i++ {
		shift[pattern[i]] = p - 1 - i
	}

	return shift
}
