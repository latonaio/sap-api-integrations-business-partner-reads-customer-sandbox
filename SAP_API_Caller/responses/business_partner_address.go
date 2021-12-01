package responses

type Address struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			BusinessPartner     string `json:"BusinessPartner"`
			AddressID           string `json:"AddressID"`
			ValidityStartDate   string `json:"ValidityStartDate"`
			ValidityEndDate     string `json:"ValidityEndDate"`
			Country             string `json:"Country"`
			Region              string `json:"Region"`
			StreetName          string `json:"StreetName"`
			CityName            string `json:"CityName"`
			PostalCode          string `json:"PostalCode"`
			Language            string `json:"Language"`
			ToEmailAddress      string `json:"to_EmailAddress"`
			ToFaxNumber         string `json:"to_FaxNumber"`
			ToMobilePhoneNumber string `json:"to_MobilePhoneNumber"`
			ToPhoneNumber       string `json:"to_PhoneNumber"`
			ToURLAddress        string `json:"to_URLAddress"`
		} `json:"results"`
	} `json:"d"`
}
