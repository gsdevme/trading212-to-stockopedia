package trading212

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestLoadTransactions(t *testing.T) {
	testDirectory, err := os.Getwd()

	if err != nil {
		t.Fatal(err)
	}

	testDirectory = fmt.Sprintf("%s/../../../../test/fixtures/trading212", testDirectory)

	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    *[]TransactionCsvGBP
		wantErr bool
	}{
		{
			name: "simple",
			args: args{
				filename: fmt.Sprintf("%s/%s", testDirectory, "simple.csv"),
			},
			want: &[]TransactionCsvGBP{
				TransactionCsvGBP{
					Action:   "Deposit",
					DateTime: "2021-08-27 09:54:04",
					Total:    "550.00",
					ID:       "1DX",
				},
				TransactionCsvGBP{
					Action:       "Market buy",
					DateTime:     "2021-09-02 07:01:55",
					ISIN:         "IE00B3XXRP09",
					Ticker:       "VUSA",
					Name:         "Vanguard S&P 500 ETF",
					NoShares:     "8.0000000000",
					Price:        "62.39",
					Currency:     "GBP",
					ExchangeRate: "1.00000",
					Total:        "499.14",
					ID:           "2DX",
				},
				TransactionCsvGBP{
					Action:       "Dividend (Ordinary)",
					DateTime:     "2021-10-06 15:41:57",
					ISIN:         "IE00B3XXRP09",
					Ticker:       "VUSA",
					Name:         "Vanguard S&P 500 ETF",
					NoShares:     "8.0000000000",
					Price:        "0.19",
					Currency:     "GBP",
					ExchangeRate: "Not available",
					Total:        "1.51",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadTransactions(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadTransactions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadTransactions() got = %v, want %v", got, tt.want)
			}
		})
	}
}
