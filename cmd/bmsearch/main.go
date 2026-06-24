// Package main содержит точку входа CLI-инструмента bmsearch.
package main

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bmsearch",
	Short: "Поиск подстроки алгоритмом Бойера — Мура",
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
