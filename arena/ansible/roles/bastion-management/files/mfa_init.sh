#!/usr/bin/env bash

# Install google-authenticator package.
if ! rpm -q google-authenticator >/dev/null 2>&1; then
  # google-authenticator is provided by EPEL
  rpm -q epel-release >/dev/null 2>&1 || amazon-linux-extras install epel -y
  yum --assumeyes --quiet install google-authenticator
  yum --assumeyes --quiet erase epel-release && amazon-linux-extras disable epel
fi

# Cleanup
sed -i '/pam_google_authenticator.so/d' /etc/pam.d/sshd
sed -i '/auth\s\+required\s\+pam_permit.so/d' /etc/pam.d/sshd

# Use pam_google_authenticator.so.
cat << E0F >> /etc/pam.d/sshd
# Use pam_google_authenticator.so
auth required pam_google_authenticator.so echo_verification_code nullok [authtok_prompt=Your secret token: ]

# Keep 'auth required pam_permit.so' the **LAST** line of this file.
auth required pam_permit.so
E0F

# Update sshd_config for MFA
sed -i 's/^ChallengeResponseAuthentication\s\{1,\}no$/#ChallengeResponseAuthentication no/g' /etc/ssh/sshd_config
sed -i 's/^\(AuthenticationMethods\s.*\)$/# DISABLED BY MFA: \1/g' /etc/ssh/sshd_config

cat << E0F >> /etc/ssh/sshd_config

# Firstboot - bastion: Required by MFA implementation
ChallengeResponseAuthentication yes
AuthenticationMethods publickey,keyboard-interactive
E0F

systemctl restart sshd