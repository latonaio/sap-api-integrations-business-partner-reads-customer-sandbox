package sap_api_input_reader

import (
	"sap-api-integrations-business-partner-creates-customer/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToGeneral() *requests.General {
	data := sdc.BusinessPartner
	return &requests.General{
		BusinessPartner: data.BusinessPartner,
		//	Customer:                       data.Customer,
		//	Supplier:                       data.Supplier,
		AcademicTitle:           data.AcademicTitle,
		AuthorizationGroup:      data.AuthorizationGroup,
		BusinessPartnerCategory: data.BusinessPartnerCategory,
		//	BusinessPartnerFullName:        data.BusinessPartnerFullName,
		BusinessPartnerGrouping: data.BusinessPartnerGrouping,
		//	BusinessPartnerName:            data.BusinessPartnerName,
		CorrespondenceLanguage: data.CorrespondenceLanguage,
		//	CreationDate:                   data.CreationDate,
		//	CreationTime:                   data.CreationTime,
		FirstName:       data.FirstName,
		Industry:        data.Industry,
		IsFemale:        data.IsFemale,
		IsMale:          data.IsMale,
		IsNaturalPerson: data.IsNaturalPerson,
		IsSexUnknown:    data.IsSexUnknown,
		GenderCodeName:  data.GenderCodeName,
		Language:        data.Language,
		//	LastChangeDate:                 data.LastChangeDate,
		//	LastChangeTime:                 data.LastChangeTime,
		LastName:                      data.LastName,
		OrganizationBPName1:           data.OrganizationBPName1,
		OrganizationBPName2:           data.OrganizationBPName2,
		OrganizationBPName3:           data.OrganizationBPName3,
		OrganizationBPName4:           data.OrganizationBPName4,
		OrganizationFoundationDate:    data.OrganizationFoundationDate,
		OrganizationLiquidationDate:   data.OrganizationLiquidationDate,
		SearchTerm1:                   data.SearchTerm1,
		SearchTerm2:                   data.SearchTerm2,
		AdditionalLastName:            data.AdditionalLastName,
		BirthDate:                     data.BirthDate,
		BusinessPartnerBirthplaceName: data.BusinessPartnerBirthplaceName,
		// BusinessPartnerDeathDate:       data.BusinessPartnerDeathDate,
		BusinessPartnerIsBlocked:  data.BusinessPartnerIsBlocked,
		BusinessPartnerType:       data.BusinessPartnerType,
		GroupBusinessPartnerName1: data.GroupBusinessPartnerName1,
		GroupBusinessPartnerName2: data.GroupBusinessPartnerName2,
		//	IndependentAddressID:           data.IndependentAddressID,
		MiddleName:                   data.MiddleName,
		NameCountry:                  data.NameCountry,
		PersonFullName:               data.PersonFullName,
		PersonNumber:                 data.PersonNumber,
		IsMarkedForArchiving:         data.IsMarkedForArchiving,
		BusinessPartnerIDByExtSystem: data.BusinessPartnerIDByExtSystem,
		TradingPartner:               data.TradingPartner,
	}
}

func (sdc *SDC) ConvertToRole() *requests.Role {
	dataBusinessPartner := sdc.BusinessPartner
	data := sdc.BusinessPartner.Role
	return &requests.Role{
		BusinessPartner:     dataBusinessPartner.BusinessPartner,
		BusinessPartnerRole: data.BusinessPartnerRole,
		ValidFrom:           data.ValidFrom,
		ValidTo:             data.ValidTo,
	}
}

func (sdc *SDC) ConvertToAddress() *requests.Address {
	dataBusinessPartner := sdc.BusinessPartner
	data := sdc.BusinessPartner.Address
	return &requests.Address{
		BusinessPartner:   dataBusinessPartner.BusinessPartner,
		AddressID:         data.AddressID,
		ValidityStartDate: data.ValidityStartDate,
		ValidityEndDate:   data.ValidityEndDate,
		Country:           data.Country,
		Region:            data.Region,
		StreetName:        data.StreetName,
		CityName:          data.CityName,
		PostalCode:        data.PostalCode,
		Language:          data.Language,
	}
}

func (sdc *SDC) ConvertToBank() *requests.Bank {
	dataBusinessPartner := sdc.BusinessPartner
	data := sdc.BusinessPartner.Bank
	return &requests.Bank{
		BusinessPartner:          dataBusinessPartner.BusinessPartner,
		BankIdentification:       data.BankIdentification,
		BankCountryKey:           data.BankCountryKey,
		BankName:                 data.BankName,
		BankNumber:               data.BankNumber,
		SWIFTCode:                data.SWIFTCode,
		BankControlKey:           data.BankControlKey,
		BankAccountHolderName:    data.BankAccountHolderName,
		BankAccountName:          data.BankAccountName,
		ValidityStartDate:        data.ValidityStartDate,
		ValidityEndDate:          data.ValidityEndDate,
		IBAN:                     data.IBAN,
		IBANValidityStartDate:    data.IBANValidityStartDate,
		BankAccount:              data.BankAccount,
		BankAccountReferenceText: data.BankAccountReferenceText,
		CollectionAuthInd:        data.CollectionAuthInd,
		CityName:                 data.CityName,
		AuthorizationGroup:       data.AuthorizationGroup,
	}
}

func (sdc *SDC) ConvertToCustomer() *requests.Customer {
	data := sdc.BusinessPartner.CustomerData
	return &requests.Customer{
		Customer:                    data.Customer,
		AuthorizationGroup:          data.AuthorizationGroup,
		BillingIsBlockedForCustomer: data.BillingIsBlockedForCustomer,
		CreationDate:                data.CreationDate,
		CustomerAccountGroup:        data.CustomerAccountGroup,
		CustomerClassification:      data.CustomerClassification,
		CustomerFullName:            data.CustomerFullName,
		CustomerName:                data.CustomerName,
		DeliveryIsBlocked:           data.DeliveryIsBlocked,
		OrderIsBlockedForCustomer:   data.OrderIsBlockedForCustomer,
		PostingIsBlocked:            data.PostingIsBlocked,
		Supplier:                    data.Supplier,
		CustomerCorporateGroup:      data.CustomerCorporateGroup,
		Industry:                    data.Industry,
		TaxNumber1:                  data.TaxNumber1,
		DeletionIndicator:           data.DeletionIndicator,
		CityCode:                    data.CityCode,
		County:                      data.County,
	}
}

func (sdc *SDC) ConvertToSalesArea() *requests.SalesArea {
	dataCustomer := sdc.BusinessPartner.CustomerData
	data := sdc.BusinessPartner.CustomerData.SalesArea
	return &requests.SalesArea{
		Customer:                       dataCustomer.Customer,
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
	}
}

func (sdc *SDC) ConvertToPartnerFunction() *requests.PartnerFunction {
	dataPartnerFunction := sdc.BusinessPartner.CustomerData
	data := sdc.BusinessPartner.CustomerData.SalesArea.PartnerFunction
	return &requests.PartnerFunction{
		Customer:                   dataPartnerFunction.Customer,
		SalesOrganization:          data.SalesOrganization,
		DistributionChannel:        data.DistributionChannel,
		Division:                   data.Division,
		PartnerCounter:             data.PartnerCounter,
		PartnerFunction:            data.PartnerFunction,
		BPCustomerNumber:           data.BPCustomerNumber,
		CustomerPartnerDescription: data.CustomerPartnerDescription,
		DefaultPartner:             data.DefaultPartner,
		Supplier:                   data.Supplier,
		AuthorizationGroup:         data.AuthorizationGroup,
	}
}

func (sdc *SDC) ConvertToCompany() *requests.Company {
	dataCustomer := sdc.BusinessPartner.CustomerData
	data := sdc.BusinessPartner.CustomerData.Company
	return &requests.Company{
		Customer:                       dataCustomer.Customer,
		CompanyCode:                    data.CompanyCode,
		APARToleranceGroup:             data.APARToleranceGroup,
		CustomerSupplierClearingIsUsed: data.CustomerSupplierClearingIsUsed,
		HouseBank:                      data.HouseBank,
		PaymentMethodsList:             data.PaymentMethodsList,
		PaymentTerms:                   data.PaymentTerms,
		ReconciliationAccount:          data.ReconciliationAccount,
		DeletionIndicator:              data.DeletionIndicator,
	}
}
