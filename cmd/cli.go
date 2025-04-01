package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	url         string
	requests    int
	concurrency int
)

var rootCmd = &cobra.Command{
	Use:   "load-tester",
	Short: "Ferramenta CLI para testes de carga HTTP",
	Run: func(cmd *cobra.Command, args []string) {
		if url == "" || requests <= 0 || concurrency <= 0 {
			fmt.Println("Parâmetros inválidos. Use --url, --requests e --concurrency corretamente.")
			os.Exit(1)
		}
		RunTest(url, requests, concurrency)
	},
}

func init() {
	rootCmd.Flags().StringVar(&url, "url", "", "URL do serviço a ser testado")
	rootCmd.Flags().IntVar(&requests, "requests", 1, "Número total de requests")
	rootCmd.Flags().IntVar(&concurrency, "concurrency", 1, "Número de chamadas simultâneas")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
