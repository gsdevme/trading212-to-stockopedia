package trading212

import (
	"github.com/google/martian/log"
	"github.com/gsdevme/trading212-to-stockopedia/internal/app/trading212-to-stockopedia/stockopedia"
	stocko "github.com/gsdevme/trading212-to-stockopedia/internal/pkg/stockopedia"
	"github.com/jszwec/csvutil"
	"os"

	"sync"
	"time"
)

type Ticker interface {
	Tick()
}

func LoadTransactions(filename string) (*[]TransactionCsvGBP, error) {
	var t []TransactionCsvGBP

	f, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	if err := csvutil.Unmarshal(f, &t); err != nil {
		return nil, err
	}

	return &t, nil
}

func Convert(t *[]TransactionCsvGBP, client *stocko.ApiClient, ticker Ticker) (*[]stockopedia.TransactionCsv, error) {
	var wg sync.WaitGroup
	var transactions []stockopedia.TransactionCsv

	for _, trans := range *t {
		wg.Add(1)

		go func(trans TransactionCsvGBP) {
			ticker.Tick()

			defer wg.Done()

			transaction := stockopedia.TransactionCsv{}

			t, err := time.Parse("2006-01-02 15:04:05", trans.DateTime)

			if err != nil {
				log.Errorf("unable to parse time: %w", err)

				return
			}

			transaction.Currency = trans.Currency

			switch trans.Action {
			case "Market buy":
				transaction.Type = "buy"
			case "Market sell", "Limit sell":
				transaction.Type = "sell"
			}

			transaction.Price = trans.Price
			transaction.Shares = trans.NoShares
			transaction.ExchangeRate = trans.ExchangeRate
			transaction.Date = t.Format("02/01/2006")
			transaction.Time = t.Format("15:04:05")

			r, err := client.SearchSecurity(trans.ISIN)

			if err != nil {
				log.Errorf("unable to match security %s: %w", trans.Ticker, err)

				return
			}

			if r == nil {
				log.Errorf("unable to match security %s", trans.Ticker)

				return
			}

			if len(r.Content.Data.Security.Result) == 0 {
				return
			}

			transaction.Ticker = r.Content.Data.Security.Result[0].GoogleTicker
			transaction.Commission = ""
			transaction.Tax = trans.StampDuty

			transactions = append(transactions, transaction)
		}(trans)
	}

	wg.Wait()

	return &transactions, nil
}
