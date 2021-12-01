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

func (c *SAPAPICaller) AsyncGetBPCustomer(BusinessPartner, BusinessPartnerRole, VaridityEndDate, SalesOrganization, DistributionChannel, Division, CompanyCode string) {
	wg := &sync.WaitGroup{}

	wg.Add(4)
	go func() {
		c.Role(BusinessPartner, BusinessPartnerRole)
		wg.Done()
	}()
	go func() {
		c.Address(BusinessPartner, VaridityEndDate)
		wg.Done()
	}()
	go func() {
		c.SalesArea(BusinessPartner, SalesOrganization, DistributionChannel, Division)
		wg.Done()
	}()
	go func() {
		c.Company(BusinessPartner, CompanyCode)
		wg.Done()
	}()
	wg.Wait()
}

func (c *SAPAPICaller) Role(BusinessPartner, BusinessPartnerRole string) {
	res, err := c.callBusinessPartnerSrvAPIRequirementRole("A_BusinessPartner('{BusinessPartner}')/to_BusinessPartnerRole", BusinessPartner, BusinessPartnerRole)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) Address(BusinessPartner, VaridityEndDate string) {
	res, err := c.callBusinessPartnerSrvAPIRequirementAddress("A_BusinessPartner('{BusinessPartner}')/to_BusinessPartnerAddress", BusinessPartner, VaridityEndDate)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) SalesArea(BusinessPartner, SalesOrganization, DistributionChannel, Division string) {
	res, err := c.callBusinessPartnerSrvAPIRequirementSalesArea("A_Customer('{Customer}')/to_CustomerSalesArea", BusinessPartner, SalesOrganization, DistributionChannel, Division)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) Company(BusinessPartner, CompanyCode string) {
	res, err := c.callBusinessPartnerSrvAPIRequirementCompany("A_Customer('{Customer}')/to_CustomerCompany", BusinessPartner, CompanyCode)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) callBusinessPartnerSrvAPIRequirementRole(api, BusinessPartner, BusinessPartnerRole, string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "BusinessPartner, BusinessPartnerRole")
	params.Add("$filter", fmt.Sprintf("BusinessPartner eq '%s' and BusinessPartnerRole eq '%s'", BusinessPartner, BusinessPartnerRole))
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

func (c *SAPAPICaller) callBusinessPartnerSrvAPIRequirementAddress(api, BusinessPartner, VaridityEndDate string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "BusinessPartner, VaridityEndDate")
	params.Add("$filter", fmt.Sprintf("BusinessPartner eq '%s' and VaridityEndDate eq '%s'", BusinessPartner, VaridityEndDate))
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

func (c *SAPAPICaller) callBusinessPartnerSrvAPIRequirementSalesArea(api, BusinessPartner, SalesOrganization, DistributionChannel, Division string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "BusinessPartner, SalesOrganization, DistributionChannel, Division")
	params.Add("$filter", fmt.Sprintf("BusinessPartner eq '%s' and SalesOrganization eq '%s' and DistributionChannel eq '%s' and Division eq '%s'", BusinessPartner, SalesOrganization, DistributionChannel, Division))
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

func (c *SAPAPICaller) callBusinessPartnerSrvAPIRequirementCompany(api, BusinessPartner, CompanyCode string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "BusinessPartner, CompanyCode")
	params.Add("$filter", fmt.Sprintf("BusinessPartner eq '%s' and CompanyCode eq '%s'", BusinessPartner, CompanyCode))
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
