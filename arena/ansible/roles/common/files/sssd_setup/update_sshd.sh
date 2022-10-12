#!/usr/bin/env bash

# Remove Amazon customized unit file for sshd.
if [ -d "/usr/lib/systemd/system/ssh.service.d/" ]; then
  rm -fr /usr/lib/systemd/system/ssh.service.d/
  systemctl daemon-reload
fi

# Remove any existing setting.
sed -i '/^AuthorizedKeysCommand /d' /etc/ssh/sshd_config
sed -i '/^AuthorizedKeysCommandUser /d' /etc/ssh/sshd_config

# Configure sshd to use SSSD for publickey source.
cat << E0F >> /etc/ssh/sshd_config
# Arena SSSD configuration by firstboot.
AuthorizedKeysCommand /usr/bin/sss_ssh_authorizedkeys %u
AuthorizedKeysCommandUser nobody
E0F

# Enforce publickey access only via ssh
sed -i '/^AuthenticationMethods\s+/d' /etc/ssh/sshd_config
cat << E0F >> /etc/ssh/sshd_config
# Firstboot: Use publickey method only
AuthenticationMethods publickey
E0F

systemctl restart sshd

# Disable password auth from PAM
sed -i '/auth\s\+substack\s\+password-auth/d' /etc/pam.d/sshd