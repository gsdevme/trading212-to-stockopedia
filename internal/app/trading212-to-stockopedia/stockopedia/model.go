package stockopedia

// TICKER,DATE,"TIME",TYPE,SHARES,PRICE,"CURRENCY","EXCHANGE RATE","COMMISSION","TAX"

const TypeSell = "SELL"
const TypeBuy = "BUY"
const TypeDeposit = "DEPOSIT"
const TypeDividend = "DIVIDEND"

type TransactionCsv struct {
	Ticker       string `csv:"TICKER"`
	Date         string `csv:"DATE"`
	Time         string `csv:"TIME"`
	Type         string `csv:"TYPE"`
	Shares       string `csv:"SHARES"`
	Price        string `csv:"PRICE"`
	Currency     string `csv:"CURRENCY"`
	ExchangeRate string `csv:"EXCHANGE RATE"`
	Commission   string `csv:"COMMISSION"`
	Tax          string `csv:"TAX"`
}
