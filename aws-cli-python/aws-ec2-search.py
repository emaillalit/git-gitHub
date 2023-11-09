import argparse
import boto3
import logging
import sys
from botocore.exceptions import ClientError

logger = logging.getLogger(__name__)

def get_profile(target_env):
    aws_profile = dict()
    if target_env == "arenagov":
        aws_profile['aws_profile'] = 'itar'
        aws_profile['aws_region'] = 'us-gov-west-1'
        aws_profile['aws_dr_region'] = 'us-gov-east-1'
        aws_profile['base_domain'] = 'arenagov.com'
    elif target_env == "awsgvc":
        aws_profile['aws_profile'] = 'itar'
        aws_profile['aws_region'] = 'us-gov-west-1'
        aws_profile['base_domain'] = 'gvc.arenagov.com'
    elif target_env == "arenaeurope":
        aws_profile['aws_profile'] = 'prod'
        aws_profile['aws_region'] = 'eu-central-1'
        aws_profile['aws_dr_region'] = 'eu-west-1'
        aws_profile['base_domain'] = 'europe.arenaplm.com'
    elif target_env == "awseuc":
        aws_profile['aws_profile'] = 'prod'
        aws_profile['aws_region'] = 'eu-central-1'
        aws_profile['base_domain'] = 'euc.arenaplm.com'
    elif target_env == "arenachina":
        aws_profile['aws_profile'] = 'prodcn'
        aws_profile['aws_region'] = 'cn-northwest-1'
        aws_profile['base_domain'] = 'arenaplm.cn'
    elif target_env == "awscnc":
        aws_profile['aws_profile'] = 'prodcn'
        aws_profile['aws_region'] = 'cn-northwest-1'
        aws_profile['aws_dr_region'] = 'cnc.arenaplm.cn'
    elif target_env == "arenaus":
        aws_profile['aws_profile'] = 'prod'
        aws_profile['aws_region'] = 'us-west-2'
        aws_profile['aws_dr_region'] = 'us-east-1'
        aws_profile['base_domain'] = 'us.arenaplm.com'
    elif target_env == "awsusc":
        aws_profile['aws_profile'] = 'prod'
        aws_profile['aws_region'] = 'us-west-2'
        aws_profile['base_domain'] = 'usc.arenaplm.com'
    else:
        aws_profile['aws_profile'] = 'eng'
        aws_profile['aws_region'] = 'us-west-2'
        aws_profile['base_domain'] = 'qa.aws.bom.com'
    
    return aws_profile

def get_resource(service, target_env):
    aws_profile = get_profile(target_env)
    
    try:
        session = boto3.Session(profile_name=aws_profile.get("aws_profile"), region_name=aws_profile.get("aws_region"))
        return session.client(service)
    except ClientError as ex:
        ex.response['Error']['Message'] = "An error occurred"
        logging.error(ex)

def filter_instance(args):
    """
    Describe instance based on the filters
    """
    ec2_resource = get_resource('ec2', args.target_env)
    if not ec2_resource:
        logging.error(f"Cannot get resources for ec2...script exited")
        sys.exit()

    if args.id is not None:
        """
        response = ec2_resource.describe_instances(
            InstanceIds = [
                args.id
                ],
                )
        """
        response = ec2_resource.describe_instances(
            Filters=[
                {
                    'Name': 'instance-id',
                    'Values': [
                        args.id,
                    ],
                },
            ],
        )
    elif args.ip is not None:
        response = ec2_resource.describe_instances(
            Filters=[
                {
                    'Name': 'private-ip-address',
                    'Values': [
                        args.ip,
                    ],
                },
            ],
        )
    elif args.eip is not None:
        response = ec2_resource.describe_instances(
            Filters=[
                {
                    'Name': 'public-ip-address',
                    'Values': [
                        args.eip,
                    ],
                },
            ],
        )
    elif args.name is not None:
        response = ec2_resource.describe_instances(
            Filters=[
                {
                    'Name': 'tag:Name',
                    'Values': [
                        args.name,
                    ],
                },
            ],
        )

    return response
     
def main():
    parser = argparse.ArgumentParser(
        description="ec2 search"
    )
    parser.add_argument('--target_env', help="Target Environment ex:(AWSQA...etc)", required=True)
    parser.add_argument('--name', help="Search instance by its name")
    parser.add_argument('--id', help="Search instance by it's id")
    parser.add_argument('--ip', help="Search instance by it's ip address")
    parser.add_argument('--eip', help="Search instance by it's eip address")
    args = parser.parse_args()

    response = filter_instance(args)
    for k, v in response.items():
        for ele in v:
            if isinstance(ele, dict):
                for k1, v1 in ele.items():
                    if k1 == 'Instances':
                        for ele2 in v1:
                            print (f"InstancId: {ele2.get('InstanceId')}")
                            print (f"InstanceType: {ele2.get('InstanceType')}")
                            print (f"PrivateIp: {ele2.get('PrivateIpAddress')}")
                            print (f"KeyName: {ele2.get('KeyName')}")
                            print (f"State: {ele2.get('State').get('Name')}")
                            print(f"LaunchTime: {ele2.get('LaunchTime')}")

if __name__ == "__main__":
    main()

    