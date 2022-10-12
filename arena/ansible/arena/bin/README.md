## Tag role Render

#### Introduction
> To render the tag role template file for repo https://git.dev.bom.com/cloud-deployment/ansible

#### Attention
> Please make sure you have set the ANSIBLE_HOME environment variable before using.

#### Usage
This tool would replace below ones from provided flags

- <<.AwsAccount>> 
- <<.Region>>
- <<.FullEnv>>
- <<.Env>>	// Split the full-env with 'aws' or 'azure' to get the short environment name.
- <<.BaseDomain>>
- <<.DbName>>
- <<.IsProd>> // bool

```
$ go run ./tag-role-render.go -h
Usage of ./tag-role-render-darwin-amd64:
  -aws-account string
    	Specify the aws account like 089875288533.
  -base-domain string
    	Specify the base domain like arenagov.com, qa.aws.bom.com, sand.aws.bom.com.
  -full-env string
    	Specify the full env like awsqa, awssand.
  -region string
    	Specify the aws region like us-east-2, us-gov-east-1.
  -db-name string
      Specify the database instance name like euawsdb
  -is-prod [true|false](string)
      Specify whether it's production "true" or "false"
```

#### Template functions
Here are some functions that could be used in the template
- toUpper (change the string to upper case)
```
<<.FullEnv | toUpper>>
```
- toLower (change the string to lower case)
```
<<.FullEnv | toLower>>
```
- trimSpace (trim the string left & right space)
```
<<.FullEnv | trimSpace>>
```

## Example of running tag-role-render.go script
### export ANSIBLE_HOME,ansible repo under $HOME path 
```
export ANSIBLE_HOME=$HOME/ansible
```
### export metadata
```
cat <<METADATA > $HOME/metadata.txt
base_domain_name:dev.aws.bom.com
aws_account_id:089875288533
aws_region:us-west-2
target_env_name:awsdev
vpc_cidr_block:10.60.0.0/16
target_db_name:devawsdb
is_production:false
syslog_enable:false
dsm_policy_id_webservers_autoscale:11
dsm_policy_id_webservers_other:11
dsm_policy_id_other:12
dsm_group_id_webservers_autoscale:2
dsm_group_id_webservers_other:7
dsm_group_id_other:1
METADATA
```
```
AWS_ACCOUNT_ID=$(grep aws_account_id $HOME/metadata.txt | cut -d: -f2)
AWS_REGION=$(grep aws_region $HOME/metadata.txt | cut -d: -f2)
PLM_ENV_NAME=$(grep target_env_name $HOME/metadata.txt | cut -d: -f2)
BASE_DOMAIN_NAME=$(grep base_domain_name $HOME/metadata.txt | cut -d: -f2)
VPC_CIDR_BLOCK=$(grep vpc_cidr_block $HOME/metadata.txt | cut -d: -f2)
DB_NAME=$(grep db_name $HOME/metadata.txt | cut -d: -f2)
IS_PROD=$(grep is_production $HOME/metadata.txt | cut -d: -f2)
SYSLOG_ENABLE=$(grep syslog_enable $HOME/metadata.txt | cut -d: -f2)
DSM_POLICY_ID_WEBSERVERS_AUTOSCALE=$(grep dsm_policy_id_webservers_autoscale $HOME/metadata.txt | cut -d: -f2)
DSM_GROUP_ID_WEBSERVERS_AUTOSCALE=$(grep dsm_group_id_webservers_autoscale $HOME/metadata.txt | cut -d: -f2)
DSM_POLICY_ID_WEBSERVERS_OTHER=$(grep dsm_policy_id_webservers_other $HOME/metadata.txt | cut -d: -f2)
DSM_GROUP_ID_WEBSERVERS_OTHER=$(grep dsm_group_id_webservers_other $HOME/metadata.txt | cut -d: -f2)
DSM_POLICY_ID_OTHER=$(grep dsm_policy_id_other $HOME/metadata.txt | cut -d: -f2)
DSM_GROUP_ID_OTHER=$(grep dsm_group_id_other $HOME/metadata.txt | cut -d: -f2)

```
### run script
- make sure golang have been installed
```
go version
```
- run with go command
```
go run $ANSIBLE_HOME/arena/bin/tag-role-render.go \
-aws-account "$AWS_ACCOUNT_ID" \
-full-env "$PLM_ENV_NAME" \
-region "$AWS_REGION" \
-base-domain "$BASE_DOMAIN_NAME" \
-cidr "$VPC_CIDR_BLOCK" \
-db-name "$DB_NAME" \
-is-prod "$IS_PROD" \
-syslog-enable "$SYSLOG_ENABLE" \
-dsm-policy-id-web-auto "$DSM_POLICY_ID_WEBSERVERS_AUTOSCALE" \
-dsm-policy-id-web-other "$DSM_POLICY_ID_WEBSERVERS_OTHER" \
-dsm-policy-id-other "$DSM_POLICY_ID_OTHER" \
-dsm-group-id-web-auto "$DSM_GROUP_ID_WEBSERVERS_AUTOSCALE" \
-dsm-group-id-web-other "$DSM_GROUP_ID_WEBSERVERS_OTHER" \
-dsm-group-id-other "$DSM_GROUP_ID_OTHER"
```
- find results in path $ANSIBLE_HOME/inventory/group_vars
```
ls -lrat $ANSIBLE_HOME/inventory/group_vars
```