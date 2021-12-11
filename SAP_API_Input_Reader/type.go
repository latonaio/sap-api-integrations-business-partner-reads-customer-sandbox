package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo      string `json:"document_no"`
		BusinessPartner string `json:"deliver_to"`
		Quantity        string `json:"quantity"`
		PickedQuantity  string `json:"picked_quantity"`
		Price           string `json:"price"`
		Batch           string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema     string `json:"api_schema"`
	MaterialCode  string `json:"material_code"`
	Plant         string `json:"plant/supplier"`
	Stock         string `json:"stock"`
	DocumentType  string `json:"document_type"`
	DocumentNo    string `json:"document_no"`
	PlannedDate   string `json:"planned_date"`
	ValidatedDate string `json:"validated_date"`
	Deleted       bool   `json:"deleted"`
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
		Address             struct {
			AddressID         string `json:"AddressID"`
			ValidityStartDate string `json:"ValidityStartDate"`
			ValidityEndDate   string `json:"ValidityEndDate"`
			Country           string `json:"Country"`
			Region            string `json:"Region"`
			StreetName        string `json:"StreetName"`
			CityName          string `json:"CityName"`
			PostalCode        string `json:"PostalCode"`
			Language          string `json:"Language"`
		} `json:"AddressID"`
		SalesArea struct {
			SalesOrganization              string `json:"SalesOrganization"`
			DistributionChannel            string `json:"DistributionChannel"`
			Division                       string `json:"Division"`
			CompleteDeliveryIsDefined      bool   `json:"CompleteDeliveryIsDefined"`
			Currency                       string `json:"Currency"`
			CustomerAccountAssignmentGroup string `json:"CustomerAccountAssignmentGroup"`
			CustomerPaymentTerms           string `json:"CustomerPaymentTerms"`
			CustomerPriceGroup             string `json:"CustomerPriceGroup"`
			CustomerPricingProcedure       string `json:"CustomerPricingProcedure"`
			DeliveryPriority               string `json:"DeliveryPriority"`
			IncotermsClassification        string `json:"IncotermsClassification"`
			InvoiceDate                    string `json:"InvoiceDate"`
			OrderCombinationIsAllowed      bool   `json:"OrderCombinationIsAllowed"`
			PartialDeliveryIsAllowed       string `json:"PartialDeliveryIsAllowed"`
			PriceListType                  string `json:"PriceListType"`
			SalesGroup                     string `json:"SalesGroup"`
			SalesOffice                    string `json:"SalesOffice"`
			ShippingCondition              string `json:"ShippingCondition"`
			SupplyingPlant                 string `json:"SupplyingPlant"`
			SalesDistrict                  string `json:"SalesDistrict"`
			InvoiceListSchedule            string `json:"InvoiceListSchedule"`
			ExchangeRateType               string `json:"ExchangeRateType"`
			OrderIsBlockedForCustomer      string `json:"OrderIsBlockedForCustomer"`
			DeliveryIsBlockedForCustomer   string `json:"DeliveryIsBlockedForCustomer"`
			BillingIsBlockedForCustomer    string `json:"BillingIsBlockedForCustomer"`
			DeletionIndicator              bool   `json:"DeletionIndicator"`
		} `json:"Sales_Area"`
		Company struct {
			CompanyCode                    string `json:"CompanyCode"`
			APARToleranceGroup             string `json:"APARToleranceGroup"`
			CustomerSupplierClearingIsUsed string `json:"CustomerSupplierClearingIsUsed"`
			HouseBank                      string `json:"HouseBank"`
			PaymentMethodsList             string `json:"PaymentMethodsList"`
			PaymentTerms                   string `json:"PaymentTerms"`
			ReconciliationAccount          string `json:"ReconciliationAccount"`
			DeletionIndicator              bool   `json:"DeletionIndicator"`
		} `json:"Company"`
	} `json:"business_partner"`
	APISchema string   `json:"api_schema"`
	Accepter  []string `json:"accepter"`
	Customer  string   `json:"business_partner_code"`
	Deleted   bool     `json:"deleted"`
}
