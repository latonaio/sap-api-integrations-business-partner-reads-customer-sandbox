package sap_api_caller

type BusinessPartnerReads struct {
	 ConnectionKey   string `json:"connection_key"`
	 Result          bool   `json:"result"`
	 RedisKey        string `json:"redis_key"`
	 Filepath        string `json:"filepath"`
	 APISchema       string `json:"api_schema"`
	 BusinessPartner string `json:"business_partner_code"`
	 Deleted         string `json:"deleted"`
}

type BusinessPartnerRole struct {
     BusinessPartner     string `json:"BusinessPartner"`
     BusinessPartnerRole string `json:"BusinessPartnerRole"`
     ValidFrom           string `json:"ValidFrom"`
     ValidTo             string `json:"ValidTo"`
}

type BusinessPartnerAddressID struct {
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
}

type BusinessPartnerSalesArea struct {
     BusinessPartner                string   `json:"BusinessPartner"`
     CustomerDesc                   string   `json:"Customer_desc"`
     SalesOrganization              string   `json:"SalesOrganization"`
     DistributionChannel            string   `json:"DistributionChannel"`
     Division                       string   `json:"Division"`
     CompleteDeliveryIsDefined      string   `json:"CompleteDeliveryIsDefined"`
     Currency                       string   `json:"Currency"`
     CustomerAccountAssignmentGroup string   `json:"CustomerAccountAssignmentGroup"`
     CustomerPaymentTerms           string   `json:"CustomerPaymentTerms"`
     CustomerPriceGroup             string   `json:"CustomerPriceGroup"`
     CustomerPricingProcedure       string   `json:"CustomerPricingProcedure"`
     DeliveryPriority               string   `json:"DeliveryPriority"`
     IncotermsClassification        string   `json:"IncotermsClassification"`
     InvoiceDate                    string   `json:"InvoiceDate"`
     OrderCombinationIsAllowed      string   `json:"OrderCombinationIsAllowed"`
     PartialDeliveryIsAllowed       string   `json:"PartialDeliveryIsAllowed"`
     PriceListType                  string   `json:"PriceListType"`
     SalesGroup                     string   `json:"SalesGroup"`
     SalesOffice                    string   `json:"SalesOffice"`
     ShippingCondition              string   `json:"ShippingCondition"`
     SupplyingPlant                 string   `json:"SupplyingPlant"`
     SalesDistrict                  string   `json:"SalesDistrict"`
     InvoiceListSchedule            string   `json:"InvoiceListSchedule"`
     ExchangeRateType               string   `json:"ExchangeRateType"`
     OrderIsBlockedForCustomer      string   `json:"OrderIsBlockedForCustomer"`
     DeliveryIsBlockedForCustomer   string   `json:"DeliveryIsBlockedForCustomer"`
     BillingIsBlockedForCustomer    string   `json:"BillingIsBlockedForCustomer"`
     DeletionIndicator              string   `json:"DeletionIndicator"`
}

type BusinessPartnerCustomerCompany struct {
     BusinessPartner                string   `json:"BusinessPartner"`
     CustomerDesc                   string   `json:"Customer_desc"`
     CompanyCode                    string   `json:"CompanyCode"`
     APARToleranceGroup             string   `json:"APARToleranceGroup"`
     CustomerSupplierClearingIsUsed string   `json:"CustomerSupplierClearingIsUsed"`
     HouseBank                      string   `json:"HouseBank"`
     PaymentMethodsList             string   `json:"PaymentMethodsList"`
     PaymentTerms                   string   `json:"PaymentTerms"`
     ReconciliationAccount          string   `json:"ReconciliationAccount"`
     DeletionIndicator              string   `json:"DeletionIndicator"`
}
