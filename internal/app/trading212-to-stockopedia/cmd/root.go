package cmd

import (
	"github.com/google/martian/log"
	"github.com/spf13/cobra"
	"os"
)

func NewRootCommand() *cobra.Command {
	c := cobra.Command{
		Use:   "Trading212 to Stockopedia",
		Short: "A simple tool to convert the CSV from Trading212 to a folio compatible CSV for Stockopedia",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	d := os.Getenv("DEBUG")

	if len(d) > 0 {
		log.SetLevel(log.Debug)
	}

	c.AddCommand(newConvertCommand())

	return &c
}
