package main

import (
	"github.com/gsdevme/trading212-to-stockopedia/internal/app/trading212-to-stockopedia/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cobra.CheckErr(cmd.NewRootCommand().Execute())
}
