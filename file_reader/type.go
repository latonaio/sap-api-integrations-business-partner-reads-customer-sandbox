package file_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo      string      `json:"document_no"`
		BusinessPartner string      `json:"deliver_to"`
		Quantity        float64     `json:"quantity"`
		PickedQuantity  float64     `json:"picked_quantity"`
		Price           float64     `json:"price"`
	    Batch           string      `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string      `json:"document_no"`
		Status               string      `json:"status"`
		DeliverTo            string      `json:"deliver_to"`
		Quantity             float64     `json:"quantity"`
		CompletedQuantity    float64     `json:"completed_quantity"`
	    PlannedStartDate     string      `json:"planned_start_date"`
	    PlannedValidatedDate string      `json:"planned_validated_date"`
	    ActualStartDate      string      `json:"actual_start_date"`
	    ActualValidatedDate  string      `json:"actual_validated_date"`
	    Batch                string      `json:"batch"`
		Work              struct {
			WorkNo                   string      `json:"work_no"`
			Quantity                 float64     `json:"quantity"`
			CompletedQuantity        float64     `json:"completed_quantity"`
			ErroredQuantity          float64     `json:"errored_quantity"`
			Component                string      `json:"component"`
			PlannedComponentQuantity float64     `json:"planned_component_quantity"`
			PlannedStartDate         string      `json:"planned_start_date"`
			PlannedStartTime         string      `json:"planned_start_time"`
			PlannedValidatedDate     string      `json:"planned_validated_date"`
			PlannedValidatedTime     string      `json:"planned_validated_time"`
			ActualStartDate          string      `json:"actual_start_date"`
			ActualStartTime          string      `json:"actual_start_time"`
			ActualValidatedDate      string      `json:"actual_validated_date"`
			ActualValidatedTime      string      `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema     string `json:"api_schema"`
	Material      string `json:"material_code"`
	Plant         string `json:"plant/supplier"`
	Stock         float64 `json:"stock"`
	DocumentType  string `json:"document_type"`
	DocumentNo    string `json:"document_no"`
	PlannedDate   string `json:"planned_date"`
	ValidatedDate string `json:"validated_date"`
	Deleted       string `json:"deleted"`
}

type SDC struct {
	ConnectionKey   string `json:"connection_key"`
	Result          bool   `json:"result"`
	RedisKey        string `json:"redis_key"`
	Filepath        string `json:"filepath"`
	BusinessPartner struct {
		BusinessPartner     string `json:"BusinessPartner"`
		BusinessPartnerRole string `json:"BusinessPartnerRole"`
		ValidFrom           string `json:"ValidFrom"`
		ValidTo             string `json:"ValidTo"`
		AddressID           struct {
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
		} `json:"AddressID"`
		SalesArea struct {
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
		} `json:"Sales_Area"`
		Company struct {
			CustomerDesc                   string `json:"Customer_desc"`
			CompanyCode                    string `json:"CompanyCode"`
			APARToleranceGroup             string `json:"APARToleranceGroup"`
			CustomerSupplierClearingIsUsed string `json:"CustomerSupplierClearingIsUsed"`
			HouseBank                      string `json:"HouseBank"`
			PaymentMethodsList             string `json:"PaymentMethodsList"`
			PaymentTerms                   string `json:"PaymentTerms"`
			ReconciliationAccount          string `json:"ReconciliationAccount"`
			DeletionIndicator              string `json:"DeletionIndicator"`
		} `json:"Company"`
	} `json:"business_partner"`
	APISchema           string `json:"api_schema"`
	BusinessPartner     string `json:"business_partner_code"`
	Deleted             string `json:"deleted"`
}