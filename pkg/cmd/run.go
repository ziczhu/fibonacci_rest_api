package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ziczhu/fibonacci_rest_api/pkg/fibonacci"
)

var number int = 0

func init() {
	runCommand.Flags().IntVarP(&number, "number", "n", 0, "The First N Fibonacci Sequence (default: 0)")
	rootCmd.AddCommand(runCommand)
}

var runCommand = &cobra.Command{
	Use:   "run",
	Short: "run fibonacci locally",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Calculating the First N Fibonacci Sequence for N: %d \n", number)
		fib := fibonacci.New(100, 1000)
		for _, val := range fib.GetSequence(number) {
			fmt.Println(val)
		}
	},
}
