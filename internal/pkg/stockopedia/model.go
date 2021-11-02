package stockopedia

// https://mholt.github.io/json-to-go/
type SearchResults struct {
	Content struct {
		Data struct {
			Security struct {
				Result []struct {
					ID                 int    `json:"id"`
					Ticker             string `json:"ticker"`
					GoogleTicker       string `json:"google_ticker"`
					ExchangeTicker     string `json:"exchange_ticker"`
					ExchangeNameTicker string `json:"exchange_name_ticker"`
					Ric                string `json:"ric"`
					Name               string `json:"name"`
					CountryCode        string `json:"country_code"`
					Exchange           string `json:"exchange"`
					Sector             string `json:"sector"`
					IndustryGroup      string `json:"industry_group"`
					Image              string `json:"image"`
					Type               string `json:"type"`
					IsPrimary          bool   `json:"is_primary"`
					IsListed           bool   `json:"is_listed"`
				} `json:"result"`
				TotalCount int `json:"total_count"`
			} `json:"security"`
		} `json:"data"`
		Type string `json:"type"`
	} `json:"content"`
	Success bool `json:"success"`
	Paging  struct {
		TotalResults int    `json:"totalResults"`
		TotalPages   int    `json:"totalPages"`
		PerPage      int    `json:"perPage"`
		CurrentPage  int    `json:"currentPage"`
		From         int    `json:"from"`
		To           int    `json:"to"`
		First        bool   `json:"_first"`
		Last         bool   `json:"_last"`
		URL          string `json:"_url"`
		Sort         string `json:"sort"`
		Order        string `json:"order"`
	} `json:"paging"`
}
