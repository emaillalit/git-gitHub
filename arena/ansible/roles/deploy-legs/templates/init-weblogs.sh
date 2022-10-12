#!/bin/bash

mkdir -p ~/.ssh
chmod 700 ~/.ssh

aws s3 cp s3://legs.{{ aws_region }}.{{ aws_account_id }}/legs ~/.ssh/legs
chmod 400 ~/.ssh/legs

LEGS_SSH_KEY_ID=$(aws iam list-ssh-public-keys --user-name APP-gitlab-to-aws-legs --query "SSHPublicKeys[*].SSHPublicKeyId" --output text)

cat <<EOF > /home/ARENA/weblogs/.ssh/config
Host git-codecommit.*.amazonaws.com
User $LEGS_SSH_KEY_ID
IdentityFile ~/.ssh/legs
StrictHostKeyChecking no
EOF

chmod 600 ~/.ssh/config

ssh-keyscan git-codecommit.{{ aws_region }}.amazonaws.com >> ~/.ssh/known_hosts

chown weblogs -R /home/ARENA/weblogs

cd ~ && git clone ssh://git-codecommit.{{ aws_region }}.amazonaws.com/v1/repos/legs

chmod +x ~/legs/refresh_and_start
chmod +x ~/legs/deploy/db/load-db.sh
chmod +x ~/legs/deploy/es2gs/load-es2gs.sh
chmod +x ~/legs/deploy/log2es/load-log2es.sh

