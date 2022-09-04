package main

import (
	sap_api_caller "sap-api-integrations-business-partner-creates-customer/SAP_API_Caller"
	sap_api_input_reader "sap-api-integrations-business-partner-creates-customer/SAP_API_Input_Reader"
	"sap-api-integrations-business-partner-creates-customer/config"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	sap_api_post_header_setup "github.com/latonaio/sap-api-post-header-setup"
)

func main() {
	l := logger.NewLogger()
	conf := config.NewConf()
	fr := sap_api_input_reader.NewFileReader()
	pc := sap_api_post_header_setup.NewSAPPostClientWithOption(conf.SAP)
	caller := sap_api_caller.NewSAPAPICaller(
		conf.SAP.BaseURL(),
		"100",
		pc,
		l,
	)
	inputSDC := fr.ReadSDC("./Inputs/SDC_Business_Partner_Customer_General_sample.json")
	accepter := getAccepter(inputSDC)
	general := inputSDC.ConvertToGeneral()
	role := inputSDC.ConvertToRole()
	address := inputSDC.ConvertToAddress()
	bank := inputSDC.ConvertToBank()
	customer := inputSDC.ConvertToCustomer()
	salesArea := inputSDC.ConvertToSalesArea()
	partnerFunction := inputSDC.ConvertToPartnerFunction()
	company := inputSDC.ConvertToCompany()

	caller.AsyncPostBPC(
		general,
		role,
		address,
		bank,
		customer,
		salesArea,
		partnerFunction,
		company,
		accepter,
	)
}

func getAccepter(sdc sap_api_input_reader.SDC) []string {
	accepter := sdc.Accepter
	if len(sdc.Accepter) == 0 {
		accepter = []string{"All"}
	}

	if accepter[0] == "All" {
		accepter = []string{
			"General", "Role", "Address", "Bank", "Customer", "SalesArea", "PartnerFunction", "Company",
		}
	}
	return accepter
}
