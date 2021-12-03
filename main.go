package main

import (
	sap_api_caller "sap-api-integrations-business-partner-reads-customer/SAP_API_Caller"
	"sap-api-integrations-business-partner-reads-customer/sap_api_input_reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs//SDC_Business_Partner_Customer_sample1.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	caller.AsyncGetBPCustomer(
		inoutSDC.BusinessPartner.BusinessPartner,
		inoutSDC.BusinessPartner.BusinessPartnerRole,
		inoutSDC.BusinessPartner.Address.AddressID,
		inoutSDC.BusinessPartner.SalesArea.SalesOrganization,
		inoutSDC.BusinessPartner.SalesArea.DistributionChannel,
		inoutSDC.BusinessPartner.SalesArea.Division,
		inoutSDC.Customer,
		inoutSDC.BusinessPartner.Company.CompanyCode,
	)
}
