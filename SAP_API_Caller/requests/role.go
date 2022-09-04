package requests

type Role struct {
	BusinessPartner      string `json:"BusinessPartner"`
	BusinessPartnerRole  string `json:"BusinessPartnerRole"`
	ValidFrom           *string `json:"ValidFrom"`
	ValidTo             *string `json:"ValidTo"`
}
