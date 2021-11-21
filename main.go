package main

import (
	sap_api_caller "sap-api-integrations-business-partner-reads-customer/SAP_API_Caller"
	"sap-api-integrations-business-partner-reads-customer/file_reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := file_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs/SDC_Business_Partner_Customer.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	caller.AsyncGetBusinessPartner(
		inoutSDC.BusinessPartner.BusinessPartnerRole,
		inoutSDC.BusinessPartner.VaridityEndDate,
		inoutSDC.BusinessPartner.SalesOrganization.DistributionChannel.Division,
		inoutSDC.BusinessPartner.CompanyCode,
	)
}
