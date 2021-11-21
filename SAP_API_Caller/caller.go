package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetBusinessPartner(BusinessPartner, BusinessPartnerRole, VaridityEndDate, SalesOrganization, DistributionChannel, Division, CompanyCode string) {
	wg := &sync.WaitGroup{}

	wg.Add(4)
	go func() {
		c.BusinessPartnerRole(BusinessPartner, BusinessPartnerRole)
		wg.Done()
	}()
	go func() {
		c.BusinessPartnerAddress(BusinessPartner, VaridityEndDate)
		wg.Done()
	}()
	go func() {
		c.BusinessPartnerSalesArea(BusinessPartner, SalesOrganization, DistributionChannel, Division)
		wg.Done()
	}()
	go func() {
		c.BusinessPartnerCompany(BusinessPartner, CompanyCode)
		wg.Done()
	}()
	wg.Wait()
}

func (c *SAPAPICaller) BusinessPartnerRole(BusinessPartner, BusinessPartnerRole string) {
	res, err := c.callBusinessPartnerSrvAPIRequirementPartnerRole("A_BusinessPartner('{BusinessPartner}')/to_BusinessPartnerRole", BusinessPartner, BusinessPartnerRole)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) BusinessPartnerAddress(BusinessPartner, VaridityEndDate string) {
	res, err := c.callBusinessPartnerSrvAPIRequirementAddress("A_BusinessPartner('{BusinessPartner}')/to_BusinessPartnerAddress", BusinessPartner, VaridityEndDate)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) BusinessPartnerSalesArea(BusinessPartner, SalesOrganization, DistributionChannel, Division string) {
	res, err := c.callBusinessPartnerSrvAPIRequirementSalesArea("A_Customer('{Customer}')/to_CustomerSalesArea", BusinessPartner, SalesOrganization, DistributionChannel, Division)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) BusinessPartnerCompany(BusinessPartner, CompanyCode string) {
	res, err := c.callBusinessPartnerSrvAPIRequirementCompany("A_Customer('{Customer}')/to_CustomerCompany", BusinessPartner, CompanyCode)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) callBusinessPartnerSrvAPIRequirement(api, BusinessPartner, BusinessPartnerRole, VaridityEndDate, SalesOrganization, DistributionChannel, Division, CompanyCode string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "BusinessPartner, BusinessPartnerRole, VaridityEndDate, SalesOrganization, DistributionChannel, Division, CompanyCode")
	params.Add("$filter", fmt.Sprintf("BusinessPartner eq '%s' and BusinessPartnerRole eq '%s' and VaridityEndDate eq '%s' and SalesOrganization eq '%s' and DistributionChannel eq '%s' and Division eq '%s' and CompanyCode eq '%s'", BusinessPartner, BusinessPartnerRole, VaridityEndDate, SalesOrganization, DistributionChannel, Division, CompanyCode))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}