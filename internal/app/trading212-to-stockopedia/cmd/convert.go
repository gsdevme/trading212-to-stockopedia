package cmd

import (
	"fmt"
	"github.com/gsdevme/trading212-to-stockopedia/internal/app/trading212-to-stockopedia/trading212"
	api_client "github.com/gsdevme/trading212-to-stockopedia/internal/pkg/stockopedia"
	"github.com/jszwec/csvutil"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	"io/fs"
	"os"
)

type ProgressBarTicker struct {
	bar *progressbar.ProgressBar
}

func (p *ProgressBarTicker) Tick() {
	err := p.bar.Add(1)
	if err != nil {
		return
	}
}

func newConvertCommand() *cobra.Command {
	c := cobra.Command{
		Use: "convert",
		RunE: func(cmd *cobra.Command, args []string) error {
			file := cmd.Flag("file").Value.String()

			client, err := api_client.NewApiClient(func() api_client.ApiClientConfig {
				return api_client.ApiClientConfig{}
			})
			if err != nil {
				return err
			}

			user := cmd.Flag("username").Value.String()
			pass := cmd.Flag("password").Value.String()

			if user == "" || pass == "" {
				return fmt.Errorf("user/pass is required")
			}

			err = client.Auth(user, pass)

			if err != nil {
				return err
			}

			t, err := trading212.LoadTransactions(file)

			if err != nil {
				return err
			}

			bar := progressbar.NewOptions(len(*t))
			progressbar.OptionSetWidth(45)

			transactions, err := trading212.Convert(t, client, &ProgressBarTicker{bar: bar})
			if err != nil {
				return err
			}

			buf, err := csvutil.Marshal(transactions)
			if err != nil {
				return fmt.Errorf("unable to marshell to csv: %w", err)
			}

			err = os.WriteFile("stockopedia.csv", buf, fs.ModeSetuid)
			if err != nil {
				return err
			}

			fmt.Println("stockopedia.csv has been written, ready to import")
			fmt.Println("https://app.stockopedia.com/portfolio")

			return nil
		},
	}

	c.Flags().String("file", "", "The path to the CSV to convert")
	c.Flags().String("username", os.Getenv("STOCKO_USER"), "Username for stockopedia")
	c.Flags().String("password", os.Getenv("STOCKO_PASS"), "Password for stockopedia")

	cobra.CheckErr(c.MarkFlagRequired("file"))

	return &c
}
