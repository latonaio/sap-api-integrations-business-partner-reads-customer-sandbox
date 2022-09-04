package requests

type Company struct {
	Customer                       string  `json:"Customer"`
	CompanyCode                    string  `json:"CompanyCode"`
	APARToleranceGroup             *string `json:"APARToleranceGroup"`
	CustomerSupplierClearingIsUsed *bool   `json:"CustomerSupplierClearingIsUsed"`
	HouseBank                      *string `json:"HouseBank"`
	PaymentMethodsList             *string `json:"PaymentMethodsList"`
	PaymentTerms                   *string `json:"PaymentTerms"`
	ReconciliationAccount          *string `json:"ReconciliationAccount"`
	DeletionIndicator              *bool   `json:"DeletionIndicator"`
}
