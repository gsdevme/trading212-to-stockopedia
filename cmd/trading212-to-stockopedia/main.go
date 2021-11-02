package main

import (
	"fmt"
	"github.com/gsdevme/trading212-to-stockopedia/internal/app/trading212-to-stockopedia/cmd"
	"os"
)

func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
