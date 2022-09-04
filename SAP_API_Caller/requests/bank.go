package requests

type Bank struct {
	BusinessPartner           string      `json:"BusinessPartner"`
	BankIdentification        string      `json:"BankIdentification"`
	BankCountryKey            string      `json:"BankCountryKey"`
	BankName                 *string      `json:"BankName"`
	BankNumber               *string      `json:"BankNumber"`
	SWIFTCode                *string      `json:"SWIFTCode"`
	BankControlKey           *string      `json:"BankControlKey"`
	BankAccountHolderName    *string      `json:"BankAccountHolderName"`
	BankAccountName          *string      `json:"BankAccountName"`
	ValidityStartDate        *string      `json:"ValidityStartDate"`
	ValidityEndDate          *string      `json:"ValidityEndDate"`
	IBAN                     *string      `json:"IBAN"`
	IBANValidityStartDate    *string      `json:"IBANValidityStartDate"`
	BankAccount              *string      `json:"BankAccount"`
	BankAccountReferenceText *string      `json:"BankAccountReferenceText"`
	CollectionAuthInd        *bool        `json:"CollectionAuthInd"`
	CityName                 *string      `json:"CityName"`
	AuthorizationGroup       *string      `json:"AuthorizationGroup"`
}
