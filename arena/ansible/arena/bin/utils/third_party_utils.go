package utils

import (
	"fmt"
	y2 "gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// Config for handle the third party config
type ThirdPartyConfig struct {
	GoodData struct {
		Host  string `yaml:"host"`
		Login string `yaml:"login"`
	} `yaml:"good_data"`
	SalesForce struct {
		LeadUrl                string `yaml:"lead_url"`
		ArenaOrgId             string `yaml:"arena_org_id"`
		CustomerPortalLoginUrl string `yaml:"customer_portal_login_url"`
		LoginId                string `yaml:"login_id"`
		LoginPassword          string `yaml:"login_password"`
		WebToCaseOrigin        string `yaml:"web_to_case_origin"`
		WebToCaseUrl           string `yaml:"web_to_case_url"`
	} `yaml:"sales_force"`
}

func ReadThirdPartyServiceConfig(ansibleHome string) map[string]ThirdPartyConfig {
	res, err := ioutil.ReadFile(ansibleHome + "/arena/etc/third-party-service-config.yml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	data := make(map[string]ThirdPartyConfig)

	err = y2.Unmarshal(res, data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return data
}
