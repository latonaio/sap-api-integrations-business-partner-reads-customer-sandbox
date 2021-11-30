package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-business-partner-reads-customer/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
)

func ConvertToBPCustomer(raw []byte, l *logger.Logger) *BPCustomer {
	pm := &responses.BPCustomer{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		l.Error(err)
		return nil
	}
	if len(pm.D.Results) == 0 {
		l.Error("Result data is not exist.")
		return nil
	}
	if len(pm.D.Results) > 1 {
		l.Error("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &BPCustomer{
		BusinessPartner                data.BusinessPartner,
		BusinessPartnerRole            data.BusinessPartnerRole,
		ValidFrom                      data.ValidFrom,
		ValidTo                        data.ValidTo,
		AddressID                      data.AddressID,
		ValidityStartDate              data.ValidityStartDate,
		ValidityEndDate                data.ValidityEndDate,
		Country                        data.Country,
		Region                         data.Region,
		StreetName                     data.StreetName,
		CityName                       data.CityName,
		PostalCode                     data.PostalCode,
		Language                       data.Language,
		ToEmailAddress                 data.ToEmailAddress,
		ToFaxNumber                    data.ToFaxNumber,
		ToMobilePhoneNumber            data.ToMobilePhoneNumber,
		ToPhoneNumber                  data.ToPhoneNumber,
		ToURLAddress                   data.ToURLAddress,
		CustomerDesc                   data.CustomerDesc,
		SalesOrganization              data.SalesOrganization,
		DistributionChannel            data.DistributionChannel,
		Division                       data.Division,
		CompleteDeliveryIsDefined      data.CompleteDeliveryIsDefined,
		Currency                       data.Currency,
		CustomerAccountAssignmentGroup data.CustomerAccountAssignmentGroup,
		CustomerPaymentTerms           data.CustomerPaymentTerms,
		CustomerPriceGroup             data.CustomerPriceGroup,
		CustomerPricingProcedure       data.CustomerPricingProcedure,
		DeliveryPriority               data.DeliveryPriority,
		IncotermsClassification        data.IncotermsClassification,
		InvoiceDate                    data.InvoiceDate,
		OrderCombinationIsAllowed      data.OrderCombinationIsAllowed,
		PartialDeliveryIsAllowed       data.PartialDeliveryIsAllowed,
		PriceListType                  data.PriceListType,
		SalesGroup                     data.SalesGroup,
		SalesOffice                    data.SalesOffice,
		ShippingCondition              data.ShippingCondition,
		SupplyingPlant                 data.SupplyingPlant,
		SalesDistrict                  data.SalesDistrict,
		InvoiceListSchedule            data.InvoiceListSchedule,
		ExchangeRateType               data.ExchangeRateType,
		OrderIsBlockedForCustomer      data.OrderIsBlockedForCustomer,
		DeliveryIsBlockedForCustomer   data.DeliveryIsBlockedForCustomer,
		BillingIsBlockedForCustomer    data.BillingIsBlockedForCustomer,
		DeletionIndicator              data.DeletionIndicator,
		CustomerDesc                   data.CustomerDesc,
		CompanyCode                    data.CompanyCode,
		APARToleranceGroup             data.APARToleranceGroup,
		CustomerSupplierClearingIsUsed data.CustomerSupplierClearingIsUsed,
		HouseBank                      data.HouseBank,
		PaymentMethodsList             data.PaymentMethodsList,
		PaymentTerms                   data.PaymentTerms,
		ReconciliationAccount          data.ReconciliationAccount,
		DeletionIndicator              data.DeletionIndicator,
	}
}
