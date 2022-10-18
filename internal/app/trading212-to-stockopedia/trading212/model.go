package trading212

const ExchangeRateEmpty = "Not available"
const PenceCurrency = "GBX"

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
	ID           string `csv:"ID"`
	Total        string `csv:"Total (GBP)"`
	// Sometimes omitted
	StampDuty   string `csv:"Stamp duty reserve tax (GBP)"`
	FinraFee    string `csv:"Finra fee (GBP)"`
	CurrencyFee string `csv:"Currency conversion fee (GBP)"`
}
