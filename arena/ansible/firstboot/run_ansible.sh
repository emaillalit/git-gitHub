#!/bin/bash

set -x

export HOME=/opt/arena
export ANSIBLE_HOME=$HOME/ansible
export ANSIBLE_INVENTORY=$ANSIBLE_HOME/inventory
export CONNECT_AS_USER=firstboot
export ANSIBLE_HOST_KEY_CHECKING="False" # disables asking "yes" for every ssh connection
export MS_TEAMS_CHANNEL_PREFIX="https://ptccloud.webhook.office.com/webhookb2/1d3f4a53-3725-4f4f-aed4-bac634e1ac63@b9921086-ff77-4d0d-828a-cb3381f678e2/IncomingWebhook/"
export MS_TEAMS_CHANNEL_SUFFIX="/c1ef2d00-7a4b-4c76-9109-eb99deae88e4"

instance_id="$1"

cd $ANSIBLE_HOME || exit

# After we get the signal from cloud-init, it takes a little bit for the ssh key to be installed
sleep 5

# Clean local Git repo
git fetch --all
git reset --hard --quiet origin/$(git branch --show-current)
git clean -f

# Pull the latest code
echo "--> Pulling the latest code"
git pull

# Get metadata
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

pushd $ANSIBLE_HOME/arena/bin
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o tag-role-render
popd

# Generate target PLM environment tag_Role files
pushd $ANSIBLE_INVENTORY/group_vars

$ANSIBLE_HOME/arena/bin/tag-role-render \
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
popd

## Get instance information
retry_instance=10
# Define a variable to determine if getting instance information was successful, 1 for failure, 0 for success
get_instance_success=1

while [ $retry_instance -gt 0 ]; do
  private_dns_name=$(aws ec2 describe-instances --instance-ids "$instance_id" | jq -r '.Reservations[].Instances[].PrivateDnsName')
  tags=$(aws ec2 describe-tags --filters "Name=resource-id,Values=$instance_id")
  exit_code=$?
  if [[ $exit_code == 0 ]]; then
    instance_role=$(echo "$tags" | jq -r '.Tags[] | select(.Key=="Role") | .Value')
    environment_tag="$(echo "$tags" | jq -r '.Tags[] | select(.Key=="Environment") | .Value')"
    if [[ "$instance_role" != "" ]]; then
      get_instance_success=0
      break
    fi
  fi

  retry_instance=$((retry_instance-1))
  sleep 5
done

if [ "$get_instance_success" == "0" ]; then
  echo "Get instance information successfully"
  echo "PRIVATE DNS NAME: $private_dns_name"
  echo "INSTANCE ROLE: $instance_role"
  echo "ENVIRONMENT TAG: $environment_tag"
else
  echo "Get instance information failed"
  exit 1
fi

# Send the deployment failed alert to individual MS Teams channel
case $environment_tag in
  awsdev)
    MS_TEAMS_CHANNEL_ID="ac3a237676d24e08ba1e9cf30c0a9f1a"
    ;;
  awsqa)
    MS_TEAMS_CHANNEL_ID="8f7aeb43f0de44e18262f2d9ba1f4eae"
    ;;
  awssand)
    MS_TEAMS_CHANNEL_ID="0c8b94540cdc45c0b1aa546f0e08d88c"
    ;;
  awsval)
    MS_TEAMS_CHANNEL_ID="b9984450bc3c46f29ca7d991fa22659d"
    ;;
  arenaeurope)
    MS_TEAMS_CHANNEL_ID="f9e489aa0a9e4559b6bdfa21a085d75d"
    ;;
  arenagov)
    MS_TEAMS_CHANNEL_ID="f21c4c43d52845749b1b5e348b486e6b"
    ;;
  awsgvc)
    MS_TEAMS_CHANNEL_ID="b84afaf4bf014f5d91537f08ce981d8f"
    ;;
  awseuc)
    MS_TEAMS_CHANNEL_ID="630e524d97a040e18ec2a046f84f097e"
    ;;
  *)
    MS_TEAMS_CHANNEL_ID="5eeaa79001484bf7b85d5b6d8eda878e"
    ;;
esac

MS_TEAMS_CHANNEL_FULL_URL="${MS_TEAMS_CHANNEL_PREFIX}${MS_TEAMS_CHANNEL_ID}${MS_TEAMS_CHANNEL_SUFFIX}"

## Trigger deployment
# Replace "-" with "_"
instance_role_underscores="${instance_role//-/_}"
# Define a variable to determine if the deployment was successful, 1 for failure, 0 for success
deployment_success=1
# Set the number of retries
retry_deployment=3

while [ $retry_deployment -gt 0 ]; do
  case "${instance_role}" in
    *firstboot*)
      make firstboot-management user=$CONNECT_AS_USER ROLE="$instance_role" ROLE_UNDERSCORES="$instance_role_underscores" LIMIT="$private_dns_name" REGION="$AWS_REGION" && deployment_success=0;;
    *bastion*)
      make bastion-management user=$CONNECT_AS_USER ROLE="$instance_role" ROLE_UNDERSCORES="$instance_role_underscores" REGION="$AWS_REGION" && deployment_success=0;;
    *redis*)
      make deploy-redis user=$CONNECT_AS_USER ROLE="$instance_role" ROLE_UNDERSCORES="$instance_role_underscores" REGION="$AWS_REGION" && deployment_success=0;;
    *rabbitmq*)
      make deploy-rabbitmq user=$CONNECT_AS_USER ROLE="$instance_role" ROLE_UNDERSCORES="$instance_role_underscores" REGION="$AWS_REGION" && deployment_success=0;;
    *mail*)
      make deploy-mail user=$CONNECT_AS_USER ROLE="$instance_role" ROLE_UNDERSCORES="$instance_role_underscores" REGION="$AWS_REGION" && deployment_success=0;;
    *metrics*)
      make deploy-metrics user=$CONNECT_AS_USER ROLE="$instance_role" ROLE_UNDERSCORES="$instance_role_underscores" REGION="$AWS_REGION" && deployment_success=0;;
    *fluentd*)
      make deploy-fluentd user=$CONNECT_AS_USER ROLE="$instance_role" ROLE_UNDERSCORES="$instance_role_underscores" REGION="$AWS_REGION" && deployment_success=0;;
    *backstop*)
      make deploy-backstop user=$CONNECT_AS_USER ROLE="$instance_role" ROLE_UNDERSCORES="$instance_role_underscores" LIMIT="$private_dns_name" REGION="$AWS_REGION" && deployment_success=0;;
    *)
      make c14n-deploy user=$CONNECT_AS_USER ROLE="$instance_role" ROLE_UNDERSCORES="$instance_role_underscores" LIMIT="$private_dns_name" REGION="$AWS_REGION" && deployment_success=0;;
  esac

  if [[ $deployment_success == 0 ]]; then
    echo "Deployment Success"
    exit 0
  fi

  retry_deployment=$((retry_deployment-1))
  sleep 60
done

echo "Deployment Failed"
# Send message to MS Teams when firstboot failed
curl -H "Content-Type: application/json" -d '{"text": "Deployment failed for '"$instance_id"' | '"$instance_role"', please take a look."}' $MS_TEAMS_CHANNEL_FULL_URL
