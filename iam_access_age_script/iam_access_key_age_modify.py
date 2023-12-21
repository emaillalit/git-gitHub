import boto3
import logging
import os
import re
import dateutil.parser
from datetime import datetime, timedelta, timezone 
from botocore.exceptions import ClientError
from botocore.config import Config
from smtplib import SMTPException, SMTP
from email.mime.multipart import MIMEMultipart
from email.mime.text import MIMEText


def get_logger():
    logger = logging.getLogger()
    if not logger.hasHandlers():
        logger.setLevel(logging.INFO) # better to have too much log than not enough
        shandler = logging.StreamHandler()
        formatter = logging.Formatter('[%(asctime)s] %(message)s')
        shandler.setFormatter(formatter)
        logger.addHandler(shandler)
    return logger

def get_client(service):
    config = Config(
        retries={
            'max_attempts': 0,
            'mode': 'standard'
        }
    )

    try:
        session = boto3.Session(profile_name='eng')
        return session.client(service, config=config)
    except ClientError as ex:
        print(f"{ex.response['Error']['Message']} = \"An error occurred\"")
        raise ex

def get_all_users(client):
    list_to_return = []
    for user in client.list_users()['Users']:
        user_name = user['UserName']
        user_info = {'username': user_name, 'arn': user['Arn']}
        # FILTER: by username
        # 
        if re.search(r'^APP|AWSAPP|myrmidon', user_name):
            continue
        # FILTER: by aws tag 'type' value
        for tag in client.list_user_tags(UserName=user_name)['Tags']:
            user_info[tag['Key'].lower()] = tag['Value']

        if 'type' in user_info and user_info['type'] in ['bot', 'robot']:
            continue
        keys_list = client.list_access_keys(UserName=user['UserName'])
        user_info['key_list'] = keys_list.get('AccessKeyMetadata')
        list_to_return.append(user_info)

    return list_to_return    

def check_user_key_age(client, user_list, key_age_limit):
    account_alias = client.list_account_aliases()['AccountAliases'][0]

    for user in user_list:
        accountnumber = user['arn'].split('::')[1].split(':')[0]
        for key in user['key_list']:
            if key.get('Status').casefold() == 'inactive':
                continue
            key_age = get_key_age(key.get('CreateDate'))
            # WARNING
            if (
                get_key_age(key.get('CreateDate')) < timedelta(key_age_limit.get('critical')) and 
                get_key_age(key.get('CreateDate')) > timedelta(key_age_limit.get('warning'))
                ):
                print(f"WARNING: {key_age} KEY -> {key}")
                # FILTER: Email tag
                email_id = user.get('email') 
                email_tag = user.get('name')
                subject_line = (f"[***--Reminder Notice--***] Remember to rotate your AWS Keys on "
                                f"{accountnumber} {account_alias} older than {key_age_limit.get('warning')} days!")
                send_notification(key, account_alias, email_id, subject_line, email_tag, accountnumber, 'warning')
            # CRITICAL
            elif get_key_age(key.get('CreateDate')) > timedelta(key_age_limit.get('critical')):
                print(f"CRITICAL: {key_age} KEY -> {key}")
                #access_key_deactivate(client, key.get('UserName'), key.get('AccessKeyId'), 'Active')
                access_key_deactivate(client, 'test', 'test', 'test')
                email_id = user.get('email') 
                email_tag = user.get('name')
                subject_line = (f"[***--Deactivation Notice--***] your AWS Access Key on {accountnumber} "
                                f"{account_alias} older than {key_age_limit.get('critical')} days has been deactivated!")
                send_notification(key, account_alias, email_id, subject_line, email_tag, accountnumber, 'critical')


def get_key_age(create_date):
    today = datetime.utcnow().replace(tzinfo=dateutil.tz.tzutc())
    time_delta = today - create_date
    return time_delta

def access_key_deactivate(client, user_name, access_key_id, status_value):
    mylogger = get_logger()
    print(f"CLIENT -> {client}")
    print(f"USERNAME -> {user_name}")
    print(f"ACCESS KEY ID -> {access_key_id}")
    print(f"STATUS VALUE -> {status_value}")
    user_name = 'test2'
    access_key_id = 'AKIARJ3H4CHK5OUSN4BX'
    status_value = 'Active'
    try:
        res = client.update_access_key (
            UserName=user_name,
            AccessKeyId=access_key_id,
            Status=status_value
        )
    except ClientError as err:
        my_logger.error(err.res['Error']['Message'])
        return False
    mylogger.info(f"Following {access_key_id} key for {user_name} user is Deactivated.")
    return True

def send_notification(key, accountid, emailid, subject, email_tag, accountnumber, key_age):
    if key_age == 'warning':
        print(f"KEY AGE -> {key_age}")
        print(f"accountid -> {accountid}")
        print(f"email id -> {emailid}")
        print(f"subject -> {subject}")
        print(f"emailtag -> {email_tag}")
        print(f"accountnumber -> {accountnumber}")
        print(f"KEY -> {key}")
        BODY_HTML = f'''\
    <p>Dear {email_tag},</p>
    <p></p>
    <p>This is a automatic reminder to rotate your AWS Access keys at least every 90 days.</p>
    <p>At the moment, you have <b><span style="color:maroon">{key.get('AccessKeyId')}</span></b> on AWS account <b>{accountnumber}</b> <b>{accountid}</b></p>
    <p>which was created <b>{key.get('CreateDate')}</b> seconds ago.</p>
    <p>To learn how to rotate your AWS Access Key, Please read the official guide at </p>
    <p>
    https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_access-keys.html#Using_RotateAccessKey
    </p>
    <p>If you have any question, please don't hesitate to contact Arena DevOps Team at arena-devops@ptc.com.</p>
    <p>This automatic reminder will be sent out again in each day, if the key(s) was not rotated.</p>
    </br>
    <p>Regards,</p>
    <p>Your Support Team</p>
    '''
        print(f"BODY HTML FOR WARNING")
        print(f"{BODY_HTML}")
        #send_email(emailid, BODY_HTML, subject)
    elif key_age == 'critical':
        print(f"KEY AGE -> {key_age}")
        print(f"accountid -> {accountid}")
        print(f"email id -> {emailid}")
        print(f"subject -> {subject}")
        print(f"emailtag -> {email_tag}")
        print(f"accountnumber -> {accountnumber}")
        print(f"KEY -> {key}")
        BODY_HTML = f'''\
    <p> Dear {email_tag},</p>
    <p></p>
    <p><b>This is a notification email to notify you that your AWS Access keys has been deactivated now.</b></p>
    <p>Following <b><span style="color:maroon">{key.get('AccessKeyId')}</span></b> key on AWS account <b>{accountnumber}</b> <b>{accountid}</b></p>
    <p>which was created <b>{key.get('CreateDate')}</b> seconds ago.</p>
    <p>Due to security reason, keys which are older than 90 days required to be rotated or deleted.</p>
    <p><b>Please create new access keys and delete old inactive keys!</b></p>
    <p>If you have any question, please don't hesitate to contact Arena DevOps Team at arena-devops@ptc.com.</p>
    <p>Regards,</p>
    <p>Your Support Team</p>
    '''
        print(f"BODY HTML FOR CRITICAL")
        print(f"{BODY_HTML}")
        #send_email(emailid, BODY_HTML, subject)
    return BODY_HTML

def check(mail):
    regex = r'\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b'
    if (re.fullmatch(regex, mail)):
        return True
    else:
        return False

def send_email(email_id, email_body, subject_line):
    HOST = os.environ.get('smtp_host') 
    PORT = os.environ.get('smtp_port') 
    SENDER = os.environ.get('sender') 
    RECIPIENT = email_id
    USER = os.environ.get('smtp_user') 
    PASSWORD = os.environ.get('smtp_password') 
    SUBJECT = subject_line
    msg = MIMEMultipart('alternative')
    msg['Subject'] = SUBJECT
    msg['From'] = SENDER
    msg['To'] = RECIPIENT

    part1 = MIMEText(email_body, 'html')
    msg.attach(part1)
    my_logger = get_logger()

    if check(RECIPIENT):
        try:
            server = SMTP(HOST, PORT)
            server.starttls()
            server.login(USER, PASSWORD)
            server.sendmail(SENDER, RECIPIENT, msg.as_string())
            server.quit()
        except SMTPException as e:
            my_logger.error(f"Error: Email sent to {RECIPIENT} failed with following error: {e}")
            return False
        my_logger.info(f"Email sent to {RECIPIENT} successfully.")
        return True
    


def main(event=None, context=None):
    client = get_client('iam')
    user_list = get_all_users(client)

    key_age_limit = {
        "warning": 30,
        "critical": 40
    }

    check_user_key_age(client, user_list, key_age_limit)

if __name__ == "__main__":
    main()