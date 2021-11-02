package trading212

type TransactionCsvGBP struct {
	Action       string `csv:"Action"`
	DateTime     string `csv:"Time"`
	ISIN         string `csv:"ISIN"`
	Ticker       string `csv:"Ticker"`
	Name         string `csv:"Name"`
	NoShares     string `csv:"No. of shares"`
	Price        string `csv:"Price / share"`
	Currency     string `csv:"Currency (Price / share)"`
	ExchangeRate string `csv:"Exchange rate"`
	Total        string `csv:"Result (GBP)Total (GBP)"`
	StampDuty    string `csv:"Stamp duty reserve tax (GBP)"`
	FinraFee     string `csv:"Finra fee (GBP)"`
	ID           string `csv:"ID"`
	CurrencyFee  string `csv:"Currency conversion fee (GBP)"`
}
