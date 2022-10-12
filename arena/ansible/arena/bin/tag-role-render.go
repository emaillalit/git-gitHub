package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"ptc.com/tag-role-render/utils"
	"strings"
	"text/template"
)

// Info struct is to render in the templates
type Info struct {
	FullEnv                          string
	Env                              string
	IsProd                           string
	AwsAccount                       string
	Region                           string
	BaseDomain                       string
	TldDomain                        string
	LdapDomain                       string
	LdapBaseDn                       string
	LdapAccessGroupPrefix            string
	CIDR                             string
	DbName                           string
	CheckIfProduction                string
	DeploymentInfrastructure         string
	EmailSubjectPrepend              string
	EmailRcptDomainsAllowed          string
	EmailSenderEmployeesAllowed      string
	ApplicationMonitoringEmail       string
	DatabaseMonitoringEmail          string
	EsCluster                        string
	FeatureTfaDisable                string
	FeedbackEmail                    string
	LoggingHost                      string
	MarketingSite                    string
	MarketingSitePathPrefix          string
	SyslogEnable                     string
	DsmPolicyIdWebServersAutoScale   string
	DsmPolicyIdWebServersOther       string
	DsmPolicyIdOther                 string
	DsmGroupIdWebServersAutoScale    string
	DsmGroupIdWebServersOther        string
	DsmGroupIdOther                  string
	EnableLanguageSetting            string
	SalesforceLeadURL                string
	SalesforceArenaOrgId             string
	SalesforceCustomerPortalLoginUrl string
	SalesforceLoginId                string
	SalesforceLoginPassword          string
	SalesForceWebToCaseOrigin        string
	SalesForceWebToCaseUrl           string
	HttpHeadersSourceIp              string
	GoodDataHost                     string
	GoodDataLogin                    string
	GoodDataPassword                 string
	GoodDataPgpPassword              string
}

func main() {
	// define some variables to be parsed from flag
	var (
		fullEnv                        string
		env                            string
		awsAccount                     string
		region                         string
		baseDomain                     string
		cidr                           string
		dbName                         string
		isProd                         string
		syslogEnable                   string
		dsmPolicyIdWebServersAutoScale string
		dsmPolicyIdWebServersOther     string
		dsmPolicyIdOther               string
		dsmGroupIdWebServersAutoScale  string
		dsmGroupIdWebServersOther      string
		dsmGroupIdOther                string
	)
	// get the ansible home variable from the environment
	ansibleHome := os.Getenv("ANSIBLE_HOME")
	if ansibleHome == "" {
		fmt.Println("Please make sure environment variable ANSIBLE_HOME has been set.")
		os.Exit(1)
	}

	// define the flags
	flag.StringVar(&fullEnv, "full-env", "", "Specify the full env like awsqa, awssand.")
	flag.StringVar(&awsAccount, "aws-account", "", "Specify the aws account like 089875288533.")
	flag.StringVar(&region, "region", "", "Specify the aws region like us-east-2, us-gov-east-1.")
	flag.StringVar(&baseDomain, "base-domain", "", "Specify the base domain like arenagov.com, qa.aws.bom.com, sand.aws.bom.com.")
	flag.StringVar(&cidr, "cidr", "", "Specify the CIDR network range for the subnets.")
	flag.StringVar(&dbName, "db-name", "", "Specify the DB instance name.")
	flag.StringVar(&isProd, "is-prod", "", "Is the environment belonging to production?")
	flag.StringVar(&syslogEnable, "syslog-enable", "", "Is the environment need syslog enabled?")
	flag.StringVar(&dsmPolicyIdWebServersAutoScale, "dsm-policy-id-web-auto", "", "Specify the DSM policy id of auto scaling web servers.")
	flag.StringVar(&dsmPolicyIdWebServersOther, "dsm-policy-id-web-other", "", "Specify the DSM policy id of other web servers.")
	flag.StringVar(&dsmPolicyIdOther, "dsm-policy-id-other", "", "Specify the DSM policy id of other servers.")
	flag.StringVar(&dsmGroupIdWebServersAutoScale, "dsm-group-id-web-auto", "", "Specify the DSM group id of auto scaling web servers.")
	flag.StringVar(&dsmGroupIdWebServersOther, "dsm-group-id-web-other", "", "Specify the DSM group id of other web servers.")
	flag.StringVar(&dsmGroupIdOther, "dsm-group-id-other", "", "Specify the DSM group id of other servers.")
	flag.Parse()

	if fullEnv == "" {
		fmt.Println("Please make sure you have offered the full env flag.")
		os.Exit(1)
	}
	if !strings.Contains(fullEnv, "aws") && !strings.Contains(fullEnv, "azure") && !strings.Contains(fullEnv, "arena") {
		fmt.Println("Please make sure 'aws' or 'azure' or 'arena' in full-env flag .")
		os.Exit(1)
	}
	if awsAccount == "" {
		fmt.Println("Pls make sure you have offered the aws account flag.")
		os.Exit(1)
	}
	if region == "" {
		fmt.Println("Pls make sure you have offered the region flag.")
		os.Exit(1)
	}
	if baseDomain == "" {
		fmt.Println("Pls make sure you have offered the base domain flag.")
		os.Exit(1)
	}
	if cidr == "" {
		fmt.Println("Pls make sure you have offered the cidr flag.")
		os.Exit(1)
	}
	if dbName == "" {
		fmt.Println("Pls make sure you have offered the db name flag.")
		os.Exit(1)
	}
	if isProd == "" {
		fmt.Println("Pls make sure you have offered the is-prod flag.")
		os.Exit(1)
	}
	if syslogEnable == "" {
		fmt.Println("Pls make sure you have offered the syslog-enable flag.")
		os.Exit(1)
	}

	if dsmPolicyIdWebServersAutoScale == "" {
		fmt.Println("Pls make sure you have offered the dsmPolicyIdWebServersAutoScale flag.")
		os.Exit(1)
	}
	if dsmPolicyIdWebServersOther == "" {
		fmt.Println("Pls make sure you have offered the dsmPolicyIdWebServersOther flag.")
		os.Exit(1)
	}
	if dsmPolicyIdOther == "" {
		fmt.Println("Pls make sure you have offered the dsmPolicyIdOther flag.")
		os.Exit(1)
	}
	if dsmGroupIdWebServersAutoScale == "" {
		fmt.Println("Pls make sure you have offered the dsmGroupIdWebServersAutoScale flag.")
		os.Exit(1)
	}
	if dsmGroupIdWebServersOther == "" {
		fmt.Println("Pls make sure you have offered the dsmGroupIdWebServersOther flag.")
		os.Exit(1)
	}
	if dsmGroupIdOther == "" {
		fmt.Println("Pls make sure you have offered the dsmGroupIdOther flag.")
		os.Exit(1)
	}

	// make sure the full-env flag contains aws or azure or arena
	if strings.Contains(fullEnv, "aws") {
		env = strings.Split(fullEnv, "aws")[1]
	} else if strings.Contains(fullEnv, "azure") {
		env = strings.Split(fullEnv, "azure")[1]
	} else if strings.Contains(fullEnv, "arena") {
		env = fullEnv
	}

	// Default value
	checkIfProduction := "0"
	deploymentInfrastructure := "AWS_REGULAR"
	emailSubjectPrepend := `\\[Non-production Site\\] `
	emailRcptDomainsAllowed := "ptc.com,mailtest.dev.bom.com,esharing.com"
	emailSenderEmployeesAllowed := strings.Join(utils.ReadFile(ansibleHome+"/arena/etc/email-sender-employees-allowed.txt"), ",")
	esCluster := "es01.util.aws.bom.com:9200,es02.util.aws.bom.com:9200,es03.util.aws.bom.com:9200"
	featureTfaDisable := fmt.Sprintf("{{ lookup('aws_ssm', '/%s/plm/feature/tfa/disable') | default('false', true) }}", env)
	feedbackEmail := "arena-dev-feedback@ptc.com"
	applicationMonitoringEmail := "arena-dev-app-monitor@ptc.com"
	databaseMonitoringEmail := "arena-dev-db-monitor@ptc.com"
	loggingHost := "syslog.util.aws.bom.com"
	marketingSite := "arena-marketing-dev.arenasolutions.com"
	marketingSitePathPrefix := "arena"
	enableLanguageSetting := "On"
	httpHeadersSourceIp := "True-Client-Ip,Cf-Connecting-Ip"
	ldapDomain := "aws.bom.com"
	ldapAccessGroupPrefix := "ArenaDev"

	// alternative value if site is Production.
	if isProd == "true" {
		checkIfProduction = "1"
		emailSubjectPrepend = ""
		emailRcptDomainsAllowed = ""
		emailSenderEmployeesAllowed = ""
		featureTfaDisable = "false"
		marketingSite = "arena-marketing.arenasolutions.com"
		enableLanguageSetting = "Off"
		applicationMonitoringEmail = "arena-app-monitor@ptc.com"
		databaseMonitoringEmail = "arena-db-monitor@ptc.com"
	}

	if fullEnv == "arenaeurope" {
		ldapAccessGroupPrefix = "ArenaEurope"
		feedbackEmail = "arena-europe-feedback@ptc.com"
		marketingSitePathPrefix = "arenaeurope"
	}

	if fullEnv == "arenagov" {
		ldapDomain = "arenagov.com"
		ldapAccessGroupPrefix = "ArenaGov"
		deploymentInfrastructure = "AWS_GOV_CLOUD"
		esCluster = "es01.util.arenagov.com:9200,es02.util.arenagov.com:9200,es03.util.arenagov.com:9200"
		feedbackEmail = "arena-itar-feedback@ptc.com"
		loggingHost = "logr.util.arenagov.com"
		marketingSitePathPrefix = "arenagov"
	}

	if fullEnv == "awsgvc" {
		ldapDomain = "arenagov.com"
	}

	// generate ldapBaseDn
	tmp_lbd := strings.Split(ldapDomain, ".")
	ldapBaseDn := "dc=" + strings.Join(tmp_lbd, ",dc=")

	// generate tldDomain
	tmp := strings.Split(baseDomain, ".")
	tldDomain := strings.Join(tmp[len(tmp)-2:], ".")

	// get the config name from ssm according to env
	configMap := utils.ReadThirdPartyServiceConfig(ansibleHome)

	// parse the third party yaml file and set the gooddata & salesforce config
	utils.DebugByFile(fmt.Sprintf("%v\n", configMap))

	// get the default config values in the yaml
	defaultThirdPartyConfig, _ := configMap["default"]

	// get the related map value of third party config yaml
	var dataThirdPartyConfig utils.ThirdPartyConfig
	var configName string

	if isProd == "true" {
		tmpConfig, ok := configMap[env]
		if !ok {
			fmt.Printf("key %s not exist in third party yml config", env)
			os.Exit(1)
		}
		dataThirdPartyConfig = tmpConfig
		configName = env
	} else {
		ssmConfigName, exit := utils.GetSsmValue(region, fmt.Sprintf("/%s/thirdparty/service/config", env))
		if !exit {
			fmt.Printf("This env %s do not have ssm key /%s/thirdparty/service/config", env, env)
			os.Exit(1)
		}
		tmpConfig, ok := configMap[ssmConfigName]
		if !ok {
			fmt.Printf("key %s not exist in third party yml config", env)
			os.Exit(1)
		}
		dataThirdPartyConfig = tmpConfig
		configName = ssmConfigName
	}

	goodDataHost := dataThirdPartyConfig.GoodData.Host
	if goodDataHost == "" {
		goodDataHost = defaultThirdPartyConfig.GoodData.Host
	}
	goodDataLogin := dataThirdPartyConfig.GoodData.Login
	if goodDataLogin == "" {
		goodDataLogin = defaultThirdPartyConfig.GoodData.Login
	}
	ssmGooddataPassword, _ := utils.GetSsmValue(region, fmt.Sprintf("/vendorinfo/gooddata/%s/password", configName))
	ssmGooddataPgpPassword, _ := utils.GetSsmValue(region, fmt.Sprintf("/vendorinfo/gooddata/%s/pgp/password", configName))

	salesforceLeadURL := dataThirdPartyConfig.SalesForce.LeadUrl
	if salesforceLeadURL == "" {
		salesforceLeadURL = defaultThirdPartyConfig.SalesForce.LeadUrl
	}
	salesforceArenaOrgId := dataThirdPartyConfig.SalesForce.ArenaOrgId
	if salesforceArenaOrgId == "" {
		salesforceArenaOrgId = defaultThirdPartyConfig.SalesForce.ArenaOrgId
	}
	salesforceCustomerPortalLoginUrl := dataThirdPartyConfig.SalesForce.CustomerPortalLoginUrl
	if salesforceCustomerPortalLoginUrl == "" {
		salesforceCustomerPortalLoginUrl = defaultThirdPartyConfig.SalesForce.CustomerPortalLoginUrl
	}

	salesforceLoginId := dataThirdPartyConfig.SalesForce.LoginId
	if salesforceLoginId == "" {
		salesforceLoginId = defaultThirdPartyConfig.SalesForce.LoginId
	}
	salesforceLoginPassword := dataThirdPartyConfig.SalesForce.LoginPassword
	if salesforceLoginPassword == "" {
		salesforceLoginPassword = defaultThirdPartyConfig.SalesForce.LoginPassword
	}
	salesForceWebToCaseOrigin := dataThirdPartyConfig.SalesForce.WebToCaseOrigin
	if salesForceWebToCaseOrigin == "" {
		salesForceWebToCaseOrigin = defaultThirdPartyConfig.SalesForce.WebToCaseOrigin
	}
	salesForceWebToCaseUrl := dataThirdPartyConfig.SalesForce.WebToCaseUrl
	if salesForceWebToCaseUrl == "" {
		salesForceWebToCaseUrl = defaultThirdPartyConfig.SalesForce.WebToCaseUrl
	}

	// new the info struct for rendering
	i := Info{
		FullEnv:                          fullEnv,
		Env:                              env,
		IsProd:                           isProd,
		AwsAccount:                       awsAccount,
		Region:                           region,
		BaseDomain:                       baseDomain,
		TldDomain:                        tldDomain,
		CIDR:                             cidr,
		LdapDomain:                       ldapDomain,
		LdapBaseDn:                       ldapBaseDn,
		LdapAccessGroupPrefix:            ldapAccessGroupPrefix,
		DbName:                           dbName,
		CheckIfProduction:                checkIfProduction,
		DeploymentInfrastructure:         deploymentInfrastructure,
		EmailSubjectPrepend:              emailSubjectPrepend,
		EmailRcptDomainsAllowed:          emailRcptDomainsAllowed,
		EmailSenderEmployeesAllowed:      emailSenderEmployeesAllowed,
		ApplicationMonitoringEmail:       applicationMonitoringEmail,
		DatabaseMonitoringEmail:          databaseMonitoringEmail,
		EsCluster:                        esCluster,
		FeatureTfaDisable:                featureTfaDisable,
		FeedbackEmail:                    feedbackEmail,
		LoggingHost:                      loggingHost,
		MarketingSite:                    marketingSite,
		MarketingSitePathPrefix:          marketingSitePathPrefix,
		SyslogEnable:                     syslogEnable,
		DsmPolicyIdWebServersAutoScale:   dsmPolicyIdWebServersAutoScale,
		DsmPolicyIdWebServersOther:       dsmPolicyIdWebServersOther,
		DsmPolicyIdOther:                 dsmPolicyIdOther,
		DsmGroupIdWebServersAutoScale:    dsmGroupIdWebServersAutoScale,
		DsmGroupIdWebServersOther:        dsmGroupIdWebServersOther,
		DsmGroupIdOther:                  dsmGroupIdOther,
		EnableLanguageSetting:            enableLanguageSetting,
		SalesforceLeadURL:                salesforceLeadURL,
		SalesforceArenaOrgId:             salesforceArenaOrgId,
		SalesforceCustomerPortalLoginUrl: salesforceCustomerPortalLoginUrl,
		SalesforceLoginId:                salesforceLoginId,
		SalesforceLoginPassword:          salesforceLoginPassword,
		SalesForceWebToCaseOrigin:        salesForceWebToCaseOrigin,
		SalesForceWebToCaseUrl:           salesForceWebToCaseUrl,
		HttpHeadersSourceIp:              httpHeadersSourceIp,
		GoodDataHost:                     goodDataHost,
		GoodDataLogin:                    goodDataLogin,
		GoodDataPassword:                 ssmGooddataPassword,
		GoodDataPgpPassword:              ssmGooddataPgpPassword,
	}

	// define the function to be used in the template, so that increase the efficiency
	fm := template.FuncMap{
		"toUpper": func(str string) string {
			return strings.ToUpper(str)
		},
		"toLower": func(str string) string {
			return strings.ToLower(str)
		},
		"trimSpace": func(str string) string {
			return strings.TrimSpace(str)
		},
	}

	// read the templates from the arena/template dir below repo
	files, err := ioutil.ReadDir(ansibleHome + "/arena/template")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// check the templates one by one
	for _, f := range files {
		if !f.IsDir() {
			if strings.Contains(f.Name(), "_template") {
				// define the templates directory prefix
				templateFile := ansibleHome + "/arena/template/" + f.Name()
				tmpl, err := template.New(f.Name()).Funcs(fm).Delims("<<", ">>").ParseFiles(templateFile)
				if err != nil {
					fmt.Println("fail to new template", err)
					os.Exit(1)
				}
				// replace the template suffix with the full-env flag
				fileName := strings.ReplaceAll(f.Name(), "template", fullEnv)
				if strings.Contains(fileName, "tag_Role") {
					// write the rendered template to group_vars
					err = tmpl.Execute(utils.WriteFile(ansibleHome, fileName), i)
					if err != nil {
						fmt.Println("fail to render", err)
						os.Exit(1)
					}
				} else {
					// write the rendered template to group_vars/all
					err = tmpl.Execute(utils.WriteGlobalFile(ansibleHome, fileName), i)
					if err != nil {
						fmt.Println("fail to render", err)
						os.Exit(1)
					}
				}

			}

		}
	}

}
