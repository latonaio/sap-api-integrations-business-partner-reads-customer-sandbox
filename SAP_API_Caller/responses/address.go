package responses

type Address struct {
	D struct {
	BusinessPartner    string `json:"BusinessPartner"`
	AddressID          string `json:"AddressID"`
	ValidityStartDate  string `json:"ValidityStartDate"`
	ValidityEndDate    string `json:"ValidityEndDate"`
	Country            string `json:"Country"`
	Region             string `json:"Region"`
	StreetName         string `json:"StreetName"`
	CityName           string `json:"CityName"`
	PostalCode         string `json:"PostalCode"`
	Language           string `json:"Language"`
	} `json:"d"`
}
