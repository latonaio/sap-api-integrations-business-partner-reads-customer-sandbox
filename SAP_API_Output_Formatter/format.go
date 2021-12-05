package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-business-partner-reads-customer/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func ConvertToRole(raw []byte, l *logger.Logger) (*Role, error) {
	pm := &responses.Role{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Role. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &Role{
		BusinessPartner:     data.BusinessPartner,
		BusinessPartnerRole: data.BusinessPartnerRole,
		ValidFrom:           data.ValidFrom,
		ValidTo:             data.ValidTo,
	}, nil
}

func ConvertToAddress(raw []byte, l *logger.Logger) (*Address, error) {
	pm := &responses.Address{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Address. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &Address{
		BusinessPartner:   data.BusinessPartner,
		AddressID:         data.AddressID,
		ValidityStartDate: data.ValidityStartDate,
		ValidityEndDate:   data.ValidityEndDate,
		Country:           data.Country,
		Region:            data.Region,
		StreetName:        data.StreetName,
		CityName:          data.CityName,
		PostalCode:        data.PostalCode,
		Language:          data.Language,
	}, nil
}

func ConvertToSalesArea(raw []byte, l *logger.Logger) (*SalesArea, error) {
	pm := &responses.SalesArea{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to SalesArea. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &SalesArea{
		Customer:                       data.Customer,
		SalesOrganization:              data.SalesOrganization,
		DistributionChannel:            data.DistributionChannel,
		Division:                       data.Division,
		CompleteDeliveryIsDefined:      data.CompleteDeliveryIsDefined,
		Currency:                       data.Currency,
		CustomerAccountAssignmentGroup: data.CustomerAccountAssignmentGroup,
		CustomerPaymentTerms:           data.CustomerPaymentTerms,
		CustomerPriceGroup:             data.CustomerPriceGroup,
		CustomerPricingProcedure:       data.CustomerPricingProcedure,
		DeliveryPriority:               data.DeliveryPriority,
		IncotermsClassification:        data.IncotermsClassification,
		InvoiceDate:                    data.InvoiceDate,
		OrderCombinationIsAllowed:      data.OrderCombinationIsAllowed,
		PartialDeliveryIsAllowed:       data.PartialDeliveryIsAllowed,
		PriceListType:                  data.PriceListType,
		SalesGroup:                     data.SalesGroup,
		SalesOffice:                    data.SalesOffice,
		ShippingCondition:              data.ShippingCondition,
		SupplyingPlant:                 data.SupplyingPlant,
		SalesDistrict:                  data.SalesDistrict,
		InvoiceListSchedule:            data.InvoiceListSchedule,
		ExchangeRateType:               data.ExchangeRateType,
		OrderIsBlockedForCustomer:      data.OrderIsBlockedForCustomer,
		DeliveryIsBlockedForCustomer:   data.DeliveryIsBlockedForCustomer,
		BillingIsBlockedForCustomer:    data.BillingIsBlockedForCustomer,
		DeletionIndicator:              data.DeletionIndicator,
	}, nil
}

func ConvertToCompany(raw []byte, l *logger.Logger) (*Company, error) {
	pm := &responses.Company{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Company. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &Company{
		Customer:                       data.Customer,
		CompanyCode:                    data.CompanyCode,
		APARToleranceGroup:             data.APARToleranceGroup,
		CustomerSupplierClearingIsUsed: data.CustomerSupplierClearingIsUsed,
		HouseBank:                      data.HouseBank,
		PaymentMethodsList:             data.PaymentMethodsList,
		PaymentTerms:                   data.PaymentTerms,
		ReconciliationAccount:          data.ReconciliationAccount,
		DeletionIndicator:              data.DeletionIndicator,
	}, nil
}
