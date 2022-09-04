package config

type SAP struct {
	baseURL         string
	user            string
	pass            string
	tokenRefreshURL string
	retryMaxCnt     int
	retryInterval   int
}

func newSAP() *SAP {
	return &SAP{
		baseURL:         getEnv("SAP_API_BASE_URL", "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/"),
		user:            getEnv("SAP_USER", ""),
		pass:            getEnv("SAP_PASS", ""),
		tokenRefreshURL: getEnv("SAP_TOKEN_REFRESH_TOKEN", ""),
		retryMaxCnt:     getEnvInt("SAP_RETRY_MAX_COUNT", 2),
		retryInterval:   getEnvInt("SAP_RETRY_INTERVAL", 1000),
	}
}

func (c *SAP) User() string {
	return c.user
}

func (c *SAP) Pass() string {
	return c.pass
}

func (c *SAP) RefreshTokenURL() string {
	return c.tokenRefreshURL
}

func (c *SAP) RetryMax() int {
	return c.retryMaxCnt
}

func (c *SAP) RetryInterval() int {
	return c.retryInterval
}

func (c *SAP) BaseURL() string {
	return c.baseURL
}
