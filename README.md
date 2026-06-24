# go-bmsearch

Реализация алгоритма поиска подстроки Boyer-Moore на Go.

Проект охватывает путь от наивного O(n*m) поиска до полного алгоритма BM
с правилами Bad Character и Good Suffix, а также сравнительные бенчмарки.

## Алгоритмы

| Алгоритм | Preprocessing | Поиск (средн.) | Поиск (худш.) | Особенности |
|----------|---------------|----------------|---------------|-------------|
| Naive    | —             | O(n*m)         | O(n*m)        | baseline для сравнения |
| Horspool | O(m + σ)      | O(n)           | O(n*m)        | только Bad Character, простая реализация |
| Boyer-Moore | O(m + σ)   | O(n/m)         | O(n*m)        | Bad Character + Good Suffix |

Здесь `n` — длина текста, `m` — длина паттерна, `σ` — размер алфавита.

## Структура проекта

```
go-bmsearch/
├── cmd/bmsearch/      # CLI на основе spf13/cobra
├── internal/
│   ├── naive/         # наивная реализация O(n·m)
│   ├── horspool/      # Boyer-Moore-Horspool
│   └── bm/            # полный Boyer-Moore
├── bench/             # сравнительные бенчмарки
└── Makefile           # команды для сборки, тестов и бенчмарков
```

Все алгоритмы реализуют общий интерфейс `internal.Searcher`, что упрощает
их взаимозаменяемость и единообразное тестирование.

## Установка

```bash
git clone https://github.com/your-username/go-bmsearch.git
cd go-bmsearch
make build
```

Бинарник появится в `./bin/bmsearch`.

## Использование

```bash
bmsearch search -p "world" -t "hello world"
bmsearch search --pattern "abc" --text "abcXabcYabc" --algo horspool
bmsearch search -p "aaaab" -t "aaaaab" -a bm
```

Доступные алгоритмы: `naive`, `horspool`, `bm`.

Полный список команд: `bmsearch --help`.

## Разработка

```bash
make help          # список всех команд
make test          # запуск тестов с race detector
make bench         # бенчмарки с замером аллокаций
make fmt vet       # форматирование и статический анализ
```

## Бенчмарки

Сравнение трёх реализаций на различных классах входных данных:

```bash
make bench-compare  # прогон с записью в bench.txt
make bench-stat     # сводная статистика через benchstat
```

Сценарии включают короткий/длинный текст, короткий/длинный паттерн,
worst case для наивного алгоритма (`aaa...ab` × `aaa...ab`),
случайный текст и случай без совпадений.

## Ссылки

- Boyer R.S., Moore J.S. — *A Fast String Searching Algorithm* (1977)
- Crochemore M., Rytter W. — *Jewels of Stringology*
- [wikipedia: Boyer-Moore](https://ru.wikipedia.org/wiki/Алгоритм_Бойера_—_Мура)
- [habr.com: Строковые алгоритмы на практике. Часть 2 — Алгоритм Бойера — Мура](https://habr.com/ru/articles/660767/)

## Лицензия

MIT