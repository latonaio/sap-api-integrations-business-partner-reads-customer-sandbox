package sap_api_caller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sap-api-integrations-business-partner-creates-customer/SAP_API_Caller/requests"
	sap_api_output_formatter "sap-api-integrations-business-partner-creates-customer/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	sap_api_post_header_setup "github.com/latonaio/sap-api-post-header-setup"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	postClient      *sap_api_post_header_setup.SAPPostClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, postClient *sap_api_post_header_setup.SAPPostClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		postClient:      postClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncPostBPC(
	general *requests.General,
	role *requests.Role,
	address *requests.Address,
	bank *requests.Bank,
	customer *requests.Customer,
	salesArea *requests.SalesArea,
	partnerFunction *requests.PartnerFunction,
	company *requests.Company,
	accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "General":
			func() {
				c.General(general)
				wg.Done()
			}()
		case "Role":
			func() {
				c.Role(role)
				wg.Done()
			}()
		case "Bank":
			func() {
				c.Bank(bank)
				wg.Done()
			}()
		case "Address":
			func() {
				c.Address(address)
				wg.Done()
			}()
		case "Customer":
			func() {
				c.Customer(customer)
				wg.Done()
			}()
		case "SalesArea":
			func() {
				c.SalesArea(salesArea)
				wg.Done()
			}()
		case "PartnerFunction":
			func() {
				c.PartnerFunction(partnerFunction)
				wg.Done()
			}()
		case "Company":
			func() {
				c.Company(company)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) General(general *requests.General) {
	outputDataGeneral, err := c.callBPSrvAPIRequirementGeneral("A_BusinessPartner", general)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataGeneral)
}

func (c *SAPAPICaller) callBPSrvAPIRequirementGeneral(api string, general *requests.General) (*sap_api_output_formatter.General, error) {
	body, err := json.Marshal(general)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.postClient.POST(url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToGeneral(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Role(role *requests.Role) {
	outputDataRole, err := c.callBPSrvAPIRequirementRole("A_BusinessPartnerRole", role)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataRole)
}

func (c *SAPAPICaller) callBPSrvAPIRequirementRole(api string, role *requests.Role) (*sap_api_output_formatter.Role, error) {
	body, err := json.Marshal(role)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.postClient.POST(url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToRole(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Address(address *requests.Address) {
	outputDataAddress, err := c.callBPSrvAPIRequirementAddress("A_BusinessPartnerAddress", address)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataAddress)
}

func (c *SAPAPICaller) callBPSrvAPIRequirementAddress(api string, address *requests.Address) (*sap_api_output_formatter.Address, error) {
	body, err := json.Marshal(address)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.postClient.POST(url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToAddress(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Bank(bank *requests.Bank) {
	outputDataBank, err := c.callBPSrvAPIRequirementBank("A_BusinessPartnerBank", bank)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataBank)
}

func (c *SAPAPICaller) callBPSrvAPIRequirementBank(api string, bank *requests.Bank) (*sap_api_output_formatter.Bank, error) {
	body, err := json.Marshal(bank)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.postClient.POST(url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToBank(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Customer(customer *requests.Customer) {
	outputDataCustomer, err := c.callBPSrvAPIRequirementCustomer("A_Customer", customer)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataCustomer)
}

func (c *SAPAPICaller) callBPSrvAPIRequirementCustomer(api string, customer *requests.Customer) (*sap_api_output_formatter.Customer, error) {
	body, err := json.Marshal(customer)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.postClient.POST(url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToCustomer(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) SalesArea(salesArea *requests.SalesArea) {
	outputDataSalesArea, err := c.callBPSrvAPIRequirementSalesArea("A_CustomerSalesArea", salesArea)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataSalesArea)
}

func (c *SAPAPICaller) callBPSrvAPIRequirementSalesArea(api string, salesArea *requests.SalesArea) (*sap_api_output_formatter.SalesArea, error) {
	body, err := json.Marshal(salesArea)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.postClient.POST(url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToSalesArea(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) PartnerFunction(partnerFunction *requests.PartnerFunction) {
	outputDataPartnerFunction, err := c.callBPSrvAPIRequirementPartnerFunction("A_CustSalesPartnerFunc", partnerFunction)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataPartnerFunction)
}

func (c *SAPAPICaller) callBPSrvAPIRequirementPartnerFunction(api string, partnerFunction *requests.PartnerFunction) (*sap_api_output_formatter.PartnerFunction, error) {
	body, err := json.Marshal(partnerFunction)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.postClient.POST(url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToPartnerFunction(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Company(company *requests.Company) {
	outputDataCompany, err := c.callBPSrvAPIRequirementCompany("A_CustomerCompany", company)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataCompany)
}

func (c *SAPAPICaller) callBPSrvAPIRequirementCompany(api string, company *requests.Company) (*sap_api_output_formatter.Company, error) {
	body, err := json.Marshal(company)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.postClient.POST(url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToCompany(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) addQuerySAPClient(params map[string]string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["sap-client"] = c.sapClientNumber
	return params
}
