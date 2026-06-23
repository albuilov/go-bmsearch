package main

import (
	"fmt"

	"github.com/albuilov/go-bmsearch/internal"
	"github.com/albuilov/go-bmsearch/internal/naive"
	"github.com/spf13/cobra"
)

// Фдаги команд search.
var (
	flagPattern string
	flagText    string
	flagAlgo    string
)

// Поддерживаемые алгоритмы поиска.
const (
	algoNaive    = "naive"
	algoHorspool = "horspool"
	algoBM       = "bm"
)

var searchCmd = &cobra.Command{
	Use:   "search [text] [pattern]",
	Short: "Выполняет поиск pattern в text выбранным алгоритмом",
	Example: `  bmsearch search -p "abc" -t "abcabcabc"
  bmsearch search --pattern "world" --text "hello world" --algo horspool`,
	RunE: runSearch,
}

func init() {
	searchCmd.Flags().StringVarP(&flagPattern, "pattern", "p", "", "искомая подстрока (обязательно)")
	searchCmd.Flags().StringVarP(&flagText, "text", "t", "", "текст для поиска (обязательно)")
	searchCmd.Flags().StringVarP(&flagAlgo, "algo", "a", algoNaive,
		fmt.Sprintf("алгоритм поиска: %s | %s | %s", algoNaive, algoHorspool, algoBM))

	_ = searchCmd.MarkFlagRequired("pattern")
	_ = searchCmd.MarkFlagRequired("text")

	rootCmd.AddCommand(searchCmd)
}

func runSearch(cmd *cobra.Command, args []string) error {
	// Выбирается реализация поисковика по флагу --algo.
	searcher, err := newSearcher(flagAlgo)
	if err != nil {
		return err
	}

	indices := searcher.Search([]byte(flagText), []byte(flagPattern))

	if len(indices) == 0 {
		fmt.Println("вхождений не найдено")
		return nil
	}

	fmt.Printf("найдено %d вхождений: %v\n", len(indices), indices)
	return nil
}

// newSearcher возвращает реализацию Searcher по имени алгоритма.
func newSearcher(algo string) (internal.Searcher, error) {
	switch algo {
	case algoNaive:
		return naive.New(), nil
	default:
		return nil, fmt.Errorf("неизвестный алгоритм: %q", algo)
	}
}
