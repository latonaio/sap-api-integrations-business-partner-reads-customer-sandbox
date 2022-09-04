package responses

type PartnerFunction struct {
	D struct {
		Customer                   string `json:"Customer"`
		SalesOrganization          string `json:"SalesOrganization"`
		DistributionChannel        string `json:"DistributionChannel"`
		Division                   string `json:"Division"`
		PartnerCounter             string `json:"PartnerCounter"`
		PartnerFunction            string `json:"PartnerFunction"`
		BPCustomerNumber           string `json:"BPCustomerNumber"`
		CustomerPartnerDescription string `json:"CustomerPartnerDescription"`
		DefaultPartner             bool   `json:"DefaultPartner"`
		Supplier                   string `json:"Supplier"`
		AuthorizationGroup         string `json:"AuthorizationGroup"`
	} `json:"d"`
}
